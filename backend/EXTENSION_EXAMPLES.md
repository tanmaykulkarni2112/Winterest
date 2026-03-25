# Factory Pattern - Extension Examples

## Adding a New Service

This guide shows how to extend the factory pattern to add new services.

### Example: Adding Email Service

#### 1. Define the Interface

In `internal/factory/interfaces.go`:

```go
// EmailService defines email operations
type EmailService interface {
    SendWelcomeEmail(username, email string) error
    SendPasswordResetEmail(username, email string) (resetToken string, err error)
    VerifyEmail(email string) error
}
```

#### 2. Implement the Service

Create `internal/factory/email_service.go`:

```go
package factory

import (
    "errors"
    "net/smtp"
)

// EmailServiceImpl implements EmailService interface
type EmailServiceImpl struct {
    dataService DataService
    smtpHost    string
    smtpPort    string
    fromEmail   string
    password    string
}

// NewEmailService creates a new EmailService
func NewEmailService(
    dataService DataService,
    smtpHost, smtpPort, fromEmail, password string,
) EmailService {
    return &EmailServiceImpl{
        dataService: dataService,
        smtpHost:    smtpHost,
        smtpPort:    smtpPort,
        fromEmail:   fromEmail,
        password:    password,
    }
}

// SendWelcomeEmail sends a welcome email to new users
func (es *EmailServiceImpl) SendWelcomeEmail(username, email string) error {
    // Implementation
    return nil
}

// SendPasswordResetEmail sends a reset email
func (es *EmailServiceImpl) SendPasswordResetEmail(username, email string) (string, error) {
    // Generate reset token from auth service
    resetToken := "generated_token"
    // Store in database
    // Send email
    return resetToken, nil
}

// VerifyEmail confirms email ownership
func (es *EmailServiceImpl) VerifyEmail(email string) error {
    // Implementation
    return nil
}
```

#### 3. Update Factory

In `internal/factory/factory.go`:

```go
type Factory struct {
    dataService  DataService
    authService  AuthService
    userService  UserService
    emailService EmailService  // Add this
    deps         *HandlerDependencies
}

func NewFactory(
    dataFilePath string,
    smtpHost, smtpPort, fromEmail, password string,
) *Factory {
    // ... existing code ...
    emailService := NewEmailService(
        dataService,
        smtpHost, smtpPort, fromEmail, password,
    )

    deps := &HandlerDependencies{
        DataService:  dataService,
        AuthService:  authService,
        UserService:  userService,
        EmailService: emailService,  // Add this
    }

    return &Factory{
        dataService:  dataService,
        authService:  authService,
        userService:  userService,
        emailService: emailService,  // Add this
        deps:         deps,
    }
}

func (f *Factory) GetEmailService() EmailService {
    return f.emailService
}
```

#### 4. Update HandlerDependencies

In `internal/factory/interfaces.go`:

```go
type HandlerDependencies struct {
    DataService  DataService
    AuthService  AuthService
    UserService  UserService
    EmailService EmailService  // Add this
}
```

#### 5. Use in Handler

Create `internal/auth/handler/emailHandler.go`:

```go
package handler

import (
    "encoding/json"
    "io"
    "net/http"

    "github.com/tanmaykulkarni2112/Winterest/backend/internal/auth/model"
    "github.com/tanmaykulkarni2112/Winterest/backend/internal/factory"
)

type VerifyEmailHandler struct {
    emailService factory.EmailService
    userService  factory.UserService
}

func NewVerifyEmailHandler(
    emailService factory.EmailService,
    userService factory.UserService,
) *VerifyEmailHandler {
    return &VerifyEmailHandler{
        emailService: emailService,
        userService:  userService,
    }
}

func (h *VerifyEmailHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
        return
    }

    var payload struct {
        Email string `json:"email"`
    }

    err := json.NewDecoder(r.Body).Decode(&payload)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    err = h.emailService.VerifyEmail(payload.Email)
    if err != nil {
        http.Error(w, "Verification failed", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "msg": "Email verified successfully",
    })
}
```

#### 6. Register in main.go

```go
func main() {
    factory := factory.NewFactory(
        "data/users.json",
        "smtp.gmail.com",
        "587",
        "your-email@gmail.com",
        "your-app-password",
    )

    // ... existing code ...

    emailService := factory.GetEmailService()
    verifyHandler := handler.NewVerifyEmailHandler(
        emailService,
        factory.GetUserService(),
    )

    http.Handle("/verify-email", verifyHandler)

    // ... rest of server setup ...
}
```

## Adding a New Handler

### Example: Password Reset Handler

#### Step 1: Create the Handler

