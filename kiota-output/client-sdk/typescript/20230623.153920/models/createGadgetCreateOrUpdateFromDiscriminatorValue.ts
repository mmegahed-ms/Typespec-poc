import {deserializeIntoGadgetCreateOrUpdate} from './deserializeIntoGadgetCreateOrUpdate';
import {GadgetCreateOrUpdate} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createGadgetCreateOrUpdateFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoGadgetCreateOrUpdate;
}
