package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"time"

	pb "ohos-grpc-testServer/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// 配置 TLS - 跳过证书校验（用于自签名证书）
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // 跳过证书校验
	}
	creds := credentials.NewTLS(tlsConfig)

	// 连接到 gRPC 服务器
	serverAddr := "localhost:50051"
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	client := pb.NewStreamServiceClient(conn)

	fmt.Printf("已连接到服务器 (TLS): %s\n", serverAddr)
	fmt.Println("=" + string(make([]byte, 50)))

	// 发送请求并获取流
	request := &pb.NumberRequest{
		Count:   10,  // 请求 10 个数字
		DelayMs: 500, // 每个数字延迟 500ms
	}

	fmt.Printf("发送请求: count=%d, delay=%dms\n", request.Count, request.DelayMs)
	fmt.Println("=" + string(make([]byte, 50)))

	// 调用服务器流式 RPC，获取 ClientReadableStream
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	stream, err := client.GetNumbers(ctx, request)
	if err != nil {
		log.Fatalf("调用 GetNumbers 失败: %v", err)
	}

	fmt.Println("开始接收流式数据...")
	fmt.Println()

	// 使用 ClientReadableStream 读取流式响应
	receivedCount := 0
	for {
		// 从流中读取响应
		response, err := stream.Recv()
		if err == io.EOF {
			// 流结束
			fmt.Println()
			fmt.Println("=" + string(make([]byte, 50)))
			fmt.Println("流式传输结束")
			break
		}
		if err != nil {
			log.Fatalf("接收数据失败: %v", err)
		}

		// 处理接收到的数据
		receivedCount++
		fmt.Printf("[%d] 收到数字: %d, 时间戳: %s\n",
			receivedCount, response.Number, response.Timestamp)
	}

	fmt.Printf("\n总共接收了 %d 个数字\n", receivedCount)
}
