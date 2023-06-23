import {CustomersRequestBuilder} from './customers/customersRequestBuilder';
import {GadgetsRequestBuilder} from './gadgets/gadgetsRequestBuilder';
import {WidgetsRequestBuilder} from './widgets/widgetsRequestBuilder';
import {BaseRequestBuilder, enableBackingStoreForSerializationWriterFactory, ParseNodeFactoryRegistry, registerDefaultDeserializer, registerDefaultSerializer, RequestAdapter, SerializationWriterFactoryRegistry} from '@microsoft/kiota-abstractions';
import {FormParseNodeFactory, FormSerializationWriterFactory} from '@microsoft/kiota-serialization-form';
import {JsonParseNodeFactory, JsonSerializationWriterFactory} from '@microsoft/kiota-serialization-json';
import {TextParseNodeFactory, TextSerializationWriterFactory} from '@microsoft/kiota-serialization-text';

/**
 * The main entry point of the SDK, exposes the configuration and the fluent API.
 */
export class ApiClient extends BaseRequestBuilder {
    /** The customers property */
    public get customers(): CustomersRequestBuilder {
        return new CustomersRequestBuilder(this.pathParameters, this.requestAdapter);
    }
    /** The gadgets property */
    public get gadgets(): GadgetsRequestBuilder {
        return new GadgetsRequestBuilder(this.pathParameters, this.requestAdapter);
    }
    /** The widgets property */
    public get widgets(): WidgetsRequestBuilder {
        return new WidgetsRequestBuilder(this.pathParameters, this.requestAdapter);
    }
    /**
     * Instantiates a new ApiClient and sets the default values.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(requestAdapter: RequestAdapter) {
        super({}, requestAdapter, "{+baseurl}");
        registerDefaultSerializer(JsonSerializationWriterFactory);
        registerDefaultSerializer(TextSerializationWriterFactory);
        registerDefaultSerializer(FormSerializationWriterFactory);
        registerDefaultDeserializer(JsonParseNodeFactory);
        registerDefaultDeserializer(TextParseNodeFactory);
        registerDefaultDeserializer(FormParseNodeFactory);
    };
}
