// Code generated by mockery. DO NOT EDIT.

package l1_parallel_sync

import (
	etherman "github.com/0xPolygonHermez/zkevm-node/etherman"
	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"
)

// synchronizerProcessBlockRangeInterfaceMock is an autogenerated mock type for the synchronizerProcessBlockRangeInterface type
type synchronizerProcessBlockRangeInterfaceMock struct {
	mock.Mock
}

type synchronizerProcessBlockRangeInterfaceMock_Expecter struct {
	mock *mock.Mock
}

func (_m *synchronizerProcessBlockRangeInterfaceMock) EXPECT() *synchronizerProcessBlockRangeInterfaceMock_Expecter {
	return &synchronizerProcessBlockRangeInterfaceMock_Expecter{mock: &_m.Mock}
}

// ProcessBlockRange provides a mock function with given fields: blocks, order
func (_m *synchronizerProcessBlockRangeInterfaceMock) ProcessBlockRange(blocks []etherman.Block, order map[common.Hash][]etherman.Order) error {
	ret := _m.Called(blocks, order)

	if len(ret) == 0 {
		panic("no return value specified for ProcessBlockRange")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]etherman.Block, map[common.Hash][]etherman.Order) error); ok {
		r0 = rf(blocks, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// synchronizerProcessBlockRangeInterfaceMock_ProcessBlockRange_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessBlockRange'
type synchronizerProcessBlockRangeInterfaceMock_ProcessBlockRange_Call struct {
	*mock.Call
}

// ProcessBlockRange is a helper method to define mock.On call
//   - blocks []etherman.Block
//   - order map[common.Hash][]etherman.Order
func (_e *synchronizerProcessBlockRangeInterfaceMock_Expecter) ProcessBlockRange(blocks interface{}, order interface{}) *synchronizerProcessBlockRangeInterfaceMock_ProcessBlockRange_Call {
	return &synchronizerProcessBlockRangeInterfaceMock_ProcessBlockRange_Call{Call: _e.mock.On("ProcessBlockRange", blocks, order)}
}

func (_c *synchronizerProcessBlockRangeInterfaceMock_ProcessBlockRange_Call) Run(run func(blocks []etherman.Block, order map[common.Hash][]etherman.Order)) *synchronizerProcessBlockRangeInterfaceMock_ProcessBlockRange_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]etherman.Block), args[1].(map[common.Hash][]etherman.Order))
	})
	return _c
}

func (_c *synchronizerProcessBlockRangeInterfaceMock_ProcessBlockRange_Call) Return(_a0 error) *synchronizerProcessBlockRangeInterfaceMock_ProcessBlockRange_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *synchronizerProcessBlockRangeInterfaceMock_ProcessBlockRange_Call) RunAndReturn(run func([]etherman.Block, map[common.Hash][]etherman.Order) error) *synchronizerProcessBlockRangeInterfaceMock_ProcessBlockRange_Call {
	_c.Call.Return(run)
	return _c
}

// newSynchronizerProcessBlockRangeInterfaceMock creates a new instance of synchronizerProcessBlockRangeInterfaceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newSynchronizerProcessBlockRangeInterfaceMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *synchronizerProcessBlockRangeInterfaceMock {
	mock := &synchronizerProcessBlockRangeInterfaceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
