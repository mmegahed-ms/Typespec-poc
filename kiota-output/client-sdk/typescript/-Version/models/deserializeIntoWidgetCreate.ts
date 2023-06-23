import {WidgetCreate} from './widgetCreate';
import {WidgetCreate_color} from './widgetCreate_color';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoWidgetCreate(widgetCreate: WidgetCreate | undefined = {} as WidgetCreate) : Record<string, (node: ParseNode) => void> {
    return {
        "color": n => { widgetCreate.color = n.getEnumValue<WidgetCreate_color>(WidgetCreate_color); },
        "weight": n => { widgetCreate.weight = n.getNumberValue(); },
    }
}
