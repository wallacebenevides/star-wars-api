// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	db "github.com/wallacebenevides/star-wars-api/db"

	mongo "go.mongodb.org/mongo-driver/mongo"
)

// CollectionHelper is an autogenerated mock type for the CollectionHelper type
type CollectionHelper struct {
	mock.Mock
}

// DeleteOne provides a mock function with given fields: ctx, filter
func (_m *CollectionHelper) DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	ret := _m.Called(ctx, filter)

	var r0 *mongo.DeleteResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) *mongo.DeleteResult); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.DeleteResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx, filter
func (_m *CollectionHelper) Find(ctx context.Context, filter interface{}) (db.CursorHelper, error) {
	ret := _m.Called(ctx, filter)

	var r0 db.CursorHelper
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) db.CursorHelper); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.CursorHelper)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: _a0, _a1
func (_m *CollectionHelper) FindOne(_a0 context.Context, _a1 interface{}) db.SingleResultHelper {
	ret := _m.Called(_a0, _a1)

	var r0 db.SingleResultHelper
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) db.SingleResultHelper); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.SingleResultHelper)
		}
	}

	return r0
}

// InsertOne provides a mock function with given fields: _a0, _a1
func (_m *CollectionHelper) InsertOne(_a0 context.Context, _a1 interface{}) (interface{}, error) {
	ret := _m.Called(_a0, _a1)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) interface{}); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
