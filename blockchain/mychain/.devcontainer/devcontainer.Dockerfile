FROM mcr.microsoft.com/devcontainers/go:1.20-bullseye
WORKDIR /mychain
USER root

RUN sudo apt-get update -y \
    && sudo apt-get install -y --no-install-recommends \
    make automake pkg-config libtool autoconf protobuf-compiler libprotobuf-dev lsb-release wget

# install grpc and grpc-gateway binaries
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0 \
    && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.13.0 \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.13.0 \
    && go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@v1.16.0 \
    && go install github.com/envoyproxy/protoc-gen-validate@v0.9.0

COPY . .
ENTRYPOINT ["sleep", "infinity"]