import {WidgetCreate_color} from './widgetCreate_color';
import {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

export interface WidgetCreate extends AdditionalDataHolder, Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    additionalData?: Record<string, unknown>;
    /** The color property */
    color?: WidgetCreate_color | undefined;
    /** The weight property */
    weight?: number | undefined;
}
