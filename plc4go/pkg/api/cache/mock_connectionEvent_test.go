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

package cache

import mock "github.com/stretchr/testify/mock"

// mockConnectionEvent is an autogenerated mock type for the connectionEvent type
type mockConnectionEvent struct {
	mock.Mock
}

type mockConnectionEvent_Expecter struct {
	mock *mock.Mock
}

func (_m *mockConnectionEvent) EXPECT() *mockConnectionEvent_Expecter {
	return &mockConnectionEvent_Expecter{mock: &_m.Mock}
}

// getConnectionContainer provides a mock function with given fields:
func (_m *mockConnectionEvent) getConnectionContainer() connectionContainer {
	ret := _m.Called()

	var r0 connectionContainer
	if rf, ok := ret.Get(0).(func() connectionContainer); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(connectionContainer)
	}

	return r0
}

// mockConnectionEvent_getConnectionContainer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getConnectionContainer'
type mockConnectionEvent_getConnectionContainer_Call struct {
	*mock.Call
}

// getConnectionContainer is a helper method to define mock.On call
func (_e *mockConnectionEvent_Expecter) getConnectionContainer() *mockConnectionEvent_getConnectionContainer_Call {
	return &mockConnectionEvent_getConnectionContainer_Call{Call: _e.mock.On("getConnectionContainer")}
}

func (_c *mockConnectionEvent_getConnectionContainer_Call) Run(run func()) *mockConnectionEvent_getConnectionContainer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mockConnectionEvent_getConnectionContainer_Call) Return(_a0 connectionContainer) *mockConnectionEvent_getConnectionContainer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockConnectionEvent_getConnectionContainer_Call) RunAndReturn(run func() connectionContainer) *mockConnectionEvent_getConnectionContainer_Call {
	_c.Call.Return(run)
	return _c
}

// newMockConnectionEvent creates a new instance of mockConnectionEvent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockConnectionEvent(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockConnectionEvent {
	mock := &mockConnectionEvent{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
