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

package utils

import mock "github.com/stretchr/testify/mock"

// MockAsciiBoxWriter is an autogenerated mock type for the AsciiBoxWriter type
type MockAsciiBoxWriter struct {
	mock.Mock
}

type MockAsciiBoxWriter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAsciiBoxWriter) EXPECT() *MockAsciiBoxWriter_Expecter {
	return &MockAsciiBoxWriter_Expecter{mock: &_m.Mock}
}

// AlignBoxes provides a mock function with given fields: asciiBoxes, desiredWith
func (_m *MockAsciiBoxWriter) AlignBoxes(asciiBoxes []AsciiBox, desiredWith int) AsciiBox {
	ret := _m.Called(asciiBoxes, desiredWith)

	var r0 AsciiBox
	if rf, ok := ret.Get(0).(func([]AsciiBox, int) AsciiBox); ok {
		r0 = rf(asciiBoxes, desiredWith)
	} else {
		r0 = ret.Get(0).(AsciiBox)
	}

	return r0
}

// MockAsciiBoxWriter_AlignBoxes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AlignBoxes'
type MockAsciiBoxWriter_AlignBoxes_Call struct {
	*mock.Call
}

// AlignBoxes is a helper method to define mock.On call
//   - asciiBoxes []AsciiBox
//   - desiredWith int
func (_e *MockAsciiBoxWriter_Expecter) AlignBoxes(asciiBoxes interface{}, desiredWith interface{}) *MockAsciiBoxWriter_AlignBoxes_Call {
	return &MockAsciiBoxWriter_AlignBoxes_Call{Call: _e.mock.On("AlignBoxes", asciiBoxes, desiredWith)}
}

func (_c *MockAsciiBoxWriter_AlignBoxes_Call) Run(run func(asciiBoxes []AsciiBox, desiredWith int)) *MockAsciiBoxWriter_AlignBoxes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]AsciiBox), args[1].(int))
	})
	return _c
}

func (_c *MockAsciiBoxWriter_AlignBoxes_Call) Return(_a0 AsciiBox) *MockAsciiBoxWriter_AlignBoxes_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAsciiBoxWriter_AlignBoxes_Call) RunAndReturn(run func([]AsciiBox, int) AsciiBox) *MockAsciiBoxWriter_AlignBoxes_Call {
	_c.Call.Return(run)
	return _c
}

// BoxBelowBox provides a mock function with given fields: box1, box2
func (_m *MockAsciiBoxWriter) BoxBelowBox(box1 AsciiBox, box2 AsciiBox) AsciiBox {
	ret := _m.Called(box1, box2)

	var r0 AsciiBox
	if rf, ok := ret.Get(0).(func(AsciiBox, AsciiBox) AsciiBox); ok {
		r0 = rf(box1, box2)
	} else {
		r0 = ret.Get(0).(AsciiBox)
	}

	return r0
}

// MockAsciiBoxWriter_BoxBelowBox_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BoxBelowBox'
type MockAsciiBoxWriter_BoxBelowBox_Call struct {
	*mock.Call
}

// BoxBelowBox is a helper method to define mock.On call
//   - box1 AsciiBox
//   - box2 AsciiBox
func (_e *MockAsciiBoxWriter_Expecter) BoxBelowBox(box1 interface{}, box2 interface{}) *MockAsciiBoxWriter_BoxBelowBox_Call {
	return &MockAsciiBoxWriter_BoxBelowBox_Call{Call: _e.mock.On("BoxBelowBox", box1, box2)}
}

func (_c *MockAsciiBoxWriter_BoxBelowBox_Call) Run(run func(box1 AsciiBox, box2 AsciiBox)) *MockAsciiBoxWriter_BoxBelowBox_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(AsciiBox), args[1].(AsciiBox))
	})
	return _c
}

func (_c *MockAsciiBoxWriter_BoxBelowBox_Call) Return(_a0 AsciiBox) *MockAsciiBoxWriter_BoxBelowBox_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAsciiBoxWriter_BoxBelowBox_Call) RunAndReturn(run func(AsciiBox, AsciiBox) AsciiBox) *MockAsciiBoxWriter_BoxBelowBox_Call {
	_c.Call.Return(run)
	return _c
}

// BoxBox provides a mock function with given fields: name, box, charWidth
func (_m *MockAsciiBoxWriter) BoxBox(name string, box AsciiBox, charWidth int) AsciiBox {
	ret := _m.Called(name, box, charWidth)

	var r0 AsciiBox
	if rf, ok := ret.Get(0).(func(string, AsciiBox, int) AsciiBox); ok {
		r0 = rf(name, box, charWidth)
	} else {
		r0 = ret.Get(0).(AsciiBox)
	}

	return r0
}

