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

// Code generated by mockery v2.42.2. DO NOT EDIT.

package bacnetip

import (
	context "context"

	model "github.com/apache/plc4x/plc4go/protocols/bacnetip/readwrite/model"
	mock "github.com/stretchr/testify/mock"

	spi "github.com/apache/plc4x/plc4go/spi"

	utils "github.com/apache/plc4x/plc4go/spi/utils"
)

// MockAPDU is an autogenerated mock type for the APDU type
type MockAPDU struct {
	mock.Mock
}

type MockAPDU_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAPDU) EXPECT() *MockAPDU_Expecter {
	return &MockAPDU_Expecter{mock: &_m.Mock}
}

// Decode provides a mock function with given fields: pdu
func (_m *MockAPDU) Decode(pdu Arg) error {
	ret := _m.Called(pdu)

	if len(ret) == 0 {
		panic("no return value specified for Decode")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Arg) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAPDU_Decode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Decode'
type MockAPDU_Decode_Call struct {
	*mock.Call
}

// Decode is a helper method to define mock.On call
//   - pdu Arg
func (_e *MockAPDU_Expecter) Decode(pdu interface{}) *MockAPDU_Decode_Call {
	return &MockAPDU_Decode_Call{Call: _e.mock.On("Decode", pdu)}
}

func (_c *MockAPDU_Decode_Call) Run(run func(pdu Arg)) *MockAPDU_Decode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Arg))
	})
	return _c
}

func (_c *MockAPDU_Decode_Call) Return(_a0 error) *MockAPDU_Decode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAPDU_Decode_Call) RunAndReturn(run func(Arg) error) *MockAPDU_Decode_Call {
	_c.Call.Return(run)
	return _c
}

// Encode provides a mock function with given fields: pdu
func (_m *MockAPDU) Encode(pdu Arg) error {
	ret := _m.Called(pdu)

	if len(ret) == 0 {
		panic("no return value specified for Encode")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Arg) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAPDU_Encode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Encode'
type MockAPDU_Encode_Call struct {
	*mock.Call
}

// Encode is a helper method to define mock.On call
//   - pdu Arg
func (_e *MockAPDU_Expecter) Encode(pdu interface{}) *MockAPDU_Encode_Call {
	return &MockAPDU_Encode_Call{Call: _e.mock.On("Encode", pdu)}
}

func (_c *MockAPDU_Encode_Call) Run(run func(pdu Arg)) *MockAPDU_Encode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Arg))
	})
	return _c
}

func (_c *MockAPDU_Encode_Call) Return(_a0 error) *MockAPDU_Encode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAPDU_Encode_Call) RunAndReturn(run func(Arg) error) *MockAPDU_Encode_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields:
func (_m *MockAPDU) Get() (byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 byte
	var r1 error
	if rf, ok := ret.Get(0).(func() (byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() byte); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(byte)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAPDU_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockAPDU_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
func (_e *MockAPDU_Expecter) Get() *MockAPDU_Get_Call {
	return &MockAPDU_Get_Call{Call: _e.mock.On("Get")}
}

func (_c *MockAPDU_Get_Call) Run(run func()) *MockAPDU_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAPDU_Get_Call) Return(_a0 byte, _a1 error) *MockAPDU_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAPDU_Get_Call) RunAndReturn(run func() (byte, error)) *MockAPDU_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetApduInvokeID provides a mock function with given fields:
func (_m *MockAPDU) GetApduInvokeID() *uint8 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetApduInvokeID")
	}

	var r0 *uint8
	if rf, ok := ret.Get(0).(func() *uint8); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*uint8)
		}
	}

	return r0
}

// MockAPDU_GetApduInvokeID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetApduInvokeID'
type MockAPDU_GetApduInvokeID_Call struct {
	*mock.Call
}

// GetApduInvokeID is a helper method to define mock.On call
func (_e *MockAPDU_Expecter) GetApduInvokeID() *MockAPDU_GetApduInvokeID_Call {
	return &MockAPDU_GetApduInvokeID_Call{Call: _e.mock.On("GetApduInvokeID")}
}

