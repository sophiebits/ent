import { Pool } from "pg";
import { QueryRecorder } from "../../testutils/db_mock";
import { ID, Ent, Viewer } from "../base";
import { AssocEdge, DefaultLimit } from "../ent";
import { EdgeQuery } from "./query";
import { EdgeQueryCtr } from "./assoc_query";
import { IDViewer, LoggedOutViewer } from "../viewer";
import { advanceBy } from "jest-date-mock";
import {
  FakeUser,
  UserToContactsQuery,
  FakeContact,
  EdgeType,
  getUserBuilder,
  UserToFriendsQuery,
  FakeEvent,
  UserToEventsAttendingQuery,
  EventToHostsQuery,
  NodeType,
  UserToCustomEdgeQuery,
  CustomEdge,
} from "../../testutils/fake_data/index";
import {
  inputs,
  getUserInput,
  createTestUser,
  createAllContacts,
  verifyUserToContactEdges,
  verifyUserToContacts,
  createEdges,
  createTestEvent,
} from "../../testutils/fake_data/test_helpers";
import { commonTests } from "./shared_test";

jest.mock("pg");
QueryRecorder.mockPool(Pool);

beforeEach(async () => {
  QueryRecorder.clear();
  await createEdges();
  QueryRecorder.clearQueries();
});

commonTests({
  newQuery(viewer: Viewer, user: FakeUser) {
    return UserToContactsQuery.query(viewer, user);
  },
  tableName: "user_to_contacts_table",
  entsLength: 2,
  where: "id1 = $1 AND edge_type = $2",
  sortCol: "time",
});

describe("custom edge", () => {
  let user1, user2: FakeUser;

  beforeEach(async () => {
    user2 = await createTestUser();

    const builder = getUserBuilder(new LoggedOutViewer(), getUserInput());
    builder.orchestrator.addOutboundEdge(
      user2.id,
      EdgeType.UserToCustomEdge,
      NodeType.FakeUser,
    );
    await builder.saveX();
    user1 = await builder.editedEntX();
  });

  test("ids", async () => {
    const ids = await UserToCustomEdgeQuery.query(
      user1.viewer,
      user1,
    ).queryIDs();
    expect(ids.length).toBe(1);
    expect(ids).toEqual([user2.id]);
  });

  test("edges", async () => {
    const edges = await UserToCustomEdgeQuery.query(
      user1.viewer,
      user1,
    ).queryEdges();
    expect(edges.length).toBe(1);
    const edge = edges![0];
    expect(edge).toBeInstanceOf(CustomEdge);
    expect(edge.id1).toBe(user1.id);
    expect(edge.id2).toBe(user2.id);
    expect(edge.edgeType).toBe(EdgeType.UserToCustomEdge);
  });

  test("id2", async () => {
    const edge = await UserToCustomEdgeQuery.query(
      user1.viewer,
      user1,
    ).queryID2(user2.id);
    expect(edge).toBeDefined();
    expect(edge).toBeInstanceOf(CustomEdge);
    expect(edge!.id1).toBe(user1.id);
    expect(edge!.id2).toBe(user2.id);
    expect(edge!.edgeType).toBe(EdgeType.UserToCustomEdge);
  });
});

const N = 2;
function firstNFilter(q: UserToContactsQuery) {
  return q.first(N);
}

function firstNEntsFilter(contacts: FakeContact[]) {
  return contacts.reverse().slice(0, N);
}

function verifyQuery({
  length = 1,
  numQueries = 1,
  limit = DefaultLimit,
  disablePaginationBump = false,
}) {
  const queries = QueryRecorder.getCurrentQueries();
  expect(queries.length).toBe(length);
  for (let i = 0; i < numQueries; i++) {
    const query = queries[i];
    let expLimit = disablePaginationBump ? limit : limit + 1;
    expect(query.qs?.whereClause, `${i}`).toBe(
      // default limit
      `id1 = $1 AND edge_type = $2 ORDER BY time DESC LIMIT ${expLimit}`,
    );
  }
}

function verifyCountQuery({ length = 1, numQueries = 1 }) {
  const queries = QueryRecorder.getCurrentQueries();
  expect(queries.length).toBe(length);
  for (let i = 0; i < numQueries; i++) {
    const query = queries[i];
    expect(query.qs?.whereClause).toBe(`id1 = $1 AND edge_type = $2`);
  }
}

