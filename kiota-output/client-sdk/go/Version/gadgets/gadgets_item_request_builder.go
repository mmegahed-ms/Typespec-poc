package gadgets

import (
    "context"
    i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9 "apisdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GadgetsItemRequestBuilder builds and executes requests for operations under \gadgets\{id}
type GadgetsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GadgetsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GadgetsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GadgetsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GadgetsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GadgetsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GadgetsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GadgetsItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GadgetsItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analyze the analyze property
func (m *GadgetsItemRequestBuilder) Analyze()(*ItemAnalyzeRequestBuilder) {
    return NewItemAnalyzeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewGadgetsItemRequestBuilderInternal instantiates a new GadgetsItemRequestBuilder and sets the default values.
func NewGadgetsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GadgetsItemRequestBuilder) {
    m := &GadgetsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/gadgets/{id}", pathParameters),
    }
    return m
}
// NewGadgetsItemRequestBuilder instantiates a new GadgetsItemRequestBuilder and sets the default values.
func NewGadgetsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GadgetsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGadgetsItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *GadgetsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GadgetsItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *GadgetsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GadgetsItemRequestBuilderGetRequestConfiguration)(i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Gadgetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.CreateGadgetFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Gadgetable), nil
}
func (m *GadgetsItemRequestBuilder) Patch(ctx context.Context, body i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.GadgetUpdateable, requestConfiguration *GadgetsItemRequestBuilderPatchRequestConfiguration)(i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Gadgetable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.CreateGadgetFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Gadgetable), nil
}
func (m *GadgetsItemRequestBuilder) Put(ctx context.Context, body i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.GadgetCreateOrUpdateable, requestConfiguration *GadgetsItemRequestBuilderPutRequestConfiguration)(i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Gadgetable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.CreateGadgetFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Gadgetable), nil
}
func (m *GadgetsItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GadgetsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
func (m *GadgetsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GadgetsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *GadgetsItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.GadgetUpdateable, requestConfiguration *GadgetsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
func (m *GadgetsItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.GadgetCreateOrUpdateable, requestConfiguration *GadgetsItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
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
