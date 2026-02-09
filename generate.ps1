# Generate gRPC code from proto file
protoc --go_out=. --go_opt=paths=source_relative `
       --go-grpc_out=. --go-grpc_opt=paths=source_relative `
       proto/stream.proto

Write-Host "Proto files generated successfully!"
