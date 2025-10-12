

# Mini REST Server

A simple REST API built in Go using only the standard library.  
It supports basic CRUD operations for user data stored in memory.

---

## Features
- **GET /users** – list all users  
- **GET /users/{id}** – get a user by ID  
- **POST /users** – create a new user  
- **DELETE /users/{id}** – delete a user by ID  
- Thread-safe operations using `sync.RWMutex`
- In-memory data storage (no database)

---

## Run the server
```bash
go run .
# or if located in cmd/server
go run ./cmd/server
```
The server will start on:
```
http://localhost:8080
```

---

## Example Usage

### Create a user
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Alice"}'
```

### Get a user
```bash
curl http://localhost:8080/users/1
```

### List all users
```bash
curl http://localhost:8080/users
```

### Delete a user
```bash
curl -X DELETE http://localhost:8080/users/1
```

---

## Example Responses

### `POST /users`
```json
{
  "id": 1,
  "user": {
    "name": "Alice"
  }
}
```

### `GET /users`
```json
[
  { "id": 1, "name": "Alice" },
  { "id": 2, "name": "Bob" }
]
```

---

## Development Notes
- Uses Go 1.22+ (new ServeMux with method patterns)
- Demonstrates JSON parsing (`encoding/json`)
- Shows concurrent-safe shared state using `sync.RWMutex`
- Ideal for learning Go’s `net/http` basics

---

## Future Improvements
- Add persistent storage (SQLite, PostgreSQL)
- Add PUT `/users/{id}` for updates
- Add unit tests for each endpoint
- Add graceful shutdown and structured logging