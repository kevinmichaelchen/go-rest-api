## Getting started
To build Docker images and spin them up, just run
```bash
make
```

To bring it all down, run
```bash
make stop
```

### Using the API
```bash
make list-users
make create-user
make list-users
```

## TODO
- Project layout
- Pagination
- Per-route JWT authentication
- DB migrations.
  - Currently I have a Make target to just create tables.
  - We should have something akin to Flyway, which runs on program start and runs migration scripts.
  - Look into [goose](https://github.com/pressly/goose) or [migrate](https://github.com/mattes/migrate).
- [Minimal Docker container](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/).
- [Harden Alpine](https://gist.github.com/jumanjiman/f9d3db977846c163df12)

## References
### basic
* [Building and Testing a REST API in GoLang using Gorilla Mux and MySQL](https://medium.com/@kelvinpfw/building-and-testing-a-rest-api-in-golang-using-gorilla-mux-and-mysql-1f0518818ff6)
* https://talks.golang.org/2014/organizeio.slide#4

### structure
* https://github.com/qiangxue/golang-restful-starter-kit

### pagination
* [Twitter Cursoring for Pagination](https://developer.twitter.com/en/docs/basics/cursoring)

### jwt
* [auth0's go-jwt-middleware](https://github.com/auth0/go-jwt-middleware)
* [How to correctly use context.Context in Go 1.7](https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39)
* [Gorilla context vs context.Context](https://github.com/gorilla/context/issues/32)
* [Revisiting context and http.Handler for Go 1.7](https://joeshaw.org/revisiting-context-and-http-handler-for-go-17/)
* [Go Context (2014), might be legacy](https://blog.golang.org/context)

### grpc
* [HTTP up front, Protobufs in the rear](https://github.com/harlow/go-micro-services)