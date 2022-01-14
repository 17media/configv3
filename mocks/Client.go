// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	configv3 "github.com/17media/configv3"
	mock "github.com/stretchr/testify/mock"

	regexp "regexp"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// AddListener provides a mock function with given fields: pathRegEx
func (_m *Client) AddListener(pathRegEx *regexp.Regexp) *chan configv3.ModifiedFile {
	ret := _m.Called(pathRegEx)

	var r0 *chan configv3.ModifiedFile
	if rf, ok := ret.Get(0).(func(*regexp.Regexp) *chan configv3.ModifiedFile); ok {
		r0 = rf(pathRegEx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chan configv3.ModifiedFile)
		}
	}

	return r0
}

// ConfigInfo provides a mock function with given fields:
func (_m *Client) ConfigInfo() configv3.ConfigInfo {
	ret := _m.Called()

	var r0 configv3.ConfigInfo
	if rf, ok := ret.Get(0).(func() configv3.ConfigInfo); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(configv3.ConfigInfo)
	}

	return r0
}

// Get provides a mock function with given fields: path
func (_m *Client) Get(path string) ([]byte, error) {
	ret := _m.Called(path)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: path
func (_m *Client) List(path string) (map[string][]byte, error) {
	ret := _m.Called(path)

	var r0 map[string][]byte
	if rf, ok := ret.Get(0).(func(string) map[string][]byte); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveListener provides a mock function with given fields: ch
func (_m *Client) RemoveListener(ch *chan configv3.ModifiedFile) {
	_m.Called(ch)
}

// Stop provides a mock function with given fields:
func (_m *Client) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Watch provides a mock function with given fields: path, callback, errChan
func (_m *Client) Watch(path string, callback func([]byte) error, errChan chan<- error) error {
	ret := _m.Called(path, callback, errChan)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, func([]byte) error, chan<- error) error); ok {
		r0 = rf(path, callback, errChan)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
