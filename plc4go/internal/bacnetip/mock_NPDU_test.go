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
	model "github.com/apache/plc4x/plc4go/protocols/bacnetip/readwrite/model"
	spi "github.com/apache/plc4x/plc4go/spi"
	mock "github.com/stretchr/testify/mock"
)

// MockNPDU is an autogenerated mock type for the NPDU type
type MockNPDU struct {
	mock.Mock
}

type MockNPDU_Expecter struct {
	mock *mock.Mock
}

func (_m *MockNPDU) EXPECT() *MockNPDU_Expecter {
	return &MockNPDU_Expecter{mock: &_m.Mock}
}

// Decode provides a mock function with given fields: pdu
func (_m *MockNPDU) Decode(pdu Arg) error {
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

// MockNPDU_Decode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Decode'
type MockNPDU_Decode_Call struct {
	*mock.Call
}

// Decode is a helper method to define mock.On call
//   - pdu Arg
func (_e *MockNPDU_Expecter) Decode(pdu interface{}) *MockNPDU_Decode_Call {
	return &MockNPDU_Decode_Call{Call: _e.mock.On("Decode", pdu)}
}

func (_c *MockNPDU_Decode_Call) Run(run func(pdu Arg)) *MockNPDU_Decode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Arg))
	})
	return _c
}

func (_c *MockNPDU_Decode_Call) Return(_a0 error) *MockNPDU_Decode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNPDU_Decode_Call) RunAndReturn(run func(Arg) error) *MockNPDU_Decode_Call {
	_c.Call.Return(run)
	return _c
}

// Encode provides a mock function with given fields: pdu
func (_m *MockNPDU) Encode(pdu Arg) error {
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

// MockNPDU_Encode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Encode'
type MockNPDU_Encode_Call struct {
	*mock.Call
}

// Encode is a helper method to define mock.On call
//   - pdu Arg
func (_e *MockNPDU_Expecter) Encode(pdu interface{}) *MockNPDU_Encode_Call {
	return &MockNPDU_Encode_Call{Call: _e.mock.On("Encode", pdu)}
}

func (_c *MockNPDU_Encode_Call) Run(run func(pdu Arg)) *MockNPDU_Encode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Arg))
	})
	return _c
}

func (_c *MockNPDU_Encode_Call) Return(_a0 error) *MockNPDU_Encode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNPDU_Encode_Call) RunAndReturn(run func(Arg) error) *MockNPDU_Encode_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields:
func (_m *MockNPDU) Get() (byte, error) {
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

// MockNPDU_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockNPDU_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
func (_e *MockNPDU_Expecter) Get() *MockNPDU_Get_Call {
	return &MockNPDU_Get_Call{Call: _e.mock.On("Get")}
}

func (_c *MockNPDU_Get_Call) Run(run func()) *MockNPDU_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNPDU_Get_Call) Return(_a0 byte, _a1 error) *MockNPDU_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNPDU_Get_Call) RunAndReturn(run func() (byte, error)) *MockNPDU_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetData provides a mock function with given fields: dlen
func (_m *MockNPDU) GetData(dlen int) ([]byte, error) {
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

// MockNPDU_GetData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetData'
type MockNPDU_GetData_Call struct {
	*mock.Call
}

// GetData is a helper method to define mock.On call
//   - dlen int
func (_e *MockNPDU_Expecter) GetData(dlen interface{}) *MockNPDU_GetData_Call {
	return &MockNPDU_GetData_Call{Call: _e.mock.On("GetData", dlen)}
}

func (_c *MockNPDU_GetData_Call) Run(run func(dlen int)) *MockNPDU_GetData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockNPDU_GetData_Call) Return(_a0 []byte, _a1 error) *MockNPDU_GetData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNPDU_GetData_Call) RunAndReturn(run func(int) ([]byte, error)) *MockNPDU_GetData_Call {
	_c.Call.Return(run)
	return _c
}

