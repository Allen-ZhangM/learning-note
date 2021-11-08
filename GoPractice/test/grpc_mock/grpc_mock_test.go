package grpc_mock

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/examples/helloworld/helloworld"
	"testing"
)

func TestName(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockGreeterClient := NewMockGreeterClient(ctrl)

	req := &helloworld.HelloRequest{Name: "unit_test"}
	mockGreeterClient.EXPECT().SayHello(
		gomock.Any(),
		&rpcMsg{msg: req},
	).Return(&helloworld.HelloReply{Message: "Mocked Interface"}, nil)

	resp, err := mockGreeterClient.SayHello(context.Background(), req)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(resp)
}
