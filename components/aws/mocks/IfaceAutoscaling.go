// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	autoscaling "deploy-helper/components/aws/autoscaling"

	mock "github.com/stretchr/testify/mock"
)

// IfaceAutoscaling is an autogenerated mock type for the IfaceAutoscaling type
type IfaceAutoscaling struct {
	mock.Mock
}

// GetTargetGroups provides a mock function with given fields: autoscalingName
func (_m *IfaceAutoscaling) GetTargetGroups(autoscalingName string) ([]autoscaling.TargetGroup, error) {
	ret := _m.Called(autoscalingName)

	var r0 []autoscaling.TargetGroup
	if rf, ok := ret.Get(0).(func(string) []autoscaling.TargetGroup); ok {
		r0 = rf(autoscalingName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]autoscaling.TargetGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(autoscalingName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