// MockAsciiBoxWriter_BoxBox_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BoxBox'
type MockAsciiBoxWriter_BoxBox_Call struct {
	*mock.Call
}

// BoxBox is a helper method to define mock.On call
//   - name string
//   - box AsciiBox
//   - charWidth int
func (_e *MockAsciiBoxWriter_Expecter) BoxBox(name interface{}, box interface{}, charWidth interface{}) *MockAsciiBoxWriter_BoxBox_Call {
	return &MockAsciiBoxWriter_BoxBox_Call{Call: _e.mock.On("BoxBox", name, box, charWidth)}
}

func (_c *MockAsciiBoxWriter_BoxBox_Call) Run(run func(name string, box AsciiBox, charWidth int)) *MockAsciiBoxWriter_BoxBox_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(AsciiBox), args[2].(int))
	})
	return _c
}

func (_c *MockAsciiBoxWriter_BoxBox_Call) Return(_a0 AsciiBox) *MockAsciiBoxWriter_BoxBox_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAsciiBoxWriter_BoxBox_Call) RunAndReturn(run func(string, AsciiBox, int) AsciiBox) *MockAsciiBoxWriter_BoxBox_Call {
	_c.Call.Return(run)
	return _c
}

// BoxSideBySide provides a mock function with given fields: box1, box2
func (_m *MockAsciiBoxWriter) BoxSideBySide(box1 AsciiBox, box2 AsciiBox) AsciiBox {
	ret := _m.Called(box1, box2)

	var r0 AsciiBox
	if rf, ok := ret.Get(0).(func(AsciiBox, AsciiBox) AsciiBox); ok {
		r0 = rf(box1, box2)
	} else {
		r0 = ret.Get(0).(AsciiBox)
	}

	return r0
}

// MockAsciiBoxWriter_BoxSideBySide_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BoxSideBySide'
type MockAsciiBoxWriter_BoxSideBySide_Call struct {
	*mock.Call
}

// BoxSideBySide is a helper method to define mock.On call
//   - box1 AsciiBox
//   - box2 AsciiBox
func (_e *MockAsciiBoxWriter_Expecter) BoxSideBySide(box1 interface{}, box2 interface{}) *MockAsciiBoxWriter_BoxSideBySide_Call {
	return &MockAsciiBoxWriter_BoxSideBySide_Call{Call: _e.mock.On("BoxSideBySide", box1, box2)}
}

func (_c *MockAsciiBoxWriter_BoxSideBySide_Call) Run(run func(box1 AsciiBox, box2 AsciiBox)) *MockAsciiBoxWriter_BoxSideBySide_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(AsciiBox), args[1].(AsciiBox))
	})
	return _c
}

func (_c *MockAsciiBoxWriter_BoxSideBySide_Call) Return(_a0 AsciiBox) *MockAsciiBoxWriter_BoxSideBySide_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAsciiBoxWriter_BoxSideBySide_Call) RunAndReturn(run func(AsciiBox, AsciiBox) AsciiBox) *MockAsciiBoxWriter_BoxSideBySide_Call {
	_c.Call.Return(run)
	return _c
}

// BoxString provides a mock function with given fields: name, data, charWidth
func (_m *MockAsciiBoxWriter) BoxString(name string, data string, charWidth int) AsciiBox {
	ret := _m.Called(name, data, charWidth)

	var r0 AsciiBox
	if rf, ok := ret.Get(0).(func(string, string, int) AsciiBox); ok {
		r0 = rf(name, data, charWidth)
	} else {
		r0 = ret.Get(0).(AsciiBox)
	}

	return r0
}

// MockAsciiBoxWriter_BoxString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BoxString'
type MockAsciiBoxWriter_BoxString_Call struct {
	*mock.Call
}

// BoxString is a helper method to define mock.On call
//   - name string
//   - data string
//   - charWidth int
func (_e *MockAsciiBoxWriter_Expecter) BoxString(name interface{}, data interface{}, charWidth interface{}) *MockAsciiBoxWriter_BoxString_Call {
	return &MockAsciiBoxWriter_BoxString_Call{Call: _e.mock.On("BoxString", name, data, charWidth)}
}

func (_c *MockAsciiBoxWriter_BoxString_Call) Run(run func(name string, data string, charWidth int)) *MockAsciiBoxWriter_BoxString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(int))
	})
	return _c
}

func (_c *MockAsciiBoxWriter_BoxString_Call) Return(_a0 AsciiBox) *MockAsciiBoxWriter_BoxString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAsciiBoxWriter_BoxString_Call) RunAndReturn(run func(string, string, int) AsciiBox) *MockAsciiBoxWriter_BoxString_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAsciiBoxWriter creates a new instance of MockAsciiBoxWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAsciiBoxWriter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAsciiBoxWriter {
	mock := &MockAsciiBoxWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
