import {deserializeIntoCustomer} from './deserializeIntoCustomer';
import {Customer} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createCustomerFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoCustomer;
}
