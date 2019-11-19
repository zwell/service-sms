package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"zwell.github/mic-server/sms/internal/config"

	"google.golang.org/grpc"
	"zwell.github/mic-server/sms/internal/database"
	"zwell.github/mic-server/sms/internal/service"
	pb "zwell.github/mic-server/sms/proto"
)

func init() {
	// init database
	database.GetDB()
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedSmsServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Send(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	// 获取发送的供应商
	var smsService service.SmsService
	templateResult, err := smsService.GetSupplier(in.GetTemplate())
	if err != nil {
		return &pb.Reply{Code: 500, Message: err.Error()}, nil
	}

	// json 数据解析
	jsonData := in.GetParams()
	var f interface{}
	b := []byte(jsonData)
	err = json.Unmarshal(b, &f)
	if err != nil {
		return &pb.Reply{Code: 500, Message: "参数解析失败params"}, nil
	}
	m := f.(map[string]interface{})

	// 发送短信
	response, err := smsService.Send(templateResult, in.GetPhone(), m)
	if err != nil {
		return &pb.Reply{Code: 500, Message: err.Error()}, nil
	}

	fmt.Println(response)

	log.Printf("Received: %v %v", in.GetTemplate(), in.GetPhone())

	return &pb.Reply{Code: response.Code, Message: response.Message}, nil
}

func main() {
	conf := config.GetConf()
	lis, err := net.Listen("tcp", ":"+conf.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSmsServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
