import {Item} from './item';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoItem(item: Item | undefined = {} as Item) : Record<string, (node: ParseNode) => void> {
    return {
        "name": n => { item.name = n.getStringValue(); },
        "price": n => { item.price = n.getNumberValue(); },
        "quantity": n => { item.quantity = n.getNumberValue(); },
    }
}
