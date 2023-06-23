import {Item} from './item';
import {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

export interface CustomerUpdate extends AdditionalDataHolder, Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    additionalData?: Record<string, unknown>;
    /** The address property */
    address?: string | undefined;
    /** The city property */
    city?: string | undefined;
    /** The items property */
    items?: Item[] | undefined;
    /** The name property */
    name?: string | undefined;
    /** The state property */
    state?: string | undefined;
    /** The zip property */
    zip?: string | undefined;
}
