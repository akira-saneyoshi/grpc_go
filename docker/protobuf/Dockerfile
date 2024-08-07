FROM golang:1.22.5

ARG PROTOBUF_VERSION=27.3
EXPOSE 8090

RUN apt-get update \
    && apt-get install -y build-essential \
    && apt-get install -y tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && dpkg-reconfigure -f noninteractive tzdata \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /tmp/protoc
RUN apt-get update \
    && apt-get install -y --no-install-recommends unzip curl \
    && curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip -o protoc.zip \
    && unzip protoc.zip \
    && mv /tmp/protoc/bin/* /usr/local/bin/ \
    && mv /tmp/protoc/include /usr/local/include \
    && rm -rf protoc protoc.zip \
    && apt-get install -y --no-install-recommends protobuf-compiler \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /opt/protoc_service
COPY ./grpc_backend/protoc_service /opt/protoc_service
RUN go mod tidy \
  && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
  && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest