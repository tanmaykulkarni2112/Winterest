# Authentication Module Documentation

## Overview

This authentication module implements a session-based authentication system with CSRF token protection. It handles user registration, login, logout, and protected resource access with secure password hashing and token validation.

---

## Folder Structure

```
auth/
├── handler/                    # HTTP request handlers
│   ├── registerUser.go         # User registration endpoint
│   ├── loginUser.go            # User login and token generation
│   ├── logoutUser.go           # User logout and token cleanup
│   └── projected.go            # Protected resource access
├── service/                    # Business logic layer
│   ├── authorize.go            # Token validation logic
│   ├── checkPassword.go        # Password verification
│   ├── generateToken.go        # Secure token generation
│   └── userExist.go            # User existence check
├── model/                      # Data structures
│   └── login.go                # User and request payload models
├── utils/                      # Utility functions
│   └── utils.go                # Password hashing utilities
└── README.md                   # This file
```

---

## Data Models

### `Login` Struct (in `model/login.go`)

```go
type Login struct {
    HashPassword string  // Bcrypt hashed password
    SessionToken string  // Session identifier
    CSRFToken    string  // Cross-Site Request Forgery token
}
```

### `RequestPayload` Struct (in `model/login.go`)

```go
type RequestPayload struct {
    Username string `json:"username"`
    Password string `json:"password"`
}
```

---

## Routes & Implementation

### 1. **Register User** - `POST /register`

#### Purpose

Creates a new user account with a hashed password stored in the database.

#### Request

```json
{
  "username": "testuser",
  "password": "password123456"
}
```

#### Response (Success)

```
Status: 200 OK
Body: {"msg":"New user created"}
```

#### Response (Error Cases)

```
Status: 406 Not Acceptable
Body: Invalid user / password
(Reason: Username or password less than 8 characters)

Status: 409 Conflict
Body: User already exist
(Reason: Username already registered)

Status: 500 Internal Server Error
Body: Failed to save user
(Reason: Database save failure)
```

#### Implementation Flow

```
1. Validate HTTP method (POST only)
2. Read request body
3. Unmarshal JSON into RequestPayload struct
4. Validate username & password length (minimum 8 characters)
5. Check if user already exists in data.Users map
6. Hash password using bcrypt
7. Create Login struct with hashed password
8. Save to in-memory map (data.Users)
9. Persist to users.json file
10. Return success message
```

#### Key Components

- **Validation**: Username and password length check
- **Security**: Bcrypt password hashing with cost factor 10
- **Storage**: Saves to in-memory map and persists to JSON
- **Conflict Handling**: Prevents duplicate usernames

---

### 2. **Login User** - `POST /login`

#### Purpose

Authenticates user credentials and generates session + CSRF tokens.

#### Request

```json
{
  "username": "testuser",
  "password": "password123456"
}
```

#### Response (Success)

```
Status: 200 OK
Body: User login successful
Headers:
  Set-Cookie: session_token=<random_token>; Expires=<+24 minutes>; HttpOnly=false
  Set-Cookie: csrf_token=<random_token>; Expires=<+24 minutes>; HttpOnly=false
```

#### Response (Error Cases)

```
Status: 401 Unauthorized
Body: Invalid user or password
(Reason: User doesn't exist OR password mismatch)
```

#### Implementation Flow

```
1. Validate HTTP method (POST only)
2. Read request body
3. Unmarshal JSON into RequestPayload struct
4. Check if user exists using service.UserExist()
5. Retrieve stored hash using data.Users[username]
6. Verify password using service.CheckPassword()
7. Generate 32-byte random session token
8. Generate 32-byte random CSRF token
9. Set both tokens as HTTP cookies (24-minute expiry)
10. Update user's tokens in data.Users map
11. Persist updated user to users.json file
12. Return success message
```

#### Token Generation

- **Algorithm**: Cryptographically secure random bytes
- **Length**: 32 bytes (Base64 encoded)
- **Expiry**: 24 minutes from login
- **Storage**: In-memory map + JSON file

