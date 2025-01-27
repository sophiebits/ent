// Generated by github.com/lolopinto/ent/ent, DO NOT EDIT.

import {
  CustomEdgeQueryBase,
  ID,
  IndexLoaderFactory,
  RawCountLoaderFactory,
  Viewer,
} from "@lolopinto/ent";
import { Event, User, eventLoader } from "src/ent/internal";

export const userToEventsCountLoaderFactory = new RawCountLoaderFactory(
  Event.loaderOptions(),
  "creator_id",
);
export const userToEventsDataLoaderFactory = new IndexLoaderFactory(
  Event.loaderOptions(),
  "creator_id",
  {
    toPrime: [eventLoader],
  },
);

export class UserToEventsQueryBase extends CustomEdgeQueryBase<Event> {
  constructor(viewer: Viewer, src: User | ID) {
    super(viewer, {
      src: src,
      countLoaderFactory: userToEventsCountLoaderFactory,
      dataLoaderFactory: userToEventsDataLoaderFactory,
      options: Event.loaderOptions(),
    });
  }

  static query<T extends UserToEventsQueryBase>(
    this: new (viewer: Viewer, src: User | ID) => T,
    viewer: Viewer,
    src: User | ID,
  ): T {
    return new this(viewer, src);
  }
}
