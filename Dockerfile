FROM golang:1.9.2 as builder
RUN go get -d -v github.com/gorilla/mux github.com/go-sql-driver/mysql
ADD . /go/src/github.com/teslagov/clarakm-projects-go/
WORKDIR /go/src/github.com/teslagov/clarakm-projects-go/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/teslagov/clarakm-projects-go/app .
CMD ["./app"]