// GetExpectingReply provides a mock function with given fields:
func (_m *MockNPDU) GetExpectingReply() bool {
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

// MockNPDU_GetExpectingReply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExpectingReply'
type MockNPDU_GetExpectingReply_Call struct {
	*mock.Call
}

// GetExpectingReply is a helper method to define mock.On call
func (_e *MockNPDU_Expecter) GetExpectingReply() *MockNPDU_GetExpectingReply_Call {
	return &MockNPDU_GetExpectingReply_Call{Call: _e.mock.On("GetExpectingReply")}
}

func (_c *MockNPDU_GetExpectingReply_Call) Run(run func()) *MockNPDU_GetExpectingReply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNPDU_GetExpectingReply_Call) Return(_a0 bool) *MockNPDU_GetExpectingReply_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNPDU_GetExpectingReply_Call) RunAndReturn(run func() bool) *MockNPDU_GetExpectingReply_Call {
	_c.Call.Return(run)
	return _c
}

// GetLong provides a mock function with given fields:
func (_m *MockNPDU) GetLong() (int64, error) {
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

// MockNPDU_GetLong_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLong'
type MockNPDU_GetLong_Call struct {
	*mock.Call
}

// GetLong is a helper method to define mock.On call
func (_e *MockNPDU_Expecter) GetLong() *MockNPDU_GetLong_Call {
	return &MockNPDU_GetLong_Call{Call: _e.mock.On("GetLong")}
}

func (_c *MockNPDU_GetLong_Call) Run(run func()) *MockNPDU_GetLong_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNPDU_GetLong_Call) Return(_a0 int64, _a1 error) *MockNPDU_GetLong_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNPDU_GetLong_Call) RunAndReturn(run func() (int64, error)) *MockNPDU_GetLong_Call {
	_c.Call.Return(run)
	return _c
}

// GetNPDUNetMessage provides a mock function with given fields:
func (_m *MockNPDU) GetNPDUNetMessage() *uint8 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNPDUNetMessage")
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

// MockNPDU_GetNPDUNetMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNPDUNetMessage'
type MockNPDU_GetNPDUNetMessage_Call struct {
	*mock.Call
}

// GetNPDUNetMessage is a helper method to define mock.On call
func (_e *MockNPDU_Expecter) GetNPDUNetMessage() *MockNPDU_GetNPDUNetMessage_Call {
	return &MockNPDU_GetNPDUNetMessage_Call{Call: _e.mock.On("GetNPDUNetMessage")}
}

func (_c *MockNPDU_GetNPDUNetMessage_Call) Run(run func()) *MockNPDU_GetNPDUNetMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNPDU_GetNPDUNetMessage_Call) Return(_a0 *uint8) *MockNPDU_GetNPDUNetMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNPDU_GetNPDUNetMessage_Call) RunAndReturn(run func() *uint8) *MockNPDU_GetNPDUNetMessage_Call {
	_c.Call.Return(run)
	return _c
}

// GetNetworkPriority provides a mock function with given fields:
func (_m *MockNPDU) GetNetworkPriority() model.NPDUNetworkPriority {
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

// MockNPDU_GetNetworkPriority_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNetworkPriority'
type MockNPDU_GetNetworkPriority_Call struct {
	*mock.Call
}

// GetNetworkPriority is a helper method to define mock.On call
func (_e *MockNPDU_Expecter) GetNetworkPriority() *MockNPDU_GetNetworkPriority_Call {
	return &MockNPDU_GetNetworkPriority_Call{Call: _e.mock.On("GetNetworkPriority")}
}

func (_c *MockNPDU_GetNetworkPriority_Call) Run(run func()) *MockNPDU_GetNetworkPriority_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNPDU_GetNetworkPriority_Call) Return(_a0 model.NPDUNetworkPriority) *MockNPDU_GetNetworkPriority_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNPDU_GetNetworkPriority_Call) RunAndReturn(run func() model.NPDUNetworkPriority) *MockNPDU_GetNetworkPriority_Call {
	_c.Call.Return(run)
	return _c
}