#### Key Components

- **Password Verification**: Using bcrypt.CompareHashAndPassword
- **Token Generation**: Secure random token from service.GenerateToken
- **Session Storage**: Tokens stored in user struct in database
- **Cookie Management**: Two separate cookies with different HttpOnly flags

---

### 3. **Protected Resource** - `POST /protected`

#### Purpose

Validates session and CSRF tokens to grant access to protected endpoints.

#### Request

```
Method: POST
URL: http://localhost:8080/protected
Headers:
  Content-Type: application/json
  X-CSRF-Token: <csrf_token_from_login>
  Cookie: session_token=<session_token_from_login>
Body:
{
  "username": "testuser",
  "password": "password123456"
}
```

#### Response (Success)

```
Status: 200 OK
Body: {"msg":"Access granted to protected resource","username":"testuser"}
```

#### Response (Error Cases)

```
Status: 405 Method Not Allowed
Body: Cannot access
(Reason: Only POST method allowed)

Status: 401 Unauthorized
Body: Unauthorized
(Reason: Missing/invalid session token OR CSRF token mismatch)
```

#### Implementation Flow

```
1. Check HTTP method (POST only)
2. Read request body
3. Unmarshal JSON to get username
4. Call service.Authorize(username, request)
   a. Retrieve user from data.Users[username]
   b. Extract session_token from request cookies
   c. Compare with stored SessionToken (must match exactly)
   d. Extract X-CSRF-Token from request headers
   e. Compare with stored CSRFToken (must match exactly)
   f. Return error if any mismatch
5. If authorized, return success with username
6. Set response headers and status code
7. Encode response as JSON
```

#### Authorization Logic

```go
func Authorize(username string, r *http.Request) error {
    // Step 1: Get user from database
    user, ok := data.Users[username]
    if !ok {
        return AuthError
    }

    // Step 2: Validate session token from cookie
    st, err := r.Cookie("session_token")
    if err != nil || st.Value != user.SessionToken {
        return AuthError
    }

    // Step 3: Validate CSRF token from header
    csrf := r.Header.Get("X-CSRF-Token")
    if csrf != user.CSRFToken || csrf == "" {
        return AuthError
    }

    return nil
}
```

#### Key Components

- **Dual Token Validation**: Both session and CSRF tokens required
- **Source Separation**: Session in cookie, CSRF in header (CSRF protection)
- **String Comparison**: Exact token value matching
- **User Lookup**: Username passed in request body

---

### 4. **Logout User** - `POST /logout`

#### Purpose

Invalidates user tokens and clears them from both database and cookies.

#### Request

```
Method: POST
URL: http://localhost:8080/logout
Headers:
  X-CSRF-Token: <csrf_token_from_login>
  Cookie: session_token=<session_token_from_login>
Body: (empty)
```

#### Response (Success)

```
Status: 200 OK
Body: {"msg":"Logged out successfully"}
Headers:
  Set-Cookie: session_token=; Expires=<past_date>
  Set-Cookie: csrf_token=; Expires=<past_date>
```

#### Response (Error Cases)

```
Status: 401 Unauthorized
Body: Unauthorized
(Reason: Missing/invalid session token)

Status: 500 Internal Server Error
Body: Failed to logout
(Reason: Database save failure)
```

#### Implementation Flow

```
1. Extract session_token from request cookie
2. Loop through all users in data.Users to find matching session
3. If no match found, return 401 Unauthorized
4. Verify authorization using service.Authorize()
5. Clear both tokens from user struct:
   a. Set SessionToken = ""
   b. Set CSRFToken = ""
6. Update user in data.Users map
7. Persist changes to users.json file
8. Set both cookies with past expiry date (browser deletion)
9. Return success message
```

#### Token Clearing Strategy

