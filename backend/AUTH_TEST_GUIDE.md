# Authentication API Test Guide

## Using the Postman Collection

### Import the Collection

1. Open Postman
2. Click "Import" in the top left
3. Select the `postmanAuthTestCollection.json` file
4. The collection will be imported with all test endpoints

### Test Flow

#### Step 1: Register a New User (Optional)

- **Endpoint**: `POST /register`
- **Body**:
  ```json
  {
    "username": "testuser123",
    "password": "password123456"
  }
  ```
- **Expected Response**:
  - Status: `200 OK`
  - Body: `{"msg":"New user created"}`

#### Step 2: Login User

- **Endpoint**: `POST /login`
- **Body**:
  ```json
  {
    "username": "testuser123",
    "password": "password123456"
  }
  ```
- **Expected Response**:
  - Status: `200 OK`
  - Body: `User login successful`
  - **Cookies Set** (check Response Headers under "Set-Cookie"):
    - `session_token=<generated_token>`
    - `csrf_token=<generated_token>`

**Important**: Postman will automatically store these cookies for the localhost domain.

#### Step 3: Access Protected Resource

- **Endpoint**: `POST /protected`
- **Headers Required**:
  - `Content-Type: application/json`
  - `X-CSRF-Token: <csrf_token_from_login>` (Copy the csrf_token value from login response Set-Cookie header)
- **Body**:
  ```json
  {
    "username": "testuser123",
    "password": "password123456"
  }
  ```
- **Expected Response**:
  - Status: `200 OK`
  - Body: `{"msg":"Access granted to protected resource","username":"testuser123"}`
  - **Cookies Automatically Sent**: session_token (Postman stores in cookies jar)

#### Step 4: Logout User

- **Endpoint**: `POST /logout`
- **Headers Required**:
  - `X-CSRF-Token: <csrf_token_from_login>`
- **Expected Response**:
  - Status: `200 OK`
  - Body: `{"msg":"Logged out successfully"}`
  - **Cookies Cleared**:
    - `session_token` expires = past date
    - `csrf_token` expires = past date

#### Step 5: Home Endpoint (No Auth Required)

- **Endpoint**: `GET /home`
- **Expected Response**:
  - Status: `200 OK`
  - Returns home page content

---

## Testing Error Cases

### Missing CSRF Token on Protected Endpoint

- Try calling `/protected` WITHOUT the `X-CSRF-Token` header
- Expected: `401 Unauthorized`

### Using Expired/Invalid Session Token

- Modify the session_token cookie value to something random
- Try calling `/protected`
- Expected: `401 Unauthorized`

### Invalid Credentials on Login

- Use wrong username or password
- Expected: `401 Unauthorized` with message "Invalid user or password"

---

## Changes Made to Support Testing

### HttpOnly Flag Update

- **Before**: `session_token` had `HttpOnly: true` (prevents client access)
- **After**: `session_token` has `HttpOnly: false` (allows Postman/curl access)
- **csrf_token**: Always had `HttpOnly: false`

This change makes testing easier while still maintaining security through CSRF token validation.

---

## Testing with curl (Alternative)

If you prefer curl instead of Postman:

```bash
# 1. Login and save cookies
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser123","password":"password123456"}' \
  -c cookies.txt

# 2. Extract CSRF token from response and use it
curl -X POST http://localhost:8080/protected \
  -H "Content-Type: application/json" \
  -H "X-CSRF-Token: <csrf_token_value>" \
  -b cookies.txt \
  -d '{"username":"testuser123","password":"password123456"}'

# 3. Logout
curl -X POST http://localhost:8080/logout \
  -H "X-CSRF-Token: <csrf_token_value>" \
  -b cookies.txt
```

---

## Server Status

The server should be running on `http://localhost:8080`

Check if running:

```powershell
netstat -ano | findstr :8080
```

Start the server (if not running):

```powershell
cd d:\Backend\backend
go run cmd/main.go
```
