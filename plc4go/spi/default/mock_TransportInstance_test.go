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

package _default

import (
	context "context"

	transports "github.com/apache/plc4x/plc4go/spi/transports"
	mock "github.com/stretchr/testify/mock"
)

// MockTransportInstance is an autogenerated mock type for the TransportInstance type
type MockTransportInstance struct {
	mock.Mock
}

type MockTransportInstance_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTransportInstance) EXPECT() *MockTransportInstance_Expecter {
	return &MockTransportInstance_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockTransportInstance) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTransportInstance_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockTransportInstance_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockTransportInstance_Expecter) Close() *MockTransportInstance_Close_Call {
	return &MockTransportInstance_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockTransportInstance_Close_Call) Run(run func()) *MockTransportInstance_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTransportInstance_Close_Call) Return(_a0 error) *MockTransportInstance_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTransportInstance_Close_Call) RunAndReturn(run func() error) *MockTransportInstance_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Connect provides a mock function with given fields:
func (_m *MockTransportInstance) Connect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTransportInstance_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type MockTransportInstance_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
func (_e *MockTransportInstance_Expecter) Connect() *MockTransportInstance_Connect_Call {
	return &MockTransportInstance_Connect_Call{Call: _e.mock.On("Connect")}
}

func (_c *MockTransportInstance_Connect_Call) Run(run func()) *MockTransportInstance_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTransportInstance_Connect_Call) Return(_a0 error) *MockTransportInstance_Connect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTransportInstance_Connect_Call) RunAndReturn(run func() error) *MockTransportInstance_Connect_Call {
	_c.Call.Return(run)
	return _c
}

// ConnectWithContext provides a mock function with given fields: ctx
func (_m *MockTransportInstance) ConnectWithContext(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTransportInstance_ConnectWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConnectWithContext'
type MockTransportInstance_ConnectWithContext_Call struct {
	*mock.Call
}

// ConnectWithContext is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockTransportInstance_Expecter) ConnectWithContext(ctx interface{}) *MockTransportInstance_ConnectWithContext_Call {
	return &MockTransportInstance_ConnectWithContext_Call{Call: _e.mock.On("ConnectWithContext", ctx)}
}

func (_c *MockTransportInstance_ConnectWithContext_Call) Run(run func(ctx context.Context)) *MockTransportInstance_ConnectWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockTransportInstance_ConnectWithContext_Call) Return(_a0 error) *MockTransportInstance_ConnectWithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTransportInstance_ConnectWithContext_Call) RunAndReturn(run func(context.Context) error) *MockTransportInstance_ConnectWithContext_Call {
	_c.Call.Return(run)
	return _c
}

// FillBuffer provides a mock function with given fields: until
func (_m *MockTransportInstance) FillBuffer(until func(uint, byte, transports.ExtendedReader) bool) error {
	ret := _m.Called(until)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(uint, byte, transports.ExtendedReader) bool) error); ok {
		r0 = rf(until)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTransportInstance_FillBuffer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FillBuffer'
type MockTransportInstance_FillBuffer_Call struct {
	*mock.Call
}

// FillBuffer is a helper method to define mock.On call
//   - until func(uint , byte , transports.ExtendedReader) bool
func (_e *MockTransportInstance_Expecter) FillBuffer(until interface{}) *MockTransportInstance_FillBuffer_Call {
	return &MockTransportInstance_FillBuffer_Call{Call: _e.mock.On("FillBuffer", until)}
}

func (_c *MockTransportInstance_FillBuffer_Call) Run(run func(until func(uint, byte, transports.ExtendedReader) bool)) *MockTransportInstance_FillBuffer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(uint, byte, transports.ExtendedReader) bool))
	})
	return _c
}

func (_c *MockTransportInstance_FillBuffer_Call) Return(_a0 error) *MockTransportInstance_FillBuffer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTransportInstance_FillBuffer_Call) RunAndReturn(run func(func(uint, byte, transports.ExtendedReader) bool) error) *MockTransportInstance_FillBuffer_Call {
	_c.Call.Return(run)
	return _c
}

// GetNumBytesAvailableInBuffer provides a mock function with given fields:
func (_m *MockTransportInstance) GetNumBytesAvailableInBuffer() (uint32, error) {
	ret := _m.Called()

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint32, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTransportInstance_GetNumBytesAvailableInBuffer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNumBytesAvailableInBuffer'
type MockTransportInstance_GetNumBytesAvailableInBuffer_Call struct {
	*mock.Call
}

// GetNumBytesAvailableInBuffer is a helper method to define mock.On call
func (_e *MockTransportInstance_Expecter) GetNumBytesAvailableInBuffer() *MockTransportInstance_GetNumBytesAvailableInBuffer_Call {
	return &MockTransportInstance_GetNumBytesAvailableInBuffer_Call{Call: _e.mock.On("GetNumBytesAvailableInBuffer")}
}

func (_c *MockTransportInstance_GetNumBytesAvailableInBuffer_Call) Run(run func()) *MockTransportInstance_GetNumBytesAvailableInBuffer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTransportInstance_GetNumBytesAvailableInBuffer_Call) Return(_a0 uint32, _a1 error) *MockTransportInstance_GetNumBytesAvailableInBuffer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTransportInstance_GetNumBytesAvailableInBuffer_Call) RunAndReturn(run func() (uint32, error)) *MockTransportInstance_GetNumBytesAvailableInBuffer_Call {
	_c.Call.Return(run)
	return _c
}

