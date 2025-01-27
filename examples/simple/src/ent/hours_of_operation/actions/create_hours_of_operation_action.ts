import { AlwaysAllowPrivacyPolicy } from "@lolopinto/ent";
import {
  CreateHoursOfOperationActionBase,
  HoursOfOperationCreateInput,
} from "src/ent/hours_of_operation/actions/generated/create_hours_of_operation_action_base";

export { HoursOfOperationCreateInput };

// we're only writing this once except with --force and packageName provided
export default class CreateHoursOfOperationAction extends CreateHoursOfOperationActionBase {
  getPrivacyPolicy() {
    return AlwaysAllowPrivacyPolicy;
  }
}
