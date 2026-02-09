# TLS 配置说明

## 证书生成

本项目包含两个证书生成脚本：

### Linux/云服务器使用（推荐）
```bash
bash generate-cert.sh
```

### Windows 本地使用
```bash
bash generate-cert.sh  # 使用 Git Bash
# 或
.\generate-cert.ps1    # 使用 PowerShell (需要 OpenSSL)
```

## 生成的文件

证书文件存放在 `certs/` 目录：
- `certs/server.key` - 服务器私钥
- `certs/server.crt` - 服务器证书（自签名，有效期365天）

## 服务器配置

服务器已配置使用 TLS 加密：
```go
creds, err := credentials.NewServerTLSFromFile("certs/server.crt", "certs/server.key")
grpcServer := grpc.NewServer(grpc.Creds(creds))
```

## 客户端配置

客户端配置为**跳过证书校验**（适用于自签名证书）：
```go
tlsConfig := &tls.Config{
    InsecureSkipVerify: true, // 跳过证书校验
}
creds := credentials.NewTLS(tlsConfig)
conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(creds))
```

## 云服务器部署

### 1. 生成证书
```bash
# SSH 登录到云服务器
ssh user@your-server

# 进入项目目录
cd /path/to/ohos-grpc-testServer

# 生成证书
bash generate-cert.sh
```

### 2. 修改证书的 CN 字段（可选）

如果需要使用域名或公网IP，修改 `generate-cert.sh` 中的 CN 字段：
```bash
# 使用域名
-subj "/C=CN/ST=Beijing/L=Beijing/O=TestOrg/OU=TestUnit/CN=yourdomain.com"

# 使用公网 IP
-subj "/C=CN/ST=Beijing/L=Beijing/O=TestOrg/OU=TestUnit/CN=1.2.3.4"
```

### 3. 启动服务器
```bash
# 编译
go build -o server main.go

# 运行
./server

# 后台运行
nohup ./server > server.log 2>&1 &
```

### 4. 客户端连接
修改客户端的服务器地址：
```go
serverAddr := "yourdomain.com:50051"  // 或使用 IP:端口
```

## 安全注意事项

⚠️ **当前配置的安全级别**：
- ✅ 传输数据已加密（TLS）
- ⚠️ 客户端跳过证书校验（`InsecureSkipVerify: true`）

**生产环境建议**：
1. 使用 CA 签名的证书（Let's Encrypt 等）
2. 客户端启用证书校验
3. 配置正确的域名和证书 CN 字段

## 测试验证

### 本地测试
```bash
# 终端1: 启动服务器
go run main.go

# 终端2: 运行客户端
go run testClient/client.go
```

### 验证 TLS 连接
客户端输出应显示：
```
已连接到服务器 (TLS): localhost:50051
```

服务器输出应显示：
```
gRPC 服务器启动 (TLS已启用)，监听端口 :50051
```

## 常见问题

### Q: Windows 上 OpenSSL 不可用？
A: 使用 Git Bash 运行 `bash generate-cert.sh`，Git for Windows 自带 OpenSSL

### Q: 证书过期怎么办？
A: 重新运行证书生成脚本即可：`bash generate-cert.sh`

### Q: 如何在客户端启用证书校验？
A: 修改客户端代码，加载证书并移除 `InsecureSkipVerify`：
```go
creds, err := credentials.NewClientTLSFromFile("certs/server.crt", "")
conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(creds))
```

### Q: 云服务器上端口被防火墙拦截？
A: 确保开放 50051 端口：
```bash
# Ubuntu/Debian
sudo ufw allow 50051/tcp

# CentOS/RHEL
sudo firewall-cmd --permanent --add-port=50051/tcp
sudo firewall-cmd --reload
```