// IsConnected provides a mock function with given fields:
func (_m *MockTransportInstance) IsConnected() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockTransportInstance_IsConnected_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsConnected'
type MockTransportInstance_IsConnected_Call struct {
	*mock.Call
}

// IsConnected is a helper method to define mock.On call
func (_e *MockTransportInstance_Expecter) IsConnected() *MockTransportInstance_IsConnected_Call {
	return &MockTransportInstance_IsConnected_Call{Call: _e.mock.On("IsConnected")}
}

func (_c *MockTransportInstance_IsConnected_Call) Run(run func()) *MockTransportInstance_IsConnected_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTransportInstance_IsConnected_Call) Return(_a0 bool) *MockTransportInstance_IsConnected_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTransportInstance_IsConnected_Call) RunAndReturn(run func() bool) *MockTransportInstance_IsConnected_Call {
	_c.Call.Return(run)
	return _c
}

// PeekReadableBytes provides a mock function with given fields: numBytes
func (_m *MockTransportInstance) PeekReadableBytes(numBytes uint32) ([]byte, error) {
	ret := _m.Called(numBytes)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(uint32) ([]byte, error)); ok {
		return rf(numBytes)
	}
	if rf, ok := ret.Get(0).(func(uint32) []byte); ok {
		r0 = rf(numBytes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(numBytes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTransportInstance_PeekReadableBytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PeekReadableBytes'
type MockTransportInstance_PeekReadableBytes_Call struct {
	*mock.Call
}

// PeekReadableBytes is a helper method to define mock.On call
//   - numBytes uint32
func (_e *MockTransportInstance_Expecter) PeekReadableBytes(numBytes interface{}) *MockTransportInstance_PeekReadableBytes_Call {
	return &MockTransportInstance_PeekReadableBytes_Call{Call: _e.mock.On("PeekReadableBytes", numBytes)}
}

func (_c *MockTransportInstance_PeekReadableBytes_Call) Run(run func(numBytes uint32)) *MockTransportInstance_PeekReadableBytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint32))
	})
	return _c
}

func (_c *MockTransportInstance_PeekReadableBytes_Call) Return(_a0 []byte, _a1 error) *MockTransportInstance_PeekReadableBytes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTransportInstance_PeekReadableBytes_Call) RunAndReturn(run func(uint32) ([]byte, error)) *MockTransportInstance_PeekReadableBytes_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: numBytes
func (_m *MockTransportInstance) Read(numBytes uint32) ([]byte, error) {
	ret := _m.Called(numBytes)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(uint32) ([]byte, error)); ok {
		return rf(numBytes)
	}
	if rf, ok := ret.Get(0).(func(uint32) []byte); ok {
		r0 = rf(numBytes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(numBytes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTransportInstance_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type MockTransportInstance_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - numBytes uint32
func (_e *MockTransportInstance_Expecter) Read(numBytes interface{}) *MockTransportInstance_Read_Call {
	return &MockTransportInstance_Read_Call{Call: _e.mock.On("Read", numBytes)}
}

func (_c *MockTransportInstance_Read_Call) Run(run func(numBytes uint32)) *MockTransportInstance_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint32))
	})
	return _c
}

func (_c *MockTransportInstance_Read_Call) Return(_a0 []byte, _a1 error) *MockTransportInstance_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTransportInstance_Read_Call) RunAndReturn(run func(uint32) ([]byte, error)) *MockTransportInstance_Read_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockTransportInstance) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockTransportInstance_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockTransportInstance_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockTransportInstance_Expecter) String() *MockTransportInstance_String_Call {
	return &MockTransportInstance_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockTransportInstance_String_Call) Run(run func()) *MockTransportInstance_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTransportInstance_String_Call) Return(_a0 string) *MockTransportInstance_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTransportInstance_String_Call) RunAndReturn(run func() string) *MockTransportInstance_String_Call {
	_c.Call.Return(run)
	return _c
}

// Write provides a mock function with given fields: data
func (_m *MockTransportInstance) Write(data []byte) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTransportInstance_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type MockTransportInstance_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - data []byte
func (_e *MockTransportInstance_Expecter) Write(data interface{}) *MockTransportInstance_Write_Call {
	return &MockTransportInstance_Write_Call{Call: _e.mock.On("Write", data)}
}

func (_c *MockTransportInstance_Write_Call) Run(run func(data []byte)) *MockTransportInstance_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MockTransportInstance_Write_Call) Return(_a0 error) *MockTransportInstance_Write_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTransportInstance_Write_Call) RunAndReturn(run func([]byte) error) *MockTransportInstance_Write_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTransportInstance creates a new instance of MockTransportInstance. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTransportInstance(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTransportInstance {
	mock := &MockTransportInstance{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
