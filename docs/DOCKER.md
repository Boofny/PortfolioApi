# Running GoLive with Docker

## Quick Start

Pull and run the image locally:
```bash
docker run -p 8000:8000 -e PORT=8000 boofny/golive-docker
```

Once running, the server is available at `http://localhost:8000`.

## Routes

| Method | Path | Description |
|--------|------|-------------|
| GET | [`/ping`](http://localhost:8000/ping) | Health check |
| GET | [`/user/{id}`](http://localhost:8000/user/1) | Fetch user by ID |
| POST | [`/posting`](http://localhost:8000/posting) | Submit JSON data |
| GET | [`/v1/ping`](http://localhost:8000/v1/ping) | Route group example |

**POST /posting**
```bash
curl -X POST http://localhost:8000/posting \
  -H "Content-Type: application/json" \
  -d '{"name": "john", "email": "johnBram@gmail.com"}'
```

## Docker Hub

Full image details available on [Docker Hub](https://hub.docker.com/repository/docker/boofny/golive-docker/general).
