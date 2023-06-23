import {Gadget_color} from './gadget_color';
import {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

export interface Gadget extends AdditionalDataHolder, Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    additionalData?: Record<string, unknown>;
    /** The color property */
    color?: Gadget_color | undefined;
    /** The id property */
    id?: string | undefined;
    /** The weight property */
    weight?: number | undefined;
}
