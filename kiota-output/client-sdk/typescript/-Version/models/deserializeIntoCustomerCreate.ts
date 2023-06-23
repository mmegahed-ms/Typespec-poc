import {createItemFromDiscriminatorValue} from './createItemFromDiscriminatorValue';
import {CustomerCreate} from './customerCreate';
import {Item} from './item';
import {serializeItem} from './serializeItem';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoCustomerCreate(customerCreate: CustomerCreate | undefined = {} as CustomerCreate) : Record<string, (node: ParseNode) => void> {
    return {
        "address": n => { customerCreate.address = n.getStringValue(); },
        "city": n => { customerCreate.city = n.getStringValue(); },
        "items": n => { customerCreate.items = n.getCollectionOfObjectValues<Item>(createItemFromDiscriminatorValue); },
        "name": n => { customerCreate.name = n.getStringValue(); },
        "state": n => { customerCreate.state = n.getStringValue(); },
        "zip": n => { customerCreate.zip = n.getStringValue(); },
    }
}
