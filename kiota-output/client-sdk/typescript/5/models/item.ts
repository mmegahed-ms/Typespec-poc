import {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

export interface Item extends AdditionalDataHolder, Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    additionalData?: Record<string, unknown>;
    /** The name property */
    name?: string | undefined;
    /** The price property */
    price?: number | undefined;
    /** The quantity property */
    quantity?: number | undefined;
}
