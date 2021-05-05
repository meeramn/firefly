// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package enginemocks

import (
	context "context"

	fftypes "github.com/kaleido-io/firefly/internal/fftypes"

	mock "github.com/stretchr/testify/mock"
)

// Engine is an autogenerated mock type for the Engine type
type Engine struct {
	mock.Mock
}

// BroadcastSchemaDefinition provides a mock function with given fields: ctx, s
func (_m *Engine) BroadcastSchemaDefinition(ctx context.Context, s *fftypes.Schema) (*fftypes.Message, error) {
	ret := _m.Called(ctx, s)

	var r0 *fftypes.Message
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Schema) *fftypes.Message); ok {
		r0 = rf(ctx, s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.Schema) error); ok {
		r1 = rf(ctx, s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *Engine) Close() {
	_m.Called()
}

// GetBatchById provides a mock function with given fields: ctx, ns, id
func (_m *Engine) GetBatchById(ctx context.Context, ns string, id string) (*fftypes.Batch, error) {
	ret := _m.Called(ctx, ns, id)

	var r0 *fftypes.Batch
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *fftypes.Batch); ok {
		r0 = rf(ctx, ns, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Batch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, ns, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDataById provides a mock function with given fields: ctx, ns, id
func (_m *Engine) GetDataById(ctx context.Context, ns string, id string) (*fftypes.Data, error) {
	ret := _m.Called(ctx, ns, id)

	var r0 *fftypes.Data
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *fftypes.Data); ok {
		r0 = rf(ctx, ns, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Data)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, ns, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMessageById provides a mock function with given fields: ctx, ns, id
func (_m *Engine) GetMessageById(ctx context.Context, ns string, id string) (*fftypes.Message, error) {
	ret := _m.Called(ctx, ns, id)

	var r0 *fftypes.Message
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *fftypes.Message); ok {
		r0 = rf(ctx, ns, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, ns, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionById provides a mock function with given fields: ctx, ns, id
func (_m *Engine) GetTransactionById(ctx context.Context, ns string, id string) (*fftypes.Transaction, error) {
	ret := _m.Called(ctx, ns, id)

	var r0 *fftypes.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *fftypes.Transaction); ok {
		r0 = rf(ctx, ns, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, ns, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields: ctx
func (_m *Engine) Init(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
