import { AlwaysAllowPrivacyPolicy } from "@lolopinto/ent";
import {
  CreateAddressActionBase,
  AddressCreateInput,
} from "src/ent/address/actions/generated/create_address_action_base";

export { AddressCreateInput };

// we're only writing this once except with --force and packageName provided
export default class CreateAddressAction extends CreateAddressActionBase {
  getPrivacyPolicy() {
    return AlwaysAllowPrivacyPolicy;
  }
}
