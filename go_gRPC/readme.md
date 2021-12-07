# go_gRPC steps

## 1. 安装Protocol Buffers3 编译器
<blockquote>https://github.com/protocolbuffers/protobuf/releases</blockquote>
安装完毕后，查看版本
<blockquote>$ protoc --version</blockquote>

## 2. 安装golang grpc插件

### 2.1 设置代理(非必须)，视网络情况而定
export GOPROXY="https://goproxy.cn,direct"

### 2.2 安装grpc golang插件
go get google.golang.org/grpc

## 3.0 安装golang版本proto编译器插件
<blockquote>go get -u github.com/golang/protobuf/{proto,protoc-gen-go}</blockquote>

## 4.0 生成(根据protoc参数调整)
<blockquote>
1：protoc -I . --go_out=plugins=grpc:. ./example/data.proto <br />
2：protoc -I . --go_out=plugins=grpc:. --go_opt=paths=source_relative ./example/data.proto
</blockquote>

