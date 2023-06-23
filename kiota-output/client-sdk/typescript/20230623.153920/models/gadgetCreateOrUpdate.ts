import {GadgetCreateOrUpdate_color} from './gadgetCreateOrUpdate_color';
import {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

export interface GadgetCreateOrUpdate extends AdditionalDataHolder, Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    additionalData?: Record<string, unknown>;
    /** The color property */
    color?: GadgetCreateOrUpdate_color | undefined;
    /** The weight property */
    weight?: number | undefined;
}
