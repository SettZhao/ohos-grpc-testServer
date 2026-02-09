package main

import (
	"fmt"
	"log"
	"net"

	// "time"

	pb "ohos-grpc-testServer/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// StreamServiceServer 实现流式服务
type StreamServiceServer struct {
	pb.UnimplementedStreamServiceServer
}

// GetNumbers 实现服务器流式 RPC
func (s *StreamServiceServer) GetNumbers(req *pb.NumberRequest, stream pb.StreamService_GetNumbersServer) error {
	count := req.GetCount()
	delayMs := req.GetDelayMs()

	log.Printf("收到请求: count=%d, delay=%dms", count, delayMs)

	// 直接返回错误，不发送任何数据
	log.Println("服务器主动关闭流并返回错误")
	return fmt.Errorf("服务器拒绝处理此请求：测试异常场景")
}

func main() {
	// 加载 TLS 证书
	certFile := "certs/server.crt"
	keyFile := "certs/server.key"
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("加载证书失败: %v", err)
	}

	// 监听端口
	port := ":50051"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	// 创建带 TLS 的 gRPC 服务器
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	// 注册服务
	pb.RegisterStreamServiceServer(grpcServer, &StreamServiceServer{})

	fmt.Printf("gRPC 服务器启动 (TLS已启用)，监听端口 %s\n", port)
	log.Printf("等待客户端连接...\n")

	// 启动服务器
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
