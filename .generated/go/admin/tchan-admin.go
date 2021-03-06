// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// @generated Code generated by thrift-gen. Do not modify.

// Package admin is generated code used to make or handle TChannel calls using Thrift.
package admin

import (
	"fmt"

	athrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/uber/tchannel-go/thrift"
)

// Interfaces for the service and client for the services defined in the IDL.

// TChanControllerHostAdmin is the interface that defines the server handler and client interface.
type TChanControllerHostAdmin interface {
	ExtentsUnreachable(ctx thrift.Context, request *ExtentsUnreachableRequest) error
}

// TChanInputHostAdmin is the interface that defines the server handler and client interface.
type TChanInputHostAdmin interface {
	DestinationsUpdated(ctx thrift.Context, request *DestinationsUpdatedRequest) error
}

// TChanOutputHostAdmin is the interface that defines the server handler and client interface.
type TChanOutputHostAdmin interface {
	ConsumerGroupsUpdated(ctx thrift.Context, request *ConsumerGroupsUpdatedRequest) error
	UnloadConsumerGroups(ctx thrift.Context, request *UnloadConsumerGroupsRequest) error
}

// Implementation of a client and service handler.

type tchanControllerHostAdminClient struct {
	thriftService string
	client        thrift.TChanClient
}

func NewTChanControllerHostAdminInheritedClient(thriftService string, client thrift.TChanClient) *tchanControllerHostAdminClient {
	return &tchanControllerHostAdminClient{
		thriftService,
		client,
	}
}

// NewTChanControllerHostAdminClient creates a client that can be used to make remote calls.
func NewTChanControllerHostAdminClient(client thrift.TChanClient) TChanControllerHostAdmin {
	return NewTChanControllerHostAdminInheritedClient("ControllerHostAdmin", client)
}

func (c *tchanControllerHostAdminClient) ExtentsUnreachable(ctx thrift.Context, request *ExtentsUnreachableRequest) error {
	var resp ControllerHostAdminExtentsUnreachableResult
	args := ControllerHostAdminExtentsUnreachableArgs{
		Request: request,
	}
	success, err := c.client.Call(ctx, c.thriftService, "extentsUnreachable", &args, &resp)
	if err == nil && !success {
	}

	return err
}

type tchanControllerHostAdminServer struct {
	handler TChanControllerHostAdmin
}

// NewTChanControllerHostAdminServer wraps a handler for TChanControllerHostAdmin so it can be
// registered with a thrift.Server.
func NewTChanControllerHostAdminServer(handler TChanControllerHostAdmin) thrift.TChanServer {
	return &tchanControllerHostAdminServer{
		handler,
	}
}

func (s *tchanControllerHostAdminServer) Service() string {
	return "ControllerHostAdmin"
}

func (s *tchanControllerHostAdminServer) Methods() []string {
	return []string{
		"extentsUnreachable",
	}
}

func (s *tchanControllerHostAdminServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "extentsUnreachable":
		return s.handleExtentsUnreachable(ctx, protocol)

	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanControllerHostAdminServer) handleExtentsUnreachable(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerHostAdminExtentsUnreachableArgs
	var res ControllerHostAdminExtentsUnreachableResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.ExtentsUnreachable(ctx, req.Request)

	if err != nil {
		return false, nil, err
	} else {
	}

	return err == nil, &res, nil
}

type tchanInputHostAdminClient struct {
	thriftService string
	client        thrift.TChanClient
}

func NewTChanInputHostAdminInheritedClient(thriftService string, client thrift.TChanClient) *tchanInputHostAdminClient {
	return &tchanInputHostAdminClient{
		thriftService,
		client,
	}
}

// NewTChanInputHostAdminClient creates a client that can be used to make remote calls.
func NewTChanInputHostAdminClient(client thrift.TChanClient) TChanInputHostAdmin {
	return NewTChanInputHostAdminInheritedClient("InputHostAdmin", client)
}

func (c *tchanInputHostAdminClient) DestinationsUpdated(ctx thrift.Context, request *DestinationsUpdatedRequest) error {
	var resp InputHostAdminDestinationsUpdatedResult
	args := InputHostAdminDestinationsUpdatedArgs{
		Request: request,
	}
	success, err := c.client.Call(ctx, c.thriftService, "destinationsUpdated", &args, &resp)
	if err == nil && !success {
	}

	return err
}