// GetPDUDestination provides a mock function with given fields:
func (_m *MockNPDU) GetPDUDestination() *Address {
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

// MockNPDU_GetPDUDestination_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPDUDestination'
type MockNPDU_GetPDUDestination_Call struct {
	*mock.Call
}

// GetPDUDestination is a helper method to define mock.On call
func (_e *MockNPDU_Expecter) GetPDUDestination() *MockNPDU_GetPDUDestination_Call {
	return &MockNPDU_GetPDUDestination_Call{Call: _e.mock.On("GetPDUDestination")}
}

func (_c *MockNPDU_GetPDUDestination_Call) Run(run func()) *MockNPDU_GetPDUDestination_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNPDU_GetPDUDestination_Call) Return(_a0 *Address) *MockNPDU_GetPDUDestination_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNPDU_GetPDUDestination_Call) RunAndReturn(run func() *Address) *MockNPDU_GetPDUDestination_Call {
	_c.Call.Return(run)
	return _c
}

// GetPDUSource provides a mock function with given fields:
func (_m *MockNPDU) GetPDUSource() *Address {
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

// MockNPDU_GetPDUSource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPDUSource'
type MockNPDU_GetPDUSource_Call struct {
	*mock.Call
}

// GetPDUSource is a helper method to define mock.On call
func (_e *MockNPDU_Expecter) GetPDUSource() *MockNPDU_GetPDUSource_Call {
	return &MockNPDU_GetPDUSource_Call{Call: _e.mock.On("GetPDUSource")}
}

func (_c *MockNPDU_GetPDUSource_Call) Run(run func()) *MockNPDU_GetPDUSource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNPDU_GetPDUSource_Call) Return(_a0 *Address) *MockNPDU_GetPDUSource_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNPDU_GetPDUSource_Call) RunAndReturn(run func() *Address) *MockNPDU_GetPDUSource_Call {
	_c.Call.Return(run)
	return _c
}

// GetPDUUserData provides a mock function with given fields:
func (_m *MockNPDU) GetPDUUserData() spi.Message {
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

// MockNPDU_GetPDUUserData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPDUUserData'
type MockNPDU_GetPDUUserData_Call struct {
	*mock.Call
}

// GetPDUUserData is a helper method to define mock.On call
func (_e *MockNPDU_Expecter) GetPDUUserData() *MockNPDU_GetPDUUserData_Call {
	return &MockNPDU_GetPDUUserData_Call{Call: _e.mock.On("GetPDUUserData")}
}

func (_c *MockNPDU_GetPDUUserData_Call) Run(run func()) *MockNPDU_GetPDUUserData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNPDU_GetPDUUserData_Call) Return(_a0 spi.Message) *MockNPDU_GetPDUUserData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNPDU_GetPDUUserData_Call) RunAndReturn(run func() spi.Message) *MockNPDU_GetPDUUserData_Call {
	_c.Call.Return(run)
	return _c
}

// GetPduData provides a mock function with given fields:
func (_m *MockNPDU) GetPduData() []byte {
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

// MockNPDU_GetPduData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPduData'
type MockNPDU_GetPduData_Call struct {
	*mock.Call
}

// GetPduData is a helper method to define mock.On call
func (_e *MockNPDU_Expecter) GetPduData() *MockNPDU_GetPduData_Call {
	return &MockNPDU_GetPduData_Call{Call: _e.mock.On("GetPduData")}
}

func (_c *MockNPDU_GetPduData_Call) Run(run func()) *MockNPDU_GetPduData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNPDU_GetPduData_Call) Return(_a0 []byte) *MockNPDU_GetPduData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNPDU_GetPduData_Call) RunAndReturn(run func() []byte) *MockNPDU_GetPduData_Call {
	_c.Call.Return(run)
	return _c
}

// GetShort provides a mock function with given fields:
func (_m *MockNPDU) GetShort() (int16, error) {
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

// MockNPDU_GetShort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetShort'
type MockNPDU_GetShort_Call struct {
	*mock.Call
}

// GetShort is a helper method to define mock.On call
func (_e *MockNPDU_Expecter) GetShort() *MockNPDU_GetShort_Call {
	return &MockNPDU_GetShort_Call{Call: _e.mock.On("GetShort")}
}

func (_c *MockNPDU_GetShort_Call) Run(run func()) *MockNPDU_GetShort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNPDU_GetShort_Call) Return(_a0 int16, _a1 error) *MockNPDU_GetShort_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNPDU_GetShort_Call) RunAndReturn(run func() (int16, error)) *MockNPDU_GetShort_Call {
	_c.Call.Return(run)
	return _c
}

