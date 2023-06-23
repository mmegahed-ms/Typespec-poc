import {createItemFromDiscriminatorValue} from './createItemFromDiscriminatorValue';
import {CustomerUpdate} from './customerUpdate';
import {Item} from './item';
import {serializeItem} from './serializeItem';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoCustomerUpdate(customerUpdate: CustomerUpdate | undefined = {} as CustomerUpdate) : Record<string, (node: ParseNode) => void> {
    return {
        "address": n => { customerUpdate.address = n.getStringValue(); },
        "city": n => { customerUpdate.city = n.getStringValue(); },
        "items": n => { customerUpdate.items = n.getCollectionOfObjectValues<Item>(createItemFromDiscriminatorValue); },
        "name": n => { customerUpdate.name = n.getStringValue(); },
        "state": n => { customerUpdate.state = n.getStringValue(); },
        "zip": n => { customerUpdate.zip = n.getStringValue(); },
    }
}
