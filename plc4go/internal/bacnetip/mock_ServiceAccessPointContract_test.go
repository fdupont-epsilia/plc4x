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

import mock "github.com/stretchr/testify/mock"

// MockServiceAccessPointContract is an autogenerated mock type for the ServiceAccessPointContract type
type MockServiceAccessPointContract struct {
	mock.Mock
}

type MockServiceAccessPointContract_Expecter struct {
	mock *mock.Mock
}

func (_m *MockServiceAccessPointContract) EXPECT() *MockServiceAccessPointContract_Expecter {
	return &MockServiceAccessPointContract_Expecter{mock: &_m.Mock}
}

// SapConfirmation provides a mock function with given fields: _a0, _a1
func (_m *MockServiceAccessPointContract) SapConfirmation(_a0 Args, _a1 KWArgs) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SapConfirmation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Args, KWArgs) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockServiceAccessPointContract_SapConfirmation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SapConfirmation'
type MockServiceAccessPointContract_SapConfirmation_Call struct {
	*mock.Call
}

// SapConfirmation is a helper method to define mock.On call
//   - _a0 Args
//   - _a1 KWArgs
func (_e *MockServiceAccessPointContract_Expecter) SapConfirmation(_a0 interface{}, _a1 interface{}) *MockServiceAccessPointContract_SapConfirmation_Call {
	return &MockServiceAccessPointContract_SapConfirmation_Call{Call: _e.mock.On("SapConfirmation", _a0, _a1)}
}

func (_c *MockServiceAccessPointContract_SapConfirmation_Call) Run(run func(_a0 Args, _a1 KWArgs)) *MockServiceAccessPointContract_SapConfirmation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Args), args[1].(KWArgs))
	})
	return _c
}

func (_c *MockServiceAccessPointContract_SapConfirmation_Call) Return(_a0 error) *MockServiceAccessPointContract_SapConfirmation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServiceAccessPointContract_SapConfirmation_Call) RunAndReturn(run func(Args, KWArgs) error) *MockServiceAccessPointContract_SapConfirmation_Call {
	_c.Call.Return(run)
	return _c
}

// SapIndication provides a mock function with given fields: _a0, _a1
func (_m *MockServiceAccessPointContract) SapIndication(_a0 Args, _a1 KWArgs) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SapIndication")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Args, KWArgs) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockServiceAccessPointContract_SapIndication_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SapIndication'
type MockServiceAccessPointContract_SapIndication_Call struct {
	*mock.Call
}

// SapIndication is a helper method to define mock.On call
//   - _a0 Args
//   - _a1 KWArgs
func (_e *MockServiceAccessPointContract_Expecter) SapIndication(_a0 interface{}, _a1 interface{}) *MockServiceAccessPointContract_SapIndication_Call {
	return &MockServiceAccessPointContract_SapIndication_Call{Call: _e.mock.On("SapIndication", _a0, _a1)}
}

func (_c *MockServiceAccessPointContract_SapIndication_Call) Run(run func(_a0 Args, _a1 KWArgs)) *MockServiceAccessPointContract_SapIndication_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Args), args[1].(KWArgs))
	})
	return _c
}

func (_c *MockServiceAccessPointContract_SapIndication_Call) Return(_a0 error) *MockServiceAccessPointContract_SapIndication_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServiceAccessPointContract_SapIndication_Call) RunAndReturn(run func(Args, KWArgs) error) *MockServiceAccessPointContract_SapIndication_Call {
	_c.Call.Return(run)
	return _c
}

// SapRequest provides a mock function with given fields: _a0, _a1
func (_m *MockServiceAccessPointContract) SapRequest(_a0 Args, _a1 KWArgs) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SapRequest")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Args, KWArgs) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockServiceAccessPointContract_SapRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SapRequest'
type MockServiceAccessPointContract_SapRequest_Call struct {
	*mock.Call
}