// Put provides a mock function with given fields: _a0
func (_m *MockNPDU) Put(_a0 byte) {
	_m.Called(_a0)
}

// MockNPDU_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type MockNPDU_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - _a0 byte
func (_e *MockNPDU_Expecter) Put(_a0 interface{}) *MockNPDU_Put_Call {
	return &MockNPDU_Put_Call{Call: _e.mock.On("Put", _a0)}
}

func (_c *MockNPDU_Put_Call) Run(run func(_a0 byte)) *MockNPDU_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(byte))
	})
	return _c
}

func (_c *MockNPDU_Put_Call) Return() *MockNPDU_Put_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockNPDU_Put_Call) RunAndReturn(run func(byte)) *MockNPDU_Put_Call {
	_c.Call.Return(run)
	return _c
}

// PutData provides a mock function with given fields: _a0
func (_m *MockNPDU) PutData(_a0 ...byte) {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockNPDU_PutData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutData'
type MockNPDU_PutData_Call struct {
	*mock.Call
}

// PutData is a helper method to define mock.On call
//   - _a0 ...byte
func (_e *MockNPDU_Expecter) PutData(_a0 ...interface{}) *MockNPDU_PutData_Call {
	return &MockNPDU_PutData_Call{Call: _e.mock.On("PutData",
		append([]interface{}{}, _a0...)...)}
}

func (_c *MockNPDU_PutData_Call) Run(run func(_a0 ...byte)) *MockNPDU_PutData_Call {
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

func (_c *MockNPDU_PutData_Call) Return() *MockNPDU_PutData_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockNPDU_PutData_Call) RunAndReturn(run func(...byte)) *MockNPDU_PutData_Call {
	_c.Call.Return(run)
	return _c
}

// PutLong provides a mock function with given fields: _a0
func (_m *MockNPDU) PutLong(_a0 int64) {
	_m.Called(_a0)
}

// MockNPDU_PutLong_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutLong'
type MockNPDU_PutLong_Call struct {
	*mock.Call
}

// PutLong is a helper method to define mock.On call
//   - _a0 int64
func (_e *MockNPDU_Expecter) PutLong(_a0 interface{}) *MockNPDU_PutLong_Call {
	return &MockNPDU_PutLong_Call{Call: _e.mock.On("PutLong", _a0)}
}

func (_c *MockNPDU_PutLong_Call) Run(run func(_a0 int64)) *MockNPDU_PutLong_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockNPDU_PutLong_Call) Return() *MockNPDU_PutLong_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockNPDU_PutLong_Call) RunAndReturn(run func(int64)) *MockNPDU_PutLong_Call {
	_c.Call.Return(run)
	return _c
}

// PutShort provides a mock function with given fields: _a0
func (_m *MockNPDU) PutShort(_a0 int16) {
	_m.Called(_a0)
}

// MockNPDU_PutShort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutShort'
type MockNPDU_PutShort_Call struct {
	*mock.Call
}

// PutShort is a helper method to define mock.On call
//   - _a0 int16
func (_e *MockNPDU_Expecter) PutShort(_a0 interface{}) *MockNPDU_PutShort_Call {
	return &MockNPDU_PutShort_Call{Call: _e.mock.On("PutShort", _a0)}
}

func (_c *MockNPDU_PutShort_Call) Run(run func(_a0 int16)) *MockNPDU_PutShort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int16))
	})
	return _c
}

func (_c *MockNPDU_PutShort_Call) Return() *MockNPDU_PutShort_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockNPDU_PutShort_Call) RunAndReturn(run func(int16)) *MockNPDU_PutShort_Call {
	_c.Call.Return(run)
	return _c
}

// SetPDUDestination provides a mock function with given fields: _a0
func (_m *MockNPDU) SetPDUDestination(_a0 *Address) {
	_m.Called(_a0)
}

// MockNPDU_SetPDUDestination_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPDUDestination'
type MockNPDU_SetPDUDestination_Call struct {
	*mock.Call
}

