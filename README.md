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
- Pagination
- Per-route JWT authentication
- DB migrations.
  - Currently I have a Make target to just create tables.
  - We should have something akin to Flyway, which runs on program start and runs migration scripts.
  - Look into [goose](https://github.com/pressly/goose) or [migrate](https://github.com/mattes/migrate).
- [Minimal Docker container](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/).
- [Harden Alpine](https://gist.github.com/jumanjiman/f9d3db977846c163df12)

## References
* [Medium article](https://medium.com/@kelvinpfw/building-and-testing-a-rest-api-in-golang-using-gorilla-mux-and-mysql-1f0518818ff6)
* [HTTP up front, Protobufs in the rear](https://github.com/harlow/go-micro-services)