class MultiIDsTestQueryFilter {
  dataz: [FakeUser, FakeContact[]][] = [];
  constructor(
    private filter: (q: UserToContactsQuery) => UserToContactsQuery,
    private ents: (contacts: FakeContact[]) => FakeContact[],
    private limit?: number,
  ) {}

  async beforeEach() {
    let [user1, user2, user3] = await Promise.all([
      createAllContacts({ firstName: "Jon", lastName: "Snow" }),
      createAllContacts({ firstName: "Aegon", lastName: "Targaryen" }),
      createAllContacts({ firstName: "Ned", lastName: "Stark" }),
    ]);
    // modify contacts as needed
    user1[1] = this.ents(user1[1]);
    user2[1] = this.ents(user2[1]);
    user3[1] = this.ents(user3[1]);
    this.dataz = [user1, user2, user3];
    QueryRecorder.clearQueries();
  }

  getQuery(viewer?: Viewer) {
    return this.filter(
      UserToContactsQuery.query(
        viewer || new LoggedOutViewer(),
        this.dataz.map((data) => data[0]),
      ),
    );
  }

  async testIDs() {
    const idsMap = await this.getQuery().queryAllIDs();

    expect(idsMap.size).toBe(this.dataz.length);

    for (let i = 0; i < this.dataz.length; i++) {
      let data = this.dataz[i];

      expect(idsMap.get(data[0].id)).toStrictEqual(
        data[1].map((contact) => contact.id),
      );
    }
    verifyQuery({
      length: this.dataz.length,
      numQueries: this.dataz.length,
      limit: this.limit || DefaultLimit,
    });
  }

  // rawCount isn't affected by filters...
  async testRawCount() {
    const countMap = await this.getQuery().queryAllRawCount();

    expect(countMap.size).toBe(this.dataz.length);

    for (let i = 0; i < this.dataz.length; i++) {
      let data = this.dataz[i];

      expect(countMap.get(data[0].id)).toStrictEqual(inputs.length);
    }
    verifyCountQuery({ numQueries: 3, length: 3 });
  }

  async testCount() {
    const countMap = await this.getQuery().queryAllCount();

    expect(countMap.size).toBe(this.dataz.length);

    for (let i = 0; i < this.dataz.length; i++) {
      let data = this.dataz[i];

      expect(countMap.get(data[0].id)).toStrictEqual(data[1].length);
    }
    verifyQuery({
      length: this.dataz.length,
      numQueries: this.dataz.length,
      limit: this.limit || DefaultLimit,
    });
  }

  async testEdges() {
    const edgesMap = await this.getQuery().queryAllEdges();

    expect(edgesMap.size).toBe(this.dataz.length);

    for (let i = 0; i < this.dataz.length; i++) {
      let data = this.dataz[i];

      verifyUserToContactEdges(
        data[0],
        edgesMap.get(data[0].id) || [],
        data[1],
      );
    }
    verifyQuery({
      length: this.dataz.length,
      numQueries: this.dataz.length,
      limit: this.limit || DefaultLimit,
    });
  }

  async testEnts() {
    // privacy...
    const entsMap = await this.getQuery().queryAllEnts();
    expect(entsMap.size).toBe(this.dataz.length);
    for (let i = 0; i < this.dataz.length; i++) {
      let data = this.dataz[i];
      verifyUserToContacts(data[0], entsMap.get(data[0].id) || [], []);
    }

    // privacy. only data for the first id is visible in this case
    const entsMap2 = await this.getQuery(
      new IDViewer(this.dataz[0][0].id),
    ).queryAllEnts();
    expect(entsMap2.size).toBe(this.dataz.length);
    for (let i = 0; i < this.dataz.length; i++) {
      let data = this.dataz[i];
      verifyUserToContacts(
        data[0],
        entsMap2.get(data[0].id) || [],
        i == 0 ? data[1] : [],
      );
    }
    verifyQuery({
      // extra query for the nodes
      // dataz.length twice to fetch the edge data
      // and then twice to fetch all the nodes for the contacts
      length: this.dataz.length + this.dataz.length + this.dataz.length * 2,
      numQueries: this.dataz.length,
      limit: this.limit || DefaultLimit,
    });
  }
}

