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

// Package controller is generated code used to make or handle TChannel calls using Thrift.
package controller

import (
	"fmt"

	athrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/uber/tchannel-go/thrift"

	"github.com/uber/cherami-thrift/.generated/go/shared"
)

var _ = shared.GoUnusedProtection__

// Interfaces for the service and client for the services defined in the IDL.

// TChanController is the interface that defines the server handler and client interface.
type TChanController interface {
	CreateConsumerGroup(ctx thrift.Context, createRequest *shared.CreateConsumerGroupRequest) (*shared.ConsumerGroupDescription, error)
	CreateDestination(ctx thrift.Context, createRequest *shared.CreateDestinationRequest) (*shared.DestinationDescription, error)
	CreateRemoteZoneExtent(ctx thrift.Context, createRequest *shared.CreateExtentRequest) (*shared.CreateExtentResult_, error)
	DeleteConsumerGroup(ctx thrift.Context, deleteRequest *shared.DeleteConsumerGroupRequest) error
	DeleteDestination(ctx thrift.Context, deleteRequest *shared.DeleteDestinationRequest) error
	GetCapacities(ctx thrift.Context, getCapacitiesRequest *GetCapacitiesRequest) (*GetCapacitiesResult_, error)
	GetInputHosts(ctx thrift.Context, getHostsRequest *GetInputHostsRequest) (*GetInputHostsResult_, error)
	GetOutputHosts(ctx thrift.Context, getHostsRequest *GetOutputHostsRequest) (*GetOutputHostsResult_, error)
	GetQueueDepthInfo(ctx thrift.Context, getQueueDepthInfoRequest *GetQueueDepthInfoRequest) (*GetQueueDepthInfoResult_, error)
	RemoveCapacities(ctx thrift.Context, removeCapacitiesRequest *RemoveCapacitiesRequest) error
	ReportConsumerGroupExtentMetric(ctx thrift.Context, reportMetricRequest *ReportConsumerGroupExtentMetricRequest) error
	ReportConsumerGroupMetric(ctx thrift.Context, reportMetricRequest *ReportConsumerGroupMetricRequest) error
	ReportDestinationExtentMetric(ctx thrift.Context, reportMetricRequest *ReportDestinationExtentMetricRequest) error
	ReportDestinationMetric(ctx thrift.Context, reportMetricRequest *ReportDestinationMetricRequest) error
	ReportNodeMetric(ctx thrift.Context, reportMetricRequest *ReportNodeMetricRequest) error
	ReportStoreExtentMetric(ctx thrift.Context, reportMetricRequest *ReportStoreExtentMetricRequest) error
	UpdateConsumerGroup(ctx thrift.Context, updateRequest *shared.UpdateConsumerGroupRequest) (*shared.ConsumerGroupDescription, error)
	UpdateDestination(ctx thrift.Context, updateRequest *shared.UpdateDestinationRequest) (*shared.DestinationDescription, error)
	UpsertInputHostCapacities(ctx thrift.Context, upsertCapacitiesRequest *UpsertInputHostCapacitiesRequest) error
	UpsertOutputHostCapacities(ctx thrift.Context, upsertCapacitiesRequest *UpsertOutputHostCapacitiesRequest) error
	UpsertStoreCapacities(ctx thrift.Context, upsertCapacitiesRequest *UpsertStoreCapacitiesRequest) error
}

// Implementation of a client and service handler.

type tchanControllerClient struct {
	thriftService string
	client        thrift.TChanClient
}

func NewTChanControllerInheritedClient(thriftService string, client thrift.TChanClient) *tchanControllerClient {
	return &tchanControllerClient{
		thriftService,
		client,
	}
}

// NewTChanControllerClient creates a client that can be used to make remote calls.
func NewTChanControllerClient(client thrift.TChanClient) TChanController {
	return NewTChanControllerInheritedClient("Controller", client)
}

