import { fail } from "assert";
import { advanceBy, advanceTo } from "jest-date-mock";
import { IDViewer, LoggedOutViewer } from "../../core/viewer";
import { Data } from "../../core/base";
import { AssocEdge, loadEdgeData } from "../../core/ent";
import { snakeCase } from "snake-case";
import { createRowForTest } from "../write";
import {
  TempDB,
  assoc_edge_config_table,
  assoc_edge_table,
} from "../db/test_db";

import {
  createUser,
  FakeUser,
  UserCreateInput,
  ContactCreateInput,
  FakeContact,
  getContactBuilder,
  EdgeType,
  SymmetricEdges,
  InverseEdges,
} from ".";
import { EventCreateInput, getEventBuilder } from "./fake_event";
import { NodeType } from "./const";

export function getContactInput(
  user: FakeUser,
  input?: Partial<ContactCreateInput>,
): ContactCreateInput {
  return {
    firstName: "Jon",
    lastName: "Snow",
    emailAddress: "foo@bar.com",
    userID: user.id,
    ...input,
  };
}

export function getUserInput(
  input?: Partial<UserCreateInput>,
): UserCreateInput {
  return {
    firstName: "Jon",
    lastName: "Snow",
    emailAddress: "foo@bar.com",
    phoneNumber: "415-212-1212",
    password: "pa$$w0rd",
    ...input,
  };
}

export async function createTestUser(
  input?: Partial<UserCreateInput>,
): Promise<FakeUser> {
  const user = await createUser(new LoggedOutViewer(), {
    firstName: "Jon",
    lastName: "Snow",
    password: "12345678",
    phoneNumber: "4152221231",
    emailAddress: "foo@bar.com",
    ...input,
  });
  if (!user) {
    fail("error creating user");
  }
  return user;
}

export const inputs: Partial<ContactCreateInput>[] = [
  {
    firstName: "Arya",
    lastName: "Stark",
  },
  {
    firstName: "Robb",
    lastName: "Stark",
  },
  {
    firstName: "Sansa",
    lastName: "Stark",
  },
  {
    firstName: "Rickon",
    lastName: "Stark",
  },
  {
    firstName: "Bran",
    lastName: "Stark",
  },
];

export async function createAllContacts(
  input?: Partial<UserCreateInput>,
  slice?: number,
): Promise<[FakeUser, FakeContact[]]> {
  const user = await createTestUser(input);

  let userInputs = inputs.slice(0, slice || inputs.length);
  const contacts = await Promise.all(
    userInputs.map(async (input) => {
      // just to make times deterministic so that tests can consistently work
      advanceBy(100);
      const builder = getContactBuilder(
        user.viewer,
        getContactInput(user, input),
      );
      // add edge from user to contact
      builder.orchestrator.addInboundEdge(
        user.id,
        EdgeType.UserToContacts,
        NodeType.FakeUser,
        {
          time: new Date(), // set time to advanceBy time
        },
      );
      await builder.saveX();
      return await builder.editedEntX();
    }),
  );
  expect(contacts.length).toBe(userInputs.length);
  return [user, contacts];
}

export function verifyUserToContactEdges(
  user: FakeUser,
  edges: AssocEdge[],
  contacts: FakeContact[],
) {
  expect(edges.length).toBe(contacts.length);

  for (let i = 0; i < contacts.length; i++) {
    const edge = edges[i];
    const expectedEdge = {
      id1: user.id,
      id1Type: NodeType.FakeUser,
      id2: contacts[i].id,
      id2Type: NodeType.FakeContact,
      data: null,
      edgeType: EdgeType.UserToContacts,
    };
    expect(edge, `${i}th index`).toMatchObject(expectedEdge);
    expect(edge.getCursor()).not.toBe("");
  }
}

export function verifyUserToContactRawData(
  user: FakeUser,
  edges: Data[],
  contacts: FakeContact[],
) {
  expect(edges.length).toBe(contacts.length);

  for (let i = 0; i < contacts.length; i++) {
    const contact = contacts[i];
    const edge = edges[i];
    const expectedEdge = {
      id: contact.id,
      created_at: contact.createdAt,
      updated_at: contact.updatedAt,
      first_name: contact.firstName,
      last_name: contact.lastName,
      email_address: contact.emailAddress,
      user_id: contact.userID,
    };
    expect(edge, `${i}th index`).toMatchObject(expectedEdge);
  }
}

export function verifyUserToContacts(
  user: FakeUser,
  ents: FakeContact[],
  contacts: FakeContact[],
) {
  expect(ents.length).toBe(contacts.length);
  const expectedContacts = contacts.map((contact) => contact.id);

  expect(ents.map((contact) => contact.id)).toStrictEqual(expectedContacts);
}

export async function createEdges() {
  // create all edges// for now all one-way
  const edgeNames = Object.keys(EdgeType);
  const edges = Object.values(EdgeType);

  for (let i = 0; i < edges.length; i++) {
    const edge = edges[i];
    await createRowForTest({
      tableName: "assoc_edge_config",
      fields: {
        edge_table: snakeCase(`${edge}_table`),
        symmetric_edge: SymmetricEdges.has(edge),
        inverse_edge_type: InverseEdges.get(edge) || null,
        edge_type: edge,
        edge_name: edgeNames[i],
        created_at: new Date(),
        updated_at: new Date(),
      },
    });
    const edgeData = await loadEdgeData(edge);
    expect(edgeData).toBeDefined();
  }
}

export function edgeTableNames() {
  const edges = Object.values(EdgeType);
  return edges.map((edge) => snakeCase(`${edge}_table`));
}

export async function createTestEvent(
  user: FakeUser,
  input?: Partial<EventCreateInput>,
) {
  const vc = new IDViewer(user.id);
  const builder = getEventBuilder(vc, {
    startTime: new Date(),
    location: "fun house",
    description: "fun fun fun",
    title: "fun time",
    userID: user.id,
    ...input,
  });
  builder.orchestrator.addOutboundEdge(user.id, EdgeType.EventToHosts, "User");

  await builder.saveX();
  return await builder.editedEntX();
}

export async function setupTempDB() {
  const tables = [
    FakeUser.getTestTable(),
    FakeContact.getTestTable(),
    assoc_edge_config_table(),
  ];
  edgeTableNames().forEach((tableName) =>
    tables.push(assoc_edge_table(tableName)),
  );

  const tdb = new TempDB(...tables);

  await tdb.beforeAll();

  // create once
  await createEdges();

  return tdb;
}