func (_c *MockAPDU_GetApduInvokeID_Call) Run(run func()) *MockAPDU_GetApduInvokeID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAPDU_GetApduInvokeID_Call) Return(_a0 *uint8) *MockAPDU_GetApduInvokeID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAPDU_GetApduInvokeID_Call) RunAndReturn(run func() *uint8) *MockAPDU_GetApduInvokeID_Call {
	_c.Call.Return(run)
	return _c
}

// GetApduType provides a mock function with given fields:
func (_m *MockAPDU) GetApduType() model.ApduType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetApduType")
	}

	var r0 model.ApduType
	if rf, ok := ret.Get(0).(func() model.ApduType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.ApduType)
	}

	return r0
}

// MockAPDU_GetApduType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetApduType'
type MockAPDU_GetApduType_Call struct {
	*mock.Call
}

// GetApduType is a helper method to define mock.On call
func (_e *MockAPDU_Expecter) GetApduType() *MockAPDU_GetApduType_Call {
	return &MockAPDU_GetApduType_Call{Call: _e.mock.On("GetApduType")}
}

func (_c *MockAPDU_GetApduType_Call) Run(run func()) *MockAPDU_GetApduType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAPDU_GetApduType_Call) Return(_a0 model.ApduType) *MockAPDU_GetApduType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAPDU_GetApduType_Call) RunAndReturn(run func() model.ApduType) *MockAPDU_GetApduType_Call {
	_c.Call.Return(run)
	return _c
}

// GetData provides a mock function with given fields: dlen
func (_m *MockAPDU) GetData(dlen int) ([]byte, error) {
	ret := _m.Called(dlen)

	if len(ret) == 0 {
		panic("no return value specified for GetData")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]byte, error)); ok {
		return rf(dlen)
	}
	if rf, ok := ret.Get(0).(func(int) []byte); ok {
		r0 = rf(dlen)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(dlen)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAPDU_GetData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetData'
type MockAPDU_GetData_Call struct {
	*mock.Call
}

// GetData is a helper method to define mock.On call
//   - dlen int
func (_e *MockAPDU_Expecter) GetData(dlen interface{}) *MockAPDU_GetData_Call {
	return &MockAPDU_GetData_Call{Call: _e.mock.On("GetData", dlen)}
}

func (_c *MockAPDU_GetData_Call) Run(run func(dlen int)) *MockAPDU_GetData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockAPDU_GetData_Call) Return(_a0 []byte, _a1 error) *MockAPDU_GetData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAPDU_GetData_Call) RunAndReturn(run func(int) ([]byte, error)) *MockAPDU_GetData_Call {
	_c.Call.Return(run)
	return _c
}

// GetExpectingReply provides a mock function with given fields:
func (_m *MockAPDU) GetExpectingReply() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetExpectingReply")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockAPDU_GetExpectingReply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExpectingReply'
type MockAPDU_GetExpectingReply_Call struct {
	*mock.Call
}

// GetExpectingReply is a helper method to define mock.On call
func (_e *MockAPDU_Expecter) GetExpectingReply() *MockAPDU_GetExpectingReply_Call {
	return &MockAPDU_GetExpectingReply_Call{Call: _e.mock.On("GetExpectingReply")}
}

func (_c *MockAPDU_GetExpectingReply_Call) Run(run func()) *MockAPDU_GetExpectingReply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAPDU_GetExpectingReply_Call) Return(_a0 bool) *MockAPDU_GetExpectingReply_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAPDU_GetExpectingReply_Call) RunAndReturn(run func() bool) *MockAPDU_GetExpectingReply_Call {
	_c.Call.Return(run)
	return _c
}

// GetLengthInBits provides a mock function with given fields: ctx
func (_m *MockAPDU) GetLengthInBits(ctx context.Context) uint16 {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetLengthInBits")
	}

	var r0 uint16
	if rf, ok := ret.Get(0).(func(context.Context) uint16); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// MockAPDU_GetLengthInBits_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLengthInBits'
type MockAPDU_GetLengthInBits_Call struct {
	*mock.Call
}