// SetPDUDestination is a helper method to define mock.On call
//   - _a0 *Address
func (_e *MockNPDU_Expecter) SetPDUDestination(_a0 interface{}) *MockNPDU_SetPDUDestination_Call {
	return &MockNPDU_SetPDUDestination_Call{Call: _e.mock.On("SetPDUDestination", _a0)}
}

func (_c *MockNPDU_SetPDUDestination_Call) Run(run func(_a0 *Address)) *MockNPDU_SetPDUDestination_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Address))
	})
	return _c
}

func (_c *MockNPDU_SetPDUDestination_Call) Return() *MockNPDU_SetPDUDestination_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockNPDU_SetPDUDestination_Call) RunAndReturn(run func(*Address)) *MockNPDU_SetPDUDestination_Call {
	_c.Call.Return(run)
	return _c
}

// SetPDUSource provides a mock function with given fields: source
func (_m *MockNPDU) SetPDUSource(source *Address) {
	_m.Called(source)
}

// MockNPDU_SetPDUSource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPDUSource'
type MockNPDU_SetPDUSource_Call struct {
	*mock.Call
}

// SetPDUSource is a helper method to define mock.On call
//   - source *Address
func (_e *MockNPDU_Expecter) SetPDUSource(source interface{}) *MockNPDU_SetPDUSource_Call {
	return &MockNPDU_SetPDUSource_Call{Call: _e.mock.On("SetPDUSource", source)}
}

func (_c *MockNPDU_SetPDUSource_Call) Run(run func(source *Address)) *MockNPDU_SetPDUSource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Address))
	})
	return _c
}

func (_c *MockNPDU_SetPDUSource_Call) Return() *MockNPDU_SetPDUSource_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockNPDU_SetPDUSource_Call) RunAndReturn(run func(*Address)) *MockNPDU_SetPDUSource_Call {
	_c.Call.Return(run)
	return _c
}

// SetPDUUserData provides a mock function with given fields: _a0
func (_m *MockNPDU) SetPDUUserData(_a0 spi.Message) {
	_m.Called(_a0)
}

// MockNPDU_SetPDUUserData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPDUUserData'
type MockNPDU_SetPDUUserData_Call struct {
	*mock.Call
}

// SetPDUUserData is a helper method to define mock.On call
//   - _a0 spi.Message
func (_e *MockNPDU_Expecter) SetPDUUserData(_a0 interface{}) *MockNPDU_SetPDUUserData_Call {
	return &MockNPDU_SetPDUUserData_Call{Call: _e.mock.On("SetPDUUserData", _a0)}
}

func (_c *MockNPDU_SetPDUUserData_Call) Run(run func(_a0 spi.Message)) *MockNPDU_SetPDUUserData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spi.Message))
	})
	return _c
}

func (_c *MockNPDU_SetPDUUserData_Call) Return() *MockNPDU_SetPDUUserData_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockNPDU_SetPDUUserData_Call) RunAndReturn(run func(spi.Message)) *MockNPDU_SetPDUUserData_Call {
	_c.Call.Return(run)
	return _c
}

// SetPduData provides a mock function with given fields: _a0
func (_m *MockNPDU) SetPduData(_a0 []byte) {
	_m.Called(_a0)
}

// MockNPDU_SetPduData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPduData'
type MockNPDU_SetPduData_Call struct {
	*mock.Call
}

// SetPduData is a helper method to define mock.On call
//   - _a0 []byte
func (_e *MockNPDU_Expecter) SetPduData(_a0 interface{}) *MockNPDU_SetPduData_Call {
	return &MockNPDU_SetPduData_Call{Call: _e.mock.On("SetPduData", _a0)}
}

func (_c *MockNPDU_SetPduData_Call) Run(run func(_a0 []byte)) *MockNPDU_SetPduData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MockNPDU_SetPduData_Call) Return() *MockNPDU_SetPduData_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockNPDU_SetPduData_Call) RunAndReturn(run func([]byte)) *MockNPDU_SetPduData_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockNPDU) String() string {
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

// MockNPDU_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockNPDU_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockNPDU_Expecter) String() *MockNPDU_String_Call {
	return &MockNPDU_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockNPDU_String_Call) Run(run func()) *MockNPDU_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNPDU_String_Call) Return(_a0 string) *MockNPDU_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNPDU_String_Call) RunAndReturn(run func() string) *MockNPDU_String_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: pci
