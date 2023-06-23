import {GadgetCreateOrUpdate} from './gadgetCreateOrUpdate';
import {GadgetCreateOrUpdate_color} from './gadgetCreateOrUpdate_color';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoGadgetCreateOrUpdate(gadgetCreateOrUpdate: GadgetCreateOrUpdate | undefined = {} as GadgetCreateOrUpdate) : Record<string, (node: ParseNode) => void> {
    return {
        "color": n => { gadgetCreateOrUpdate.color = n.getEnumValue<GadgetCreateOrUpdate_color>(GadgetCreateOrUpdate_color); },
        "weight": n => { gadgetCreateOrUpdate.weight = n.getNumberValue(); },
    }
}
