// Code generated by mockery v1.0.0

package mocks

import (
	context "context"

	"github.com/stretchr/testify/mock"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// SecretsClient is an autogenerated mock type for the SecretsClient type
type SecretsClient struct {
	mock.Mock
}

// CreateSecret provides a mock function with given fields: _a0, _a1
func (_m *SecretsClient) CreateSecret(_a0 context.Context, _a1 *v1.Secret) (*v1.Secret, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.Secret
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Secret) *v1.Secret); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Secret)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.Secret) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSecret provides a mock function with given fields: ctx, secretName
func (_m *SecretsClient) DeleteSecret(ctx context.Context, secretName string) error {
	ret := _m.Called(ctx, secretName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, secretName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSecret provides a mock function with given fields: ctx, secretName
func (_m *SecretsClient) GetSecret(ctx context.Context, secretName string) (*v1.Secret, error) {
	ret := _m.Called(ctx, secretName)

	var r0 *v1.Secret
	if rf, ok := ret.Get(0).(func(context.Context, string) *v1.Secret); ok {
		r0 = rf(ctx, secretName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Secret)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, secretName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSecrets provides a mock function with given fields: _a0
func (_m *SecretsClient) ListSecrets(_a0 context.Context) ([]v1.Secret, error) {
	ret := _m.Called(_a0)

	var r0 []v1.Secret
	if rf, ok := ret.Get(0).(func(context.Context) []v1.Secret); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1.Secret)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSecret provides a mock function with given fields: _a0, _a1
func (_m *SecretsClient) UpdateSecret(_a0 context.Context, _a1 *v1.Secret) (*v1.Secret, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.Secret
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Secret) *v1.Secret); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Secret)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.Secret) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