// GetLengthInBits is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockAPDU_Expecter) GetLengthInBits(ctx interface{}) *MockAPDU_GetLengthInBits_Call {
	return &MockAPDU_GetLengthInBits_Call{Call: _e.mock.On("GetLengthInBits", ctx)}
}

func (_c *MockAPDU_GetLengthInBits_Call) Run(run func(ctx context.Context)) *MockAPDU_GetLengthInBits_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockAPDU_GetLengthInBits_Call) Return(_a0 uint16) *MockAPDU_GetLengthInBits_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAPDU_GetLengthInBits_Call) RunAndReturn(run func(context.Context) uint16) *MockAPDU_GetLengthInBits_Call {
	_c.Call.Return(run)
	return _c
}

// GetLengthInBytes provides a mock function with given fields: ctx
func (_m *MockAPDU) GetLengthInBytes(ctx context.Context) uint16 {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetLengthInBytes")
	}

	var r0 uint16
	if rf, ok := ret.Get(0).(func(context.Context) uint16); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// MockAPDU_GetLengthInBytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLengthInBytes'
type MockAPDU_GetLengthInBytes_Call struct {
	*mock.Call
}

// GetLengthInBytes is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockAPDU_Expecter) GetLengthInBytes(ctx interface{}) *MockAPDU_GetLengthInBytes_Call {
	return &MockAPDU_GetLengthInBytes_Call{Call: _e.mock.On("GetLengthInBytes", ctx)}
}

func (_c *MockAPDU_GetLengthInBytes_Call) Run(run func(ctx context.Context)) *MockAPDU_GetLengthInBytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockAPDU_GetLengthInBytes_Call) Return(_a0 uint16) *MockAPDU_GetLengthInBytes_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAPDU_GetLengthInBytes_Call) RunAndReturn(run func(context.Context) uint16) *MockAPDU_GetLengthInBytes_Call {
	_c.Call.Return(run)
	return _c
}

// GetLong provides a mock function with given fields:
func (_m *MockAPDU) GetLong() (int64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLong")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAPDU_GetLong_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLong'
type MockAPDU_GetLong_Call struct {
	*mock.Call
}

// GetLong is a helper method to define mock.On call
func (_e *MockAPDU_Expecter) GetLong() *MockAPDU_GetLong_Call {
	return &MockAPDU_GetLong_Call{Call: _e.mock.On("GetLong")}
}

func (_c *MockAPDU_GetLong_Call) Run(run func()) *MockAPDU_GetLong_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAPDU_GetLong_Call) Return(_a0 int64, _a1 error) *MockAPDU_GetLong_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAPDU_GetLong_Call) RunAndReturn(run func() (int64, error)) *MockAPDU_GetLong_Call {
	_c.Call.Return(run)
	return _c
}

// GetNetworkPriority provides a mock function with given fields:
func (_m *MockAPDU) GetNetworkPriority() model.NPDUNetworkPriority {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNetworkPriority")
	}

	var r0 model.NPDUNetworkPriority
	if rf, ok := ret.Get(0).(func() model.NPDUNetworkPriority); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.NPDUNetworkPriority)
	}

	return r0
}

// MockAPDU_GetNetworkPriority_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNetworkPriority'
type MockAPDU_GetNetworkPriority_Call struct {
	*mock.Call
}

// GetNetworkPriority is a helper method to define mock.On call
func (_e *MockAPDU_Expecter) GetNetworkPriority() *MockAPDU_GetNetworkPriority_Call {
	return &MockAPDU_GetNetworkPriority_Call{Call: _e.mock.On("GetNetworkPriority")}
}

func (_c *MockAPDU_GetNetworkPriority_Call) Run(run func()) *MockAPDU_GetNetworkPriority_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAPDU_GetNetworkPriority_Call) Return(_a0 model.NPDUNetworkPriority) *MockAPDU_GetNetworkPriority_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAPDU_GetNetworkPriority_Call) RunAndReturn(run func() model.NPDUNetworkPriority) *MockAPDU_GetNetworkPriority_Call {
	_c.Call.Return(run)
	return _c
}

