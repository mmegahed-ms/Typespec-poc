import {Gadget} from './gadget';
import {Gadget_color} from './gadget_color';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeGadget(writer: SerializationWriter, gadget: Gadget | undefined = {} as Gadget) : void {
        writer.writeEnumValue<Gadget_color>("color", gadget.color);
        writer.writeStringValue("id", gadget.id);
        writer.writeNumberValue("weight", gadget.weight);
        writer.writeAdditionalData(gadget.additionalData);
}
