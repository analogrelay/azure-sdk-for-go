// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v7"
	"net/http"
	"net/url"
	"regexp"
)

// MaintenanceConfigurationsServer is a fake server for instances of the armcontainerservice.MaintenanceConfigurationsClient type.
type MaintenanceConfigurationsServer struct {
	// CreateOrUpdate is the fake for method MaintenanceConfigurationsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, resourceName string, configName string, parameters armcontainerservice.MaintenanceConfiguration, options *armcontainerservice.MaintenanceConfigurationsClientCreateOrUpdateOptions) (resp azfake.Responder[armcontainerservice.MaintenanceConfigurationsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method MaintenanceConfigurationsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, resourceName string, configName string, options *armcontainerservice.MaintenanceConfigurationsClientDeleteOptions) (resp azfake.Responder[armcontainerservice.MaintenanceConfigurationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method MaintenanceConfigurationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, configName string, options *armcontainerservice.MaintenanceConfigurationsClientGetOptions) (resp azfake.Responder[armcontainerservice.MaintenanceConfigurationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByManagedClusterPager is the fake for method MaintenanceConfigurationsClient.NewListByManagedClusterPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByManagedClusterPager func(resourceGroupName string, resourceName string, options *armcontainerservice.MaintenanceConfigurationsClientListByManagedClusterOptions) (resp azfake.PagerResponder[armcontainerservice.MaintenanceConfigurationsClientListByManagedClusterResponse])
}

// NewMaintenanceConfigurationsServerTransport creates a new instance of MaintenanceConfigurationsServerTransport with the provided implementation.
// The returned MaintenanceConfigurationsServerTransport instance is connected to an instance of armcontainerservice.MaintenanceConfigurationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMaintenanceConfigurationsServerTransport(srv *MaintenanceConfigurationsServer) *MaintenanceConfigurationsServerTransport {
	return &MaintenanceConfigurationsServerTransport{
		srv:                          srv,
		newListByManagedClusterPager: newTracker[azfake.PagerResponder[armcontainerservice.MaintenanceConfigurationsClientListByManagedClusterResponse]](),
	}
}

// MaintenanceConfigurationsServerTransport connects instances of armcontainerservice.MaintenanceConfigurationsClient to instances of MaintenanceConfigurationsServer.
// Don't use this type directly, use NewMaintenanceConfigurationsServerTransport instead.
type MaintenanceConfigurationsServerTransport struct {
	srv                          *MaintenanceConfigurationsServer
	newListByManagedClusterPager *tracker[azfake.PagerResponder[armcontainerservice.MaintenanceConfigurationsClientListByManagedClusterResponse]]
}

// Do implements the policy.Transporter interface for MaintenanceConfigurationsServerTransport.
func (m *MaintenanceConfigurationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToMethodFake(req, method)
}

func (m *MaintenanceConfigurationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if maintenanceConfigurationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = maintenanceConfigurationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "MaintenanceConfigurationsClient.CreateOrUpdate":
				res.resp, res.err = m.dispatchCreateOrUpdate(req)
			case "MaintenanceConfigurationsClient.Delete":
				res.resp, res.err = m.dispatchDelete(req)
			case "MaintenanceConfigurationsClient.Get":
				res.resp, res.err = m.dispatchGet(req)
			case "MaintenanceConfigurationsClient.NewListByManagedClusterPager":
				res.resp, res.err = m.dispatchNewListByManagedClusterPager(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (m *MaintenanceConfigurationsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/managedClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/maintenanceConfigurations/(?P<configName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcontainerservice.MaintenanceConfiguration](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	configNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, resourceNameParam, configNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).MaintenanceConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MaintenanceConfigurationsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if m.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/managedClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/maintenanceConfigurations/(?P<configName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	configNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Delete(req.Context(), resourceGroupNameParam, resourceNameParam, configNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MaintenanceConfigurationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/managedClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/maintenanceConfigurations/(?P<configName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	configNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, configNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).MaintenanceConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MaintenanceConfigurationsServerTransport) dispatchNewListByManagedClusterPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListByManagedClusterPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByManagedClusterPager not implemented")}
	}
	newListByManagedClusterPager := m.newListByManagedClusterPager.get(req)
	if newListByManagedClusterPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerService/managedClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/maintenanceConfigurations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListByManagedClusterPager(resourceGroupNameParam, resourceNameParam, nil)
		newListByManagedClusterPager = &resp
		m.newListByManagedClusterPager.add(req, newListByManagedClusterPager)
		server.PagerResponderInjectNextLinks(newListByManagedClusterPager, req, func(page *armcontainerservice.MaintenanceConfigurationsClientListByManagedClusterResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByManagedClusterPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListByManagedClusterPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByManagedClusterPager) {
		m.newListByManagedClusterPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to MaintenanceConfigurationsServerTransport
var maintenanceConfigurationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
