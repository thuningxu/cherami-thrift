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

// Package store is generated code used to make or handle TChannel calls using Thrift.
package store

import (
	"fmt"

	athrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/uber/tchannel-go/thrift"

	"github.com/uber/cherami-thrift/.generated/go/cherami"

	"github.com/uber/cherami-thrift/.generated/go/shared"
)

var _ = cherami.GoUnusedProtection__

var _ = shared.GoUnusedProtection__

// Interfaces for the service and client for the services defined in the IDL.

// TChanBStore is the interface that defines the server handler and client interface.
type TChanBStore interface {
	GetAddressFromTimestamp(ctx thrift.Context, getAddressRequest *GetAddressFromTimestampRequest) (*GetAddressFromTimestampResult_, error)
	GetExtentInfo(ctx thrift.Context, extentInfoRequest *GetExtentInfoRequest) (*ExtentInfo, error)
	PurgeMessages(ctx thrift.Context, purgeRequest *PurgeMessagesRequest) (*PurgeMessagesResult_, error)
	ReadMessages(ctx thrift.Context, readMessagesRequest *ReadMessagesRequest) (*ReadMessagesResult_, error)
	RemoteReplicateExtent(ctx thrift.Context, request *RemoteReplicateExtentRequest) error
	ReplicateExtent(ctx thrift.Context, replicateExtentRequest *ReplicateExtentRequest) error
	SealExtent(ctx thrift.Context, sealRequest *SealExtentRequest) error
}

// Implementation of a client and service handler.

type tchanBStoreClient struct {
	thriftService string
	client        thrift.TChanClient
}

func NewTChanBStoreInheritedClient(thriftService string, client thrift.TChanClient) *tchanBStoreClient {
	return &tchanBStoreClient{
		thriftService,
		client,
	}
}

// NewTChanBStoreClient creates a client that can be used to make remote calls.
func NewTChanBStoreClient(client thrift.TChanClient) TChanBStore {
	return NewTChanBStoreInheritedClient("BStore", client)
}

