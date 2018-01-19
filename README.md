## Getting started
To build Docker images and spin them up, just run
```bash
make
```

To bring it all down, run
```bash
make stop
```

## Using the API
```bash
make list-users
make create-user
make list-users
```

## Why should we use Golang?
Our Golang server has a much lighter memory footprint than our Spring Boot apps, which do more or less the same thing.

| Language | Framework   | Memory Footprint | Video (asciinema)                                        |
|:--------:|:-----------:|-----------------:|:--------------------------------------------------------:|
| Java     | Spring Boot | 500 MB           | [vid](https://asciinema.org/a/qlAlCexwOj3hygKDrSE6noHpN) |
| Golang   | net/http    | 5 MB             | [vid](https://asciinema.org/a/72mpi0VXUF9K65oX5bZYqUWa1) |

See [this gist](https://gist.github.com/kevinmichaelchen/22ac37452979b05f78e99f775e249659)
for a fuller explanation. At bast, embedded Tomcat runs at 120MB, which is still 24x higher
than what Golang consumes. Out of the box, embedded Tomcat runs 100x higher than Golang.