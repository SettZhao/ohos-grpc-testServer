#!/bin/bash
# 生成自签名证书用于 gRPC TLS

CERT_DIR="./certs"
DAYS=365

# 创建证书目录
mkdir -p $CERT_DIR

echo "开始生成自签名证书..."

# 生成私钥
openssl genrsa -out $CERT_DIR/server.key 2048

# 生成证书签名请求（CSR）
openssl req -new -key $CERT_DIR/server.key -out $CERT_DIR/server.csr \
    -subj "/C=CN/ST=Beijing/L=Beijing/O=TestOrg/OU=TestUnit/CN=localhost"

# 生成自签名证书
openssl x509 -req -days $DAYS -in $CERT_DIR/server.csr \
    -signkey $CERT_DIR/server.key -out $CERT_DIR/server.crt

# 清理 CSR 文件
rm $CERT_DIR/server.csr

echo "证书生成完成！"
echo "私钥: $CERT_DIR/server.key"
echo "证书: $CERT_DIR/server.crt"
echo "有效期: $DAYS 天"

# 显示证书信息
openssl x509 -in $CERT_DIR/server.crt -text -noout | grep -E "Subject:|Not Before|Not After"