func (_m *MockNPDU) Update(pci Arg) error {
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

// MockNPDU_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockNPDU_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - pci Arg
func (_e *MockNPDU_Expecter) Update(pci interface{}) *MockNPDU_Update_Call {
	return &MockNPDU_Update_Call{Call: _e.mock.On("Update", pci)}
}

func (_c *MockNPDU_Update_Call) Run(run func(pci Arg)) *MockNPDU_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Arg))
	})
	return _c
}

func (_c *MockNPDU_Update_Call) Return(_a0 error) *MockNPDU_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNPDU_Update_Call) RunAndReturn(run func(Arg) error) *MockNPDU_Update_Call {
	_c.Call.Return(run)
	return _c
}

// setAPDU provides a mock function with given fields: apdu
func (_m *MockNPDU) setAPDU(apdu model.APDU) {
	_m.Called(apdu)
}

// MockNPDU_setAPDU_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'setAPDU'
type MockNPDU_setAPDU_Call struct {
	*mock.Call
}

// setAPDU is a helper method to define mock.On call
//   - apdu model.APDU
func (_e *MockNPDU_Expecter) setAPDU(apdu interface{}) *MockNPDU_setAPDU_Call {
	return &MockNPDU_setAPDU_Call{Call: _e.mock.On("setAPDU", apdu)}
}

func (_c *MockNPDU_setAPDU_Call) Run(run func(apdu model.APDU)) *MockNPDU_setAPDU_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.APDU))
	})
	return _c
}

func (_c *MockNPDU_setAPDU_Call) Return() *MockNPDU_setAPDU_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockNPDU_setAPDU_Call) RunAndReturn(run func(model.APDU)) *MockNPDU_setAPDU_Call {
	_c.Call.Return(run)
	return _c
}

// setNLM provides a mock function with given fields: nlm
func (_m *MockNPDU) setNLM(nlm model.NLM) {
	_m.Called(nlm)
}

// MockNPDU_setNLM_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'setNLM'
type MockNPDU_setNLM_Call struct {
	*mock.Call
}

// setNLM is a helper method to define mock.On call
//   - nlm model.NLM
func (_e *MockNPDU_Expecter) setNLM(nlm interface{}) *MockNPDU_setNLM_Call {
	return &MockNPDU_setNLM_Call{Call: _e.mock.On("setNLM", nlm)}
}

func (_c *MockNPDU_setNLM_Call) Run(run func(nlm model.NLM)) *MockNPDU_setNLM_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.NLM))
	})
	return _c
}

func (_c *MockNPDU_setNLM_Call) Return() *MockNPDU_setNLM_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockNPDU_setNLM_Call) RunAndReturn(run func(model.NLM)) *MockNPDU_setNLM_Call {
	_c.Call.Return(run)
	return _c
}

// setNPDU provides a mock function with given fields: npdu
func (_m *MockNPDU) setNPDU(npdu model.NPDU) {
	_m.Called(npdu)
}

// MockNPDU_setNPDU_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'setNPDU'
type MockNPDU_setNPDU_Call struct {
	*mock.Call
}

// setNPDU is a helper method to define mock.On call
//   - npdu model.NPDU
func (_e *MockNPDU_Expecter) setNPDU(npdu interface{}) *MockNPDU_setNPDU_Call {
	return &MockNPDU_setNPDU_Call{Call: _e.mock.On("setNPDU", npdu)}
}

func (_c *MockNPDU_setNPDU_Call) Run(run func(npdu model.NPDU)) *MockNPDU_setNPDU_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.NPDU))
	})
	return _c
}

func (_c *MockNPDU_setNPDU_Call) Return() *MockNPDU_setNPDU_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockNPDU_setNPDU_Call) RunAndReturn(run func(model.NPDU)) *MockNPDU_setNPDU_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockNPDU creates a new instance of MockNPDU. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockNPDU(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockNPDU {
	mock := &MockNPDU{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}