func (c *tchanControllerClient) CreateConsumerGroup(ctx thrift.Context, createRequest *shared.CreateConsumerGroupRequest) (*shared.ConsumerGroupDescription, error) {
	var resp ControllerCreateConsumerGroupResult
	args := ControllerCreateConsumerGroupArgs{
		CreateRequest: createRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "createConsumerGroup", &args, &resp)
	if err == nil && !success {
		if e := resp.EntityExistsError; e != nil {
			err = e
		}
		if e := resp.RequestError; e != nil {
			err = e
		}
		if e := resp.InternalError; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanControllerClient) CreateDestination(ctx thrift.Context, createRequest *shared.CreateDestinationRequest) (*shared.DestinationDescription, error) {
	var resp ControllerCreateDestinationResult
	args := ControllerCreateDestinationArgs{
		CreateRequest: createRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "createDestination", &args, &resp)
	if err == nil && !success {
		if e := resp.EntityExistsError; e != nil {
			err = e
		}
		if e := resp.RequestError; e != nil {
			err = e
		}
		if e := resp.InternalError; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanControllerClient) CreateRemoteZoneExtent(ctx thrift.Context, createRequest *shared.CreateExtentRequest) (*shared.CreateExtentResult_, error) {
	var resp ControllerCreateRemoteZoneExtentResult
	args := ControllerCreateRemoteZoneExtentArgs{
		CreateRequest: createRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "createRemoteZoneExtent", &args, &resp)
	if err == nil && !success {
		if e := resp.EntityExistsError; e != nil {
			err = e
		}
		if e := resp.RequestError; e != nil {
			err = e
		}
		if e := resp.InternalError; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanControllerClient) DeleteConsumerGroup(ctx thrift.Context, deleteRequest *shared.DeleteConsumerGroupRequest) error {
	var resp ControllerDeleteConsumerGroupResult
	args := ControllerDeleteConsumerGroupArgs{
		DeleteRequest: deleteRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "deleteConsumerGroup", &args, &resp)
	if err == nil && !success {
		if e := resp.EntityError; e != nil {
			err = e
		}
		if e := resp.RequestError; e != nil {
			err = e
		}
		if e := resp.InternalError; e != nil {
			err = e
		}
	}

	return err
}

func (c *tchanControllerClient) DeleteDestination(ctx thrift.Context, deleteRequest *shared.DeleteDestinationRequest) error {
	var resp ControllerDeleteDestinationResult
	args := ControllerDeleteDestinationArgs{
		DeleteRequest: deleteRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "deleteDestination", &args, &resp)
	if err == nil && !success {
		if e := resp.EntityError; e != nil {
			err = e
		}
		if e := resp.RequestError; e != nil {
			err = e
		}
		if e := resp.InternalError; e != nil {
			err = e
		}
	}

	return err
}

func (c *tchanControllerClient) GetCapacities(ctx thrift.Context, getCapacitiesRequest *GetCapacitiesRequest) (*GetCapacitiesResult_, error) {
	var resp ControllerGetCapacitiesResult
	args := ControllerGetCapacitiesArgs{
		GetCapacitiesRequest: getCapacitiesRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "getCapacities", &args, &resp)
	if err == nil && !success {
	}

	return resp.GetSuccess(), err
}

func (c *tchanControllerClient) GetInputHosts(ctx thrift.Context, getHostsRequest *GetInputHostsRequest) (*GetInputHostsResult_, error) {
	var resp ControllerGetInputHostsResult
	args := ControllerGetInputHostsArgs{
		GetHostsRequest: getHostsRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "getInputHosts", &args, &resp)
	if err == nil && !success {
		if e := resp.RequestError; e != nil {
			err = e
		}
		if e := resp.InternalError; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanControllerClient) GetOutputHosts(ctx thrift.Context, getHostsRequest *GetOutputHostsRequest) (*GetOutputHostsResult_, error) {
	var resp ControllerGetOutputHostsResult
	args := ControllerGetOutputHostsArgs{
		GetHostsRequest: getHostsRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "getOutputHosts", &args, &resp)
	if err == nil && !success {
		if e := resp.RequestError; e != nil {
			err = e
		}
		if e := resp.InternalError; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanControllerClient) GetQueueDepthInfo(ctx thrift.Context, getQueueDepthInfoRequest *GetQueueDepthInfoRequest) (*GetQueueDepthInfoResult_, error) {
	var resp ControllerGetQueueDepthInfoResult
	args := ControllerGetQueueDepthInfoArgs{
		GetQueueDepthInfoRequest: getQueueDepthInfoRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "getQueueDepthInfo", &args, &resp)
	if err == nil && !success {
		if e := resp.CacheMissError; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanControllerClient) RemoveCapacities(ctx thrift.Context, removeCapacitiesRequest *RemoveCapacitiesRequest) error {
	var resp ControllerRemoveCapacitiesResult
	args := ControllerRemoveCapacitiesArgs{
		RemoveCapacitiesRequest: removeCapacitiesRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "removeCapacities", &args, &resp)
	if err == nil && !success {
	}

	return err
}

func (c *tchanControllerClient) ReportConsumerGroupExtentMetric(ctx thrift.Context, reportMetricRequest *ReportConsumerGroupExtentMetricRequest) error {
	var resp ControllerReportConsumerGroupExtentMetricResult
	args := ControllerReportConsumerGroupExtentMetricArgs{
		ReportMetricRequest: reportMetricRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "reportConsumerGroupExtentMetric", &args, &resp)
	if err == nil && !success {
	}

	return err
}

func (c *tchanControllerClient) ReportConsumerGroupMetric(ctx thrift.Context, reportMetricRequest *ReportConsumerGroupMetricRequest) error {
	var resp ControllerReportConsumerGroupMetricResult
	args := ControllerReportConsumerGroupMetricArgs{
		ReportMetricRequest: reportMetricRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "reportConsumerGroupMetric", &args, &resp)
	if err == nil && !success {
	}

	return err
}

func (c *tchanControllerClient) ReportDestinationExtentMetric(ctx thrift.Context, reportMetricRequest *ReportDestinationExtentMetricRequest) error {
	var resp ControllerReportDestinationExtentMetricResult
	args := ControllerReportDestinationExtentMetricArgs{
		ReportMetricRequest: reportMetricRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "reportDestinationExtentMetric", &args, &resp)
	if err == nil && !success {
	}

	return err
}

func (c *tchanControllerClient) ReportDestinationMetric(ctx thrift.Context, reportMetricRequest *ReportDestinationMetricRequest) error {
	var resp ControllerReportDestinationMetricResult
	args := ControllerReportDestinationMetricArgs{
		ReportMetricRequest: reportMetricRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "reportDestinationMetric", &args, &resp)
	if err == nil && !success {
	}

	return err
}

func (c *tchanControllerClient) ReportNodeMetric(ctx thrift.Context, reportMetricRequest *ReportNodeMetricRequest) error {
	var resp ControllerReportNodeMetricResult
	args := ControllerReportNodeMetricArgs{
		ReportMetricRequest: reportMetricRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "reportNodeMetric", &args, &resp)
	if err == nil && !success {
	}

	return err
}

func (c *tchanControllerClient) ReportStoreExtentMetric(ctx thrift.Context, reportMetricRequest *ReportStoreExtentMetricRequest) error {
	var resp ControllerReportStoreExtentMetricResult
	args := ControllerReportStoreExtentMetricArgs{
		ReportMetricRequest: reportMetricRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "reportStoreExtentMetric", &args, &resp)
	if err == nil && !success {
	}

	return err
}

func (c *tchanControllerClient) UpdateConsumerGroup(ctx thrift.Context, updateRequest *shared.UpdateConsumerGroupRequest) (*shared.ConsumerGroupDescription, error) {
	var resp ControllerUpdateConsumerGroupResult
	args := ControllerUpdateConsumerGroupArgs{
		UpdateRequest: updateRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "updateConsumerGroup", &args, &resp)
	if err == nil && !success {
		if e := resp.EntityError; e != nil {
			err = e
		}
		if e := resp.RequestError; e != nil {
			err = e
		}
		if e := resp.InternalError; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanControllerClient) UpdateDestination(ctx thrift.Context, updateRequest *shared.UpdateDestinationRequest) (*shared.DestinationDescription, error) {
	var resp ControllerUpdateDestinationResult
	args := ControllerUpdateDestinationArgs{
		UpdateRequest: updateRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "updateDestination", &args, &resp)
	if err == nil && !success {
		if e := resp.EntityError; e != nil {
			err = e
		}
		if e := resp.RequestError; e != nil {
			err = e
		}
		if e := resp.InternalError; e != nil {
			err = e
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanControllerClient) UpsertInputHostCapacities(ctx thrift.Context, upsertCapacitiesRequest *UpsertInputHostCapacitiesRequest) error {
	var resp ControllerUpsertInputHostCapacitiesResult
	args := ControllerUpsertInputHostCapacitiesArgs{
		UpsertCapacitiesRequest: upsertCapacitiesRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "upsertInputHostCapacities", &args, &resp)
	if err == nil && !success {
	}

	return err
}

func (c *tchanControllerClient) UpsertOutputHostCapacities(ctx thrift.Context, upsertCapacitiesRequest *UpsertOutputHostCapacitiesRequest) error {
	var resp ControllerUpsertOutputHostCapacitiesResult
	args := ControllerUpsertOutputHostCapacitiesArgs{
		UpsertCapacitiesRequest: upsertCapacitiesRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "upsertOutputHostCapacities", &args, &resp)
	if err == nil && !success {
	}

	return err
}

func (c *tchanControllerClient) UpsertStoreCapacities(ctx thrift.Context, upsertCapacitiesRequest *UpsertStoreCapacitiesRequest) error {
	var resp ControllerUpsertStoreCapacitiesResult
	args := ControllerUpsertStoreCapacitiesArgs{
		UpsertCapacitiesRequest: upsertCapacitiesRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "upsertStoreCapacities", &args, &resp)
	if err == nil && !success {
	}

	return err
}

type tchanControllerServer struct {
	handler TChanController
}

// NewTChanControllerServer wraps a handler for TChanController so it can be
// registered with a thrift.Server.
func NewTChanControllerServer(handler TChanController) thrift.TChanServer {
	return &tchanControllerServer{
		handler,
	}
}

func (s *tchanControllerServer) Service() string {
	return "Controller"
}

func (s *tchanControllerServer) Methods() []string {
	return []string{
		"createConsumerGroup",
		"createDestination",
		"createRemoteZoneExtent",
		"deleteConsumerGroup",
		"deleteDestination",
		"getCapacities",
		"getInputHosts",
		"getOutputHosts",
		"getQueueDepthInfo",
		"removeCapacities",
		"reportConsumerGroupExtentMetric",
		"reportConsumerGroupMetric",
		"reportDestinationExtentMetric",
		"reportDestinationMetric",
		"reportNodeMetric",
		"reportStoreExtentMetric",
		"updateConsumerGroup",
		"updateDestination",
		"upsertInputHostCapacities",
		"upsertOutputHostCapacities",
		"upsertStoreCapacities",
	}
}

func (s *tchanControllerServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "createConsumerGroup":
		return s.handleCreateConsumerGroup(ctx, protocol)
	case "createDestination":
		return s.handleCreateDestination(ctx, protocol)
	case "createRemoteZoneExtent":
		return s.handleCreateRemoteZoneExtent(ctx, protocol)
	case "deleteConsumerGroup":
		return s.handleDeleteConsumerGroup(ctx, protocol)
	case "deleteDestination":
		return s.handleDeleteDestination(ctx, protocol)
	case "getCapacities":
		return s.handleGetCapacities(ctx, protocol)
	case "getInputHosts":
		return s.handleGetInputHosts(ctx, protocol)
	case "getOutputHosts":
		return s.handleGetOutputHosts(ctx, protocol)
	case "getQueueDepthInfo":
		return s.handleGetQueueDepthInfo(ctx, protocol)
	case "removeCapacities":
		return s.handleRemoveCapacities(ctx, protocol)
	case "reportConsumerGroupExtentMetric":
		return s.handleReportConsumerGroupExtentMetric(ctx, protocol)
	case "reportConsumerGroupMetric":
		return s.handleReportConsumerGroupMetric(ctx, protocol)
	case "reportDestinationExtentMetric":
		return s.handleReportDestinationExtentMetric(ctx, protocol)
	case "reportDestinationMetric":
		return s.handleReportDestinationMetric(ctx, protocol)
	case "reportNodeMetric":
		return s.handleReportNodeMetric(ctx, protocol)
	case "reportStoreExtentMetric":
		return s.handleReportStoreExtentMetric(ctx, protocol)
	case "updateConsumerGroup":
		return s.handleUpdateConsumerGroup(ctx, protocol)
	case "updateDestination":
		return s.handleUpdateDestination(ctx, protocol)
	case "upsertInputHostCapacities":
		return s.handleUpsertInputHostCapacities(ctx, protocol)
	case "upsertOutputHostCapacities":
		return s.handleUpsertOutputHostCapacities(ctx, protocol)
	case "upsertStoreCapacities":
		return s.handleUpsertStoreCapacities(ctx, protocol)

	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanControllerServer) handleCreateConsumerGroup(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerCreateConsumerGroupArgs
	var res ControllerCreateConsumerGroupResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.CreateConsumerGroup(ctx, req.CreateRequest)

	if err != nil {
		switch v := err.(type) {
		case *shared.EntityAlreadyExistsError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for entityExistsError returned non-nil error type *shared.EntityAlreadyExistsError but nil value")
			}
			res.EntityExistsError = v
		case *shared.BadRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for requestError returned non-nil error type *shared.BadRequestError but nil value")
			}
			res.RequestError = v
		case *shared.InternalServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for internalError returned non-nil error type *shared.InternalServiceError but nil value")
			}
			res.InternalError = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleCreateDestination(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerCreateDestinationArgs
	var res ControllerCreateDestinationResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.CreateDestination(ctx, req.CreateRequest)

	if err != nil {
		switch v := err.(type) {
		case *shared.EntityAlreadyExistsError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for entityExistsError returned non-nil error type *shared.EntityAlreadyExistsError but nil value")
			}
			res.EntityExistsError = v
		case *shared.BadRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for requestError returned non-nil error type *shared.BadRequestError but nil value")
			}
			res.RequestError = v
		case *shared.InternalServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for internalError returned non-nil error type *shared.InternalServiceError but nil value")
			}
			res.InternalError = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleCreateRemoteZoneExtent(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerCreateRemoteZoneExtentArgs
	var res ControllerCreateRemoteZoneExtentResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.CreateRemoteZoneExtent(ctx, req.CreateRequest)

	if err != nil {
		switch v := err.(type) {
		case *shared.EntityAlreadyExistsError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for entityExistsError returned non-nil error type *shared.EntityAlreadyExistsError but nil value")
			}
			res.EntityExistsError = v
		case *shared.BadRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for requestError returned non-nil error type *shared.BadRequestError but nil value")
			}
			res.RequestError = v
		case *shared.InternalServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for internalError returned non-nil error type *shared.InternalServiceError but nil value")
			}
			res.InternalError = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleDeleteConsumerGroup(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerDeleteConsumerGroupArgs
	var res ControllerDeleteConsumerGroupResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.DeleteConsumerGroup(ctx, req.DeleteRequest)

	if err != nil {
		switch v := err.(type) {
		case *shared.EntityNotExistsError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for entityError returned non-nil error type *shared.EntityNotExistsError but nil value")
			}
			res.EntityError = v
		case *shared.BadRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for requestError returned non-nil error type *shared.BadRequestError but nil value")
			}
			res.RequestError = v
		case *shared.InternalServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for internalError returned non-nil error type *shared.InternalServiceError but nil value")
			}
			res.InternalError = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleDeleteDestination(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerDeleteDestinationArgs
	var res ControllerDeleteDestinationResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.DeleteDestination(ctx, req.DeleteRequest)

	if err != nil {
		switch v := err.(type) {
		case *shared.EntityNotExistsError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for entityError returned non-nil error type *shared.EntityNotExistsError but nil value")
			}
			res.EntityError = v
		case *shared.BadRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for requestError returned non-nil error type *shared.BadRequestError but nil value")
			}
			res.RequestError = v
		case *shared.InternalServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for internalError returned non-nil error type *shared.InternalServiceError but nil value")
			}
			res.InternalError = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleGetCapacities(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerGetCapacitiesArgs
	var res ControllerGetCapacitiesResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.GetCapacities(ctx, req.GetCapacitiesRequest)

	if err != nil {
		return false, nil, err
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleGetInputHosts(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerGetInputHostsArgs
	var res ControllerGetInputHostsResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.GetInputHosts(ctx, req.GetHostsRequest)

	if err != nil {
		switch v := err.(type) {
		case *shared.BadRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for requestError returned non-nil error type *shared.BadRequestError but nil value")
			}
			res.RequestError = v
		case *shared.InternalServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for internalError returned non-nil error type *shared.InternalServiceError but nil value")
			}
			res.InternalError = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleGetOutputHosts(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerGetOutputHostsArgs
	var res ControllerGetOutputHostsResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.GetOutputHosts(ctx, req.GetHostsRequest)

	if err != nil {
		switch v := err.(type) {
		case *shared.BadRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for requestError returned non-nil error type *shared.BadRequestError but nil value")
			}
			res.RequestError = v
		case *shared.InternalServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for internalError returned non-nil error type *shared.InternalServiceError but nil value")
			}
			res.InternalError = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleGetQueueDepthInfo(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerGetQueueDepthInfoArgs
	var res ControllerGetQueueDepthInfoResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.GetQueueDepthInfo(ctx, req.GetQueueDepthInfoRequest)

	if err != nil {
		switch v := err.(type) {
		case *QueueCacheMissError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for cacheMissError returned non-nil error type *QueueCacheMissError but nil value")
			}
			res.CacheMissError = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleRemoveCapacities(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerRemoveCapacitiesArgs
	var res ControllerRemoveCapacitiesResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.RemoveCapacities(ctx, req.RemoveCapacitiesRequest)

	if err != nil {
		return false, nil, err
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleReportConsumerGroupExtentMetric(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerReportConsumerGroupExtentMetricArgs
	var res ControllerReportConsumerGroupExtentMetricResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.ReportConsumerGroupExtentMetric(ctx, req.ReportMetricRequest)

	if err != nil {
		return false, nil, err
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleReportConsumerGroupMetric(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerReportConsumerGroupMetricArgs
	var res ControllerReportConsumerGroupMetricResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.ReportConsumerGroupMetric(ctx, req.ReportMetricRequest)

	if err != nil {
		return false, nil, err
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleReportDestinationExtentMetric(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerReportDestinationExtentMetricArgs
	var res ControllerReportDestinationExtentMetricResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.ReportDestinationExtentMetric(ctx, req.ReportMetricRequest)

	if err != nil {
		return false, nil, err
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleReportDestinationMetric(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerReportDestinationMetricArgs
	var res ControllerReportDestinationMetricResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.ReportDestinationMetric(ctx, req.ReportMetricRequest)

	if err != nil {
		return false, nil, err
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleReportNodeMetric(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerReportNodeMetricArgs
	var res ControllerReportNodeMetricResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.ReportNodeMetric(ctx, req.ReportMetricRequest)

	if err != nil {
		return false, nil, err
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleReportStoreExtentMetric(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerReportStoreExtentMetricArgs
	var res ControllerReportStoreExtentMetricResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.ReportStoreExtentMetric(ctx, req.ReportMetricRequest)

	if err != nil {
		return false, nil, err
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleUpdateConsumerGroup(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerUpdateConsumerGroupArgs
	var res ControllerUpdateConsumerGroupResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.UpdateConsumerGroup(ctx, req.UpdateRequest)

	if err != nil {
		switch v := err.(type) {
		case *shared.EntityNotExistsError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for entityError returned non-nil error type *shared.EntityNotExistsError but nil value")
			}
			res.EntityError = v
		case *shared.BadRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for requestError returned non-nil error type *shared.BadRequestError but nil value")
			}
			res.RequestError = v
		case *shared.InternalServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for internalError returned non-nil error type *shared.InternalServiceError but nil value")
			}
			res.InternalError = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleUpdateDestination(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerUpdateDestinationArgs
	var res ControllerUpdateDestinationResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.UpdateDestination(ctx, req.UpdateRequest)

	if err != nil {
		switch v := err.(type) {
		case *shared.EntityNotExistsError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for entityError returned non-nil error type *shared.EntityNotExistsError but nil value")
			}
			res.EntityError = v
		case *shared.BadRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for requestError returned non-nil error type *shared.BadRequestError but nil value")
			}
			res.RequestError = v
		case *shared.InternalServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for internalError returned non-nil error type *shared.InternalServiceError but nil value")
			}
			res.InternalError = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleUpsertInputHostCapacities(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerUpsertInputHostCapacitiesArgs
	var res ControllerUpsertInputHostCapacitiesResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.UpsertInputHostCapacities(ctx, req.UpsertCapacitiesRequest)

	if err != nil {
		return false, nil, err
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleUpsertOutputHostCapacities(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerUpsertOutputHostCapacitiesArgs
	var res ControllerUpsertOutputHostCapacitiesResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.UpsertOutputHostCapacities(ctx, req.UpsertCapacitiesRequest)

	if err != nil {
		return false, nil, err
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanControllerServer) handleUpsertStoreCapacities(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req ControllerUpsertStoreCapacitiesArgs
	var res ControllerUpsertStoreCapacitiesResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.UpsertStoreCapacities(ctx, req.UpsertCapacitiesRequest)

	if err != nil {
		return false, nil, err
	} else {
	}

	return err == nil, &res, nil
}
