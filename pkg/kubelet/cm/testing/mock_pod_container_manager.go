/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by mockery v2.40.3. DO NOT EDIT.

package testing

import (
	mock "github.com/stretchr/testify/mock"
	cm "k8s.io/kubernetes/pkg/kubelet/cm"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/api/core/v1"
)

// MockPodContainerManager is an autogenerated mock type for the PodContainerManager type
type MockPodContainerManager struct {
	mock.Mock
}

type MockPodContainerManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPodContainerManager) EXPECT() *MockPodContainerManager_Expecter {
	return &MockPodContainerManager_Expecter{mock: &_m.Mock}
}

// Destroy provides a mock function with given fields: name
func (_m *MockPodContainerManager) Destroy(name cm.CgroupName) error {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Destroy")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(cm.CgroupName) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPodContainerManager_Destroy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Destroy'
type MockPodContainerManager_Destroy_Call struct {
	*mock.Call
}

// Destroy is a helper method to define mock.On call
//   - name cm.CgroupName
func (_e *MockPodContainerManager_Expecter) Destroy(name interface{}) *MockPodContainerManager_Destroy_Call {
	return &MockPodContainerManager_Destroy_Call{Call: _e.mock.On("Destroy", name)}
}

func (_c *MockPodContainerManager_Destroy_Call) Run(run func(name cm.CgroupName)) *MockPodContainerManager_Destroy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(cm.CgroupName))
	})
	return _c
}

func (_c *MockPodContainerManager_Destroy_Call) Return(_a0 error) *MockPodContainerManager_Destroy_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPodContainerManager_Destroy_Call) RunAndReturn(run func(cm.CgroupName) error) *MockPodContainerManager_Destroy_Call {
	_c.Call.Return(run)
	return _c
}

// EnsureExists provides a mock function with given fields: _a0
func (_m *MockPodContainerManager) EnsureExists(_a0 *v1.Pod) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for EnsureExists")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1.Pod) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPodContainerManager_EnsureExists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnsureExists'
type MockPodContainerManager_EnsureExists_Call struct {
	*mock.Call
}

// EnsureExists is a helper method to define mock.On call
//   - _a0 *v1.Pod
func (_e *MockPodContainerManager_Expecter) EnsureExists(_a0 interface{}) *MockPodContainerManager_EnsureExists_Call {
	return &MockPodContainerManager_EnsureExists_Call{Call: _e.mock.On("EnsureExists", _a0)}
}

func (_c *MockPodContainerManager_EnsureExists_Call) Run(run func(_a0 *v1.Pod)) *MockPodContainerManager_EnsureExists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*v1.Pod))
	})
	return _c
}

func (_c *MockPodContainerManager_EnsureExists_Call) Return(_a0 error) *MockPodContainerManager_EnsureExists_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPodContainerManager_EnsureExists_Call) RunAndReturn(run func(*v1.Pod) error) *MockPodContainerManager_EnsureExists_Call {
	_c.Call.Return(run)
	return _c
}

// Exists provides a mock function with given fields: _a0
func (_m *MockPodContainerManager) Exists(_a0 *v1.Pod) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Exists")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(*v1.Pod) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPodContainerManager_Exists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exists'
type MockPodContainerManager_Exists_Call struct {
	*mock.Call
}

// Exists is a helper method to define mock.On call
//   - _a0 *v1.Pod
func (_e *MockPodContainerManager_Expecter) Exists(_a0 interface{}) *MockPodContainerManager_Exists_Call {
	return &MockPodContainerManager_Exists_Call{Call: _e.mock.On("Exists", _a0)}
}

func (_c *MockPodContainerManager_Exists_Call) Run(run func(_a0 *v1.Pod)) *MockPodContainerManager_Exists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*v1.Pod))
	})
	return _c
}

func (_c *MockPodContainerManager_Exists_Call) Return(_a0 bool) *MockPodContainerManager_Exists_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPodContainerManager_Exists_Call) RunAndReturn(run func(*v1.Pod) bool) *MockPodContainerManager_Exists_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllPodsFromCgroups provides a mock function with given fields:
func (_m *MockPodContainerManager) GetAllPodsFromCgroups() (map[types.UID]cm.CgroupName, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllPodsFromCgroups")
	}

	var r0 map[types.UID]cm.CgroupName
	var r1 error
	if rf, ok := ret.Get(0).(func() (map[types.UID]cm.CgroupName, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() map[types.UID]cm.CgroupName); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[types.UID]cm.CgroupName)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPodContainerManager_GetAllPodsFromCgroups_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllPodsFromCgroups'
