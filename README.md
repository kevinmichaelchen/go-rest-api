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

#### Golang
Our Golang server undergoes 500 POSTs to its `/user` endpoint.
This triggers 500 INSERTs into a MySQL table.

The server starts at just under 4MB and, under load, crawls up to just under 6MB.

#### Spring Boot
Contrast that with Rampart, our Spring Boot app,
which undergoes 500 POSTs to its `/person` endpoint and crawls up from 500MB to ~530MB.

With a [Gradle task](https://stackoverflow.com/a/38058671/1780216), we know Rampart only has 40MB of dependencies.

Embedded Tomcat is only supposed to be around 120MB, according to this [Java Code Geeks post](https://examples.javacodegeeks.com/enterprise-java/spring/tomcat-vs-jetty-vs-undertow-comparison-of-spring-boot-embedded-servlet-containers/).
Even if that is the case, Spring Boot still requires 24 times more memory.

This [post](https://www.marccostello.com/memory-analysis-of-a-spring-boot-application-in-docker-lessons-learnt/) explains how to use `-Xmx56m` to lower the memory footprint of Spring Boot but still ends up with 188.5MB of usage (about 38 times more memory than our Go server).

According to this [Spring Boot post](https://spring.io/blog/2015/12/10/spring-boot-memory-performance),
vanilla Spring Boot has a 6MB heap, but once you add Eureka it balloons to 80MB.
So adding Spring libraries can probably quickly cause memory footprint to rise.

#### Conclusion
Golang is 2 orders of magnitude lighter than Spring Boot.

We could run 100 replicas of a Go server before its memory exceeds that of its Spring Boot counterpart.

On my Macbook Pro with 16BG of memory, I can only run 3 replicas of Rampart with the default Docker memory usage of 2GB.
That's not enough to run a microservices setup on your local environment.
With Golang, you can run hundreds of replicas.

Memory is not cheap!