// GetPDUDestination provides a mock function with given fields:
func (_m *MockAPDU) GetPDUDestination() *Address {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPDUDestination")
	}

	var r0 *Address
	if rf, ok := ret.Get(0).(func() *Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Address)
		}
	}

	return r0
}

// MockAPDU_GetPDUDestination_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPDUDestination'
type MockAPDU_GetPDUDestination_Call struct {
	*mock.Call
}

// GetPDUDestination is a helper method to define mock.On call
func (_e *MockAPDU_Expecter) GetPDUDestination() *MockAPDU_GetPDUDestination_Call {
	return &MockAPDU_GetPDUDestination_Call{Call: _e.mock.On("GetPDUDestination")}
}

func (_c *MockAPDU_GetPDUDestination_Call) Run(run func()) *MockAPDU_GetPDUDestination_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAPDU_GetPDUDestination_Call) Return(_a0 *Address) *MockAPDU_GetPDUDestination_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAPDU_GetPDUDestination_Call) RunAndReturn(run func() *Address) *MockAPDU_GetPDUDestination_Call {
	_c.Call.Return(run)
	return _c
}

// GetPDUSource provides a mock function with given fields:
func (_m *MockAPDU) GetPDUSource() *Address {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPDUSource")
	}

	var r0 *Address
	if rf, ok := ret.Get(0).(func() *Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Address)
		}
	}

	return r0
}

// MockAPDU_GetPDUSource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPDUSource'
type MockAPDU_GetPDUSource_Call struct {
	*mock.Call
}

// GetPDUSource is a helper method to define mock.On call
func (_e *MockAPDU_Expecter) GetPDUSource() *MockAPDU_GetPDUSource_Call {
	return &MockAPDU_GetPDUSource_Call{Call: _e.mock.On("GetPDUSource")}
}

func (_c *MockAPDU_GetPDUSource_Call) Run(run func()) *MockAPDU_GetPDUSource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAPDU_GetPDUSource_Call) Return(_a0 *Address) *MockAPDU_GetPDUSource_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAPDU_GetPDUSource_Call) RunAndReturn(run func() *Address) *MockAPDU_GetPDUSource_Call {
	_c.Call.Return(run)
	return _c
}

// GetPDUUserData provides a mock function with given fields:
func (_m *MockAPDU) GetPDUUserData() spi.Message {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPDUUserData")
	}

	var r0 spi.Message
	if rf, ok := ret.Get(0).(func() spi.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spi.Message)
		}
	}

	return r0
}

// MockAPDU_GetPDUUserData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPDUUserData'
type MockAPDU_GetPDUUserData_Call struct {
	*mock.Call
}

// GetPDUUserData is a helper method to define mock.On call
func (_e *MockAPDU_Expecter) GetPDUUserData() *MockAPDU_GetPDUUserData_Call {
	return &MockAPDU_GetPDUUserData_Call{Call: _e.mock.On("GetPDUUserData")}
}

func (_c *MockAPDU_GetPDUUserData_Call) Run(run func()) *MockAPDU_GetPDUUserData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAPDU_GetPDUUserData_Call) Return(_a0 spi.Message) *MockAPDU_GetPDUUserData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAPDU_GetPDUUserData_Call) RunAndReturn(run func() spi.Message) *MockAPDU_GetPDUUserData_Call {
	_c.Call.Return(run)
	return _c
}

// GetPduData provides a mock function with given fields:
func (_m *MockAPDU) GetPduData() []byte {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPduData")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// MockAPDU_GetPduData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPduData'
type MockAPDU_GetPduData_Call struct {
	*mock.Call
}

// GetPduData is a helper method to define mock.On call
func (_e *MockAPDU_Expecter) GetPduData() *MockAPDU_GetPduData_Call {
	return &MockAPDU_GetPduData_Call{Call: _e.mock.On("GetPduData")}
}

func (_c *MockAPDU_GetPduData_Call) Run(run func()) *MockAPDU_GetPduData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAPDU_GetPduData_Call) Return(_a0 []byte) *MockAPDU_GetPduData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAPDU_GetPduData_Call) RunAndReturn(run func() []byte) *MockAPDU_GetPduData_Call {
	_c.Call.Return(run)
	return _c
}

