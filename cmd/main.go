package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"github.com/zwell/service-sms/internal/config"
	"github.com/zwell/service-sms/internal/service"

	"google.golang.org/grpc"
	pb "github.com/zwell/service-sms/proto"
)

type server struct {
	pb.UnimplementedSmsServer
}

func (s *server) Send(ctx context.Context, in *pb.Request) (*pb.Reply, error) {

	// json 数据解析
	jsonData := in.GetParams()
	var f interface{}
	b := []byte(jsonData)
	err := json.Unmarshal(b, &f)
	if err != nil {
		return &pb.Reply{Code: 500, Message: "参数解析失败params"}, nil
	}
	m := f.(map[string]interface{})

	// 发送短信
	var smsService service.SmsService
	response, err := smsService.Send(in.GetTemplate(), in.GetPhone(), m)
	if err != nil {
		return &pb.Reply{Code: 500, Message: err.Error()}, nil
	}

	fmt.Println(in, response)

	return &pb.Reply{Code: response.Code, Message: response.Message}, nil
}

func main() {
	// 监听端口
	conf := config.GetConf()
	lis, err := net.Listen("tcp", ":"+conf.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 创建服务
	s := grpc.NewServer()
	pb.RegisterSmsServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
