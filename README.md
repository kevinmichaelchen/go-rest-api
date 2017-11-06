## Getting started
```bash
mysql -uroot -e 'CREATE DATABASE IF NOT EXISTS rest_api_example;'

go test -v

make run
```

## TODO
- [goose](https://github.com/pressly/goose) or [migrate](https://github.com/mattes/migrate) for DB migrations.
- [Minimal Docker container](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/).

## References
* [Medium article](https://medium.com/@kelvinpfw/building-and-testing-a-rest-api-in-golang-using-gorilla-mux-and-mysql-1f0518818ff6)