// GetRootMessage provides a mock function with given fields:
func (_m *MockAPDU) GetRootMessage() spi.Message {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetRootMessage")
	}

	var r0 spi.Message
	if rf, ok := ret.Get(0).(func() spi.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spi.Message)
		}
	}

	return r0
}

// MockAPDU_GetRootMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRootMessage'
type MockAPDU_GetRootMessage_Call struct {
	*mock.Call
}

// GetRootMessage is a helper method to define mock.On call
func (_e *MockAPDU_Expecter) GetRootMessage() *MockAPDU_GetRootMessage_Call {
	return &MockAPDU_GetRootMessage_Call{Call: _e.mock.On("GetRootMessage")}
}

func (_c *MockAPDU_GetRootMessage_Call) Run(run func()) *MockAPDU_GetRootMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAPDU_GetRootMessage_Call) Return(_a0 spi.Message) *MockAPDU_GetRootMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAPDU_GetRootMessage_Call) RunAndReturn(run func() spi.Message) *MockAPDU_GetRootMessage_Call {
	_c.Call.Return(run)
	return _c
}

// GetShort provides a mock function with given fields:
func (_m *MockAPDU) GetShort() (int16, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetShort")
	}

	var r0 int16
	var r1 error
	if rf, ok := ret.Get(0).(func() (int16, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int16)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAPDU_GetShort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetShort'
type MockAPDU_GetShort_Call struct {
	*mock.Call
}

// GetShort is a helper method to define mock.On call
func (_e *MockAPDU_Expecter) GetShort() *MockAPDU_GetShort_Call {
	return &MockAPDU_GetShort_Call{Call: _e.mock.On("GetShort")}
}

func (_c *MockAPDU_GetShort_Call) Run(run func()) *MockAPDU_GetShort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAPDU_GetShort_Call) Return(_a0 int16, _a1 error) *MockAPDU_GetShort_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAPDU_GetShort_Call) RunAndReturn(run func() (int16, error)) *MockAPDU_GetShort_Call {
	_c.Call.Return(run)
	return _c
}

// Put provides a mock function with given fields: _a0
func (_m *MockAPDU) Put(_a0 byte) {
	_m.Called(_a0)
}

// MockAPDU_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type MockAPDU_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - _a0 byte
func (_e *MockAPDU_Expecter) Put(_a0 interface{}) *MockAPDU_Put_Call {
	return &MockAPDU_Put_Call{Call: _e.mock.On("Put", _a0)}
}

func (_c *MockAPDU_Put_Call) Run(run func(_a0 byte)) *MockAPDU_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(byte))
	})
	return _c
}

func (_c *MockAPDU_Put_Call) Return() *MockAPDU_Put_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockAPDU_Put_Call) RunAndReturn(run func(byte)) *MockAPDU_Put_Call {
	_c.Call.Return(run)
	return _c
}

// PutData provides a mock function with given fields: _a0
func (_m *MockAPDU) PutData(_a0 ...byte) {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockAPDU_PutData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutData'
type MockAPDU_PutData_Call struct {
	*mock.Call
}

// PutData is a helper method to define mock.On call
//   - _a0 ...byte
func (_e *MockAPDU_Expecter) PutData(_a0 ...interface{}) *MockAPDU_PutData_Call {
	return &MockAPDU_PutData_Call{Call: _e.mock.On("PutData",
		append([]interface{}{}, _a0...)...)}
}

func (_c *MockAPDU_PutData_Call) Run(run func(_a0 ...byte)) *MockAPDU_PutData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]byte, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(byte)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockAPDU_PutData_Call) Return() *MockAPDU_PutData_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockAPDU_PutData_Call) RunAndReturn(run func(...byte)) *MockAPDU_PutData_Call {
	_c.Call.Return(run)
	return _c
}

