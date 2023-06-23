import {CustomerUpdate} from './customerUpdate';
import {Item} from './item';
import {serializeItem} from './serializeItem';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeCustomerUpdate(writer: SerializationWriter, customerUpdate: CustomerUpdate | undefined = {} as CustomerUpdate) : void {
        writer.writeStringValue("address", customerUpdate.address);
        writer.writeStringValue("city", customerUpdate.city);
        writer.writeCollectionOfObjectValues<Item>("items", customerUpdate.items, serializeItem);
        writer.writeStringValue("name", customerUpdate.name);
        writer.writeStringValue("state", customerUpdate.state);
        writer.writeStringValue("zip", customerUpdate.zip);
        writer.writeAdditionalData(customerUpdate.additionalData);
}
