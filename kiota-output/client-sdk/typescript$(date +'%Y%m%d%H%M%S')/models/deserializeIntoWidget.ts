import {Widget} from './widget';
import {Widget_color} from './widget_color';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoWidget(widget: Widget | undefined = {} as Widget) : Record<string, (node: ParseNode) => void> {
    return {
        "color": n => { widget.color = n.getEnumValue<Widget_color>(Widget_color); },
        "id": n => { widget.id = n.getStringValue(); },
        "weight": n => { widget.weight = n.getNumberValue(); },
    }
}