describe("multi-ids", () => {
  const filter = new MultiIDsTestQueryFilter(
    (q: UserToContactsQuery) => {
      // no filters
      return q;
    },
    (contacts: FakeContact[]) => {
      // nothing to do here
      // reverse because edges are most recent first
      return contacts.reverse();
    },
  );

  beforeEach(async () => {
    await filter.beforeEach();
  });

  test("ids", async () => {
    await filter.testIDs();
  });

  test("rawCount", async () => {
    await filter.testRawCount();
  });

  test("count", async () => {
    await filter.testCount();
  });

  test("edges", async () => {
    await filter.testEdges();
  });

  test("ents", async () => {
    await filter.testEnts();
  });

  test("id2", async () => {
    //    const users = filter.dataz.map((data) => data[0]);
    for (let i = 0; i < filter.dataz.length; i++) {
      const id1 = filter.dataz[i][0].id;
      // get user that corresponds to position
      const query = filter.getQuery(new IDViewer(id1));

      const id2 = filter.dataz[i][1][0].id;
      //get first contact for user
      const m = await query.queryAllID2(id2);

      for (let j = 0; j < filter.dataz.length; j++) {
        const edge = m.get(filter.dataz[j][0].id);
        if (i == j) {
          expect(edge).toBeDefined();
          expect(edge).toBeInstanceOf(AssocEdge);
          expect(edge!.id1).toBe(id1);
          expect(edge!.id2).toBe(id2);
          expect(edge!.edgeType).toBe(EdgeType.UserToContacts);
        } else {
          expect(edge).toBeUndefined();
        }
      }
    }
  });
});

describe("multi-ids. firstN", () => {
  const filter = new MultiIDsTestQueryFilter(firstNFilter, firstNEntsFilter, 2);

  beforeEach(async () => {
    await filter.beforeEach();
  });

  test("ids", async () => {
    await filter.testIDs();
  });

  test("rawCount", async () => {
    await filter.testRawCount();
  });

  test("count", async () => {
    await filter.testCount();
  });

  test("edges", async () => {
    await filter.testEdges();
  });

  test("ents", async () => {
    await filter.testEnts();
  });
});

class ChainTestQueryFilter {
  user: FakeUser;
  event: FakeEvent;
  event2: FakeEvent;
  friends: FakeUser[];

  constructor(
    private initialQuery: EdgeQueryCtr<Ent, AssocEdge>,
    private subsequentQueries: EdgeQueryCtr<Ent, AssocEdge>[],
    private filter: (q: EdgeQuery<Ent, AssocEdge>) => EdgeQuery<Ent, AssocEdge>,
    private lastHopFilter?: (
      q: EdgeQuery<Ent, AssocEdge>,
    ) => EdgeQuery<Ent, AssocEdge>,
  ) {}

  async beforeEach() {
    this.user = await createTestUser();
    this.event = await createTestEvent(this.user);
    this.event2 = await createTestEvent(this.user, { title: "Red Wedding" });

    this.friends = await Promise.all(
      inputs.map(async (input) => {
        // just to make times deterministic so that tests can consistently work
        advanceBy(100);
        const builder = getUserBuilder(this.user.viewer, getUserInput(input));
        // add edge from user to contact
        builder.orchestrator.addOutboundEdge(
          this.user.id,
          EdgeType.UserToFriends,
          NodeType.FakeUser,
          {
            time: new Date(), // set time to advanceBy time
          },
        );
        // all invited and all attending
        builder.orchestrator.addInboundEdge(
          this.event.id,
          EdgeType.EventToInvited,
          "Event",
        );
        builder.orchestrator.addInboundEdge(
          this.event.id,
          EdgeType.EventToAttendees,
          NodeType.FakeEvent,
        );
        // Robb also attending the red wedding
        if (input.firstName === "Robb") {
          builder.orchestrator.addInboundEdge(
            this.event2.id,
            EdgeType.EventToInvited,
            NodeType.FakeEvent,
          );
          builder.orchestrator.addInboundEdge(
            this.event2.id,
            EdgeType.EventToAttendees,
            NodeType.FakeEvent,
          );
        }
        await builder.saveX();
        return await builder.editedEntX();
      }),
    );

    expect(this.friends.length).toBe(inputs.length);

    const count = await UserToFriendsQuery.query(
      new IDViewer(this.user.id),
      this.user.id,
    ).queryCount();
    expect(count).toStrictEqual(inputs.length);
  }

