FROM golang:1.15

WORKDIR /go/src
COPY external/ external/
COPY hypercloudurl/ hypercloudurl/
COPY hyperclusterresource/ hyperclusterresource/
COPY remotecluster/ remotecluster/
COPY util/ util/
COPY go.mod go.mod 
COPY go.sum go.sum
COPY main.go main.go
COPY start.sh start.sh

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o main .
RUN chmod +777 main
RUN chmod +777 start.sh
ENTRYPOINT ["/go/src/start.sh"]
