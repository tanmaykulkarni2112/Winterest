```
Backend
│
├── cmd
│   └── server
│       └── main.go
│
├── internal
│   │
│   ├── auth
│   ├── user
│   ├── pin
│   ├── board
│   ├── feed
│   │
│   ├── middleware
│   ├── database
│   └── routes
```

## Current authentication functionality

### Register

data/user.json -> append the user and hashed password
-> assign the cookie to the user and store the cookie as well for validation?

### Logout

        --> Clear the cookie so user has to login eachtime

### Login

        --> Validate by reading the json -- users.json
        if user exists.. assign the cookie
