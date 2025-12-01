# Go GraphQL Authentication API

A complete Go application with GraphQL, PostgreSQL, and JWT authentication for user registration and login.

## Features

- User registration and login with JWT authentication
- GraphQL API with playground
- PostgreSQL database with user management
- Password hashing with bcrypt
- Middleware for authentication
- Environment-based configuration

## Prerequisites

- Go 1.21 or higher
- PostgreSQL database
- Git

## Setup

### 1. Clone and navigate to the project
```bash
cd last-go
```

### 2. Install dependencies
```bash
go mod tidy
```

### 3. Generate GraphQL code
```bash
go run github.com/99designs/gqlgen generate
```

### 4. Set up PostgreSQL database
Create a PostgreSQL database named `last_go` (or update the environment variables to match your setup).

### 5. Configure environment variables
Copy the example environment file and configure your settings:

```bash
cp env.example .env
```

Or set these environment variables manually:

```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=biruhken
export DB_PASSWORD=ayana
export DB_NAME=last_go
export PORT=8080
export JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
export JWT_EXPIRY_HOURS=24
export BCRYPT_COST=12
```

**Important**: Change the `JWT_SECRET` to a secure random string in production!

### 6. Run the application
```bash
go run main.go
```

The server will start on `http://localhost:8080` with the GraphQL playground available at the root URL.

## API Usage

### GraphQL Playground
Visit `http://localhost:8080/` to access the GraphQL playground where you can test the API.

### Register a new user
```graphql
mutation Register {
  register(input: {
    email: "user@example.com"
    username: "testuser"
    password: "password123"
  }) {
    user {
      id
      email
      username
      createdAt
    }
    token
  }
}
```

### Login
```graphql
mutation Login {
  login(input: {
    email: "user@example.com"
    password: "password123"
  }) {
    user {
      id
      email
      username
      createdAt
    }
    token
  }
}
```

### Get current user (requires authentication)
```graphql
query Me {
  me {
    id
    email
    username
    createdAt
    updatedAt
  }
}
```

### Get all users
```graphql
query Users {
  users {
    id
    email
    username
    createdAt
    updatedAt
  }
}
```

## Authentication

To access protected endpoints, include the JWT token in the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

## Project Structure

```
last-go/
├── graphql/
│   ├── schema.graphqls      # GraphQL schema definition
│   ├── resolvers.go         # GraphQL resolvers
│   ├── generated.go         # Generated GraphQL code
│   └── models_gen.go        # Generated models
├── internal/
│   ├── auth/
│   │   └── jwt.go           # JWT utilities
│   ├── database/
│   │   └── database.go      # Database operations
│   └── middleware/
│       └── auth.go          # Authentication middleware
├── go.mod                   # Go module file
├── gqlgen.yml              # GraphQL code generation config
├── main.go                 # Main application entry point
└── README.md               # This file
```

## Environment Variables

The application uses the following environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `DB_HOST` | `localhost` | PostgreSQL host |
| `DB_PORT` | `5432` | PostgreSQL port |
| `DB_USER` | `biruhken` | PostgreSQL username |
| `DB_PASSWORD` | `ayana` | PostgreSQL password |
| `DB_NAME` | `last_go` | PostgreSQL database name |
| `PORT` | `8080` | Server port |
| `JWT_SECRET` | `your-secret-key-change-this-in-production` | JWT signing secret |
| `JWT_EXPIRY_HOURS` | `24` | JWT token expiry time in hours |
| `BCRYPT_COST` | `10` | Bcrypt hashing cost |

## Security Notes

- **Always change the JWT_SECRET** to a secure random string in production
- Use environment variables for all sensitive configuration
- Consider adding rate limiting and input validation
- Use HTTPS in production
- The application will panic if JWT_SECRET is not set in production mode

## Development

To regenerate GraphQL code after schema changes:
```bash
go run github.com/99designs/gqlgen generate
```

## License

This project is open source and available under the MIT License.

