import schema from "src/graphql/schema";
import {
  queryRootConfig,
  expectQueryFromRoot,
  expectMutation,
} from "@lolopinto/ent-graphql-tests";
import { DB, LoggedOutViewer } from "@lolopinto/ent";
import CreateUserAction, {
  UserCreateInput,
} from "src/ent/user/actions/create_user_action";
import { randomEmail, random, randomPhoneNumber } from "src/util/random";
import { clearAuthHandlers } from "@lolopinto/ent/auth";
import { User } from "src/ent/";
import { Express } from "express";
import { PassportStrategyHandler } from "@lolopinto/ent-passport";
import supertest from "supertest";
import jwt from "jsonwebtoken";
import { encodeGQLID } from "@lolopinto/ent/graphql";

// TODO we need something that does this by default for all tests
afterAll(async () => {
  await DB.getInstance().endPool();
});

afterEach(() => {
  clearAuthHandlers();
});

function getUserRootConfig(
  user: User,
  partialConfig?: Partial<queryRootConfig>,
): queryRootConfig {
  return {
    schema: schema,
    root: "node",
    args: {
      id: encodeGQLID(user),
    },
    inlineFragmentRoot: "User",
    ...partialConfig,
  };
}

const loggedOutViewer = new LoggedOutViewer();
async function createUser(input?: Partial<UserCreateInput>): Promise<User> {
  return await CreateUserAction.create(loggedOutViewer, {
    firstName: "first",
    lastName: "last",
    emailAddress: randomEmail(),
    password: random(),
    phoneNumber: randomPhoneNumber(),
    ...input,
  }).saveX();
}

test("no viewer", async () => {
  const user = await createUser();

  await expectQueryFromRoot(
    getUserRootConfig(user, {
      rootQueryNull: true,
    }),
    ["id", null],
  );
});

test("wrong login credentials", async () => {
  const user = await createUser();

  await expectMutation(
    {
      mutation: "userAuthJWT",
      schema,
      args: {
        emailAddress: user.emailAddress,
        password: random(),
      },
      expectedError: /invalid login credentials/,
    },
    ["token", null],
    ["viewerID", null],
  );
});

test("right credentials", async () => {
  const pw = random();
  const user = await createUser({
    password: pw,
  });

  let st: supertest.SuperTest<supertest.Test>;
  let bearerToken: string | undefined;

  st = await expectMutation(
    {
      // pass a function that takes a server that keeps track of cookies etc
      // and use that for this request
      test: (app: Express) => {
        return supertest.agent(app);
      },
      init: PassportStrategyHandler.testInitJWTFunction({
        secretOrKey: "secret",
        loaderOptions: User.loaderOptions(),
        authOptions: {
          session: false,
        },
      }),
      mutation: "userAuthJWT",
      schema,
      args: {
        emailAddress: user.emailAddress,
        password: pw,
      },
    },
    [
      "token",
      (token) => {
        const decoded = jwt.decode(token);
        expect(decoded).not.toBe(null);
        expect(decoded!["viewerID"]).toBe(user.id);

        bearerToken = token;
      },
    ],
    ["viewerID", encodeGQLID(user)],
  );

  let headers = {};
  if (bearerToken) {
    headers["Authorization"] = `Bearer ${bearerToken}`;
  }
  // send to authed server from above
  // and user is logged in and can make queries!
  await expectQueryFromRoot(
    getUserRootConfig(user, {
      // pass the agent used above to the same server and user is authed!
      test: st,
      // also pass the token as a bearer token for authorization
      headers: headers,
    }),
    ["id", encodeGQLID(user)],
    ["emailAddress", user.emailAddress],
  );

  // same server, no token, user isn't logged in
  await expectQueryFromRoot(
    getUserRootConfig(user, {
      // pass the agent used above to the same server and user is authed!
      test: st,
      rootQueryNull: true,
    }),
    ["id", null],
    ["emailAddress", null],
  );

  // independent server, nothing is saved + no token. user isn't logged in
  await expectQueryFromRoot(
    getUserRootConfig(user, {
      rootQueryNull: true,
    }),
    ["id", null],
  );
});
