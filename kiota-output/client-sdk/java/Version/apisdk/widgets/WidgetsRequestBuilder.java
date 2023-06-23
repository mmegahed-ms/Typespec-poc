package apisdk.widgets;

import apisdk.models.Widget;
import apisdk.models.WidgetCreate;
import apisdk.widgets.item.WidgetsItemRequestBuilder;
import com.microsoft.kiota.BaseRequestBuilder;
import com.microsoft.kiota.BaseRequestConfiguration;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.RequestAdapter;
import com.microsoft.kiota.RequestInformation;
import com.microsoft.kiota.RequestOption;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParsableFactory;
import java.net.URISyntaxException;
import java.util.Collection;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
/**
 * Builds and executes requests for operations under /widgets
 */
public class WidgetsRequestBuilder extends BaseRequestBuilder {
    /**
     * Gets an item from the ApiSdk.widgets.item collection
     * @param id Unique identifier of the item
     * @return a WidgetsItemRequestBuilder
     */
    @javax.annotation.Nonnull
    public WidgetsItemRequestBuilder byId(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final HashMap<String, Object> urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("id", id);
        return new WidgetsItemRequestBuilder(urlTplParams, requestAdapter);
    }
    /**
     * Instantiates a new WidgetsRequestBuilder and sets the default values.
     * @param pathParameters Path parameters for the request
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    @javax.annotation.Nullable
    public WidgetsRequestBuilder(@javax.annotation.Nonnull final HashMap<String, Object> pathParameters, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/widgets", pathParameters);
    }
    /**
     * Instantiates a new WidgetsRequestBuilder and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    @javax.annotation.Nullable
    public WidgetsRequestBuilder(@javax.annotation.Nonnull final String rawUrl, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/widgets", rawUrl);
    }
    /**
     * @return a CompletableFuture of Widget
     */
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<java.util.List<Widget>> get() {
        try {
            final RequestInformation requestInfo = toGetRequestInformation(null);
            return this.requestAdapter.sendCollectionAsync(requestInfo, Widget::createFromDiscriminatorValue, null);
        } catch (URISyntaxException ex) {
            final java.util.concurrent.CompletableFuture<java.util.List<Widget>> executionException = new java.util.concurrent.CompletableFuture<java.util.List<Widget>>();
            executionException.completeExceptionally(ex);
            return executionException;
        }
    }
    /**
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return a CompletableFuture of Widget
     */
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<java.util.List<Widget>> get(@javax.annotation.Nullable final java.util.function.Consumer<GetRequestConfiguration> requestConfiguration) {
        try {
            final RequestInformation requestInfo = toGetRequestInformation(requestConfiguration);
            return this.requestAdapter.sendCollectionAsync(requestInfo, Widget::createFromDiscriminatorValue, null);
        } catch (URISyntaxException ex) {
            final java.util.concurrent.CompletableFuture<java.util.List<Widget>> executionException = new java.util.concurrent.CompletableFuture<java.util.List<Widget>>();
            executionException.completeExceptionally(ex);
            return executionException;
        }
    }
    /**
     * @param body The request body
     * @return a CompletableFuture of Widget
     */
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<Widget> post(@javax.annotation.Nonnull final WidgetCreate body) {
        try {
            final RequestInformation requestInfo = toPostRequestInformation(body, null);
            return this.requestAdapter.sendAsync(requestInfo, Widget::createFromDiscriminatorValue, null);
        } catch (URISyntaxException ex) {
            final java.util.concurrent.CompletableFuture<Widget> executionException = new java.util.concurrent.CompletableFuture<Widget>();
            executionException.completeExceptionally(ex);
            return executionException;
        }
    }
    /**
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return a CompletableFuture of Widget
     */
    @javax.annotation.Nonnull
    public java.util.concurrent.CompletableFuture<Widget> post(@javax.annotation.Nonnull final WidgetCreate body, @javax.annotation.Nullable final java.util.function.Consumer<PostRequestConfiguration> requestConfiguration) {
        Objects.requireNonNull(body);
        try {
            final RequestInformation requestInfo = toPostRequestInformation(body, requestConfiguration);
            return this.requestAdapter.sendAsync(requestInfo, Widget::createFromDiscriminatorValue, null);
        } catch (URISyntaxException ex) {
            final java.util.concurrent.CompletableFuture<Widget> executionException = new java.util.concurrent.CompletableFuture<Widget>();
            executionException.completeExceptionally(ex);
            return executionException;
        }
    }
    /**
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation toGetRequestInformation() throws URISyntaxException {
        return toGetRequestInformation(null);
    }
    /**
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation toGetRequestInformation(@javax.annotation.Nullable final java.util.function.Consumer<GetRequestConfiguration> requestConfiguration) throws URISyntaxException {
        final RequestInformation requestInfo = new RequestInformation();
        requestInfo.httpMethod = HttpMethod.GET;
        requestInfo.urlTemplate = urlTemplate;
        requestInfo.pathParameters = pathParameters;
        requestInfo.headers.add("Accept", "application/json");
        if (requestConfiguration != null) {
            final GetRequestConfiguration requestConfig = new GetRequestConfiguration();
            requestConfiguration.accept(requestConfig);
            requestInfo.headers.putAll(requestConfig.headers);
            requestInfo.addRequestOptions(requestConfig.options);
        }
        return requestInfo;
    }
    /**
     * @param body The request body
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation toPostRequestInformation(@javax.annotation.Nonnull final WidgetCreate body) throws URISyntaxException {
        return toPostRequestInformation(body, null);
    }
    /**
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation toPostRequestInformation(@javax.annotation.Nonnull final WidgetCreate body, @javax.annotation.Nullable final java.util.function.Consumer<PostRequestConfiguration> requestConfiguration) throws URISyntaxException {
        Objects.requireNonNull(body);
        final RequestInformation requestInfo = new RequestInformation();
        requestInfo.httpMethod = HttpMethod.POST;
        requestInfo.urlTemplate = urlTemplate;
        requestInfo.pathParameters = pathParameters;
        requestInfo.headers.add("Accept", "application/json");
        requestInfo.setContentFromParsable(requestAdapter, "application/json", body);
        if (requestConfiguration != null) {
            final PostRequestConfiguration requestConfig = new PostRequestConfiguration();
            requestConfiguration.accept(requestConfig);
            requestInfo.headers.putAll(requestConfig.headers);
            requestInfo.addRequestOptions(requestConfig.options);
        }
        return requestInfo;
    }
    /**
     * Configuration for the request such as headers, query parameters, and middleware options.
     */
    public class GetRequestConfiguration extends BaseRequestConfiguration {
    }
    /**
     * Configuration for the request such as headers, query parameters, and middleware options.
     */
    public class PostRequestConfiguration extends BaseRequestConfiguration {
    }
}
