// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	pkg "github.com/vektra/mockery/v2/pkg"

	testing "testing"
)

// WalkerVisitor is an autogenerated mock type for the WalkerVisitor type
type WalkerVisitor struct {
	mock.Mock
}

// VisitWalk provides a mock function with given fields: _a0, _a1
func (_m *WalkerVisitor) VisitWalk(_a0 context.Context, _a1 *pkg.Interface) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *pkg.Interface) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewWalkerVisitor creates a new instance of WalkerVisitor. It also registers a cleanup function to assert the mocks expectations.
func NewWalkerVisitor(t testing.TB) *WalkerVisitor {
	mock := &WalkerVisitor{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
