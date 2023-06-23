package gadgets

import (
    "context"
    i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9 "apisdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GadgetsRequestBuilder builds and executes requests for operations under \gadgets
type GadgetsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GadgetsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GadgetsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ById gets an item from the ApiSdk.gadgets.item collection
func (m *GadgetsRequestBuilder) ById(id string)(*GadgetsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewGadgetsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewGadgetsRequestBuilderInternal instantiates a new GadgetsRequestBuilder and sets the default values.
func NewGadgetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GadgetsRequestBuilder) {
    m := &GadgetsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/gadgets", pathParameters),
    }
    return m
}
// NewGadgetsRequestBuilder instantiates a new GadgetsRequestBuilder and sets the default values.
func NewGadgetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GadgetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGadgetsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a list of resources
func (m *GadgetsRequestBuilder) Get(ctx context.Context, requestConfiguration *GadgetsRequestBuilderGetRequestConfiguration)([]i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Gadgetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.CreateGadgetFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Gadgetable, len(res))
    for i, v := range res {
        val[i] = v.(i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Gadgetable)
    }
    return val, nil
}
// ToGetRequestInformation get a list of resources
func (m *GadgetsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GadgetsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