```
Before Logout:
  SessionToken: "VAxqemnrl5XYPtugj8VHFcQ1vQ-LCHbQhkUIwKmy8Vo="
  CSRFToken: "lU8y4bS28ZICcZI7iLQNeDdFS9RxjUluedL4CLsNAHk="

After Logout:
  SessionToken: ""
  CSRFToken: ""
  (Database persisted)
```

#### Key Components

- **Token Lookup**: Finds user by session token
- **Dual Clearing**: Clears both tokens in database
- **Cookie Invalidation**: Sets expiry to past date
- **Database Persistence**: Changes saved to JSON file

---

## Service Layer Functions

### `authorize.go` - Token Validation

```go
func Authorize(username string, r *http.Request) error
```

- Validates both session and CSRF tokens
- Returns AuthError if validation fails
- Used by Protected and Logout endpoints

### `checkPassword.go` - Password Verification

```go
func CheckPassword(password string, hash string) bool
```

- Compares plaintext password with bcrypt hash
- Returns true if password matches
- Used by Login endpoint

### `generateToken.go` - Secure Token Generation

```go
func GenerateToken(length int) string
```

- Generates cryptographically secure random bytes
- Base64 encodes the result
- Used by Login endpoint for both session and CSRF tokens

### `userExist.go` - User Existence Check

```go
func UserExist(username string) bool
```

- Checks if username exists in data.Users map
- Used by Login endpoint

---

## Utility Functions

### `utils.go` - Password Hashing

```go
func HashPassword(password string) (string, error)
```

- Uses bcrypt with cost factor 10
- Returns hashed password string
- Used by Register endpoint

```go
func VerifyPassword(hash, password string) error
```

- Compares hash with plaintext password
- Part of CheckPassword service logic

---

## Authentication Flow Diagrams

### Registration Flow

```
User Request
    ↓
Validate Method (POST)
    ↓
Parse JSON Body
    ↓
Validate Length (≥8 chars)
    ↓
Check Username Exists?
    ├─ YES → Return 409 Conflict
    └─ NO → Continue
    ↓
Hash Password (Bcrypt)
    ↓
Create Login Struct
    ↓
Save to Memory (data.Users)
    ↓
Persist to JSON (users.json)
    ↓
Return 200 Success
```

### Login Flow

```
User Request
    ↓
Validate Method (POST)
    ↓
Parse JSON Body
    ↓
User Exists?
    ├─ NO → Return 401 Unauthorized
    └─ YES → Continue
    ↓
Password Correct?
    ├─ NO → Return 401 Unauthorized
    └─ YES → Continue
    ↓
Generate Session Token (32 bytes)
    ↓
Generate CSRF Token (32 bytes)
    ↓
Set Cookies (24 min expiry)
    ↓
Update User in Memory
    ↓
Persist to JSON
    ↓
Return 200 Success + Tokens in Cookies
```

### Protected Access Flow

```
User Request
    ↓
Validate Method (POST)
    ↓
Parse JSON Body
    ↓
Extract Username from Body
    ↓
Has Session Cookie?
    ├─ NO → Return 401 Unauthorized
    └─ YES → Continue
    ↓
Has X-CSRF-Token Header?
    ├─ NO → Return 401 Unauthorized
    └─ YES → Continue
    ↓
Get User from Database
    ├─ NOT FOUND → Return 401 Unauthorized
    └─ FOUND → Continue
    ↓
Session Token Matches?
    ├─ NO → Return 401 Unauthorized
    └─ YES → Continue
    ↓
CSRF Token Matches?
    ├─ NO → Return 401 Unauthorized
    └─ YES → Continue
    ↓
Return 200 Success + Resource
```

### Logout Flow

```
User Request
    ↓
Extract Session Token from Cookie
    ├─ NOT FOUND → Return 401 Unauthorized
    └─ FOUND → Continue
    ↓
Loop Users to Find Matching Session
    ├─ NOT FOUND → Return 401 Unauthorized
    └─ FOUND → Continue
    ↓
Authorize User (Same as Protected)
    ├─ FAILED → Return 401 Unauthorized
    └─ SUCCESS → Continue
    ↓
Clear SessionToken (set to "")
    ↓
Clear CSRFToken (set to "")
    ↓
Update User in Memory
    ↓
Persist to JSON
    ↓
Set Cookies with Past Expiry
    ↓
Return 200 Success
```

