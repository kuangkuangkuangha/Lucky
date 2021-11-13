FROM golang:latest

ENV GO111MODULE=on \
	GOPROXY="https://goproxy.cn,direct" \
	server_host="0.0.0.0" \
 	server_port="5000" \
 	redirect_url="https://lucky.itoken.team/" \
 	db_server="localhost" \
 	db_port="3306" \
 	db_name="lucky" \
 	db_username="mariadb" \
 	db_password=""

WORKDIR /build

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN go build lucky 

CMD ["/build/lucky"]