  getQuery(vc: Viewer) {
    return this.filter(new this.initialQuery(vc, this.user.id));
  }

  private async compare(
    fn: (q: EdgeQuery<Ent, AssocEdge>) => any,
    comparer?: (oneHop: any, allHops: any) => any,
  ) {
    const vc = new IDViewer(this.user.id);
    const oneHopResult = await fn(this.getQuery(vc));

    const queries = [this.initialQuery, ...this.subsequentQueries];
    let last: ID[] = [this.user.id];
    let allHopsResult: any;
    for (let i = 0; i < queries.length; i++) {
      let queryCtr = queries[i];

      let query = new queryCtr(vc, last);
      if (this.lastHopFilter && i + 1 == queries.length - 1) {
        query = this.lastHopFilter(query);
      }
      if (i === queries.length - 1) {
        allHopsResult = await fn(query);
        break;
      }

      let result = await query.queryAllIDs();
      last = [];
      for (const [_, ids] of result) {
        last.push(...ids);
      }
    }
    if (comparer) {
      comparer(oneHopResult, allHopsResult);
    } else {
      expect(oneHopResult).toStrictEqual(allHopsResult);
    }
  }

  async testIDs() {
    await this.compare((q) => q.queryAllIDs());
  }

  async testCount() {
    await this.compare((q) => q.queryAllCount());
  }

  async testRawCount() {
    await this.compare((q) => q.queryAllRawCount());
  }

  async testEdges() {
    await this.compare((q) => q.queryAllEdges());
  }

  async testEnts() {
    function compare(oneHop: Map<ID, Ent[]>, allHops: Map<ID, Ent[]>) {
      expect(allHops.size).toEqual(oneHop.size);
      for (const [key, ents] of allHops) {
        expect(oneHop.has(key)).toEqual(true);

        const oneHopEnts = oneHop.get(key)!;
        for (let i = 0; i < ents.length; i++) {
          expect(oneHopEnts[i].id).toEqual(ents[i].id);
        }
      }
    }
    await this.compare((q) => q.queryAllEnts(), compare);
  }
}

describe("chained queries 2 steps", () => {
  const filter = new ChainTestQueryFilter(
    UserToFriendsQuery,
    [UserToEventsAttendingQuery],
    (q: UserToFriendsQuery) => {
      return q.queryEventsAttending();
    },
  );

  beforeEach(async () => {
    await filter.beforeEach();
  });

  test("ids", async () => {
    await filter.testIDs();
  });

  test("count", async () => {
    await filter.testCount();
  });

  test("rawCount", async () => {
    await filter.testRawCount();
  });

  test("edges", async () => {
    await filter.testEdges();
  });

  test("ents", async () => {
    await filter.testEnts();
  });
});

describe("chained queries 2 steps w/ filter", () => {
  const filter = new ChainTestQueryFilter(
    UserToFriendsQuery,
    [UserToEventsAttendingQuery],
    (q: UserToFriendsQuery) => {
      return q.first(2).queryEventsAttending();
    },
    (q: UserToFriendsQuery) => {
      return q.first(2);
    },
  );

  beforeEach(async () => {
    await filter.beforeEach();
  });

  test("ids", async () => {
    await filter.testIDs();
  });

  test("count", async () => {
    await filter.testCount();
  });

  test("rawCount", async () => {
    await filter.testRawCount();
  });

  test("edges", async () => {
    await filter.testEdges();
  });

  test("ents", async () => {
    await filter.testEnts();
  });
});

describe("chained queries 3 steps", () => {
  const filter = new ChainTestQueryFilter(
    UserToFriendsQuery,
    [UserToEventsAttendingQuery, EventToHostsQuery],
    (q: UserToFriendsQuery) => {
      return q.queryEventsAttending().queryHosts();
    },
  );

  beforeEach(async () => {
    await filter.beforeEach();
  });

  test("ids", async () => {
    await filter.testIDs();
  });

  test("count", async () => {
    await filter.testCount();
  });

  test("rawCount", async () => {
    await filter.testRawCount();
  });

  test("edges", async () => {
    await filter.testEdges();
  });

  test("ents", async () => {
    await filter.testEnts();
  });
});

// TODO need to figure out a better way to test time. we had ms here
// for times but we needed Date object comparions
// tests work for both but production only works with Date comparisons
// flaw with nosql parse_sql implementation
