FROM golang
RUN mkdir /serv
ADD . /serv
WORKDIR /serv
RUN go build  ./cmd/main.go
CMD ["/serv/main"]