type tchanInputHostAdminServer struct {
	handler TChanInputHostAdmin
}

// NewTChanInputHostAdminServer wraps a handler for TChanInputHostAdmin so it can be
// registered with a thrift.Server.
func NewTChanInputHostAdminServer(handler TChanInputHostAdmin) thrift.TChanServer {
	return &tchanInputHostAdminServer{
		handler,
	}
}

func (s *tchanInputHostAdminServer) Service() string {
	return "InputHostAdmin"
}

func (s *tchanInputHostAdminServer) Methods() []string {
	return []string{
		"destinationsUpdated",
	}
}

func (s *tchanInputHostAdminServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "destinationsUpdated":
		return s.handleDestinationsUpdated(ctx, protocol)

	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanInputHostAdminServer) handleDestinationsUpdated(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req InputHostAdminDestinationsUpdatedArgs
	var res InputHostAdminDestinationsUpdatedResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.DestinationsUpdated(ctx, req.Request)

	if err != nil {
		return false, nil, err
	} else {
	}

	return err == nil, &res, nil
}

type tchanOutputHostAdminClient struct {
	thriftService string
	client        thrift.TChanClient
}

func NewTChanOutputHostAdminInheritedClient(thriftService string, client thrift.TChanClient) *tchanOutputHostAdminClient {
	return &tchanOutputHostAdminClient{
		thriftService,
		client,
	}
}

// NewTChanOutputHostAdminClient creates a client that can be used to make remote calls.
func NewTChanOutputHostAdminClient(client thrift.TChanClient) TChanOutputHostAdmin {
	return NewTChanOutputHostAdminInheritedClient("OutputHostAdmin", client)
}

func (c *tchanOutputHostAdminClient) ConsumerGroupsUpdated(ctx thrift.Context, request *ConsumerGroupsUpdatedRequest) error {
	var resp OutputHostAdminConsumerGroupsUpdatedResult
	args := OutputHostAdminConsumerGroupsUpdatedArgs{
		Request: request,
	}
	success, err := c.client.Call(ctx, c.thriftService, "consumerGroupsUpdated", &args, &resp)
	if err == nil && !success {
	}

	return err
}

func (c *tchanOutputHostAdminClient) UnloadConsumerGroups(ctx thrift.Context, request *UnloadConsumerGroupsRequest) error {
	var resp OutputHostAdminUnloadConsumerGroupsResult
	args := OutputHostAdminUnloadConsumerGroupsArgs{
		Request: request,
	}
	success, err := c.client.Call(ctx, c.thriftService, "unloadConsumerGroups", &args, &resp)
	if err == nil && !success {
	}

	return err
}

type tchanOutputHostAdminServer struct {
	handler TChanOutputHostAdmin
}

// NewTChanOutputHostAdminServer wraps a handler for TChanOutputHostAdmin so it can be
// registered with a thrift.Server.
func NewTChanOutputHostAdminServer(handler TChanOutputHostAdmin) thrift.TChanServer {
	return &tchanOutputHostAdminServer{
		handler,
	}
}

func (s *tchanOutputHostAdminServer) Service() string {
	return "OutputHostAdmin"
}

func (s *tchanOutputHostAdminServer) Methods() []string {
	return []string{
		"consumerGroupsUpdated",
		"unloadConsumerGroups",
	}
}

func (s *tchanOutputHostAdminServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "consumerGroupsUpdated":
		return s.handleConsumerGroupsUpdated(ctx, protocol)
	case "unloadConsumerGroups":
		return s.handleUnloadConsumerGroups(ctx, protocol)

	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanOutputHostAdminServer) handleConsumerGroupsUpdated(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req OutputHostAdminConsumerGroupsUpdatedArgs
	var res OutputHostAdminConsumerGroupsUpdatedResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.ConsumerGroupsUpdated(ctx, req.Request)

	if err != nil {
		return false, nil, err
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanOutputHostAdminServer) handleUnloadConsumerGroups(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req OutputHostAdminUnloadConsumerGroupsArgs
	var res OutputHostAdminUnloadConsumerGroupsResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.UnloadConsumerGroups(ctx, req.Request)

	if err != nil {
		return false, nil, err
	} else {
	}

	return err == nil, &res, nil
}
