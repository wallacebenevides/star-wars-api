// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import db "github.com/wallacebenevides/star-wars-api/db"
import mock "github.com/stretchr/testify/mock"

// DatabaseHelper is an autogenerated mock type for the DatabaseHelper type
type DatabaseHelper struct {
	mock.Mock
}

// Client provides a mock function with given fields:
func (_m *DatabaseHelper) Client() db.ClientHelper {
	ret := _m.Called()

	var r0 db.ClientHelper
	if rf, ok := ret.Get(0).(func() db.ClientHelper); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.ClientHelper)
		}
	}

	return r0
}

// Collection provides a mock function with given fields: name
func (_m *DatabaseHelper) Collection(name string) db.CollectionHelper {
	ret := _m.Called(name)

	var r0 db.CollectionHelper
	if rf, ok := ret.Get(0).(func(string) db.CollectionHelper); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.CollectionHelper)
		}
	}

	return r0
}
