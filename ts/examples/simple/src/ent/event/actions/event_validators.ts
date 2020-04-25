import { EventBuilder } from "./event_builder";
import { Validator } from "ent/action";

export class EventTimeValidator implements Validator {
  async validate(builder: EventBuilder): Promise<void> {
    const input = builder.getInput();

    const existingEnt = builder.existingEnt;
    const startTime = input.startTime || existingEnt?.startTime;
    const endTime = input.endTime || existingEnt?.endTime;

    if (!startTime) {
      throw new Error("startTime required");
    }

    if (!endTime) {
      // nothing to do here
      return;
    }

    if (startTime.getTime() > endTime.getTime()) {
      throw new Error("start time cannot be after end time");
    }
  }
}

export const SharedValidators = [new EventTimeValidator()];