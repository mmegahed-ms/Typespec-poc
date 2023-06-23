package widgets

import (
    "context"
    i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9 "apisdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WidgetsRequestBuilder builds and executes requests for operations under \widgets
type WidgetsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WidgetsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WidgetsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WidgetsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WidgetsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ById gets an item from the ApiSdk.widgets.item collection
func (m *WidgetsRequestBuilder) ById(id string)(*WidgetsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewWidgetsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWidgetsRequestBuilderInternal instantiates a new WidgetsRequestBuilder and sets the default values.
func NewWidgetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WidgetsRequestBuilder) {
    m := &WidgetsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/widgets", pathParameters),
    }
    return m
}
// NewWidgetsRequestBuilder instantiates a new WidgetsRequestBuilder and sets the default values.
func NewWidgetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WidgetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWidgetsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WidgetsRequestBuilder) Get(ctx context.Context, requestConfiguration *WidgetsRequestBuilderGetRequestConfiguration)([]i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Widgetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.CreateWidgetFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Widgetable, len(res))
    for i, v := range res {
        val[i] = v.(i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Widgetable)
    }
    return val, nil
}
func (m *WidgetsRequestBuilder) Post(ctx context.Context, body i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.WidgetCreateable, requestConfiguration *WidgetsRequestBuilderPostRequestConfiguration)(i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Widgetable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.CreateWidgetFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Widgetable), nil
}
func (m *WidgetsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WidgetsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *WidgetsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.WidgetCreateable, requestConfiguration *WidgetsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
