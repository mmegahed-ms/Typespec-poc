import {GadgetCreateOrUpdate} from './gadgetCreateOrUpdate';
import {GadgetCreateOrUpdate_color} from './gadgetCreateOrUpdate_color';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeGadgetCreateOrUpdate(writer: SerializationWriter, gadgetCreateOrUpdate: GadgetCreateOrUpdate | undefined = {} as GadgetCreateOrUpdate) : void {
        writer.writeEnumValue<GadgetCreateOrUpdate_color>("color", gadgetCreateOrUpdate.color);
        writer.writeNumberValue("weight", gadgetCreateOrUpdate.weight);
        writer.writeAdditionalData(gadgetCreateOrUpdate.additionalData);
}
