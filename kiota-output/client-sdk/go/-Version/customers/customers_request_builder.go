package customers

import (
    "context"
    i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9 "apisdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CustomersRequestBuilder builds and executes requests for operations under \customers
type CustomersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CustomersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CustomersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ById gets an item from the ApiSdk.customers.item collection
func (m *CustomersRequestBuilder) ById(id string)(*CustomersItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewCustomersItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCustomersRequestBuilderInternal instantiates a new CustomersRequestBuilder and sets the default values.
func NewCustomersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomersRequestBuilder) {
    m := &CustomersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/customers", pathParameters),
    }
    return m
}
// NewCustomersRequestBuilder instantiates a new CustomersRequestBuilder and sets the default values.
func NewCustomersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustomersRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *CustomersRequestBuilder) Get(ctx context.Context, requestConfiguration *CustomersRequestBuilderGetRequestConfiguration)([]i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Customerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.CreateCustomerFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Customerable, len(res))
    for i, v := range res {
        val[i] = v.(i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Customerable)
    }
    return val, nil
}
func (m *CustomersRequestBuilder) Post(ctx context.Context, body i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.CustomerCreateable, requestConfiguration *CustomersRequestBuilderPostRequestConfiguration)(i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Customerable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.CreateCustomerFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Customerable), nil
}
func (m *CustomersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CustomersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
func (m *CustomersRequestBuilder) ToPostRequestInformation(ctx context.Context, body i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.CustomerCreateable, requestConfiguration *CustomersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
