import {deserializeIntoCustomerCreate} from './deserializeIntoCustomerCreate';
import {CustomerCreate} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createCustomerCreateFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoCustomerCreate;
}
