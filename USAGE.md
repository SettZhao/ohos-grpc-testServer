# 使用说明

## 快速开始

### 1. 启动服务器

```bash
go run main.go
```

或者运行编译好的可执行文件：
```bash
.\server.exe
```

### 2. 运行客户端

在另一个终端运行：
```bash
go run testClient/client.go
```

或者：
```bash
.\client.exe
```

## 测试结果

客户端会请求 10 个数字，每个数字间隔 500ms，通过 ClientReadableStream 接收服务器的流式响应。

预期输出：
```
已连接到服务器: localhost:50051
发送请求: count=10, delay=500ms
开始接收流式数据...

[1] 收到数字: 1, 时间戳: 2026-02-09 12:03:35.667
[2] 收到数字: 2, 时间戳: 2026-02-09 12:03:36.168
...
[10] 收到数字: 10, 时间戳: 2026-02-09 12:03:40.173

流式传输结束
总共接收了 10 个数字
```

## 自定义参数

修改 [testClient/client.go](testClient/client.go#L27-L30) 中的请求参数：

```go
request := &pb.NumberRequest{
    Count:   10,     // 修改数字数量
    DelayMs: 500,    // 修改延迟时间（毫秒）
}
```
