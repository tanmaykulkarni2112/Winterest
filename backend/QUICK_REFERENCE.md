# Factory Pattern - Quick Reference

## TL;DR - Core Concepts

| Concept             | Purpose                          | Location                     |
| ------------------- | -------------------------------- | ---------------------------- |
| **Factory**         | Creates and manages all services | `factory/factory.go`         |
| **Interfaces**      | Define service contracts         | `factory/interfaces.go`      |
| **Implementations** | Concrete service implementations | `factory/*_service.go`       |
| **Handlers**        | HTTP request handlers            | `auth/handler/*.go`, `home/` |
| **Dependencies**    | Group of services for handlers   | `factory/interfaces.go`      |

## Service Interfaces at a Glance

```go
// Data persistence
type DataService interface {
    GetUser(username string) (Login, bool)
    UserExists(username string) bool
    SaveUser(username string, user Login) error
    UpdateUser(username string, user Login) error
    LoadUsersFromFile() error
    SaveUsersToFile() error
}

// Authentication & security
type AuthService interface {
    HashPassword(password string) (string, error)
    CheckPassword(password, hash string) bool
    GenerateToken(length int) string
    Authorize(username string, r *http.Request) error
}

// User operations
type UserService interface {
    RegisterUser(username, password string) error
    LoginUser(username, password string) (sessionToken, csrfToken string, err error)
    LogoutUser(username string) error
    GetUserBySessionToken(sessionToken string) (string, bool)
}
```

## Creating Handlers - Template

```go
// Step 1: Define handler struct
type MyHandler struct {
    service factory.SomeService
}

// Step 2: Constructor function
func NewMyHandler(service factory.SomeService) *MyHandler {
    return &MyHandler{service: service}
}

// Step 3: Implement ServeHTTP
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Use h.service to handle the request
}
```

## Using Handlers in main.go

```go
// Before: func-based (old way)
http.HandleFunc("/endpoint", handler.HandlerFunction)

// After: struct-based (new way)
http.Handle("/endpoint", handler.NewHandlerName(factory.GetService()))
```

## Common Handler Patterns

### Read & Parse Request

```go
var payload Model
err := json.NewDecoder(r.Body).Decode(&payload)
if err != nil {
    http.Error(w, "Invalid request", http.StatusBadRequest)
    return
}
```

### Call Service

```go
result, err := h.service.DoSomething(data)
if err != nil {
    http.Error(w, "Operation failed", http.StatusInternalServerError)
    return
}
```

### Return JSON Response

```go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(map[string]interface{}{
    "key": result,
})
```

## Adding a New Service - Checklist

- [ ] Define interface in `internal/factory/interfaces.go`
- [ ] Create implementation in `internal/factory/<service>_service.go`
- [ ] Add to Factory struct in `internal/factory/factory.go`
- [ ] Add getter method in Factory
- [ ] Update `HandlerDependencies` struct
- [ ] Create handler(s) that use the service
- [ ] Register handler(s) in `main.go`
- [ ] Add tests in `internal/factory/factory_test.go`

## Error Handling

### Define Custom Errors

```go
var (
    ErrCustomError = errors.New("descriptive message")
)
```

### Use in Service

```go
if condition {
    return ErrCustomError
}
```

### Handle in Handler

```go
err := h.service.Operation()
if err != nil {
    map[error]int{
        factory.ErrUserExists: http.StatusConflict,
        factory.ErrUnauthorized: http.StatusUnauthorized,
    }
}
```

## Testing Pattern

```go
// Create mock
mock := &MockService{}

// Create handler with mock
handler := NewHandler(mock)

// Test with handler
recorder := httptest.NewRecorder()
req := httptest.NewRequest("POST", "/test", body)
handler.ServeHTTP(recorder, req)

// Assert result
if recorder.Code != http.StatusOK {
    t.Fatalf("Expected 200, got %d", recorder.Code)
}
```

## File Organization