type MockPodContainerManager_GetAllPodsFromCgroups_Call struct {
	*mock.Call
}

// GetAllPodsFromCgroups is a helper method to define mock.On call
func (_e *MockPodContainerManager_Expecter) GetAllPodsFromCgroups() *MockPodContainerManager_GetAllPodsFromCgroups_Call {
	return &MockPodContainerManager_GetAllPodsFromCgroups_Call{Call: _e.mock.On("GetAllPodsFromCgroups")}
}

func (_c *MockPodContainerManager_GetAllPodsFromCgroups_Call) Run(run func()) *MockPodContainerManager_GetAllPodsFromCgroups_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPodContainerManager_GetAllPodsFromCgroups_Call) Return(_a0 map[types.UID]cm.CgroupName, _a1 error) *MockPodContainerManager_GetAllPodsFromCgroups_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPodContainerManager_GetAllPodsFromCgroups_Call) RunAndReturn(run func() (map[types.UID]cm.CgroupName, error)) *MockPodContainerManager_GetAllPodsFromCgroups_Call {
	_c.Call.Return(run)
	return _c
}

// GetPodCgroupConfig provides a mock function with given fields: pod, resource
func (_m *MockPodContainerManager) GetPodCgroupConfig(pod *v1.Pod, resource v1.ResourceName) (*cm.ResourceConfig, error) {
	ret := _m.Called(pod, resource)

	if len(ret) == 0 {
		panic("no return value specified for GetPodCgroupConfig")
	}

	var r0 *cm.ResourceConfig
	var r1 error
	if rf, ok := ret.Get(0).(func(*v1.Pod, v1.ResourceName) (*cm.ResourceConfig, error)); ok {
		return rf(pod, resource)
	}
	if rf, ok := ret.Get(0).(func(*v1.Pod, v1.ResourceName) *cm.ResourceConfig); ok {
		r0 = rf(pod, resource)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cm.ResourceConfig)
		}
	}

	if rf, ok := ret.Get(1).(func(*v1.Pod, v1.ResourceName) error); ok {
		r1 = rf(pod, resource)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPodContainerManager_GetPodCgroupConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPodCgroupConfig'
type MockPodContainerManager_GetPodCgroupConfig_Call struct {
	*mock.Call
}

// GetPodCgroupConfig is a helper method to define mock.On call
//   - pod *v1.Pod
//   - resource v1.ResourceName
func (_e *MockPodContainerManager_Expecter) GetPodCgroupConfig(pod interface{}, resource interface{}) *MockPodContainerManager_GetPodCgroupConfig_Call {
	return &MockPodContainerManager_GetPodCgroupConfig_Call{Call: _e.mock.On("GetPodCgroupConfig", pod, resource)}
}

func (_c *MockPodContainerManager_GetPodCgroupConfig_Call) Run(run func(pod *v1.Pod, resource v1.ResourceName)) *MockPodContainerManager_GetPodCgroupConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*v1.Pod), args[1].(v1.ResourceName))
	})
	return _c
}

func (_c *MockPodContainerManager_GetPodCgroupConfig_Call) Return(_a0 *cm.ResourceConfig, _a1 error) *MockPodContainerManager_GetPodCgroupConfig_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPodContainerManager_GetPodCgroupConfig_Call) RunAndReturn(run func(*v1.Pod, v1.ResourceName) (*cm.ResourceConfig, error)) *MockPodContainerManager_GetPodCgroupConfig_Call {
	_c.Call.Return(run)
	return _c
}

