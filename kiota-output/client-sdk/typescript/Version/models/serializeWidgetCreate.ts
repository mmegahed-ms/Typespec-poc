import {WidgetCreate} from './widgetCreate';
import {WidgetCreate_color} from './widgetCreate_color';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeWidgetCreate(writer: SerializationWriter, widgetCreate: WidgetCreate | undefined = {} as WidgetCreate) : void {
        writer.writeEnumValue<WidgetCreate_color>("color", widgetCreate.color);
        writer.writeNumberValue("weight", widgetCreate.weight);
        writer.writeAdditionalData(widgetCreate.additionalData);
}