```
internal/factory/
├── interfaces.go         ← Service contracts
├── factory.go            ← Main Factory class
├── data_service.go       ← DataService impl
├── auth_service.go       ← AuthService impl
├── user_service.go       ← UserService impl
└── factory_test.go       ← Tests & mocks

cmd/
└── main.go              ← Factory initialization & wiring

internal/auth/handler/
├── registerUser.go      ← Refactored handlers
├── loginUser.go
├── logoutUser.go
└── projected.go
```

## Dependency Injection Flow

```
main.go
  ↓
Factory.NewFactory(config)
  ├→ Create DataService
  ├→ Create AuthService (inject DataService)
  ├→ Create UserService (inject DataService, AuthService)
  ↓
Service instances ready
  ↓
Create handlers, inject services
  ↓
Register HTTP routes
  ↓
Start server
```

## Common Mistakes to Avoid

| ❌ Wrong                     | ✅ Right                                |
| ---------------------------- | --------------------------------------- |
| `global var service`         | `service := factory.Get...()`           |
| Handler creates service      | Handler receives service in constructor |
| Mix HTTP & business logic    | Separate into handler + service         |
| Tightly coupled services     | Dependency injection via interfaces     |
| Direct global state access   | DataService for all data access         |
| Testing needs setup/teardown | Mock services for testing               |

## Key Files Summary

| File              | Purpose        | Key Classes                           |
| ----------------- | -------------- | ------------------------------------- |
| `interfaces.go`   | Contracts      | DataService, AuthService, UserService |
| `factory.go`      | Creation       | Factory                               |
| `data_service.go` | Persistence    | DataServiceImpl                       |
| `auth_service.go` | Security       | AuthServiceImpl                       |
| `user_service.go` | Business logic | UserServiceImpl                       |
| `factory_test.go` | Testing        | MockDataService, test functions       |
| `cmd/main.go`     | Setup          | Factory init, handler registration    |

## Migration Path (If Needed)

### Phase 1: Parallel Implementation ✓ DONE

- Keep old code, add factory alongside
- New handlers use factory
- Old handlers still work

### Phase 2: Gradual Migration (Next)

- Convert remaining dependencies to factory
- Add tests as you go
- Update documentation

### Phase 3: Cleanup (Future)

- Remove old code when fully migrated
- Consolidate duplicate logic
- Optimize performance

## Performance Notes

✅ Factory pattern adds no performance overhead

- Services created once at startup
- No repeated instantiation
- Mutex-protected data access is efficient
- Token generation is fast

## Scaling Notes

The factory pattern scales well:

- **More services?** Add to factory
- **New handlers?** Follow same pattern
- **Different storage?** Swap DataService impl
- **Multiple environments?** Factory can accept config
- **Testing?** Mock entire service layer

## Quick Start: Add New Handler

```bash
# 1. Define the handler
cat > internal/auth/handler/newHandler.go << 'EOF'
package handler

type NewHandler struct {
    service factory.SomeService
}

func NewNewHandler(service factory.SomeService) *NewHandler {
    return &NewHandler{service: service}
}

func (h *NewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Implementation
}
EOF

# 2. Use in main.go
# handler := handler.NewNewHandler(factory.GetSomeService())
# http.Handle("/endpoint", handler)
```

## Getting Help

1. **Understand flow**: Read `FACTORY_PATTERN_GUIDE.md`
2. **See architecture**: Check `ARCHITECTURE_DIAGRAM.md`
3. **Examples**: Look at `EXTENSION_EXAMPLES.md`
4. **Reference**: Use this `QUICK_REFERENCE.md`
5. **See implementation**: Check actual files in `internal/factory/`

## Architecture at a Glance

```
HTTP Request
    ↓
Handler (HTTP logic only)
    ↓
Service Layer (calls appropriate service)
    ├→ UserService (orchestrates user ops)
    ├→ AuthService (password, token, auth)
    └→ DataService (persistence)
    ↓
Data Storage (file/database)
```