// PutLong provides a mock function with given fields: _a0
func (_m *MockAPDU) PutLong(_a0 uint32) {
	_m.Called(_a0)
}

// MockAPDU_PutLong_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutLong'
type MockAPDU_PutLong_Call struct {
	*mock.Call
}

// PutLong is a helper method to define mock.On call
//   - _a0 uint32
func (_e *MockAPDU_Expecter) PutLong(_a0 interface{}) *MockAPDU_PutLong_Call {
	return &MockAPDU_PutLong_Call{Call: _e.mock.On("PutLong", _a0)}
}

func (_c *MockAPDU_PutLong_Call) Run(run func(_a0 uint32)) *MockAPDU_PutLong_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint32))
	})
	return _c
}

func (_c *MockAPDU_PutLong_Call) Return() *MockAPDU_PutLong_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockAPDU_PutLong_Call) RunAndReturn(run func(uint32)) *MockAPDU_PutLong_Call {
	_c.Call.Return(run)
	return _c
}

// PutShort provides a mock function with given fields: _a0
func (_m *MockAPDU) PutShort(_a0 uint16) {
	_m.Called(_a0)
}

// MockAPDU_PutShort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutShort'
type MockAPDU_PutShort_Call struct {
	*mock.Call
}

// PutShort is a helper method to define mock.On call
//   - _a0 uint16
func (_e *MockAPDU_Expecter) PutShort(_a0 interface{}) *MockAPDU_PutShort_Call {
	return &MockAPDU_PutShort_Call{Call: _e.mock.On("PutShort", _a0)}
}

func (_c *MockAPDU_PutShort_Call) Run(run func(_a0 uint16)) *MockAPDU_PutShort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint16))
	})
	return _c
}

func (_c *MockAPDU_PutShort_Call) Return() *MockAPDU_PutShort_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockAPDU_PutShort_Call) RunAndReturn(run func(uint16)) *MockAPDU_PutShort_Call {
	_c.Call.Return(run)
	return _c
}

// Serialize provides a mock function with given fields:
func (_m *MockAPDU) Serialize() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Serialize")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAPDU_Serialize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Serialize'
type MockAPDU_Serialize_Call struct {
	*mock.Call
}

// Serialize is a helper method to define mock.On call
func (_e *MockAPDU_Expecter) Serialize() *MockAPDU_Serialize_Call {
	return &MockAPDU_Serialize_Call{Call: _e.mock.On("Serialize")}
}

func (_c *MockAPDU_Serialize_Call) Run(run func()) *MockAPDU_Serialize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAPDU_Serialize_Call) Return(_a0 []byte, _a1 error) *MockAPDU_Serialize_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAPDU_Serialize_Call) RunAndReturn(run func() ([]byte, error)) *MockAPDU_Serialize_Call {
	_c.Call.Return(run)
	return _c
}

// SerializeWithWriteBuffer provides a mock function with given fields: ctx, writeBuffer
func (_m *MockAPDU) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	ret := _m.Called(ctx, writeBuffer)

	if len(ret) == 0 {
		panic("no return value specified for SerializeWithWriteBuffer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, utils.WriteBuffer) error); ok {
		r0 = rf(ctx, writeBuffer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAPDU_SerializeWithWriteBuffer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SerializeWithWriteBuffer'
type MockAPDU_SerializeWithWriteBuffer_Call struct {
	*mock.Call
}

// SerializeWithWriteBuffer is a helper method to define mock.On call
//   - ctx context.Context
//   - writeBuffer utils.WriteBuffer
func (_e *MockAPDU_Expecter) SerializeWithWriteBuffer(ctx interface{}, writeBuffer interface{}) *MockAPDU_SerializeWithWriteBuffer_Call {
	return &MockAPDU_SerializeWithWriteBuffer_Call{Call: _e.mock.On("SerializeWithWriteBuffer", ctx, writeBuffer)}
}

func (_c *MockAPDU_SerializeWithWriteBuffer_Call) Run(run func(ctx context.Context, writeBuffer utils.WriteBuffer)) *MockAPDU_SerializeWithWriteBuffer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(utils.WriteBuffer))
	})
	return _c
}

