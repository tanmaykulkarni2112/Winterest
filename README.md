## 📌 Commit Message Guidelines

To maintain a clean and understandable project history, all commit messages should follow a structured format.

### 🔹 Format

```
<type>(optional-scope): <short description>

[optional body]

[optional footer]
```

---

### 🔹 Common Commit Types

| Type       | Description                                               |
| ---------- | --------------------------------------------------------- |
| `feat`     | A new feature                                             |
| `fix`      | A bug fix                                                 |
| `docs`     | Documentation changes                                     |
| `style`    | Code style changes (formatting, missing semicolons, etc.) |
| `refactor` | Code changes that neither fix a bug nor add a feature     |
| `test`     | Adding or updating tests                                  |
| `chore`    | Maintenance tasks (build, dependencies, configs)          |
| `perf`     | Performance improvements                                  |
| `ci`       | Changes to CI/CD pipelines                                |

---

### 🔹 Examples

#### ✅ Feature

```
feat(auth): add JWT-based authentication
```

#### ✅ Bug Fix

```
fix(api): handle null response in user service
```

#### ✅ Documentation

```
docs(readme): update setup instructions
```

#### ✅ Refactor

```
refactor(db): optimize query structure for user lookup
```

#### ✅ Style

```
style(ui): fix indentation in navbar component
```

#### ✅ Test

```
test(auth): add unit tests for login flow
```

#### ✅ Chore

```
chore(deps): update npm packages
```
