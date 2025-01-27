import {
  EditEventActionBase,
  EventEditInput,
} from "src/ent/event/actions/generated/edit_event_action_base";
import { Validator } from "@lolopinto/ent/action";
import { SharedValidators } from "./event_validators";
import { Event } from "src/ent/";
import {
  AllowIfViewerIsRule,
  AlwaysDenyRule,
  PrivacyPolicy,
} from "@lolopinto/ent";

export { EventEditInput };

// we're only writing this once except with --force and packageName provided
export default class EditEventAction extends EditEventActionBase {
  validators: Validator<Event>[] = [...SharedValidators];

  getPrivacyPolicy(): PrivacyPolicy {
    return {
      rules: [new AllowIfViewerIsRule("creatorID"), AlwaysDenyRule],
    };
  }
}
