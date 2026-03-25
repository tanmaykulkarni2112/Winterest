# Postman Authentication API Test Guide

## Quick Start: Import the Collection

1. Open **Postman**
2. Click **File** → **Import** (or use the Import button)
3. Select the `postmanAuthTestCollection.json` file from `d:\Backend\backend\`
4. The collection will appear in your left sidebar

---

## Complete Test Execution Flow

### **Prerequisite: Start the Server**

```powershell
cd d:\Backend\backend
go run cmd/main.go
```

Server should be running on `http://localhost:8080`

---

## Test Sequence

### **REQUEST 1: Register User**

**Name:** `1. Register User`  
**Method:** `POST`  
**URL:** `http://localhost:8080/register`

**Headers:**

```
Content-Type: application/json
```

**Body (raw JSON):**

```json
{
  "username": "postmanuser",
  "password": "postmantest123"
}
```

**Expected Response:**

```
Status: 200 OK
Body: {"msg":"New user created"}
```

---

### **REQUEST 2: Login User**

**Name:** `2. Login User`  
**Method:** `POST`  
**URL:** `http://localhost:8080/login`

**Headers:**

```
Content-Type: application/json
```

**Body (raw JSON):**

```json
{
  "username": "postmanuser",
  "password": "postmantest123"
}
```

**Expected Response:**

```
Status: 200 OK
Body: User login successful
```

**⚠️ IMPORTANT - Check the Cookies:**

1. After this request, go to **Headers** tab → scroll down to see **Set-Cookie**
2. You should see:
   - `session_token=<long_token_value>`
   - `csrf_token=<long_token_value>`
3. **Copy the `csrf_token` value** - you'll need it for the next request

**Postman Automatically Saves:**

- Postman will automatically store these cookies in its cookie jar for localhost:8080
- These will be sent with subsequent requests automatically

---

### **REQUEST 3: Access Protected Resource**

**Name:** `3. Access Protected Resource`  
**Method:** `POST`  
**URL:** `http://localhost:8080/protected`

**Headers:**

```
Content-Type: application/json
X-CSRF-Token: <paste_csrf_token_from_login_here>
```

_Replace `<paste_csrf_token_from_login_here>` with the actual token from the Login response_

**Body (raw JSON):**

```json
{
  "username": "postmanuser",
  "password": "postmantest123"
}
```

**Expected Response:**

```
Status: 200 OK
Body: {"msg":"Access granted to protected resource","username":"postmanuser"}
```

**How It Works:**

- The `session_token` cookie is automatically sent by Postman (from the login response)
- You manually provide the `X-CSRF-Token` header
- Both are validated by the server before granting access

---

### **REQUEST 4: Logout User**

**Name:** `4. Logout User`  
**Method:** `POST`  
**URL:** `http://localhost:8080/logout`

**Headers:**

```
X-CSRF-Token: <paste_csrf_token_from_login_here>
```

**Body:** (empty - no request body needed)

**Expected Response:**

```
Status: 200 OK
Body: {"msg":"Logged out successfully"}
```

**After Logout:**

- Both `session_token` and `csrf_token` will be cleared from the database
- If you try to access `/protected` again, you'll get `401 Unauthorized`

---

### **REQUEST 5: Home Endpoint**

**Name:** `5. Home Endpoint`  
**Method:** `GET`  
**URL:** `http://localhost:8080/home`

**Headers:** (none required)

**Body:** (none)

**Expected Response:**

```
Status: 200 OK
Body: (home page content)
```

**Note:** This endpoint requires no authentication

---

## How Postman Handles Cookies

Postman automatically:

1. **Receives** cookies from `Set-Cookie` headers in responses
2. **Stores** them in the cookie jar for that domain (localhost:8080)
3. **Sends** cookies in subsequent requests (if the request matches the domain/path)

You can view/manage cookies:

- Click the **Cookies** button in Postman (next to the Send button)
- You'll see all cookies stored for localhost:8080

---

## Step-by-Step Walkthrough

### **Step 1: Register**

- Send Request 1 (Register User)
- Expected: 200 OK with success message

### **Step 2: Login**

- Send Request 2 (Login User)
- Expected: 200 OK with message and Set-Cookie headers
- **Action:** Copy the `csrf_token` value from the Set-Cookie header

### **Step 3: Try Protected Without Token** (Optional - to test error)

- Send Request 3 WITHOUT the X-CSRF-Token header
- Expected: 401 Unauthorized

### **Step 4: Protected With Token**

- Send Request 3 WITH the X-CSRF-Token header filled in
- Expected: 200 OK with access granted message

### **Step 5: Logout**

- Send Request 4 (Logout User) with the X-CSRF-Token header
- Expected: 200 OK with logout message

### **Step 6: Try Protected After Logout** (Optional - to test token cleared)

- Send Request 3 again
- Expected: 401 Unauthorized (because tokens were cleared)

---

## Adding Automated Tests (Optional)

If you want Postman to automatically validate responses, add test scripts:

### For Login Request (Request 2):

1. Click the **Tests** tab
2. Add this code:

```javascript
pm.test("Login successful", function () {
  pm.response.to.have.status(200);
  pm.response.to.include("User login successful");
});

// Store tokens in variables for later use
var cookies = pm.cookies.jar().cookies;
cookies.forEach(function (cookie) {
  if (cookie.name === "csrf_token") {
    pm.environment.set("csrf_token", cookie.value);
  }
});
```

### For Protected Request (Request 3):

```javascript
pm.test("Protected access successful", function () {
  pm.response.to.have.status(200);
  pm.response.to.include("Access granted");
});
```

### For Logout Request (Request 4):

```javascript
pm.test("Logout successful", function () {
  pm.response.to.have.status(200);
  pm.response.to.include("Logged out successfully");
});
```

---

## Troubleshooting

### "401 Unauthorized" on Protected Endpoint

- **Check:** Is the X-CSRF-Token header filled with the correct token from login?
- **Check:** Are you using Postman's cookie jar? (Cookies from login should be sent automatically)
- **Check:** Is the server still running?

### Cookies Not Showing

- **Solution:** Make sure you're checking the "Headers" tab in the response, not the main request headers
- **Solution:** The Set-Cookie headers appear in **Response Headers**, not request headers

### Can't Copy CSRF Token

- **Solution:** In PostMan, click on the response
- **Solution:** Go to **Headers** tab
- **Solution:** Look for `Set-Cookie: csrf_token=...`
- **Solution:** Copy the value between `csrf_token=` and the next semicolon

---

## Collection Variables (Optional Advanced)

The collection includes variables that can be used:

- `{{base_url}}` - http://localhost:8080
- `{{session_token}}` - For storing session token
- `{{csrf_token}}` - For storing CSRF token

You can reference these in URLs or headers like: `{{base_url}}/protected`

---

## Summary

| Request      | Method | Auth Required | Action                                |
| ------------ | ------ | ------------- | ------------------------------------- |
| 1. Register  | POST   | No            | Create new user                       |
| 2. Login     | POST   | No            | Get session & CSRF tokens             |
| 3. Protected | POST   | YES           | Requires session cookie + CSRF header |
| 4. Logout    | POST   | YES           | Clear tokens from database            |
| 5. Home      | GET    | No            | Public endpoint                       |

Tokens are valid for **24 minutes** from login time.
