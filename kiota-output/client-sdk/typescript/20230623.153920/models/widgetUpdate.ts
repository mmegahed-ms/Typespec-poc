import {WidgetUpdate_color} from './widgetUpdate_color';
import {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

export interface WidgetUpdate extends AdditionalDataHolder, Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    additionalData?: Record<string, unknown>;
    /** The color property */
    color?: WidgetUpdate_color | undefined;
    /** The weight property */
    weight?: number | undefined;
}
