FROM golang:1.17.10-bullseye

RUN apt update \
    && apt install -y \
    build-essential \
    curl \
    git \
    # golang-goprotobuf-dev \
    libtool \
    libssl-dev \
    make \
    # protobuf-compiler \
    zip

# ARG NODEJS_VERSION="v16.15.0"
# RUN curl -o node-linux-x64.tar.xz "https://nodejs.org/dist/${NODEJS_VERSION}/node-${NODEJS_VERSION}-linux-x64.tar.xz" \
#     && tar -xf node-linux-x64.tar.xz \
#     && cd "node-${NODEJS_VERSION}-linux-x64" \
#     && cp -r lib /usr \
#     && cp -r share /usr \
#     && cp -r include /usr \
#     && cp -r bin /usr \
#     && cd .. \
#     && rm -fr node-*

WORKDIR /usr/local/go/src/github.com/nicholasaiello/go-http2-api-example

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
# RUN go mod tidy && go mod verify
RUN go mod download && go mod verify

#RUN go get -u github.com/golang/protobuf/proto
#RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN GO111MODULE=on go get -u google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc
        # google.golang.org/protobuf/proto \
        # google.golang.org/protobuf/protoc-gen-go \
        # google.golang.org/protobuf/protoc-gen-go-grpc && \
        # mv /go/bin/protoc-gen-go* /usr/local/bin/

# ENV PATH="${PATH}:$(go env GOPATH)/bin"

#RUN echo 'export PATH=$PATH:$GOPATH/bin' >> $HOME/.bashrc

# Copy app
COPY . .

# RUN cd client \
#     && npm ci \
#     && npm run build:client

# RUN protoc --go_out=./ --go_opt=paths=source_relative \
#     --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
#     proto/holdings.proto

# RUN protoc -I=./proto holdings.proto  \
#     --js_out=import_style=commonjs:./client/js \
#     --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./client/js

# RUN go build -v -o /usr/local/bin/app ./...

EXPOSE 8080 8443 9090
CMD ["go", "run", "server.go"]
