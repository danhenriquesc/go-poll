# GO Poll

Simple voting application to deal with large amount of requests per second.

Inspired by 1x1 voting duels like Big Brother Brasil and The Voice Battle. (Can be easily used for N candidates)

# Post on Medium
https://hackernoon.com/making-a-web-voting-app-handling-245-million-votes-with-0-12-e59d5ec6a030

## Deps
- Docker Compose

## Go Deps
- github.com/gorilla/mux
- github.com/go-redis/redis

## Stack
- Golang 1.11
- Redis (with AOF persistence)

## Running Webserver

Just run:
```
docker-compose up
```

## Voting and get results

To vote in candidate 1:
```
POST http://localhost:8000/vote/1
```

To vote in candidate 2:
```
POST http://localhost:8000/vote/2
```

To get current result:
```
GET http://localhost:8000/vote
```

## Benchmarking

Benchmarks were made in DigitalOcean simple droplets with Ubuntu 16.04.4 x64

### $5/mo Droplet (~5800 req/sec)

```
RAM: 1 GB
CPU: 1 vCPU 1.8GHz
```

***Execution 1***
```
root@ubuntu-s-1vcpu-1gb-nyc1-01:~# wrk -t 1000 -c 10000 -d 1m -s ./post.lua http://68.183.137.138:8000/vote/2
Running 1m test @ http://68.183.137.138:8000/vote/2
  1000 threads and 10000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   607.99ms  157.48ms   1.99s    82.22%
    Req/Sec    24.63     22.40   292.00     75.02%
  345301 requests in 1.00m, 52.72MB read
Requests/sec:   5755.02
Transfer/sec:      0.88MB
```

***Execution 2***
```
root@ubuntu-s-1vcpu-1gb-nyc1-01:~# wrk -t 1000 -c 20000 -d 1m -s ./post.lua http://68.183.137.138:8000/vote/2
Running 1m test @ http://68.183.137.138:8000/vote/2
  1000 threads and 20000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   860.65ms  296.29ms   2.00s    65.63%
    Req/Sec    39.72     45.10   686.00     83.99%
  358947 requests in 1.00m, 54.94MB read
Requests/sec:   5982.46
Transfer/sec:      0.91MB
```

### $15/mo Droplet (~12000 req/sec)

```
RAM: 1 GB
CPU: 3 vCPUs 1.8GHz
```

***Execution 1***
```
root@ubuntu-s-3vcpu-1gb-nyc1-01:~# wrk -t 7500 -c 10000 -d 1m -s ./post.lua http://204.48.20.138:8000/vote/2                                                                                                                                   
Running 1m test @ http://204.48.20.138:8000/vote/2
  7500 threads and 10000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   631.39ms  107.29ms 798.00ms   91.50%
    Req/Sec     1.25      2.28   272.00     97.79%
  735058 requests in 1.01m, 112.16MB read
Requests/sec:  12091.66
Transfer/sec:      1.85MB
```

***Execution 2***
```
root@ubuntu-s-3vcpu-1gb-nyc1-01:~# wrk -t 1000 -c 13000 -d 1m -s ./post.lua http://204.48.20.138:8000/vote/1                                                                                                                                   
Running 1m test @ http://204.48.20.138:8000/vote/1
  1000 threads and 13000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.08s    89.31ms   1.91s    80.95%
    Req/Sec    22.67     29.20   650.00     87.31%
  714698 requests in 1.00m, 113.13MB read
Requests/sec:  11884.94
Transfer/sec:      1.88MB
```

***Execution 3***
```
root@ubuntu-s-3vcpu-1gb-nyc1-01:~# wrk -t 1000 -c 15000 -d 1m -s ./post.lua http://204.48.20.138:8000/vote/1                                                                                                                                   
Running 1m test @ http://204.48.20.138:8000/vote/1
  1000 threads and 15000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.32s   128.68ms   1.99s    77.60%
    Req/Sec    24.45     33.76     1.09k    85.97%
  671964 requests in 1.00m, 102.71MB read
Requests/sec:  11199.40
Transfer/sec:      1.71MB
```

### $80/mo Droplet (~68630 req/sec)

```
RAM: 16 GB
CPU: 6 vCPUs 1.8GHz
```

***Execution 1***
```
root@ubuntu-s-6vcpu-16gb-nyc1-01:~# wrk -t 1000 -c 70000 -d 1m -s ./post.lua http://204.48.20.138:8000/vote/1                                                                                                                                   
Running 1m test @ http://204.48.20.138:8000/vote/1
  1000 threads and 70000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   902.54ms  103.71ms   2.00s    78.21%
    Req/Sec    79.04     71.04     3.75k    84.72%
  4117831 requests in 1.01m, 637.73MB read
Requests/sec:  68630.52
Transfer/sec:     10.54MB
```

***Execution 2***
```
root@ubuntu-s-6vcpu-16gb-nyc1-01:~# wrk -t 1000 -c 30000 -d 1m -s ./post.lua http://204.48.20.138:8000/vote/1 
Running 1m test @ http://204.48.20.138:8000/vote/1
  1000 threads and 30000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   942.14ms  127.26ms   2.00s    85.55%
    Req/Sec    47.32     55.14     1.33k    88.25%
  1830829 requests in 1.00m, 281.99MB read
Requests/sec:  30513.82
Transfer/sec:      4.68MB
```

***Execution 3***
```
root@ubuntu-s-6vcpu-16gb-nyc1-01:~# wrk -t 1000 -c 20000 -d 1m -s ./post.lua http://204.48.20.138:8000/vote/2
Running 1m test @ http://204.48.20.138:8000/vote/2
  1000 threads and 20000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   656.35ms   89.49ms   1.61s    84.31%
    Req/Sec    40.59     41.40     1.13k    87.53%
  1828994 requests in 1.00m, 292.16MB read
Requests/sec:  30483.21
Transfer/sec:      4.85MB
```
