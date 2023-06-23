import {GadgetUpdate} from './gadgetUpdate';
import {GadgetUpdate_color} from './gadgetUpdate_color';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeGadgetUpdate(writer: SerializationWriter, gadgetUpdate: GadgetUpdate | undefined = {} as GadgetUpdate) : void {
        writer.writeEnumValue<GadgetUpdate_color>("color", gadgetUpdate.color);
        writer.writeNumberValue("weight", gadgetUpdate.weight);
        writer.writeAdditionalData(gadgetUpdate.additionalData);
}