func (_c *MockAPDU_SerializeWithWriteBuffer_Call) Return(_a0 error) *MockAPDU_SerializeWithWriteBuffer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAPDU_SerializeWithWriteBuffer_Call) RunAndReturn(run func(context.Context, utils.WriteBuffer) error) *MockAPDU_SerializeWithWriteBuffer_Call {
	_c.Call.Return(run)
	return _c
}

// SetPDUDestination provides a mock function with given fields: _a0
func (_m *MockAPDU) SetPDUDestination(_a0 *Address) {
	_m.Called(_a0)
}

// MockAPDU_SetPDUDestination_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPDUDestination'
type MockAPDU_SetPDUDestination_Call struct {
	*mock.Call
}

// SetPDUDestination is a helper method to define mock.On call
//   - _a0 *Address
func (_e *MockAPDU_Expecter) SetPDUDestination(_a0 interface{}) *MockAPDU_SetPDUDestination_Call {
	return &MockAPDU_SetPDUDestination_Call{Call: _e.mock.On("SetPDUDestination", _a0)}
}

func (_c *MockAPDU_SetPDUDestination_Call) Run(run func(_a0 *Address)) *MockAPDU_SetPDUDestination_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Address))
	})
	return _c
}

func (_c *MockAPDU_SetPDUDestination_Call) Return() *MockAPDU_SetPDUDestination_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockAPDU_SetPDUDestination_Call) RunAndReturn(run func(*Address)) *MockAPDU_SetPDUDestination_Call {
	_c.Call.Return(run)
	return _c
}

// SetPDUSource provides a mock function with given fields: source
func (_m *MockAPDU) SetPDUSource(source *Address) {
	_m.Called(source)
}

// MockAPDU_SetPDUSource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPDUSource'
type MockAPDU_SetPDUSource_Call struct {
	*mock.Call
}

// SetPDUSource is a helper method to define mock.On call
//   - source *Address
func (_e *MockAPDU_Expecter) SetPDUSource(source interface{}) *MockAPDU_SetPDUSource_Call {
	return &MockAPDU_SetPDUSource_Call{Call: _e.mock.On("SetPDUSource", source)}
}

func (_c *MockAPDU_SetPDUSource_Call) Run(run func(source *Address)) *MockAPDU_SetPDUSource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Address))
	})
	return _c
}

func (_c *MockAPDU_SetPDUSource_Call) Return() *MockAPDU_SetPDUSource_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockAPDU_SetPDUSource_Call) RunAndReturn(run func(*Address)) *MockAPDU_SetPDUSource_Call {
	_c.Call.Return(run)
	return _c
}

// SetPDUUserData provides a mock function with given fields: _a0
func (_m *MockAPDU) SetPDUUserData(_a0 spi.Message) {
	_m.Called(_a0)
}

// MockAPDU_SetPDUUserData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPDUUserData'
type MockAPDU_SetPDUUserData_Call struct {
	*mock.Call
}

// SetPDUUserData is a helper method to define mock.On call
//   - _a0 spi.Message
func (_e *MockAPDU_Expecter) SetPDUUserData(_a0 interface{}) *MockAPDU_SetPDUUserData_Call {
	return &MockAPDU_SetPDUUserData_Call{Call: _e.mock.On("SetPDUUserData", _a0)}
}

func (_c *MockAPDU_SetPDUUserData_Call) Run(run func(_a0 spi.Message)) *MockAPDU_SetPDUUserData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spi.Message))
	})
	return _c
}

func (_c *MockAPDU_SetPDUUserData_Call) Return() *MockAPDU_SetPDUUserData_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockAPDU_SetPDUUserData_Call) RunAndReturn(run func(spi.Message)) *MockAPDU_SetPDUUserData_Call {
	_c.Call.Return(run)
	return _c
}

// SetPduData provides a mock function with given fields: _a0
func (_m *MockAPDU) SetPduData(_a0 []byte) {
	_m.Called(_a0)
}

