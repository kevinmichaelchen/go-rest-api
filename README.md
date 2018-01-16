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

### Why should we use Golang?
Our Golang server has a much lighter memory footprint than our Spring Boot apps, which do more or less the same thing.

| Language | Framework   | Memory Footprint | Video (asciinema)                                        |
|:--------:|:-----------:|-----------------:|:--------------------------------------------------------:|
| Java     | Spring Boot | 500 MB           | [vid](https://asciinema.org/a/qlAlCexwOj3hygKDrSE6noHpN) |
| Golang   | net/http    | 5 MB             | [vid](https://asciinema.org/a/72mpi0VXUF9K65oX5bZYqUWa1) |

Our Golang server undergoes 500 POSTs to its `/user` endpoint.
This triggers 500 INSERTs into a MySQL table.

The server starts at just under 4MB and, under load, crawls up to just under 6MB.

Contrast that with Rampart, our Spring Boot app,
which undergoes 500 POSTs to its `/person` endpoint and crawls up from 500MB to ~530MB.

That's 2 orders of magnitude difference!!!

We could run 100 replicas of a Go server before its memory exceeds that of its Spring Boot counterpart.

Memory is not cheap!
