import {WidgetUpdate} from './widgetUpdate';
import {WidgetUpdate_color} from './widgetUpdate_color';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoWidgetUpdate(widgetUpdate: WidgetUpdate | undefined = {} as WidgetUpdate) : Record<string, (node: ParseNode) => void> {
    return {
        "color": n => { widgetUpdate.color = n.getEnumValue<WidgetUpdate_color>(WidgetUpdate_color); },
        "weight": n => { widgetUpdate.weight = n.getNumberValue(); },
    }
}
