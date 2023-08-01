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

// Code generated by mockery v2.32.2. DO NOT EDIT.

package plc4go

import mock "github.com/stretchr/testify/mock"

// MockPlcConnectionCloseResult is an autogenerated mock type for the PlcConnectionCloseResult type
type MockPlcConnectionCloseResult struct {
	mock.Mock
}

type MockPlcConnectionCloseResult_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcConnectionCloseResult) EXPECT() *MockPlcConnectionCloseResult_Expecter {
	return &MockPlcConnectionCloseResult_Expecter{mock: &_m.Mock}
}

// GetConnection provides a mock function with given fields:
func (_m *MockPlcConnectionCloseResult) GetConnection() PlcConnection {
	ret := _m.Called()

	var r0 PlcConnection
	if rf, ok := ret.Get(0).(func() PlcConnection); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcConnection)
		}
	}

	return r0
}

// MockPlcConnectionCloseResult_GetConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnection'
type MockPlcConnectionCloseResult_GetConnection_Call struct {
	*mock.Call
}

// GetConnection is a helper method to define mock.On call
func (_e *MockPlcConnectionCloseResult_Expecter) GetConnection() *MockPlcConnectionCloseResult_GetConnection_Call {
	return &MockPlcConnectionCloseResult_GetConnection_Call{Call: _e.mock.On("GetConnection")}
}

func (_c *MockPlcConnectionCloseResult_GetConnection_Call) Run(run func()) *MockPlcConnectionCloseResult_GetConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnectionCloseResult_GetConnection_Call) Return(_a0 PlcConnection) *MockPlcConnectionCloseResult_GetConnection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnectionCloseResult_GetConnection_Call) RunAndReturn(run func() PlcConnection) *MockPlcConnectionCloseResult_GetConnection_Call {
	_c.Call.Return(run)
	return _c
}

// GetErr provides a mock function with given fields:
func (_m *MockPlcConnectionCloseResult) GetErr() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcConnectionCloseResult_GetErr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetErr'
type MockPlcConnectionCloseResult_GetErr_Call struct {
	*mock.Call
}

// GetErr is a helper method to define mock.On call
func (_e *MockPlcConnectionCloseResult_Expecter) GetErr() *MockPlcConnectionCloseResult_GetErr_Call {
	return &MockPlcConnectionCloseResult_GetErr_Call{Call: _e.mock.On("GetErr")}
}

func (_c *MockPlcConnectionCloseResult_GetErr_Call) Run(run func()) *MockPlcConnectionCloseResult_GetErr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnectionCloseResult_GetErr_Call) Return(_a0 error) *MockPlcConnectionCloseResult_GetErr_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnectionCloseResult_GetErr_Call) RunAndReturn(run func() error) *MockPlcConnectionCloseResult_GetErr_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockPlcConnectionCloseResult) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcConnectionCloseResult_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockPlcConnectionCloseResult_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockPlcConnectionCloseResult_Expecter) String() *MockPlcConnectionCloseResult_String_Call {
	return &MockPlcConnectionCloseResult_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockPlcConnectionCloseResult_String_Call) Run(run func()) *MockPlcConnectionCloseResult_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnectionCloseResult_String_Call) Return(_a0 string) *MockPlcConnectionCloseResult_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnectionCloseResult_String_Call) RunAndReturn(run func() string) *MockPlcConnectionCloseResult_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcConnectionCloseResult creates a new instance of MockPlcConnectionCloseResult. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcConnectionCloseResult(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcConnectionCloseResult {
	mock := &MockPlcConnectionCloseResult{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
