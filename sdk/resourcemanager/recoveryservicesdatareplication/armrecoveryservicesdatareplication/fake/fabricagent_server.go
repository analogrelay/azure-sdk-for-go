// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
	"net/http"
	"net/url"
	"regexp"
)

// FabricAgentServer is a fake server for instances of the armrecoveryservicesdatareplication.FabricAgentClient type.
type FabricAgentServer struct {
	// BeginCreate is the fake for method FabricAgentClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, fabricName string, fabricAgentName string, resource armrecoveryservicesdatareplication.FabricAgentModel, options *armrecoveryservicesdatareplication.FabricAgentClientBeginCreateOptions) (resp azfake.PollerResponder[armrecoveryservicesdatareplication.FabricAgentClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method FabricAgentClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, fabricName string, fabricAgentName string, options *armrecoveryservicesdatareplication.FabricAgentClientBeginDeleteOptions) (resp azfake.PollerResponder[armrecoveryservicesdatareplication.FabricAgentClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method FabricAgentClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, fabricName string, fabricAgentName string, options *armrecoveryservicesdatareplication.FabricAgentClientGetOptions) (resp azfake.Responder[armrecoveryservicesdatareplication.FabricAgentClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method FabricAgentClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, fabricName string, options *armrecoveryservicesdatareplication.FabricAgentClientListOptions) (resp azfake.PagerResponder[armrecoveryservicesdatareplication.FabricAgentClientListResponse])
}

// NewFabricAgentServerTransport creates a new instance of FabricAgentServerTransport with the provided implementation.
// The returned FabricAgentServerTransport instance is connected to an instance of armrecoveryservicesdatareplication.FabricAgentClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFabricAgentServerTransport(srv *FabricAgentServer) *FabricAgentServerTransport {
	return &FabricAgentServerTransport{
		srv:          srv,
		beginCreate:  newTracker[azfake.PollerResponder[armrecoveryservicesdatareplication.FabricAgentClientCreateResponse]](),
		beginDelete:  newTracker[azfake.PollerResponder[armrecoveryservicesdatareplication.FabricAgentClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armrecoveryservicesdatareplication.FabricAgentClientListResponse]](),
	}
}

// FabricAgentServerTransport connects instances of armrecoveryservicesdatareplication.FabricAgentClient to instances of FabricAgentServer.
// Don't use this type directly, use NewFabricAgentServerTransport instead.
type FabricAgentServerTransport struct {
	srv          *FabricAgentServer
	beginCreate  *tracker[azfake.PollerResponder[armrecoveryservicesdatareplication.FabricAgentClientCreateResponse]]
	beginDelete  *tracker[azfake.PollerResponder[armrecoveryservicesdatareplication.FabricAgentClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armrecoveryservicesdatareplication.FabricAgentClientListResponse]]
}

// Do implements the policy.Transporter interface for FabricAgentServerTransport.
func (f *FabricAgentServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return f.dispatchToMethodFake(req, method)
}

func (f *FabricAgentServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if fabricAgentServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = fabricAgentServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "FabricAgentClient.BeginCreate":
				res.resp, res.err = f.dispatchBeginCreate(req)
			case "FabricAgentClient.BeginDelete":
				res.resp, res.err = f.dispatchBeginDelete(req)
			case "FabricAgentClient.Get":
				res.resp, res.err = f.dispatchGet(req)
			case "FabricAgentClient.NewListPager":
				res.resp, res.err = f.dispatchNewListPager(req)
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

func (f *FabricAgentServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if f.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := f.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fabricAgents/(?P<fabricAgentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armrecoveryservicesdatareplication.FabricAgentModel](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
		if err != nil {
			return nil, err
		}
		fabricAgentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricAgentName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginCreate(req.Context(), resourceGroupNameParam, fabricNameParam, fabricAgentNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		f.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		f.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		f.beginCreate.remove(req)
	}

	return resp, nil
}

func (f *FabricAgentServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if f.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := f.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fabricAgents/(?P<fabricAgentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
		if err != nil {
			return nil, err
		}
		fabricAgentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricAgentName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginDelete(req.Context(), resourceGroupNameParam, fabricNameParam, fabricAgentNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		f.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		f.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		f.beginDelete.remove(req)
	}

	return resp, nil
}

func (f *FabricAgentServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fabricAgents/(?P<fabricAgentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
	if err != nil {
		return nil, err
	}
	fabricAgentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricAgentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Get(req.Context(), resourceGroupNameParam, fabricNameParam, fabricAgentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FabricAgentModel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FabricAgentServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := f.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataReplication/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fabricAgents`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
		if err != nil {
			return nil, err
		}
		resp := f.srv.NewListPager(resourceGroupNameParam, fabricNameParam, nil)
		newListPager = &resp
		f.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armrecoveryservicesdatareplication.FabricAgentClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		f.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to FabricAgentServerTransport
var fabricAgentServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
