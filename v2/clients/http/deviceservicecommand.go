//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"path"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"
)

type deviceServiceCommandClient struct {
	baseUrl string
}

// NewDeviceServiceCommandClient creates an instance of deviceServiceCommandClient
func NewDeviceServiceCommandClient(baseUrl string) interfaces.DeviceServiceCommandClient {
	return &deviceServiceCommandClient{
		baseUrl: baseUrl,
	}
}

func (client *deviceServiceCommandClient) ReadCommand(ctx context.Context, deviceName string, commandName string, pushEvent string, returnEvent string) (responses.EventResponse, errors.EdgeX) {
	var response responses.EventResponse
	requestPath := path.Join(v2.ApiDeviceRoute, v2.Name, url.QueryEscape(deviceName), url.QueryEscape(commandName))
	queryParams := url.Values{}
	queryParams.Set(v2.PushEvent, pushEvent)
	queryParams.Set(v2.ReturnEvent, returnEvent)
	err := utils.GetRequest(ctx, &response, client.baseUrl, requestPath, queryParams)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

func (client *deviceServiceCommandClient) WriteCommand(ctx context.Context, deviceName string, commandName string, settings map[string]string) (common.BaseResponse, errors.EdgeX) {
	var response common.BaseResponse
	requestPath := path.Join(v2.ApiDeviceRoute, v2.Name, url.QueryEscape(deviceName), url.QueryEscape(commandName))
	err := utils.PutRequest(ctx, &response, client.baseUrl+requestPath, settings)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}
