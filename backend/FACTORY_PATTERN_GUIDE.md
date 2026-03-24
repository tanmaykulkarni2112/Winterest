# Factory Pattern Refactoring Documentation

## Overview

The backend has been refactored to implement the **Factory Pattern**, a creational design pattern that provides an interface for creating objects without specifying their concrete classes. This refactoring improves code maintainability, testability, and follows SOLID principles.

## Architecture Changes

### Before Refactoring

- Tight coupling between handlers and global data stores
- Direct dependency on package-level variables (`data.Users`)
- Service functions scattered across multiple packages
- Difficult to test handlers in isolation
- Mixed concerns between HTTP handling and business logic

### After Refactoring

- **Service Layer**: Centralized business logic
- **Factory Pattern**: Single point for creating service instances
- **Dependency Injection**: Handlers receive their dependencies explicitly
- **Interfaces**: Service behavior defined through contracts
- **Testability**: Easy to mock services for testing

## Directory Structure

```
internal/
├── factory/                    # NEW: Factory and service implementations
│   ├── interfaces.go          # Service interfaces and dependencies
│   ├── factory.go             # Main factory for creating services
│   ├── data_service.go        # DataService implementation
│   ├── auth_service.go        # AuthService implementation
│   └── user_service.go        # UserService implementation
├── auth/
│   ├── handler/               # HTTP handlers (refactored)
│   │   ├── registerUser.go
│   │   ├── loginUser.go
│   │   ├── logoutUser.go
│   │   └── projected.go
│   ├── model/
│   ├── service/               # OLD: Keep for backwards compatibility
│   └── utils/
└── home/
    └── handler.go             # Refactored with factory pattern
```

## Service Interfaces

### DataService

Handles all user data persistence:

```go
type DataService interface {
    GetUser(username string) (Login, bool)
    UserExists(username string) bool
    SaveUser(username string, user Login) error
    LoadUsersFromFile() error
    SaveUsersToFile() error
    UpdateUser(username string, user Login) error
}
```

### AuthService

Handles authentication and security operations:

```go
type AuthService interface {
    HashPassword(password string) (string, error)
    CheckPassword(password, hash string) bool
    GenerateToken(length int) string
    Authorize(username string, r *http.Request) error
}
```

### UserService

Coordinates business logic for user operations:

```go
type UserService interface {
    RegisterUser(username, password string) error
    LoginUser(username, password string) (sessionToken, csrfToken string, err error)
    LogoutUser(username string) error
    GetUserBySessionToken(sessionToken string) (string, bool)
}
```

## Factory Implementation

The `Factory` struct creates and manages all service instances:

```go
type Factory struct {
    dataService DataService
    authService AuthService
    userService UserService
    deps        *HandlerDependencies
}

// Create factory
factory := factory.NewFactory("data/users.json")

// Initialize (load from file)
factory.Initialize()

// Get specific services
authService := factory.GetAuthService()
userService := factory.GetUserService()
dataService := factory.GetDataService()
```

## Handler Pattern

All handlers now follow a consistent pattern with dependency injection:

```go
type RegisterUserHandler struct {
    userService factory.UserService
}

func NewRegisterUserHandler(userService factory.UserService) *RegisterUserHandler {
    return &RegisterUserHandler{userService: userService}
}

func (h *RegisterUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Handler implementation
}
```

## Main.go Example

```go
func main() {
    // Initialize factory
    factory := factory.NewFactory("data/users.json")
    factory.Initialize()

    // Create handlers with dependencies
    registerHandler := handler.NewRegisterUserHandler(factory.GetUserService())
    loginHandler := handler.NewLoginUserHandler(
        factory.GetUserService(),
        factory.GetAuthService(),
    )

    // Register handlers
    http.Handle("/register", registerHandler)
    http.Handle("/login", loginHandler)

    http.ListenAndServe(":8080", nil)
}
```

## Benefits

### 1. **Testability**

Mock services can be injected for unit testing:

```go
type MockUserService struct {}

func (m *MockUserService) RegisterUser(u, p string) error {
    // Mock implementation
}

// In tests:
handler := handler.NewRegisterUserHandler(mockUserService)
```

### 2. **Loose Coupling**

Handlers don't depend on implementation details, only interfaces:

- Easy to swap implementations
- Change data storage without modifying handlers
- Add new auth methods independently

### 3. **Single Responsibility**

Each service has one clear purpose:

- `DataService`: Persistence
- `AuthService`: Security/Tokens
- `UserService`: User operations coordination

### 4. **Maintainability**

- Clear dependency flow
- Service creation centralized in factory
- Easy to trace code paths

### 5. **Scalability**

- Add new services easily
- Implement new handlers following same pattern
- Cache services at factory level

## Error Handling

Services return well-defined errors:

```go
var (
    ErrUserExists      = errors.New("user already exists")
    ErrInvalidPassword = errors.New("invalid username or password")
    ErrUnauthorized    = errors.New("unauthorized")
)
```

Handlers can map errors to appropriate HTTP status codes:

```go
err := userService.RegisterUser(u, p)
if err != nil {
    switch err {
    case factory.ErrInvalidLength:
        http.Error(w, "Invalid user/password", http.StatusNotAcceptable)
    case factory.ErrUserExists:
        http.Error(w, "User exists", http.StatusConflict)
    }
}
```

## Migration Guide

If you have existing code using old patterns:

### Old Pattern

```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    user := data.Users[username]
    hash, _ := utils.HashPassword(password)
}
```

### New Pattern

```go
func NewMyHandler(userService factory.UserService) *MyHandler {
    return &MyHandler{userService: userService}
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    _, err := h.userService.RegisterUser(username, password)
}
```

## Next Steps

1. **Add more services**: Implement additional business logic as services
2. **Configuration management**: Move file paths to configuration
3. **Logging**: Integrate structured logging across services
4. **Middleware**: Add authentication middleware using factory
5. **Database**: Replace file-based storage with database service

## References

- Factory Pattern: https://refactoring.guru/design-patterns/factory-method
- Dependency Injection: https://en.wikipedia.org/wiki/Dependency_injection
- SOLID Principles: https://en.wikipedia.org/wiki/SOLID
