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
Our Golang server has a light memory footprint.
See this [asciinema video](https://asciinema.org/a/72mpi0VXUF9K65oX5bZYqUWa1) for Docker stats for 500 INSERTs.
The server starts at just under 4MB and, under load, crawls up to just under 6MB.

Contrast that with [our Spring Boot memory footprint for Rampart](https://asciinema.org/a/qlAlCexwOj3hygKDrSE6noHpN),
which undergoes 500 POSTs to its `/person` endpoint and crawls up from 500MB to ~530MB.

That's 2 orders of magnitude difference!!!

Memory is not cheap!
