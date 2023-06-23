import {GadgetUpdate} from './gadgetUpdate';
import {GadgetUpdate_color} from './gadgetUpdate_color';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoGadgetUpdate(gadgetUpdate: GadgetUpdate | undefined = {} as GadgetUpdate) : Record<string, (node: ParseNode) => void> {
    return {
        "color": n => { gadgetUpdate.color = n.getEnumValue<GadgetUpdate_color>(GadgetUpdate_color); },
        "weight": n => { gadgetUpdate.weight = n.getNumberValue(); },
    }
}
