# 使用基础的 Golang 镜像作为构建环境
FROM golang:1.20.1
LABEL authors="lk"

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

# 设置镜像源为阿里源
RUN sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list \
    && sed -i 's/security.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list


# 安装 git 和 shell
RUN apt-get update \
    && apt-get install -y git \
    && rm -rf /var/lib/apt/lists/*

# 项目的工作路径
WORKDIR /data

# 复制 go.mod 和 go.sum 文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

# 将代码复制到容器中
COPY . .

# 允许使用 shell 命令
SHELL ["/bin/bash", "-c"]
#暴露端口
EXPOSE 6688