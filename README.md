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

### Golang
Our Golang server undergoes 500 POSTs to its `/user` endpoint.
This triggers 500 INSERTs into a MySQL table.

The server starts at just under 4MB and, under load, crawls up to just under 6MB.

### Spring Boot
Contrast that with Rampart, our Spring Boot app,
which undergoes 500 POSTs to its `/person` endpoint and crawls up from 500MB to ~530MB.

With a [Gradle task](https://stackoverflow.com/a/38058671/1780216), we know Rampart only has 40MB of dependencies.

Embedded Tomcat is only supposed to be around 120MB, according to this [Java Code Geeks post](https://examples.javacodegeeks.com/enterprise-java/spring/tomcat-vs-jetty-vs-undertow-comparison-of-spring-boot-embedded-servlet-containers/).
Even if that is the case, Spring Boot still requires 24 times more memory.

This [post](https://www.marccostello.com/memory-analysis-of-a-spring-boot-application-in-docker-lessons-learnt/) explains how to use `-Xmx56m` to lower the memory footprint of Spring Boot but still ends up with 188.5MB of usage (about 38 times more memory than our Go server).

According to this [Spring Boot post](https://spring.io/blog/2015/12/10/spring-boot-memory-performance),
vanilla Spring Boot has a 6MB heap, but once you add Eureka it balloons to 80MB.
So adding Spring libraries can probably quickly cause memory footprint to rise.

### Conclusion
Golang is 2 orders of magnitude lighter than Spring Boot.

We could run 100 replicas of a Go server before its memory exceeds that of its Spring Boot counterpart.

#### It matters for local
It would be nice to run everything on our local.
However, my Macbook Pro with 16BG of memory can only run 3 replicas of Rampart with 
the default Docker memory usage of 2GB.
Even if I crank up Docker's consumption of host memory, my computer will crap out 
well before I can spin up 32 Spring Boot apps.
In short, I don't have enough memory to run a full-blown microservices setup on my local environment.
With Golang, I can run hundreds of replicas.

#### Memory is not cheap!

| Instance Type | $ per Year | Memory (GiB) |
|:-------------:|:----------:|:------------:|
| t2.nano       | $23        | 0.5          |
| t2.micro      | $46        | 1            |
| t2.small      | $92        | 2            |
| t2.medium     | $187       | 4            |
| t2.large      | $374       | 8            |
| t2.xlarge     | $748       | 16           |
| t2.2xlarge    | $1496      | 32           |

According to our µServices table, we plan on having at least 20 Spring Boot backends for
[MVP](https://en.wikipedia.org/wiki/Minimum_viable_product).
After MVP, we could easily be targeting 30 backends,
which is about 15GiB of memory consumption.

Acquiring 24GiB wouldn't be unreasonable.
That gives us a 9GiB cushion.
The total cost for that will be $1122 per year.

If we transitioned to Go backends, 30 REST APIs would have a base memory footprint of 
150MiB. That could easily be covered by a t2.micro for only $46 per year.

So I've just saved nearly $1000 per year.

That's not a big deal if we were planning to spend way more than that anyway.

However, it is a big deal when you take replication into account.

Our REST backends are stateless and can easily be scaled. 

The following table shows cost of various replica scenarios assuming 1 GiB costs $46 per year.
The replica scenarios assume that all backend services have equal number of replicas.
So the first row assumes that our 30 backend services each have 1 replica.

| # of replicas per API | Spring Boot Footprint (GiB) | Spring Boot Yearly Cost | Golang Footprint (GiB) | Golang Yearly Cost |
|:---------------------:|:---------------------------:|:-----------------------:|:----------------------:|:------------------:|
| 1                     | 15                          | $690                    | 0.15                   | $6.90              |
| 3                     | 45                          | $2070                   | 0.45                   | $20.70             |
| 5                     | 75                          | $3450                   | 0.75                   | $34.50             |
| 10                    | 150                         | $6900                   | 1.5                    | $69                |

Would you rather have 1 Spring Boot API, or 100 replicas of a Golang API?

You get the picture.