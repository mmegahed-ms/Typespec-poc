import {Customer} from './customer';
import {Item} from './item';
import {serializeItem} from './serializeItem';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeCustomer(writer: SerializationWriter, customer: Customer | undefined = {} as Customer) : void {
        writer.writeStringValue("address", customer.address);
        writer.writeStringValue("city", customer.city);
        writer.writeStringValue("id", customer.id);
        writer.writeCollectionOfObjectValues<Item>("items", customer.items, serializeItem);
        writer.writeStringValue("name", customer.name);
        writer.writeStringValue("state", customer.state);
        writer.writeStringValue("zip", customer.zip);
        writer.writeAdditionalData(customer.additionalData);
}
