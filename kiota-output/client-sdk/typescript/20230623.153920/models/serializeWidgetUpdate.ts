import {WidgetUpdate} from './widgetUpdate';
import {WidgetUpdate_color} from './widgetUpdate_color';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeWidgetUpdate(writer: SerializationWriter, widgetUpdate: WidgetUpdate | undefined = {} as WidgetUpdate) : void {
        writer.writeEnumValue<WidgetUpdate_color>("color", widgetUpdate.color);
        writer.writeNumberValue("weight", widgetUpdate.weight);
        writer.writeAdditionalData(widgetUpdate.additionalData);
}
