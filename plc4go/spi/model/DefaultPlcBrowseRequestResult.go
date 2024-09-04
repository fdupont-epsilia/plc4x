/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import apiModel "github.com/apache/plc4x/plc4go/pkg/api/model"

var _ apiModel.PlcBrowseRequestResult = &DefaultPlcBrowseRequestResult{}

//go:generate go run ../../tools/plc4xGenerator/main.go -type=DefaultPlcBrowseRequestResult
type DefaultPlcBrowseRequestResult struct {
	Request  apiModel.PlcBrowseRequest
	Response apiModel.PlcBrowseResponse
	Err      error
}

func NewDefaultPlcBrowseRequestResult(Request apiModel.PlcBrowseRequest, Response apiModel.PlcBrowseResponse, Err error) apiModel.PlcBrowseRequestResult {
	return &DefaultPlcBrowseRequestResult{Request, Response, Err}
}

func (d *DefaultPlcBrowseRequestResult) GetRequest() apiModel.PlcBrowseRequest {
	return d.Request
}

func (d *DefaultPlcBrowseRequestResult) GetResponse() apiModel.PlcBrowseResponse {
	return d.Response
}

func (d *DefaultPlcBrowseRequestResult) GetErr() error {
	return d.Err
}
