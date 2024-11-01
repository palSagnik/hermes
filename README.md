# Hermes

## Go Authentication Service Template 🔐

A production-ready authentication service template built with Go, featuring user registration, login, and session management. This template uses Docker for containerization and follows modern security practices.

## Features ✨

- User registration and login
- JWT-based authentication
- Password hashing with bcrypt
- Email verification
- Docker containerization
- RESTful API endpoints
- Rate limiting by IP
- Input validation


## Tech Stack 🛠

- Go 1.21+
- PostgreSQL
- Docker & Docker Compose
- JWT for token management
- Go-Fiber Web Framework

## Prerequisites 📋

- Docker and Docker Compose installed
- Go 1.21 or higher (for local development)
- Make (optional, for using Makefile commands)

## Quick Start 🚀

1. Clone the repository:
```bash
git clone https://github.com/palSagnik/hermes
cd go-auth-template
```

2. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your configurations
```

3. Start the services using Docker Compose:
```bash
docker-compose up -d
```

The service will be available at `http://localhost:9000`

## API Endpoints 🌐

### Authentication

```
POST /auth/signup
POST /auth/login
GET  /auth/verify/{token}
```
### Health Check
```
GET  /alive
```
### User Management

```
GET    /api/users
GET    /api/userdetails
```

## Environment Variables 🔧

```
export POSTGRES_USER="postgres"
export POSTGRES_PASSWORD="password"

export POSTGRES_HOST="postgres-db"
export POSTGRES_DATABASE="hermes"
export POSTGRES_SSLMODE="disable"

export APP_PORT="9000"

export PUBLIC_URL="localhost:9000"
export EMAIL_ID="s4ych33se.ctf@gmail.com"
export EMAIL_AUTH="API TOKEN"
export TOKEN_SECRET="SECRET"
```

## Directory Structure 📁

```
hermes/
├── Dockerfile
├── README.md
├── backend
│   ├── config
│   │   ├── db.go
│   │   ├── general.go
│   │   └── mailsecrets.go
│   ├── database
│   │   ├── database.go
│   │   ├── pgdata
│   │   │   └── global
│   │   ├── queries.go
│   │   └── schemas.go
│   ├── go.mod
│   ├── go.sum
│   ├── handler
│   │   ├── api.go
│   │   ├── auth.go
│   │   └── misc.go
│   ├── main.go
│   ├── middleware
│   │   ├── middleware.go
│   │   └── ratelimiter.go
│   ├── models
│   │   └── authModels.go
│   ├── router
│   │   └── routes.go
│   ├── template
│   │   └── mail.html
│   └── utils
│       ├── mail.go
│       └── utils.go
├── compose.yml
└── scripts
    ├── init.sh
    ├── reset.sh
    └── run.sh
```

## Development 💻

### Running Locally

```bash
# Install dependencies
go mod download

# Run the server
# Uncomment POSTGRES_HOST = localhost
./scripts/run.sh
```


## Docker Commands 🐳

```bash
# Run with Docker Compose
docker compose up -d

# View logs
docker compose logs -f
```

## Security Considerations 🔒

- Passwords are hashed using bcrypt
- JWT tokens for authentication
- Rate limiting on authentication endpoints
- Input validation and sanitization

## Contributing 🤝

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License 📝

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
