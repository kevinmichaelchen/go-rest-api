## Getting started
```bash
# Rebuild the Docker image and start it
make rebuild start

# Create our DB table 
make seed

# List users
make list-users

# Create a user
make create-user

# Run tests
make test
```

To run tests, use `go test -v`.

## TODO
- [goose](https://github.com/pressly/goose) or [migrate](https://github.com/mattes/migrate) for DB migrations.
- [Minimal Docker container](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/).
- [Harden Alpine](https://gist.github.com/jumanjiman/f9d3db977846c163df12)

## References
* [Medium article](https://medium.com/@kelvinpfw/building-and-testing-a-rest-api-in-golang-using-gorilla-mux-and-mysql-1f0518818ff6)
* [HTTP up front, Protobufs in the rear](https://github.com/harlow/go-micro-services)
