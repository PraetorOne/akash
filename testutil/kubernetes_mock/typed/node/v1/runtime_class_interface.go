// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	nodev1 "k8s.io/api/node/v1"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/client-go/applyconfigurations/node/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// RuntimeClassInterface is an autogenerated mock type for the RuntimeClassInterface type
type RuntimeClassInterface struct {
	mock.Mock
}

// Apply provides a mock function with given fields: ctx, runtimeClass, opts
func (_m *RuntimeClassInterface) Apply(ctx context.Context, runtimeClass *v1.RuntimeClassApplyConfiguration, opts metav1.ApplyOptions) (*nodev1.RuntimeClass, error) {
	ret := _m.Called(ctx, runtimeClass, opts)

	var r0 *nodev1.RuntimeClass
	if rf, ok := ret.Get(0).(func(context.Context, *v1.RuntimeClassApplyConfiguration, metav1.ApplyOptions) *nodev1.RuntimeClass); ok {
		r0 = rf(ctx, runtimeClass, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*nodev1.RuntimeClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.RuntimeClassApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, runtimeClass, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, runtimeClass, opts
func (_m *RuntimeClassInterface) Create(ctx context.Context, runtimeClass *nodev1.RuntimeClass, opts metav1.CreateOptions) (*nodev1.RuntimeClass, error) {
	ret := _m.Called(ctx, runtimeClass, opts)

	var r0 *nodev1.RuntimeClass
	if rf, ok := ret.Get(0).(func(context.Context, *nodev1.RuntimeClass, metav1.CreateOptions) *nodev1.RuntimeClass); ok {
		r0 = rf(ctx, runtimeClass, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*nodev1.RuntimeClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *nodev1.RuntimeClass, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, runtimeClass, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *RuntimeClassInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *RuntimeClassInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *RuntimeClassInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*nodev1.RuntimeClass, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *nodev1.RuntimeClass
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *nodev1.RuntimeClass); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*nodev1.RuntimeClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, opts
func (_m *RuntimeClassInterface) List(ctx context.Context, opts metav1.ListOptions) (*nodev1.RuntimeClassList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *nodev1.RuntimeClassList
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *nodev1.RuntimeClassList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*nodev1.RuntimeClassList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *RuntimeClassInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*nodev1.RuntimeClass, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *nodev1.RuntimeClass
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *nodev1.RuntimeClass); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*nodev1.RuntimeClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, runtimeClass, opts
func (_m *RuntimeClassInterface) Update(ctx context.Context, runtimeClass *nodev1.RuntimeClass, opts metav1.UpdateOptions) (*nodev1.RuntimeClass, error) {
	ret := _m.Called(ctx, runtimeClass, opts)

	var r0 *nodev1.RuntimeClass
	if rf, ok := ret.Get(0).(func(context.Context, *nodev1.RuntimeClass, metav1.UpdateOptions) *nodev1.RuntimeClass); ok {
		r0 = rf(ctx, runtimeClass, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*nodev1.RuntimeClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *nodev1.RuntimeClass, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, runtimeClass, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *RuntimeClassInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, opts)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
