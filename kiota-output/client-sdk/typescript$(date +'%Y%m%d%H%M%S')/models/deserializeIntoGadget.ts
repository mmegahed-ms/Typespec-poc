import {Gadget} from './gadget';
import {Gadget_color} from './gadget_color';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoGadget(gadget: Gadget | undefined = {} as Gadget) : Record<string, (node: ParseNode) => void> {
    return {
        "color": n => { gadget.color = n.getEnumValue<Gadget_color>(Gadget_color); },
        "id": n => { gadget.id = n.getStringValue(); },
        "weight": n => { gadget.weight = n.getNumberValue(); },
    }
}