---

## Beginner Concepts Explained

### What is Authentication?

Authentication is the process of verifying that a user is who they claim to be. It typically involves checking a username and password.

### What is a Session Token?

A session token is a unique identifier assigned to a user after successful login. It proves the user is authenticated without needing to send the password again on every request.

**Example:**

```
First Request (Login):
  Username: john
  Password: secret123

Server Response:
  session_token: abc123def456...

Second Request (Protected Resource):
  Cookie: session_token=abc123def456...
  (No need to send password again)
```

### What is CSRF Protection?

CSRF (Cross-Site Request Forgery) is an attack where a malicious website tricks your browser into making unwanted requests to another site where you're logged in.

**Example Attack:**

```
You're logged into your bank (bank.com)
You visit evil.com
Evil.com secretly sends:
  POST /transfer HTTP/1.1
  Host: bank.com
  Cookie: session_token=your_token
  (Your browser automatically includes your session cookie)

Result: Your money gets transferred!
```

**CSRF Prevention:**
We use a CSRF token that:

1. Is generated on the server
2. Is sent to the client
3. Must be included in request headers (NOT automatically sent by browser)
4. Server validates both session token AND CSRF token

```
Protected Request:
  POST /transfer HTTP/1.1
  Host: bank.com
  Cookie: session_token=abc123... (automatic)
  X-CSRF-Token: xyz789... (must be manually added by legitimate client)

evil.com cannot:
  - Access the CSRF token (not in cookie, stored in header)
  - Trick your browser into sending it (browser doesn't auto-send custom headers)
```

### What is Password Hashing?

Hashing converts a password into a fixed-length string using a one-way algorithm. Even if someone accesses the database, they cannot reverse the hash to get the original password.

**Example:**

```
Original Password: MyPassword123
Hashed: $2a$10$BFGvjZBTd5YEU15dhEsiuH.CEHL95XO1jd...

What happens if database is stolen?
- Attacker sees the hash, not the password
- Hashing is one-way (cannot be reversed)
- Even same password → different hash (if using salt)
- Attacker must try millions of combinations (brute force)
```

### What is Bcrypt?

Bcrypt is a password hashing algorithm that:

- Is deliberately slow (prevents fast brute-force attacks)
- Uses salt (random data mixed with password)
- Uses cost factor (number of rounds = higher security)

```go
// Cost factor 10 = 2^10 = 1024 rounds
// More rounds = slower but more secure
cost := 10
hash, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
```

### What is HttpOnly Flag?

The HttpOnly flag on cookies prevents JavaScript from accessing the cookie, protecting against XSS (Cross-Site Scripting) attacks.

```
Without HttpOnly:
  Set-Cookie: token=abc123
  (JavaScript can access: document.cookie)

With HttpOnly:
  Set-Cookie: token=abc123; HttpOnly
  (JavaScript cannot access the token)
  (Browser automatically sends it with requests)
```

**In our implementation:**

```go
// session_token: HttpOnly=false (can be tested manually)
http.SetCookie(w, &http.Cookie{
    Name: "session_token",
    HttpOnly: false,
})

// csrf_token: HttpOnly=false (needs to be sent in header)
http.SetCookie(w, &http.Cookie{
    Name: "csrf_token",
    HttpOnly: false,
})
```

---

## Current Shortcomings & Improvements

### 1. **In-Memory Storage**

- **Issue**: All data stored in RAM, lost on server restart
- **Impact**: Users logged out when server restarts, no persistent history
- **Solution**: Use a proper database (PostgreSQL, MongoDB, MySQL)

### 2. **No Rate Limiting**

- **Issue**: No protection against brute force attacks
- **Impact**: Attacker can try millions of password combinations
- **Solution**: Implement rate limiting (max 5 login attempts per minute)