func (c *tchanBStoreClient) GetAddressFromTimestamp(ctx thrift.Context, getAddressRequest *GetAddressFromTimestampRequest) (*GetAddressFromTimestampResult_, error) {
	var resp BStoreGetAddressFromTimestampResult
	args := BStoreGetAddressFromTimestampArgs{
		GetAddressRequest: getAddressRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "getAddressFromTimestamp", &args, &resp)
	if err == nil && !success {
		if e := resp.NotFoundError; e != nil {
			err = e
		}
		if e := resp.RequestError; e != nil {
			err = e
		}
		if e := resp.ServiceError; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanBStoreClient) GetExtentInfo(ctx thrift.Context, extentInfoRequest *GetExtentInfoRequest) (*ExtentInfo, error) {
	var resp BStoreGetExtentInfoResult
	args := BStoreGetExtentInfoArgs{
		ExtentInfoRequest: extentInfoRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "getExtentInfo", &args, &resp)
	if err == nil && !success {
		if e := resp.NotFoundError; e != nil {
			err = e
		}
		if e := resp.RequestError; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanBStoreClient) PurgeMessages(ctx thrift.Context, purgeRequest *PurgeMessagesRequest) (*PurgeMessagesResult_, error) {
	var resp BStorePurgeMessagesResult
	args := BStorePurgeMessagesArgs{
		PurgeRequest: purgeRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "purgeMessages", &args, &resp)
	if err == nil && !success {
		if e := resp.NotFoundError; e != nil {
			err = e
		}
		if e := resp.RequestError; e != nil {
			err = e
		}
		if e := resp.ServiceError; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanBStoreClient) ReadMessages(ctx thrift.Context, readMessagesRequest *ReadMessagesRequest) (*ReadMessagesResult_, error) {
	var resp BStoreReadMessagesResult
	args := BStoreReadMessagesArgs{
		ReadMessagesRequest: readMessagesRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "readMessages", &args, &resp)
	if err == nil && !success {
		if e := resp.ExtentNotFoundError; e != nil {
			err = e
		}
		if e := resp.RequestError; e != nil {
			err = e
		}
		if e := resp.ServiceError; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanBStoreClient) RemoteReplicateExtent(ctx thrift.Context, request *RemoteReplicateExtentRequest) error {
	var resp BStoreRemoteReplicateExtentResult
	args := BStoreRemoteReplicateExtentArgs{
		Request: request,
	}
	success, err := c.client.Call(ctx, c.thriftService, "remoteReplicateExtent", &args, &resp)
	if err == nil && !success {
		if e := resp.ExtentNotFoundError; e != nil {
			err = e
		}
		if e := resp.RequestError; e != nil {
			err = e
		}
		if e := resp.ServiceError; e != nil {
			err = e
		}
	}

	return err
}

func (c *tchanBStoreClient) ReplicateExtent(ctx thrift.Context, replicateExtentRequest *ReplicateExtentRequest) error {
	var resp BStoreReplicateExtentResult
	args := BStoreReplicateExtentArgs{
		ReplicateExtentRequest: replicateExtentRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "replicateExtent", &args, &resp)
	if err == nil && !success {
		if e := resp.ExtentNotFoundError; e != nil {
			err = e
		}
		if e := resp.RequestError; e != nil {
			err = e
		}
		if e := resp.ServiceError; e != nil {
			err = e
		}
	}

	return err
}

func (c *tchanBStoreClient) SealExtent(ctx thrift.Context, sealRequest *SealExtentRequest) error {
	var resp BStoreSealExtentResult
	args := BStoreSealExtentArgs{
		SealRequest: sealRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "sealExtent", &args, &resp)
	if err == nil && !success {
		if e := resp.SealedError; e != nil {
			err = e
		}
		if e := resp.FailedError; e != nil {
			err = e
		}
		if e := resp.RequestError; e != nil {
			err = e
		}
		if e := resp.ServiceError; e != nil {
			err = e
		}
	}

	return err
}

type tchanBStoreServer struct {
	handler TChanBStore
}

// NewTChanBStoreServer wraps a handler for TChanBStore so it can be
// registered with a thrift.Server.
func NewTChanBStoreServer(handler TChanBStore) thrift.TChanServer {
	return &tchanBStoreServer{
		handler,
	}
}

func (s *tchanBStoreServer) Service() string {
	return "BStore"
}

func (s *tchanBStoreServer) Methods() []string {
	return []string{
		"getAddressFromTimestamp",
		"getExtentInfo",
		"purgeMessages",
		"readMessages",
		"remoteReplicateExtent",
		"replicateExtent",
		"sealExtent",
	}
}

func (s *tchanBStoreServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "getAddressFromTimestamp":
		return s.handleGetAddressFromTimestamp(ctx, protocol)
	case "getExtentInfo":
		return s.handleGetExtentInfo(ctx, protocol)
	case "purgeMessages":
		return s.handlePurgeMessages(ctx, protocol)
	case "readMessages":
		return s.handleReadMessages(ctx, protocol)
	case "remoteReplicateExtent":
		return s.handleRemoteReplicateExtent(ctx, protocol)
	case "replicateExtent":
		return s.handleReplicateExtent(ctx, protocol)
	case "sealExtent":
		return s.handleSealExtent(ctx, protocol)

	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanBStoreServer) handleGetAddressFromTimestamp(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req BStoreGetAddressFromTimestampArgs
	var res BStoreGetAddressFromTimestampResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.GetAddressFromTimestamp(ctx, req.GetAddressRequest)

	if err != nil {
		switch v := err.(type) {
		case *ExtentNotFoundError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for notFoundError returned non-nil error type *ExtentNotFoundError but nil value")
			}
			res.NotFoundError = v
		case *BadStoreRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for requestError returned non-nil error type *BadStoreRequestError but nil value")
			}
			res.RequestError = v
		case *StoreServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for serviceError returned non-nil error type *StoreServiceError but nil value")
			}
			res.ServiceError = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanBStoreServer) handleGetExtentInfo(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req BStoreGetExtentInfoArgs
	var res BStoreGetExtentInfoResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.GetExtentInfo(ctx, req.ExtentInfoRequest)

	if err != nil {
		switch v := err.(type) {
		case *ExtentNotFoundError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for notFoundError returned non-nil error type *ExtentNotFoundError but nil value")
			}
			res.NotFoundError = v
		case *BadStoreRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for requestError returned non-nil error type *BadStoreRequestError but nil value")
			}
			res.RequestError = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanBStoreServer) handlePurgeMessages(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req BStorePurgeMessagesArgs
	var res BStorePurgeMessagesResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.PurgeMessages(ctx, req.PurgeRequest)

	if err != nil {
		switch v := err.(type) {
		case *ExtentNotFoundError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for notFoundError returned non-nil error type *ExtentNotFoundError but nil value")
			}
			res.NotFoundError = v
		case *BadStoreRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for requestError returned non-nil error type *BadStoreRequestError but nil value")
			}
			res.RequestError = v
		case *StoreServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for serviceError returned non-nil error type *StoreServiceError but nil value")
			}
			res.ServiceError = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanBStoreServer) handleReadMessages(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req BStoreReadMessagesArgs
	var res BStoreReadMessagesResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.ReadMessages(ctx, req.ReadMessagesRequest)

	if err != nil {
		switch v := err.(type) {
		case *ExtentNotFoundError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for extentNotFoundError returned non-nil error type *ExtentNotFoundError but nil value")
			}
			res.ExtentNotFoundError = v
		case *BadStoreRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for requestError returned non-nil error type *BadStoreRequestError but nil value")
			}
			res.RequestError = v
		case *StoreServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for serviceError returned non-nil error type *StoreServiceError but nil value")
			}
			res.ServiceError = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanBStoreServer) handleRemoteReplicateExtent(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req BStoreRemoteReplicateExtentArgs
	var res BStoreRemoteReplicateExtentResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.RemoteReplicateExtent(ctx, req.Request)

	if err != nil {
		switch v := err.(type) {
		case *ExtentNotFoundError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for extentNotFoundError returned non-nil error type *ExtentNotFoundError but nil value")
			}
			res.ExtentNotFoundError = v
		case *BadStoreRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for requestError returned non-nil error type *BadStoreRequestError but nil value")
			}
			res.RequestError = v
		case *StoreServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for serviceError returned non-nil error type *StoreServiceError but nil value")
			}
			res.ServiceError = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanBStoreServer) handleReplicateExtent(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req BStoreReplicateExtentArgs
	var res BStoreReplicateExtentResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.ReplicateExtent(ctx, req.ReplicateExtentRequest)

	if err != nil {
		switch v := err.(type) {
		case *ExtentNotFoundError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for extentNotFoundError returned non-nil error type *ExtentNotFoundError but nil value")
			}
			res.ExtentNotFoundError = v
		case *BadStoreRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for requestError returned non-nil error type *BadStoreRequestError but nil value")
			}
			res.RequestError = v
		case *StoreServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for serviceError returned non-nil error type *StoreServiceError but nil value")
			}
			res.ServiceError = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanBStoreServer) handleSealExtent(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req BStoreSealExtentArgs
	var res BStoreSealExtentResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.SealExtent(ctx, req.SealRequest)

	if err != nil {
		switch v := err.(type) {
		case *ExtentSealedError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for sealedError returned non-nil error type *ExtentSealedError but nil value")
			}
			res.SealedError = v
		case *ExtentFailedToSealError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for failedError returned non-nil error type *ExtentFailedToSealError but nil value")
			}
			res.FailedError = v
		case *BadStoreRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for requestError returned non-nil error type *BadStoreRequestError but nil value")
			}
			res.RequestError = v
		case *StoreServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for serviceError returned non-nil error type *StoreServiceError but nil value")
			}
			res.ServiceError = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}
