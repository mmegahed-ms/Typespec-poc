import {Item} from './item';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeItem(writer: SerializationWriter, item: Item | undefined = {} as Item) : void {
        writer.writeStringValue("name", item.name);
        writer.writeNumberValue("price", item.price);
        writer.writeNumberValue("quantity", item.quantity);
        writer.writeAdditionalData(item.additionalData);
}
