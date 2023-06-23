import {Widget} from './widget';
import {Widget_color} from './widget_color';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeWidget(writer: SerializationWriter, widget: Widget | undefined = {} as Widget) : void {
        writer.writeEnumValue<Widget_color>("color", widget.color);
        writer.writeStringValue("id", widget.id);
        writer.writeNumberValue("weight", widget.weight);
        writer.writeAdditionalData(widget.additionalData);
}
