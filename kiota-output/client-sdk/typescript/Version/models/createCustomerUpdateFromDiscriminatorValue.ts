import {deserializeIntoCustomerUpdate} from './deserializeIntoCustomerUpdate';
import {CustomerUpdate} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createCustomerUpdateFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoCustomerUpdate;
}
