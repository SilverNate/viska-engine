# Viska Engine

A scalable, production-ready backend system using **Go (Golang)**, **GIN GONIC**, **gRPC**, **GORM**, **Redis**, **PostgreSQL**, and **Docker**, built with **Clean Architecture**, **SOLID Principles**, and **CI/CD with GitHub Actions**.

## 🔧 Tech Stack

| Layer         | Stack                                      |
|---------------|---------------------------------------------|
| Language      | Go (Golang) v1.21+                          |
| Framework     | Gin                                         |
| Database      | PostgreSQL, Redis                           |
| ORM           | GORM                                        |
| Communication | gRPC, Protobuf                              |
| Architecture  | Clean Architecture + SOLID principles       |
| Auth          | JWT-based auth with token validation via gRPC |
| Rate Limiting | Token Bucket + Redis                        |
| CI/CD         | GitHub Actions + Docker Compose             |
| Testing       | Testify, GoMock, Gherkin + Godog            |

---

## 🧱 Project Structure (Monorepo)

```bash
viska-engine/
│
├── proto/                        # Submodule for protobuf definitions
│   └── gen/                      # Generated gRPC stubs
│
├── auth-service/                # Authentication microservice
│   ├── cmd/                     # Main entry point
|   ├── config/                  # config
│   ├── internal/                # Business logic
│   │   ├── authentication/      # domain auth
│   │   ├── db/                  # database/migration
│   │   ├── middleware/          # middleware
│   │   └── grpc/                # gRPC server setup
│   ├── feature/                 # Gherkin Unit test
│   ├── Dockerfile
│   └── go.mod
│
├── data-service/                # Book, author, publisher microservice
│   ├── cmd/
│   ├── config/
│   ├── internal/
│   │   ├── database/
│   │   ├── library/             # domain library book, author, publisher
│   │   ├── redis/
│   │   ├── middleware/          # gRPC token validation
│   │   └── grpc/
│   ├── Dockerfile
│   └── go.mod
│
├── deployment/
|   ├── docker-compose.yml           # Global service orchestration
└── .github/workflows/ci.yml     # GitHub Actions pipeline


## 🚀 Getting Started

### 1. Clone with Submodules

```bash
git clone --recurse-submodules https://github.com/yourusername/viska-engine.git
cd viska-engine
```

---

### 2. Build and Run All Services

```bash
docker-compose up --build
```

- Auth Service → `localhost:8080`
- Data Service → `localhost:8081`
- PostgreSQL → `localhost:5432`
- Redis → `localhost:6379`

---

### 3. API Testing with Postman or Curl

#### ✅ Login

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email": "admin@viska.io", "password": "admin123"}'
```

#### 📚 Access Protected Endpoint

```bash
curl http://localhost:8081/api/books \
  -H "Authorization: Bearer <your_token>"
```

---

## 🔐 Token Validation (gRPC Flow)

1. Auth Service issues JWT during login.
2. Data Service intercepts requests via middleware.
3. Middleware extracts token from `Authorization` header.
4. Token is validated via gRPC call to `auth-service`.
5. If valid → continue request; otherwise → 401 Unauthorized.

---

## 🐳 Docker Compose Services

- **auth-service**: JWT login, token validation (gRPC)
- **data-service**: Protected endpoints with Redis cache + token validation
- **Postgres**: Shared DB
- **Redis**: For rate limiting and caching

---

## 🚥 Rate Limiting + Redis Caching

- **Rate Limiting** (Login):
  - `1 req/sec per client` using token bucket in Redis.

- **Caching**:
  - `GetBooks`, `GetAuthors`, `GetPublishers` are cached for 5 minutes.
  - Cache is invalidated on insert/update.

---

## 🧪 Testing

### Unit Tests (Testify + GoMock)

```bash
go test ./...
```

### Gherkin Integration Tests (Godog)

```bash
cd data-service
godog features/
```

Tests include:
- gRPC token validation
- Auth login and protected route access
- Redis-backed caching

---

## 🛠 Development

### Proto Generation

```bash
cd proto
protoc --go_out=. --go-grpc_out=. proto/auth.proto
```

> Or use [Buf](https://buf.build/) if configured.

---

## ⚙️ GitHub Actions (CI/CD)

### Location:
```bash
.github/workflows/ci.yml
```

### What it does:
- Installs dependencies
- Runs unit & integration tests
- Builds Docker images

### To enable Docker image push:
Set GitHub Secrets:
- `DOCKER_USERNAME`
- `DOCKER_PASSWORD`

---

## 📦 Build Without Docker Compose (Optional)

```bash
cd auth-service
go run cmd/main.go
```

---

## 📘 Tips

- Use `go work init` in monorepo root to manage workspaces:

```bash
go work init ./auth-service ./data-service
```

- Update `go.work` to include `replace` for proto:

```go
replace viska/proto => ./proto
```

---

## 🤝 Contributing

1. Fork
2. Create feature branch
3. Submit pull request

---

## 🧠 Author

Built with ❤️ by [@yourusername](https://github.com/yourusername)

```

---

Let me know if you want this saved as an actual file or added as a downloadable asset, or if you want a version in Bahasa Indonesia too!