// SapRequest is a helper method to define mock.On call
//   - _a0 Args
//   - _a1 KWArgs
func (_e *MockServiceAccessPointContract_Expecter) SapRequest(_a0 interface{}, _a1 interface{}) *MockServiceAccessPointContract_SapRequest_Call {
	return &MockServiceAccessPointContract_SapRequest_Call{Call: _e.mock.On("SapRequest", _a0, _a1)}
}

func (_c *MockServiceAccessPointContract_SapRequest_Call) Run(run func(_a0 Args, _a1 KWArgs)) *MockServiceAccessPointContract_SapRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Args), args[1].(KWArgs))
	})
	return _c
}

func (_c *MockServiceAccessPointContract_SapRequest_Call) Return(_a0 error) *MockServiceAccessPointContract_SapRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServiceAccessPointContract_SapRequest_Call) RunAndReturn(run func(Args, KWArgs) error) *MockServiceAccessPointContract_SapRequest_Call {
	_c.Call.Return(run)
	return _c
}

// SapResponse provides a mock function with given fields: _a0, _a1
func (_m *MockServiceAccessPointContract) SapResponse(_a0 Args, _a1 KWArgs) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SapResponse")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Args, KWArgs) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockServiceAccessPointContract_SapResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SapResponse'
type MockServiceAccessPointContract_SapResponse_Call struct {
	*mock.Call
}

// SapResponse is a helper method to define mock.On call
//   - _a0 Args
//   - _a1 KWArgs
func (_e *MockServiceAccessPointContract_Expecter) SapResponse(_a0 interface{}, _a1 interface{}) *MockServiceAccessPointContract_SapResponse_Call {
	return &MockServiceAccessPointContract_SapResponse_Call{Call: _e.mock.On("SapResponse", _a0, _a1)}
}

func (_c *MockServiceAccessPointContract_SapResponse_Call) Run(run func(_a0 Args, _a1 KWArgs)) *MockServiceAccessPointContract_SapResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Args), args[1].(KWArgs))
	})
	return _c
}

func (_c *MockServiceAccessPointContract_SapResponse_Call) Return(_a0 error) *MockServiceAccessPointContract_SapResponse_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServiceAccessPointContract_SapResponse_Call) RunAndReturn(run func(Args, KWArgs) error) *MockServiceAccessPointContract_SapResponse_Call {
	_c.Call.Return(run)
	return _c
}

// _setServiceElement provides a mock function with given fields: serviceElement
func (_m *MockServiceAccessPointContract) _setServiceElement(serviceElement ApplicationServiceElementContract) {
	_m.Called(serviceElement)
}

// MockServiceAccessPointContract__setServiceElement_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method '_setServiceElement'
type MockServiceAccessPointContract__setServiceElement_Call struct {
	*mock.Call
}

// _setServiceElement is a helper method to define mock.On call
//   - serviceElement ApplicationServiceElementContract
func (_e *MockServiceAccessPointContract_Expecter) _setServiceElement(serviceElement interface{}) *MockServiceAccessPointContract__setServiceElement_Call {
	return &MockServiceAccessPointContract__setServiceElement_Call{Call: _e.mock.On("_setServiceElement", serviceElement)}
}

func (_c *MockServiceAccessPointContract__setServiceElement_Call) Run(run func(serviceElement ApplicationServiceElementContract)) *MockServiceAccessPointContract__setServiceElement_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(ApplicationServiceElementContract))
	})
	return _c
}

func (_c *MockServiceAccessPointContract__setServiceElement_Call) Return() *MockServiceAccessPointContract__setServiceElement_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockServiceAccessPointContract__setServiceElement_Call) RunAndReturn(run func(ApplicationServiceElementContract)) *MockServiceAccessPointContract__setServiceElement_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockServiceAccessPointContract creates a new instance of MockServiceAccessPointContract. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockServiceAccessPointContract(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockServiceAccessPointContract {
	mock := &MockServiceAccessPointContract{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}