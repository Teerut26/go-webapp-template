# Go WebApp Template

This project is a template for a full-stack web application using Go for the backend and a modern JavaScript framework for the frontend. It is designed to help you kickstart your development with a pre-configured structure for building web applications.

## Project Structure

- **main.go** - The entry point for the Go backend.
- **go.mod / go.sum** - Go module files.
- **frontend/** - Contains the frontend code:
  - **index.html** - The main HTML file.
  - **src/** - Source code for the application.
  - **package.json** - Node dependencies and scripts.
  - **vite.config.ts** - Vite configuration for development and build.
- **Dockerfile** - Docker configuration for containerization.
- **.gitignore, .dockerignore** - Files to exclude from version control and Docker builds.

## Prerequisites

- [Go](https://golang.org/doc/install)
- [Node.js](https://nodejs.org/en/download/)
- [pnpm](https://pnpm.io/installation) (or npm/yarn as configured in your project)

## Getting Started

1. **Clone the repository:**

   ```bash
   git clone <repository-url>
   cd go-webapp-template
   ```

2. **Set up the Go backend:**

   - Install Go dependencies:
   
     ```bash
     go mod download
     ```
     
   - Run the backend server:
   
     ```bash
     go run main.go
     ```

3. **Set up the Frontend:**

   - Navigate to the frontend directory:
   
     ```bash
     cd frontend
     ```
     
   - Install Node dependencies:
   
     ```bash
     pnpm install
     ```
     
   - Run the development server:
   
     ```bash
     pnpm run dev
     ```

## Docker

A `Dockerfile` is included for containerizing the application. To build and run the container:

```bash
docker build -t go-webapp-template .
docker run -p 8080:8080 go-webapp-template
```

## Contributing

Contributions are welcome! Please fork this repository and submit a pull request with your proposed changes.

## License

This project is licensed under the MIT License.
