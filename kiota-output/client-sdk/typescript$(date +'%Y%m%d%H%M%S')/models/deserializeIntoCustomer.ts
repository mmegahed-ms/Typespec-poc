import {createItemFromDiscriminatorValue} from './createItemFromDiscriminatorValue';
import {Customer} from './customer';
import {Item} from './item';
import {serializeItem} from './serializeItem';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoCustomer(customer: Customer | undefined = {} as Customer) : Record<string, (node: ParseNode) => void> {
    return {
        "address": n => { customer.address = n.getStringValue(); },
        "city": n => { customer.city = n.getStringValue(); },
        "id": n => { customer.id = n.getStringValue(); },
        "items": n => { customer.items = n.getCollectionOfObjectValues<Item>(createItemFromDiscriminatorValue); },
        "name": n => { customer.name = n.getStringValue(); },
        "state": n => { customer.state = n.getStringValue(); },
        "zip": n => { customer.zip = n.getStringValue(); },
    }
}
