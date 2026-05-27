# GoLive

A lightweight HTTP framework built in Go.

## Live API

Base URL:
```
https://golive-production.up.railway.app subject to change
```

### Routes

| Method | Path | Description |
|--------|------|-------------|
| GET | `/ping` | Health check |
| GET | `/user/{id}` | Fetch user by ID |
| POST | `/posting` | Submit JSON data |
| GET | `/v1/ping` | Route group example |

**POST /posting**
```bash
curl -X POST https://golive-production.up.railway.app/posting \
  -H "Content-Type: application/json" \
  -d '{"name": "john", "email": "johnBram@gmail.com"}'
```

---

> For running locally with Docker, see [DOCKER.md](./DOCKER.md)
