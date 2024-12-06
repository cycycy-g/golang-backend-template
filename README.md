# Go Project Template

A template for Go projects with Docker, PostgreSQL, SQLC, and common development tools.

## Prerequisites

- Go 1.22 or later
- Docker and Docker Compose
- Make
- [SQLC](https://sqlc.dev/) (`go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`)

## Getting Started

1. Clone the template:
```bash
git clone <repository_url> your_project_name
cd your_project_name
```
2. Initialize Go module:
```bash
# Replace 'your_project_name' with your desired module name
go mod init your_project_name

# Install dependencies
go mod tidy
```

3. Configure environment:
```bash
# Copy example env file
cp app.env.example app.env

# Edit app.env with your configuration
```

4. Start services:
```bash
# Start all services with docker-compose
make up

# Run database migrations
make migrateup

# Generate SQLC code
make sqlc
```

## Project Structure
```
.
├── cmd/
│   └── main.go          # Application entry point
├── config/
│   └── config.go        # Configuration management
├── db/
│   ├── migrations/      # Database migrations
│   ├── queries/         # SQLC queries
│   └── sqlc/           # Generated SQLC code
├── docker/
│   ├── Dockerfile
│   └── wait-for.sh
├── internal/
│   └── server/         # Server implementation
├── app.env             # Environment configuration
├── docker-compose.yaml
├── sqlc.yaml           # SQLC configuration
└── Makefile
```

## Database and SQLC
### SQLC Configuration
The project uses SQLC for type-safe SQL queries. Configuration in `sqlc.yaml`:
```yaml
version: "2"
sql:
  - engine: "postgresql"
    schema: "sqlc/migrations"
    queries: "sqlc/queries"
    gen:
      go:
        package: "db"
        out: "db"
        sql_package: "pgx/v5"
        emit_json_tags: true
        emit_interface: true
        emit_empty_slices: true
        overrides:
          - db_type: "text[]"
            go_type: "github.com/lib/pq.StringArray"
          - db_type: "uuid"
            go_type: "github.com/google/uuid.UUID"
          - db_type: "timestamptz"
            go_type: "time.Time"
```
### Writing Queries
1. Create SQL migrations in db/migrations/
2. Write queries in db/queries/
3. Generate code:
```bash
make sqlc
```
Example query file (db/queries/users.sql):
```sql
-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;
```

### Using Generated Code
```go
import "your_project_name/db"

// Create a new queries object
queries := db.New(dbConn)

// Use generated methods
user, err := queries.GetUser(ctx, userID)
users, err := queries.ListUsers(ctx)
```

## Available Commands
### Docker Commands
```bash
make up      # Start all services
make down    # Stop all services
make logs    # View service logs
```
### Database Commands
```bash
make postgres    # Start PostgreSQL container
make createdb    # Create new database
make dropdb      # Drop database
```
### Migration Commands
```bash
make migrateup     # Run all migrations
make migratedown   # Revert all migrations
make migrateup1    # Run one migration forward
make migratedown1  # Revert one migration
```
### Development Commands
```bash
make server    # Run server locally
make test      # Run tests
make mock      # Generate mocks
make sqlc      # Generate SQLC code
```

## Configuration
### Environment Variables
Default configuration in app.env:
```env
# Database
DB_DRIVER=postgres
DB_SOURCE=postgresql://user:password@localhost:5432/dbname?sslmode=disable

# Server
SERVER_ADDRESS=:8080
ENVIRONMENT=development

# JWT
JWT_SECRET=your-secret-key
JWT_DURATION=24
```

## Troubleshooting
### Common Issues
1. Module Initialization

   * If you see import errors, ensure you've run go mod init and go mod tidy
   * Check that import paths match your module name

2. Container Issues
   * Port 5432 already in use:
   ```bash
   # Stop existing container
   docker stop postgres_container
   # Remove container
   docker rm postgres_container
    ```
3. Database Connection
   * Verify PostgreSQL is running: docker ps
   * Check logs: make logs
   * Verify credentials in app.env
4. SQLC Issues
   * Ensure SQLC is installed: `sqlc version`
   * Verify migration files are in correct format
   * Check query syntax in .sql files

## Development Notes

* SSL mode is disabled by default for local development
* Environment variables take precedence over app.env
* Database credentials can be modified through environment variables or app.env
* SQLC-generated code should not be edited manually

## Development Workflow

1. Create database migrations in `db/migrations/`
2. Write SQL queries in` `db/queries/`
3. Generate SQLC code: `make sqlc`
4. Implement business logic using generated code
5. Run tests: `make test`
6. Start server: `make server`

## Contributing
1. Fork the repository
2. Create your feature branch: git checkout -b feature/my-feature
3. Commit your changes: git commit -am 'Add my feature'
4. Push to the branch: git push origin feature/my-feature
5. Submit a pull request