// GetPodCgroupMemoryUsage provides a mock function with given fields: pod
func (_m *MockPodContainerManager) GetPodCgroupMemoryUsage(pod *v1.Pod) (uint64, error) {
	ret := _m.Called(pod)

	if len(ret) == 0 {
		panic("no return value specified for GetPodCgroupMemoryUsage")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(*v1.Pod) (uint64, error)); ok {
		return rf(pod)
	}
	if rf, ok := ret.Get(0).(func(*v1.Pod) uint64); ok {
		r0 = rf(pod)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(*v1.Pod) error); ok {
		r1 = rf(pod)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPodContainerManager_GetPodCgroupMemoryUsage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPodCgroupMemoryUsage'
type MockPodContainerManager_GetPodCgroupMemoryUsage_Call struct {
	*mock.Call
}

// GetPodCgroupMemoryUsage is a helper method to define mock.On call
//   - pod *v1.Pod
func (_e *MockPodContainerManager_Expecter) GetPodCgroupMemoryUsage(pod interface{}) *MockPodContainerManager_GetPodCgroupMemoryUsage_Call {
	return &MockPodContainerManager_GetPodCgroupMemoryUsage_Call{Call: _e.mock.On("GetPodCgroupMemoryUsage", pod)}
}

func (_c *MockPodContainerManager_GetPodCgroupMemoryUsage_Call) Run(run func(pod *v1.Pod)) *MockPodContainerManager_GetPodCgroupMemoryUsage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*v1.Pod))
	})
	return _c
}

func (_c *MockPodContainerManager_GetPodCgroupMemoryUsage_Call) Return(_a0 uint64, _a1 error) *MockPodContainerManager_GetPodCgroupMemoryUsage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPodContainerManager_GetPodCgroupMemoryUsage_Call) RunAndReturn(run func(*v1.Pod) (uint64, error)) *MockPodContainerManager_GetPodCgroupMemoryUsage_Call {
	_c.Call.Return(run)
	return _c
}

// GetPodContainerName provides a mock function with given fields: _a0
func (_m *MockPodContainerManager) GetPodContainerName(_a0 *v1.Pod) (cm.CgroupName, string) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetPodContainerName")
	}

	var r0 cm.CgroupName
	var r1 string
	if rf, ok := ret.Get(0).(func(*v1.Pod) (cm.CgroupName, string)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*v1.Pod) cm.CgroupName); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cm.CgroupName)
		}
	}

	if rf, ok := ret.Get(1).(func(*v1.Pod) string); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// MockPodContainerManager_GetPodContainerName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPodContainerName'
type MockPodContainerManager_GetPodContainerName_Call struct {
	*mock.Call
}

// GetPodContainerName is a helper method to define mock.On call
//   - _a0 *v1.Pod
func (_e *MockPodContainerManager_Expecter) GetPodContainerName(_a0 interface{}) *MockPodContainerManager_GetPodContainerName_Call {
	return &MockPodContainerManager_GetPodContainerName_Call{Call: _e.mock.On("GetPodContainerName", _a0)}
}

func (_c *MockPodContainerManager_GetPodContainerName_Call) Run(run func(_a0 *v1.Pod)) *MockPodContainerManager_GetPodContainerName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*v1.Pod))
	})
	return _c
}

func (_c *MockPodContainerManager_GetPodContainerName_Call) Return(_a0 cm.CgroupName, _a1 string) *MockPodContainerManager_GetPodContainerName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPodContainerManager_GetPodContainerName_Call) RunAndReturn(run func(*v1.Pod) (cm.CgroupName, string)) *MockPodContainerManager_GetPodContainerName_Call {
	_c.Call.Return(run)
	return _c
}

// IsPodCgroup provides a mock function with given fields: cgroupfs
func (_m *MockPodContainerManager) IsPodCgroup(cgroupfs string) (bool, types.UID) {
	ret := _m.Called(cgroupfs)

	if len(ret) == 0 {
		panic("no return value specified for IsPodCgroup")
	}

	var r0 bool
	var r1 types.UID
	if rf, ok := ret.Get(0).(func(string) (bool, types.UID)); ok {
		return rf(cgroupfs)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(cgroupfs)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) types.UID); ok {
		r1 = rf(cgroupfs)
	} else {
		r1 = ret.Get(1).(types.UID)
	}

	return r0, r1
}

// MockPodContainerManager_IsPodCgroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsPodCgroup'
type MockPodContainerManager_IsPodCgroup_Call struct {
	*mock.Call
}

