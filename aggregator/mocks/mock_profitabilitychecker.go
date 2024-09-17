// Code generated by mockery v2.39.0. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	mock "github.com/stretchr/testify/mock"
)

// ProfitabilityCheckerMock is an autogenerated mock type for the aggregatorTxProfitabilityChecker type
type ProfitabilityCheckerMock struct {
	mock.Mock
}

// IsProfitable provides a mock function with given fields: _a0, _a1
func (_m *ProfitabilityCheckerMock) IsProfitable(_a0 context.Context, _a1 *big.Int) (bool, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for IsProfitable")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) (bool, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProfitabilityCheckerMock creates a new instance of ProfitabilityCheckerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProfitabilityCheckerMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProfitabilityCheckerMock {
	mock := &ProfitabilityCheckerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
