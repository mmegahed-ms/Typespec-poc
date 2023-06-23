import {RequestOption} from '@microsoft/kiota-abstractions';

export interface WidgetsItemRequestBuilderPatchRequestConfiguration {
    /** Request headers */
    headers?: Record<string, string[]> | undefined;
    /** Request options */
    options?: RequestOption[] | undefined;
}
