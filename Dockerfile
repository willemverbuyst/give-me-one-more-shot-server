FROM golang:1.19-alpine

WORKDIR /server

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY *.go ./

RUN go build -o /go-server
 
EXPOSE 9090

CMD [ "/go-server" ]