# Pinterest-Style Backend (Go) — Project README

## Overview

This project implements the backend of a Pinterest-style application using the Go programming language. The goal is to practice backend architecture, REST API design, and scalable service organization while modeling a real-world social media system.

The backend exposes endpoints for user authentication, content creation (pins), collections (boards), discovery feeds, and social interactions such as saving, liking, and following.

This repository focuses on:

- REST API design
- Clean backend architecture
- Database modeling
- Image/media handling
- Feed and discovery systems

---

# Tech Stack

### Backend Language

- Go

### Web Framework

- Gin (recommended)

### Database

Choose one:

- PostgreSQL
  or
- MongoDB

### ORM / Database Library

- GORM
  or
- sqlx

### Image Storage

> Attempt will be to store the images on localstorage at first

Pinterest-style apps require external storage for media:

- Amazon S3
- Cloudinary

### Authentication

- JWT tokens
  Libraries:
- golang-jwt/jwt

---

# Core Features

The backend supports the following major features:

### Authentication

Users can create accounts and authenticate.

Capabilities:

- Register
- Login
- Logout
- Token refresh

---

### Users

Users represent creators and viewers on the platform.

Capabilities:

- View public profile
- Update profile
- View uploaded images
- View boards images uploaded by other user
- Follow / unfollow other users

---

### Images

Capabilities:

- Create images
- Delete images
- View images

Images can be:

- Public
- Private

Each pin contains:

- title
- description
- creator
- link (optional)

---

### Likes

Users can like images.

Capabilities:

- Like pin
- Unlike pin
- View pin like count

---

### Comments

Users can comment on pins.

Capabilities:

- Add comment
- View comments
- Delete comment

---

### Feed / Discovery

The homepage displays recommended images.

Capabilities:

- Randomized feed
- Search images by keyword

---

### Notifications (Future Implementation)

Users receive notifications for:

- New followers
- Likes
- Comments
- Saved pins

---

# API Resource Design

The backend API is structured around the following resource groups:

### Authentication

Handles account creation and login.

### Users

Handles user profiles and relationships.

### Images

Handles collections of image.

## Feed

Handles discovery and recommendation.

### Media Upload

Handles image uploads.

---

# Data Models

The system revolves around the following core entities:

### User

Represents a registered platform user.

Typical fields include:

- id
- username
- email
- password hash
- bio
- profile image
- created timestamp

---

### Image

Represents a content item.

Typical fields include:

- id
- title
- description
- image url
- creator id
- like count
- created timestamp

---

### Comment

Represents discussion on pins.

Typical fields include:

- id
- pin id
- user id
- content
- created timestamp

---

### Follow

Represents user relationships.

Typical fields include:

- follower id
- following id

---

# System Architecture

The backend follows a layered architecture:

### Router Layer

Defines API endpoints and groups routes.

### Handler Layer

Handles HTTP request/response processing.

### Service Layer

Contains business logic.

### Repository Layer

Handles database interaction.

### Model Layer

Defines application data structures.

This separation allows the system to remain maintainable as complexity grows.

---

# Project Structure

A typical project layout may look like:

```
cmd/
server/

internal/
handlers/
services/
repositories/
models/
routes/

pkg/
database/
config/
utils/
```

Responsibilities:

- **cmd** → entry point for application
- **internal** → core application logic
- **pkg** → reusable packages
- **database** → database connection setup

---

# Feed System

The feed is a key feature.

> FUTURE SCOPE
> Possible ranking factors:

- number of likes
- number of saves
- recency
- user interests

Example ranking idea:

score = likes + saves + freshness

More advanced approaches may include:

- collaborative filtering
- user behavior modeling
- topic clustering

---

# Media Upload Flow

> Future optimasation
> Images should not be stored directly in the application server.

Recommended flow:

1. User uploads image
2. Backend uploads to storage service
3. Storage returns URL
4. URL stored in database
5. Clients render image using CDN

---

# Pagination Strategy

> Frontend Optimasation

Feeds and search results must be paginated.

Two common approaches:

### Offset Pagination

Used in simple systems.

Example concept:

- page
- limit

---

### Cursor Pagination

Used in high-scale feeds.

Benefits:

- better performance
- consistent scrolling
- supports infinite scroll

---

# Security Considerations

> Important

Important backend protections include:

- Password hashing (bcrypt)
- JWT authentication
- Rate limiting
- Input validation
- Authorization checks

---

# Development Goals

This project aims to help developers practice:

- REST API design
- Backend architecture
- database relationships
- feed systems
- scalable service organization

---

# Possible Extensions

After completing the basic backend, the system can be extended with:

- real-time notifications
- recommendation algorithms
- image similarity search
- tag systems
- board collaboration
- analytics for creators

---

# Learning Outcomes

By completing this project you will gain experience with:

- building production-style Go backends
- structuring large backend projects
- designing social media APIs
- handling media storage
- building scalable feed systems

---

# Status

This project is intended as a backend practice system inspired by the functionality of the platform **Pinterest**.
