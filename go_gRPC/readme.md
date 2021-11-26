# go_gRPC steps

## 1. Protocol Buffers3 编译器
<blockquote>https://github.com/protocolbuffers/protobuf/releases</blockquote>
安装完毕后，查看版本
<blockquote>$ protoc --version</blockquote>

## 2. golang grpc插件

### 2.1 设置下代理(非必须)，视网络情况而定
export GOPROXY="https://goproxy.cn,direct"

### 2.2 grpc插件 安装
go get google.golang.org/grpc

## 3.0 golang版本proto编译器插件
<blockquote>go get -u github.com/golang/protobuf/{proto,protoc-gen-go}</blockquote>

## 4.0 生成(根据protoc参数调整)
<blockquote>protoc -I . --go_out=plugins=grpc:. ./example/data.proto</blockquote>

