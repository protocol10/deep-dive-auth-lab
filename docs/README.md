# Deep Dive Auth Lab

A comprehensive lab environment designed to explore, build, and demonstrate multiple ways to authenticate and authorize users. Built with Go.

## 🚀 Features & Roadmap

This project progressively builds up from basic authentication to advanced authorization mechanisms. 

- [ ] **Simple Registration & Login:** Basic email and password authentication with secure hashing (e.g., bcrypt).
- [ ] **JWT Integration:** Stateless authentication using JSON Web Tokens.
- [ ] **OTP / TOTP:** One-Time Passwords and Time-based One-Time Passwords (like Google Authenticator).
- [ ] **MFA (Multi-Factor Authentication):** Upgrading the login flow to require both a password and a TOTP code.
- [ ] **RBAC (Role-Based Access Control):** Granular authorization based on user roles (e.g., Admin, User) and permissions.

## 🛠️ Prerequisites

- [Go](https://go.dev/) 1.20 or higher
- A Database (e.g., PostgreSQL, MySQL, or SQLite)

## ⚙️ Getting Started

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/deep-dive-auth-lab.git
   cd deep-dive-auth-lab
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the application:**
   ```bash
   go run main.go
   ```

## 📚 Learnings & Notes

*(As you build out each feature, document your learnings, design decisions, and API endpoints here. It serves as a great reference for the future!)*

## 🛡️ Security Disclaimer

This is an educational lab environment. When implementing these patterns in production, always ensure you follow security best practices (secure password hashing, enforcing HTTPS, secure cookie flags, and proper secret management).