```go
package handler

import (
    "encoding/json"
    "io"
    "net/http"

    "github.com/tanmaykulkarni2112/Winterest/backend/internal/factory"
)

type ResetPasswordHandler struct {
    userService  factory.UserService
    emailService factory.EmailService
}

func NewResetPasswordHandler(
    userService factory.UserService,
    emailService factory.EmailService,
) *ResetPasswordHandler {
    return &ResetPasswordHandler{
        userService:  userService,
        emailService: emailService,
    }
}

func (h *ResetPasswordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
        return
    }

    var payload struct {
        Email string `json:"email"`
    }

    err := json.NewDecoder(r.Body).Decode(&payload)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Send reset email
    resetToken, err := h.emailService.SendPasswordResetEmail("user", payload.Email)
    if err != nil {
        http.Error(w, "Failed to send reset email", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "msg": "Reset email sent",
        "token": resetToken,
    })
}
```

#### Step 2: Register in main.go

```go
resetHandler := handler.NewResetPasswordHandler(
    factory.GetUserService(),
    factory.GetEmailService(),
)
http.Handle("/reset-password", resetHandler)
```

## Adding a Middleware

### Example: Authentication Middleware

```go
package middleware

import (
    "net/http"

    "github.com/tanmaykulkarni2112/Winterest/backend/internal/factory"
)

// AuthMiddleware checks authentication
func AuthMiddleware(authService factory.AuthService) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            username := r.Header.Get("X-Username")
            if username == "" {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }

            err := authService.Authorize(username, r)
            if err != nil {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }

            next.ServeHTTP(w, r)
        })
    }
}
```

Usage in main.go:

```go
protectedHandler := handler.NewProtectedHandler(factory.GetAuthService())
wrappedHandler := middleware.AuthMiddleware(factory.GetAuthService())(protectedHandler)
http.Handle("/protected", wrappedHandler)
```

## Database Service Example

### Replacing File Storage with Database

Example of how to implement a DatabaseService:

```go
// internal/factory/db_service.go
package factory

import (
    "database/sql"
    "github.com/tanmaykulkarni2112/Winterest/backend/internal/auth/model"
)

// DatabaseServiceImpl implements DataService with a database
type DatabaseServiceImpl struct {
    db *sql.DB
}

func NewDatabaseService(db *sql.DB) DataService {
    return &DatabaseServiceImpl{db: db}
}

func (ds *DatabaseServiceImpl) GetUser(username string) (model.Login, bool) {
    var user model.Login
    err := ds.db.QueryRow(
        "SELECT hash_password, session_token, csrf_token FROM users WHERE username = ?",
        username,
    ).Scan(&user.HashPassword, &user.SessionToken, &user.CSRFToken)

    if err != nil {
        return model.Login{}, false
    }
    return user, true
}

// ... implement other methods ...
```

Update Factory factory to choose implementation:

```go
func NewFactory(useDatabase bool, ...) *Factory {
    var dataService DataService

    if useDatabase {
        db := connectToDatabase()
        dataService = NewDatabaseService(db)
    } else {
        dataService = NewDataService("data/users.json")
    }

    // Rest of factory initialization
}
```

## Testing Pattern

### Example: Testing New Service

```go
func TestEmailService(t *testing.T) {
    mockData := &MockDataService{
        Users: make(map[string]model.Login),
    }

    emailService := NewEmailService(
        mockData,
        "localhost", "1025", "test@test.com", "",
    )

    err := emailService.SendWelcomeEmail("testuser", "test@example.com")
    if err != nil {
        t.Fatalf("Failed: %v", err)
    }
}
```

## Key Patterns for Extension

1. **Define Interface First**: Start with interface in `interfaces.go`
2. **Implement Service**: Create implementation in new file
3. **Update Factory**: Add service creation to Factory
4. **Update Dependencies**: Add to HandlerDependencies
5. **Create Handler**: Create handler with injected service
6. **Register in main.go**: Wire everything together

## Avoiding Common Mistakes

❌ **Don't**: Create global variables for new services

```go
// WRONG!
var emailService EmailService
```

✅ **Do**: Pass through factory

```go
// CORRECT
factory.GetEmailService()
```

❌ **Don't**: Let handlers create services

```go
// WRONG!
func (h *Handler) ServeHTTP(...) {
    service := NewEmailService()  // Direct creation
}
```

✅ **Do**: Inject at creation time

```go
// CORRECT
NewHandler(factory.GetEmailService())
```

❌ **Don't**: Mix concerns in handlers

```go
// WRONG!
func handler() {
    // File I/O
    // Email sending
    // Authentication
}
```

✅ **Do**: Delegate to services

```go
// CORRECT
func handler() {
    userService.RegisterUser(...)
    emailService.SendEmail(...)
}
```
