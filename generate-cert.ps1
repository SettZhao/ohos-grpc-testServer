# PowerShell 脚本 - 生成自签名证书用于 gRPC TLS

$CERT_DIR = ".\certs"
$DAYS = 365

# 创建证书目录
if (-not (Test-Path $CERT_DIR)) {
    New-Item -ItemType Directory -Path $CERT_DIR | Out-Null
}

Write-Host "开始生成自签名证书..." -ForegroundColor Green

# 检查 openssl 是否可用
try {
    $null = Get-Command openssl -ErrorAction Stop
} catch {
    Write-Host "错误: 未找到 openssl 命令" -ForegroundColor Red
    Write-Host "请安装 OpenSSL 或使用 Git Bash 运行 generate-cert.sh" -ForegroundColor Yellow
    exit 1
}

# 生成私钥
Write-Host "生成私钥..."
& openssl genrsa -out "$CERT_DIR\server.key" 2048

# 生成证书签名请求（CSR）
Write-Host "生成证书签名请求..."
& openssl req -new -key "$CERT_DIR\server.key" -out "$CERT_DIR\server.csr" `
    -subj "/C=CN/ST=Beijing/L=Beijing/O=TestOrg/OU=TestUnit/CN=localhost"

# 生成自签名证书
Write-Host "生成自签名证书..."
& openssl x509 -req -days $DAYS -in "$CERT_DIR\server.csr" `
    -signkey "$CERT_DIR\server.key" -out "$CERT_DIR\server.crt"

# 清理 CSR 文件
Remove-Item "$CERT_DIR\server.csr" -ErrorAction SilentlyContinue

Write-Host "`n证书生成完成！" -ForegroundColor Green
Write-Host "私钥: $CERT_DIR\server.key"
Write-Host "证书: $CERT_DIR\server.crt"
Write-Host "有效期: $DAYS 天"

# 显示证书信息
Write-Host "`n证书信息:" -ForegroundColor Cyan
& openssl x509 -in "$CERT_DIR\server.crt" -text -noout | Select-String -Pattern "Subject:|Not Before|Not After"
