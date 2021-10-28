FROM golang:latest

ENV GO111MODULE=on \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /build

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN go build lucky 

CMD ["/build/lucky"]
