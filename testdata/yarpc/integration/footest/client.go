// Code generated by thriftrw-plugin-yarpc
// @generated

package footest

import (
	"context"
	"go.uber.org/yarpc"
	"github.com/golang/mock/gomock"
	"github.com/yarpc/yab/testdata/yarpc/integration/fooclient"
)

// MockClient implements a gomock-compatible mock client for service
// Foo.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

var _ fooclient.Interface = (*MockClient)(nil)

type _MockClientRecorder struct {
	mock *MockClient
}

// Build a new mock client for service Foo.
//
// 	mockCtrl := gomock.NewController(t)
// 	client := footest.NewMockClient(mockCtrl)
//
// Use EXPECT() to set expectations on the mock.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

// EXPECT returns an object that allows you to define an expectation on the
// Foo mock client.
func (m *MockClient) EXPECT() *_MockClientRecorder {
	return m.recorder
}

// Bar responds to a Bar call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().Bar(gomock.Any(), ...).Return(...)
// 	... := client.Bar(...)
func (m *MockClient) Bar(
	ctx context.Context,
	_Arg *int32,
	opts ...yarpc.CallOption,
) (success int32, err error) {

	args := []interface{}{ctx, _Arg}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "Bar", args...)
	success, _ = ret[i].(int32)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) Bar(
	ctx interface{},
	_Arg interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Arg}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "Bar", args...)
}
