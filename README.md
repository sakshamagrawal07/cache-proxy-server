<!-- <img src="banner.jpg" alt="Description" style="width:100%;"> -->

## Caching Proxy Server
Cache Proxy is a simple, bare-bones caching proxy server implemented in Go. It serves as a basic solution for caching HTTP requests, using an in-memory database, redis, to store the cache.

## Features
- Caching:
  - The server uses an in-memory database to store cached responses. This approach allows for fast lookups and reduces the time needed to retrieve cached data.

- Proxy Functionality:
  - The server forwards client requests to the target server and caches the responses. If the same request is made again, the server returns the cached response, saving the time and resources of making a new request to the target server.

## Usage

### Start the server:

```bash
go run main.go --port <PORT> --origin <ORIGIN_URL>
```
### Start the redis-stack-server:

```bash
docker run -d --name redis-stack-server -p 6379:6379 redis/redis-stack-server:latest
```

### To clear the cache, start the redis server and run the following command:

```bash
go run main.go --clear-cache
```

### To view the available commands:

```bash
go run main.go -h
```

## Extras
This Repo serves as a solution to [Roadmap.sh Caching Server Problem](https://roadmap.sh/projects/caching-server)