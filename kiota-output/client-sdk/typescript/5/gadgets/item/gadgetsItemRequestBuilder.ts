import {createGadgetFromDiscriminatorValue} from '../../models/createGadgetFromDiscriminatorValue';
import {deserializeIntoGadget} from '../../models/deserializeIntoGadget';
import {deserializeIntoGadgetCreateOrUpdate} from '../../models/deserializeIntoGadgetCreateOrUpdate';
import {deserializeIntoGadgetUpdate} from '../../models/deserializeIntoGadgetUpdate';
import {Gadget} from '../../models/gadget';
import {GadgetCreateOrUpdate} from '../../models/gadgetCreateOrUpdate';
import {GadgetUpdate} from '../../models/gadgetUpdate';
import {serializeGadget} from '../../models/serializeGadget';
import {serializeGadgetCreateOrUpdate} from '../../models/serializeGadgetCreateOrUpdate';
import {serializeGadgetUpdate} from '../../models/serializeGadgetUpdate';
import {AnalyzeRequestBuilder} from './analyze/analyzeRequestBuilder';
import {GadgetsItemRequestBuilderDeleteRequestConfiguration} from './gadgetsItemRequestBuilderDeleteRequestConfiguration';
import {GadgetsItemRequestBuilderGetRequestConfiguration} from './gadgetsItemRequestBuilderGetRequestConfiguration';
import {GadgetsItemRequestBuilderPatchRequestConfiguration} from './gadgetsItemRequestBuilderPatchRequestConfiguration';
import {GadgetsItemRequestBuilderPutRequestConfiguration} from './gadgetsItemRequestBuilderPutRequestConfiguration';
import {BaseRequestBuilder, HttpMethod, Parsable, ParsableFactory, RequestAdapter, RequestInformation, RequestOption, ResponseHandler} from '@microsoft/kiota-abstractions';

/**
 * Builds and executes requests for operations under /gadgets/{id}
 */
export class GadgetsItemRequestBuilder extends BaseRequestBuilder {
    /** The analyze property */
    public get analyze(): AnalyzeRequestBuilder {
        return new AnalyzeRequestBuilder(this.pathParameters, this.requestAdapter);
    }
    /**
     * Instantiates a new GadgetsItemRequestBuilder and sets the default values.
     * @param pathParameters The raw url or the Url template parameters for the request.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(pathParameters: Record<string, unknown> | string | undefined, requestAdapter: RequestAdapter) {
        super(pathParameters, requestAdapter, "{+baseurl}/gadgets/{id}");
    };
    /**
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     */
    public delete(requestConfiguration?: GadgetsItemRequestBuilderDeleteRequestConfiguration | undefined, responseHandler?: ResponseHandler | undefined) : Promise<void> {
        const requestInfo = this.toDeleteRequestInformation(
            requestConfiguration
        );
        return this.requestAdapter?.sendNoResponseContentAsync(requestInfo, responseHandler, undefined) ?? Promise.reject(new Error('request adapter is null'));
    };
    /**
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of Gadget
     */
    public get(requestConfiguration?: GadgetsItemRequestBuilderGetRequestConfiguration | undefined, responseHandler?: ResponseHandler | undefined) : Promise<Gadget | undefined> {
        const requestInfo = this.toGetRequestInformation(
            requestConfiguration
        );
        return this.requestAdapter?.sendAsync<Gadget>(requestInfo, createGadgetFromDiscriminatorValue, responseHandler, undefined) ?? Promise.reject(new Error('request adapter is null'));
    };
    /**
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of Gadget
     */
    public patch(body: GadgetUpdate | undefined, requestConfiguration?: GadgetsItemRequestBuilderPatchRequestConfiguration | undefined, responseHandler?: ResponseHandler | undefined) : Promise<Gadget | undefined> {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = this.toPatchRequestInformation(
            body, requestConfiguration
        );
        return this.requestAdapter?.sendAsync<Gadget>(requestInfo, createGadgetFromDiscriminatorValue, responseHandler, undefined) ?? Promise.reject(new Error('request adapter is null'));
    };
    /**
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of Gadget
     */
    public put(body: GadgetCreateOrUpdate | undefined, requestConfiguration?: GadgetsItemRequestBuilderPutRequestConfiguration | undefined, responseHandler?: ResponseHandler | undefined) : Promise<Gadget | undefined> {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = this.toPutRequestInformation(
            body, requestConfiguration
        );
        return this.requestAdapter?.sendAsync<Gadget>(requestInfo, createGadgetFromDiscriminatorValue, responseHandler, undefined) ?? Promise.reject(new Error('request adapter is null'));
    };
    /**
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toDeleteRequestInformation(requestConfiguration?: GadgetsItemRequestBuilderDeleteRequestConfiguration | undefined) : RequestInformation {
        const requestInfo = new RequestInformation();
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.DELETE;
        if (requestConfiguration) {
            requestInfo.addRequestHeaders(requestConfiguration.headers);
            requestInfo.addRequestOptions(requestConfiguration.options);
        }
        return requestInfo;
    };
    /**
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toGetRequestInformation(requestConfiguration?: GadgetsItemRequestBuilderGetRequestConfiguration | undefined) : RequestInformation {
        const requestInfo = new RequestInformation();
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.GET;
        requestInfo.headers["Accept"] = ["application/json"];
        if (requestConfiguration) {
            requestInfo.addRequestHeaders(requestConfiguration.headers);
            requestInfo.addRequestOptions(requestConfiguration.options);
        }
        return requestInfo;
    };
    /**
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toPatchRequestInformation(body: GadgetUpdate | undefined, requestConfiguration?: GadgetsItemRequestBuilderPatchRequestConfiguration | undefined) : RequestInformation {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = new RequestInformation();
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.PATCH;
        requestInfo.headers["Accept"] = ["application/json"];
        if (requestConfiguration) {
            requestInfo.addRequestHeaders(requestConfiguration.headers);
            requestInfo.addRequestOptions(requestConfiguration.options);
        }
        requestInfo.setContentFromParsable(this.requestAdapter, "application/json", body as any, serializeGadgetUpdate);
        return requestInfo;
    };
    /**
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toPutRequestInformation(body: GadgetCreateOrUpdate | undefined, requestConfiguration?: GadgetsItemRequestBuilderPutRequestConfiguration | undefined) : RequestInformation {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = new RequestInformation();
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.PUT;
        requestInfo.headers["Accept"] = ["application/json"];
        if (requestConfiguration) {
            requestInfo.addRequestHeaders(requestConfiguration.headers);
            requestInfo.addRequestOptions(requestConfiguration.options);
        }
        requestInfo.setContentFromParsable(this.requestAdapter, "application/json", body as any, serializeGadgetCreateOrUpdate);
        return requestInfo;
    };
}
