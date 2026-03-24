# Factory Pattern Architecture Diagram

## Complete Architecture Flow

```
┌─────────────────────────────────────────────────────────────┐
│                       main.go                               │
│                                                             │
│  • Initialize Factory("data/users.json")                   │
│  • Load users from file                                    │
│  • Create handlers with dependencies                       │
│  • Register HTTP routes                                    │
│  • Start server on :8080                                   │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
        ┌──────────────────────────────┐
        │   Factory (Factory.go)       │
        │                              │
        │  ┌────────────────────┐     │
        │  │  DataService       │     │
        │  │  AuthService       │     │
        │  │  UserService       │     │
        │  │  Dependencies      │     │
        │  └────────────────────┘     │
        └──────────────┬───────────────┘
                       │
        ┌──────────────┼──────────────────────────────────────┐
        │              │                                      │
        ▼              ▼                                      ▼
┌────────────────┐ ┌────────────────┐ ┌────────────────┐ ┌──────────────┐
│ DataService    │ │ AuthService    │ │ UserService    │ │ Dependencies │
│                │ │                │ │                │ │              │
│ • GetUser()    │ │ • HashPassword │ │ • Register()   │ │ Groups:      │
│ • SaveUser()   │ │ • CheckPassword│ │ • Login()      │ │ • Data       │
│ • UserExists() │ │ • GenerateToken│ │ • Logout()     │ │ • Auth       │
│ • Update()     │ │ • Authorize()  │ │ • GetByToken() │ │ • User       │
│ • LoadFromFile │ └────────────────┘ │                │ └──────────────┘
│ • SaveToFile   │                   └────────────────┘
└────────────────┘
        │                ▲
        │                │
        └────────────────┘
      (Uses)       (Uses)

                       ▲
                       │
        ┌──────────────┼─────────────────────────────────┐
        │              │                                 │
        ▼              ▼                                 ▼
┌──────────────────┐ ┌──────────────────┐ ┌──────────────────┐
│   Handlers       │ │   Handlers       │ │   Handlers       │
│                  │ │                  │ │                  │
│ RegisterHandler  │ │ LoginHandler     │ │ LogoutHandler    │
│ • UserService   │ │ • UserService    │ │ • UserService    │
│                  │ │ • AuthService    │ │ • AuthService    │
│ HomeHandler      │ │                  │ │                  │
│ (no deps)       │ │ ProtectedHandler │ │                  │
│                  │ │ • AuthService    │ │                  │
└──────────────────┘ └──────────────────┘ └──────────────────┘
        │                     │                     │
        └─────────────────────┼─────────────────────┘
                              │
                              ▼
                    ┌──────────────────┐
                    │   HTTP Routes    │
                    │                  │
                    │ /home            │
                    │ /register        │
                    │ /login           │
                    │ /logout          │
                    │ /protected       │
                    └──────────────────┘
```

## Service Layer Architecture

```
┌─────────────────────────────────────────────────────────┐
│                   HTTP Handlers                         │
│                                                         │
│  • RegisterUserHandler                                  │
│  • LoginUserHandler                                     │
│  • LogoutUserHandler                                    │
│  • ProtectedHandler                                     │
│  • HomeHandler                                          │
└────────────────────┬────────────────────────────────────┘
                     │ (Uses)
                     ▼
┌─────────────────────────────────────────────────────────┐
│              Service Interfaces                         │
│                                                         │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐ │
│  │ UserService  │  │ AuthService  │  │DataService   │ │
│  └──────────────┘  └──────────────┘  └──────────────┘ │
└────────────────────┬────────────────────────────────────┘
                     │ (Implements)
                     ▼
┌─────────────────────────────────────────────────────────┐
│            Service Implementations                      │
│                                                         │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐ │
│  │UserService   │  │AuthService   │  │DataService   │ │
│  │Impl          │  │Impl          │  │Impl          │ │
│  │              │  │              │  │              │ │
│  │Orchestrates: │  │Manages:      │  │Handles:      │ │
│  │• Register    │  │• Passwords   │  │• User CRUD   │ │
│  │• Login       │  │• Tokens      │  │• File I/O    │ │
│  │• Logout      │  │• Authorization  │• Persistence │ │
│  └──────────────┘  └──────────────┘  └──────────────┘ │
└────────────────────┬────────────────────────────────────┘
                     │ (Persists to)
                     ▼
┌─────────────────────────────────────────────────────────┐
│           Data Layer (File Storage)                     │
│                                                         │
│  data/users.json                                        │
│  {                                                      │
│    "username": {                                        │
│      "HashPassword": "...",                             │
│      "SessionToken": "...",                             │
│      "CSRFToken": "..."                                 │
│    }                                                    │
│  }                                                      │
└─────────────────────────────────────────────────────────┘
```

## Dependency Injection Flow

```
Factory.NewFactory(filePath)
    │
    ├─→ NewDataService(filePath)
    │       Initializes with:
    │       - In-memory map: map[string]Login
    │       - Mutex for thread safety
    │       - File path reference
    │
    ├─→ NewAuthService(dataService)
    │       Initializes with:
    │       - Reference to DataService
    │
    ├─→ NewUserService(dataService, authService)
    │       Initializes with:
    │       - Reference to DataService
    │       - Reference to AuthService
    │
    └─→ Create HandlerDependencies
            - DataService
            - AuthService
            - UserService

main.go receives Factory
    │
    ├─→ Get services
    │       authService := factory.GetAuthService()
    │       userService := factory.GetUserService()
    │
    └─→ Create handlers
            handler := NewRegisterUserHandler(userService)
            Inject UserService into handler
```

## Data Flow: User Registration

```
1. HTTP Request
   POST /register
   └─→ RegisterUserHandler.ServeHTTP()
       │
       ├─ Parse JSON body
       └─ Extract username & password
           │
           ▼
2. Handler delegates to UserService
   userService.RegisterUser(username, password)
           │
           ├─ Validate length (min 8 chars)
           ├─ Check if user exists via DataService
           │
           ▼
3. AuthService generates hash
   authService.HashPassword(password)
   └─ bcrypt hash generation
           │
           ▼
4. DataService persists user
   dataService.SaveUser(username, user)
           │
           ├─ Lock mutex
           ├─ Update in-memory map
           ├─ Marshal to JSON
           └─ Write to file
           │
           ▼
5. Handler returns response
   200 OK + JSON response
```

## Key Patterns

### Pattern 1: Handler Creation

```go
// In main.go
handler := handler.NewRegisterUserHandler(factory.GetUserService())
http.Handle("/register", handler)
```

### Pattern 2: Service Usage

```go
// In handler
func (h *RegisterUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    err := h.userService.RegisterUser(username, password)
    if err != nil {
        // Handle specific error types
    }
}
```

### Pattern 3: Service Composition

```go
// UserService uses other services
type UserServiceImpl struct {
    dataService DataService  // For persistence
    authService AuthService  // For hashing
}
```

## Benefits of This Architecture

✅ **Separation of Concerns**

- Handlers: HTTP logic only
- Services: Business logic
- Data: Persistence logic

✅ **Testability**

- Mock services for unit testing
- No global state to manage
- Easy to test in isolation

✅ **Maintainability**

- Clear dependency flow
- Easy to understand relationships
- Simple to add new services

✅ **Scalability**

- Easy to swap implementations
- Add new handlers following pattern
- Central factory for consistency

✅ **Concurrency Safety**

- Mutex protection in DataService
- No race conditions
- Thread-safe operations
