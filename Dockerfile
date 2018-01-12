FROM golang:1.9.2 as builder
RUN go get -d -v github.com/gorilla/mux github.com/go-sql-driver/mysql
ADD . /go/src/github.com/teslagov/clarakm-projects-go/
WORKDIR /go/src/github.com/teslagov/clarakm-projects-go/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
#ENV CLARA_USER="clara" \
#    CLARA_UID="8080" \
#    CLARA_GROUP="clara" \
#    CLARA_GID="8080"
RUN apk --no-cache add ca-certificates
#&& \
#    addgroup -S -g $CLARA_GID $CLARA_GROUP && \
#    adduser -S -u $CLARA_UID -G $CLARA_GROUP $CLARA_USER
WORKDIR /root/
COPY --from=builder /go/src/github.com/teslagov/clarakm-projects-go/app .
#USER $CLARA_USER
CMD ["./app"]