// IsPodCgroup is a helper method to define mock.On call
//   - cgroupfs string
func (_e *MockPodContainerManager_Expecter) IsPodCgroup(cgroupfs interface{}) *MockPodContainerManager_IsPodCgroup_Call {
	return &MockPodContainerManager_IsPodCgroup_Call{Call: _e.mock.On("IsPodCgroup", cgroupfs)}
}

func (_c *MockPodContainerManager_IsPodCgroup_Call) Run(run func(cgroupfs string)) *MockPodContainerManager_IsPodCgroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPodContainerManager_IsPodCgroup_Call) Return(_a0 bool, _a1 types.UID) *MockPodContainerManager_IsPodCgroup_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPodContainerManager_IsPodCgroup_Call) RunAndReturn(run func(string) (bool, types.UID)) *MockPodContainerManager_IsPodCgroup_Call {
	_c.Call.Return(run)
	return _c
}

// ReduceCPULimits provides a mock function with given fields: name
func (_m *MockPodContainerManager) ReduceCPULimits(name cm.CgroupName) error {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for ReduceCPULimits")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(cm.CgroupName) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPodContainerManager_ReduceCPULimits_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReduceCPULimits'
type MockPodContainerManager_ReduceCPULimits_Call struct {
	*mock.Call
}

// ReduceCPULimits is a helper method to define mock.On call
//   - name cm.CgroupName
func (_e *MockPodContainerManager_Expecter) ReduceCPULimits(name interface{}) *MockPodContainerManager_ReduceCPULimits_Call {
	return &MockPodContainerManager_ReduceCPULimits_Call{Call: _e.mock.On("ReduceCPULimits", name)}
}

func (_c *MockPodContainerManager_ReduceCPULimits_Call) Run(run func(name cm.CgroupName)) *MockPodContainerManager_ReduceCPULimits_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(cm.CgroupName))
	})
	return _c
}

func (_c *MockPodContainerManager_ReduceCPULimits_Call) Return(_a0 error) *MockPodContainerManager_ReduceCPULimits_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPodContainerManager_ReduceCPULimits_Call) RunAndReturn(run func(cm.CgroupName) error) *MockPodContainerManager_ReduceCPULimits_Call {
	_c.Call.Return(run)
	return _c
}

// SetPodCgroupConfig provides a mock function with given fields: pod, resourceConfig
func (_m *MockPodContainerManager) SetPodCgroupConfig(pod *v1.Pod, resourceConfig *cm.ResourceConfig) error {
	ret := _m.Called(pod, resourceConfig)

	if len(ret) == 0 {
		panic("no return value specified for SetPodCgroupConfig")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1.Pod, *cm.ResourceConfig) error); ok {
		r0 = rf(pod, resourceConfig)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPodContainerManager_SetPodCgroupConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPodCgroupConfig'
type MockPodContainerManager_SetPodCgroupConfig_Call struct {
	*mock.Call
}

// SetPodCgroupConfig is a helper method to define mock.On call
//   - pod *v1.Pod
//   - resourceConfig *cm.ResourceConfig
func (_e *MockPodContainerManager_Expecter) SetPodCgroupConfig(pod interface{}, resourceConfig interface{}) *MockPodContainerManager_SetPodCgroupConfig_Call {
	return &MockPodContainerManager_SetPodCgroupConfig_Call{Call: _e.mock.On("SetPodCgroupConfig", pod, resourceConfig)}
}

func (_c *MockPodContainerManager_SetPodCgroupConfig_Call) Run(run func(pod *v1.Pod, resourceConfig *cm.ResourceConfig)) *MockPodContainerManager_SetPodCgroupConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*v1.Pod), args[1].(*cm.ResourceConfig))
	})
	return _c
}

func (_c *MockPodContainerManager_SetPodCgroupConfig_Call) Return(_a0 error) *MockPodContainerManager_SetPodCgroupConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPodContainerManager_SetPodCgroupConfig_Call) RunAndReturn(run func(*v1.Pod, *cm.ResourceConfig) error) *MockPodContainerManager_SetPodCgroupConfig_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPodContainerManager creates a new instance of MockPodContainerManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPodContainerManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPodContainerManager {
	mock := &MockPodContainerManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
