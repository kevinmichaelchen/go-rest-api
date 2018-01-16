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

### Docker stats
See this [asciinema video](https://asciinema.org/a/72mpi0VXUF9K65oX5bZYqUWa1) for Docker stats for 500 INSERTs.
Notice that a Golang server sits at 4MB.

Contrast that with [our Spring Boot memory footprint for Rampart](https://asciinema.org/a/qlAlCexwOj3hygKDrSE6noHpN),
which sits at 500MB of memory.

That's 2 orders of magnitude difference!!!

Memory is not cheap!
