import {deserializeIntoGadget} from './deserializeIntoGadget';
import {Gadget} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createGadgetFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoGadget;
}
