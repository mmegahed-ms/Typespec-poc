import {CustomerCreate} from './customerCreate';
import {Item} from './item';
import {serializeItem} from './serializeItem';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeCustomerCreate(writer: SerializationWriter, customerCreate: CustomerCreate | undefined = {} as CustomerCreate) : void {
        writer.writeStringValue("address", customerCreate.address);
        writer.writeStringValue("city", customerCreate.city);
        writer.writeCollectionOfObjectValues<Item>("items", customerCreate.items, serializeItem);
        writer.writeStringValue("name", customerCreate.name);
        writer.writeStringValue("state", customerCreate.state);
        writer.writeStringValue("zip", customerCreate.zip);
        writer.writeAdditionalData(customerCreate.additionalData);
}