### 3. **No Token Refresh**

- **Issue**: Tokens expire but no way to refresh without re-login
- **Impact**: Poor user experience for long sessions
- **Solution**: Implement refresh tokens with longer expiry

### 4. **No HTTPS**

- **Issue**: Tokens sent over HTTP (plaintext)
- **Impact**: Tokens could be intercepted by man-in-the-middle attacks
- **Solution**: Always use HTTPS in production

### 5. **No Secure Cookie Flags**

- **Issue**: Cookies not marked as Secure or SameSite
- **Impact**: Could be sent over HTTP, vulnerable to CSRF
- **Solution**: Add Secure flag (HTTPS only) and SameSite=Strict

```go
// Should be:
http.SetCookie(w, &http.Cookie{
    Name: "session_token",
    Value: sessionToken,
    HttpOnly: true,  // Prevent JavaScript access
    Secure: true,    // HTTPS only
    SameSite: http.SameSiteStrictMode,  // Prevent CSRF
    Path: "/",
})
```

### 6. **No Login Audit Log**

- **Issue**: No record of who logged in when
- **Impact**: Cannot detect suspicious activity
- **Solution**: Log all authentication events to database

### 7. **Token Stored with User**

- **Issue**: If user object is exposed, tokens are compromised
- **Impact**: Tokens are sensitive data mixed with user data
- **Solution**: Store tokens in separate table with expiry time

### 8. **No Email Verification**

- **Issue**: Anyone can register with any username
- **Impact**: Users can impersonate or spam
- **Solution**: Send verification email before account activation

### 9. **Weak Password Requirements**

- **Issue**: Only checks minimum length (8 characters)
- **Impact**: Users can use weak passwords like "12345678"
- **Solution**: Require uppercase, lowercase, numbers, special characters

### 10. **No Account Lockout**

- **Issue**: No protection after multiple failed attempts
- **Impact**: Brute force attacks possible
- **Solution**: Temporarily lock account after 5 failed attempts

---

## Testing the Authentication System

### Using Postman

**1. Register**

```
POST /register
Body: {"username":"testuser","password":"password123456"}
Expected: 200 OK
```

**2. Login**

```
POST /login
Body: {"username":"testuser","password":"password123456"}
Expected: 200 OK, Set-Cookie headers with tokens
```

**3. Protected Access**

```
POST /protected
Headers: X-CSRF-Token: <token_from_login>
Cookie: session_token=<token_from_login>
Body: {"username":"testuser","password":"password123456"}
Expected: 200 OK, {"msg":"Access granted..."}
```

**4. Logout**

```
POST /logout
Headers: X-CSRF-Token: <token_from_login>
Cookie: session_token=<token_from_login>
Expected: 200 OK, tokens cleared
```

---

## Database Schema (Future Implementation)

```sql
-- Users Table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT true
);

-- Sessions Table
CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    session_token VARCHAR(255) UNIQUE NOT NULL,
    csrf_token VARCHAR(255) UNIQUE NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Audit Log Table
CREATE TABLE auth_audit (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    action VARCHAR(50),  -- 'login', 'logout', 'failed_login'
    ip_address VARCHAR(45),
    user_agent TEXT,
    success BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## Security Best Practices Used

✅ **Password Hashing**: Bcrypt with cost factor 10  
✅ **Dual Token System**: Session (cookie) + CSRF (header)  
✅ **Token Randomness**: 32 bytes cryptographically secure  
✅ **Token Validation**: Exact string matching on every request  
✅ **Logout Cleanup**: Tokens cleared from database  
✅ **HTTP Method Validation**: POST only for write operations  
✅ **Input Validation**: Length checks, username uniqueness

---

## Summary

This authentication module provides:

- User registration with secure password hashing
- Session-based login with token generation
- Protected resource access with dual token validation
- Secure logout with token invalidation

While functional for learning purposes, production systems should implement the improvements listed in the shortcomings section.