// MockAPDU_SetPduData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPduData'
type MockAPDU_SetPduData_Call struct {
	*mock.Call
}

// SetPduData is a helper method to define mock.On call
//   - _a0 []byte
func (_e *MockAPDU_Expecter) SetPduData(_a0 interface{}) *MockAPDU_SetPduData_Call {
	return &MockAPDU_SetPduData_Call{Call: _e.mock.On("SetPduData", _a0)}
}

func (_c *MockAPDU_SetPduData_Call) Run(run func(_a0 []byte)) *MockAPDU_SetPduData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MockAPDU_SetPduData_Call) Return() *MockAPDU_SetPduData_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockAPDU_SetPduData_Call) RunAndReturn(run func([]byte)) *MockAPDU_SetPduData_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockAPDU) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockAPDU_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockAPDU_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockAPDU_Expecter) String() *MockAPDU_String_Call {
	return &MockAPDU_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockAPDU_String_Call) Run(run func()) *MockAPDU_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAPDU_String_Call) Return(_a0 string) *MockAPDU_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAPDU_String_Call) RunAndReturn(run func() string) *MockAPDU_String_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: pci
func (_m *MockAPDU) Update(pci Arg) error {
	ret := _m.Called(pci)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Arg) error); ok {
		r0 = rf(pci)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAPDU_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockAPDU_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - pci Arg
func (_e *MockAPDU_Expecter) Update(pci interface{}) *MockAPDU_Update_Call {
	return &MockAPDU_Update_Call{Call: _e.mock.On("Update", pci)}
}

func (_c *MockAPDU_Update_Call) Run(run func(pci Arg)) *MockAPDU_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Arg))
	})
	return _c
}

func (_c *MockAPDU_Update_Call) Return(_a0 error) *MockAPDU_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAPDU_Update_Call) RunAndReturn(run func(Arg) error) *MockAPDU_Update_Call {
	_c.Call.Return(run)
	return _c
}

// getAPDU provides a mock function with given fields:
func (_m *MockAPDU) getAPDU() model.APDU {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for getAPDU")
	}

	var r0 model.APDU
	if rf, ok := ret.Get(0).(func() model.APDU); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.APDU)
		}
	}

	return r0
}

// MockAPDU_getAPDU_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getAPDU'
type MockAPDU_getAPDU_Call struct {
	*mock.Call
}

// getAPDU is a helper method to define mock.On call
func (_e *MockAPDU_Expecter) getAPDU() *MockAPDU_getAPDU_Call {
	return &MockAPDU_getAPDU_Call{Call: _e.mock.On("getAPDU")}
}

func (_c *MockAPDU_getAPDU_Call) Run(run func()) *MockAPDU_getAPDU_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAPDU_getAPDU_Call) Return(_a0 model.APDU) *MockAPDU_getAPDU_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAPDU_getAPDU_Call) RunAndReturn(run func() model.APDU) *MockAPDU_getAPDU_Call {
	_c.Call.Return(run)
	return _c
}

// setAPDU provides a mock function with given fields: _a0
func (_m *MockAPDU) setAPDU(_a0 model.APDU) {
	_m.Called(_a0)
}

// MockAPDU_setAPDU_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'setAPDU'
type MockAPDU_setAPDU_Call struct {
	*mock.Call
}

// setAPDU is a helper method to define mock.On call
//   - _a0 model.APDU
func (_e *MockAPDU_Expecter) setAPDU(_a0 interface{}) *MockAPDU_setAPDU_Call {
	return &MockAPDU_setAPDU_Call{Call: _e.mock.On("setAPDU", _a0)}
}

func (_c *MockAPDU_setAPDU_Call) Run(run func(_a0 model.APDU)) *MockAPDU_setAPDU_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.APDU))
	})
	return _c
}

func (_c *MockAPDU_setAPDU_Call) Return() *MockAPDU_setAPDU_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockAPDU_setAPDU_Call) RunAndReturn(run func(model.APDU)) *MockAPDU_setAPDU_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAPDU creates a new instance of MockAPDU. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAPDU(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAPDU {
	mock := &MockAPDU{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
