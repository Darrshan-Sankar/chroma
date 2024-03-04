// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import (
	context "context"

	dbmodel "github.com/chroma-core/chroma/go/internal/metastore/db/dbmodel"
	mock "github.com/stretchr/testify/mock"
)

// IMetaDomain is an autogenerated mock type for the IMetaDomain type
type IMetaDomain struct {
	mock.Mock
}

// CollectionDb provides a mock function with given fields: ctx
func (_m *IMetaDomain) CollectionDb(ctx context.Context) dbmodel.ICollectionDb {
	ret := _m.Called(ctx)

	var r0 dbmodel.ICollectionDb
	if rf, ok := ret.Get(0).(func(context.Context) dbmodel.ICollectionDb); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dbmodel.ICollectionDb)
		}
	}

	return r0
}

// CollectionMetadataDb provides a mock function with given fields: ctx
func (_m *IMetaDomain) CollectionMetadataDb(ctx context.Context) dbmodel.ICollectionMetadataDb {
	ret := _m.Called(ctx)

	var r0 dbmodel.ICollectionMetadataDb
	if rf, ok := ret.Get(0).(func(context.Context) dbmodel.ICollectionMetadataDb); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dbmodel.ICollectionMetadataDb)
		}
	}

	return r0
}

// DatabaseDb provides a mock function with given fields: ctx
func (_m *IMetaDomain) DatabaseDb(ctx context.Context) dbmodel.IDatabaseDb {
	ret := _m.Called(ctx)

	var r0 dbmodel.IDatabaseDb
	if rf, ok := ret.Get(0).(func(context.Context) dbmodel.IDatabaseDb); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dbmodel.IDatabaseDb)
		}
	}

	return r0
}

// NotificationDb provides a mock function with given fields: ctx
func (_m *IMetaDomain) NotificationDb(ctx context.Context) dbmodel.INotificationDb {
	ret := _m.Called(ctx)

	var r0 dbmodel.INotificationDb
	if rf, ok := ret.Get(0).(func(context.Context) dbmodel.INotificationDb); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dbmodel.INotificationDb)
		}
	}

	return r0
}

// SegmentDb provides a mock function with given fields: ctx
func (_m *IMetaDomain) SegmentDb(ctx context.Context) dbmodel.ISegmentDb {
	ret := _m.Called(ctx)

	var r0 dbmodel.ISegmentDb
	if rf, ok := ret.Get(0).(func(context.Context) dbmodel.ISegmentDb); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dbmodel.ISegmentDb)
		}
	}

	return r0
}

// SegmentMetadataDb provides a mock function with given fields: ctx
func (_m *IMetaDomain) SegmentMetadataDb(ctx context.Context) dbmodel.ISegmentMetadataDb {
	ret := _m.Called(ctx)

	var r0 dbmodel.ISegmentMetadataDb
	if rf, ok := ret.Get(0).(func(context.Context) dbmodel.ISegmentMetadataDb); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dbmodel.ISegmentMetadataDb)
		}
	}

	return r0
}

// TenantDb provides a mock function with given fields: ctx
func (_m *IMetaDomain) TenantDb(ctx context.Context) dbmodel.ITenantDb {
	ret := _m.Called(ctx)

	var r0 dbmodel.ITenantDb
	if rf, ok := ret.Get(0).(func(context.Context) dbmodel.ITenantDb); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dbmodel.ITenantDb)
		}
	}

	return r0
}

func (_m *IMetaDomain) RecordLogDb(ctx context.Context) dbmodel.IRecordLogDb {
	ret := _m.Called(ctx)

	var r0 dbmodel.IRecordLogDb
	if rf, ok := ret.Get(0).(func(context.Context) dbmodel.IRecordLogDb); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dbmodel.IRecordLogDb)
		}
	}

	return r0
}

// NewIMetaDomain creates a new instance of IMetaDomain. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIMetaDomain(t interface {
	mock.TestingT
	Cleanup(func())
}) *IMetaDomain {
	mock := &IMetaDomain{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
