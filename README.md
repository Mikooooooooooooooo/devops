
# Final Project: Full-Stack Application with Docker, Go Backend, and Vue.js Frontend

## Project Overview

This project is a simple full-stack application that uses:

- **Go** for the backend API.
- **Vue.js** for the frontend.
- **Docker** for containerization of both frontend and backend services.
- **GitHub Actions** for Continuous Integration and Continuous Deployment (CI/CD).

The goal of this project is to build a basic application that includes a backend API and a frontend interface while leveraging Docker for containerization and GitHub Actions for CI/CD.

---

## Features

- **Backend (Go)**: A simple Go API that exposes a RESTful endpoint.
- **Frontend (Vue.js)**: A Vue.js app that communicates with the backend API.
- **Docker**: Docker containers for both frontend and backend services.
- **CI/CD**: Automated builds and tests with GitHub Actions.

---

## Prerequisites

Before running the project, make sure you have the following installed:

- [Docker](https://www.docker.com/products/docker-desktop) (for containerization)
- [Docker Compose](https://docs.docker.com/compose/) (for managing multi-container Docker applications)
- [Node.js](https://nodejs.org/) (required for building the Vue.js frontend)
- [Go](https://golang.org/) (for backend API)
- [jq](https://stedolan.github.io/jq/) (for parsing JSON in GitHub Actions, only needed for CI/CD step)

---

## Project Structure

```
.
├── backend/              # Go backend API
│   ├── Dockerfile        # Dockerfile for Go API
│   └── main.go           # Main Go application
├── frontend/             # Vue.js frontend
│   ├── Dockerfile        # Dockerfile for Vue.js app
│   └── src/              # Vue.js source files
├── docker-compose.yml    # Docker Compose configuration file
├── .github/              # GitHub Actions CI/CD pipeline
│   └── workflows/        # CI/CD workflow files
│       └── ci.yml        # GitHub Actions workflow file
└── README.md             # This file
```

---

## Setup and Run Locally

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/final-project.git
   cd final-project
   ```

2. Build and run the project using Docker Compose:

   ```bash
   docker-compose up --build
   ```

3. Access the application:

   - Backend API: `http://localhost:8080`
   - Frontend: `http://localhost:8081`

   The frontend should be able to communicate with the backend API at `http://localhost:8080`.

---

## Docker Setup

This project uses Docker to containerize both the backend and frontend. The Docker configuration is specified in the `docker-compose.yml` file.

1. **Backend (Go)**: The backend service is built using a Go Docker image. It exposes the backend API on port `8080`.

2. **Frontend (Vue.js)**: The frontend service uses a multi-stage build to compile the Vue.js app and then serve it with Nginx. It is accessible on port `8081`.

---

## CI/CD Setup (GitHub Actions)

This project uses GitHub Actions for CI/CD. It builds the Docker images for both backend and frontend whenever code is pushed to the `master` branch.

### GitHub Actions Workflow

The workflow is defined in `.github/workflows/ci.yml`. It consists of the following steps:

1. **Checkout Code**: Retrieves the latest code from the repository.
2. **Set up Docker Buildx**: Configures Docker Buildx for multi-platform builds.
3. **Install Docker Compose**: Installs Docker Compose on the CI runner.
4. **Build Docker Containers**: Builds the Docker containers for both frontend and backend.
5. **Run Tests** (if any are added later): Ensures everything works before deploying.

---

## Docker Compose Configuration

The `docker-compose.yml` file defines the services for both the frontend and backend.

```yaml
version: '3'
services:
  backend:
    build: ./backend
    ports:
      - "8080:8080"
    networks: 
      - devops

  frontend:
    build: ./frontend
    ports:
      - "8081:80"
    depends_on:
      - backend
    networks: 
      - devops

networks:
  devops:
```

- **Backend**: The Go API runs on port `8080` and is accessible by the frontend.
- **Frontend**: The Vue.js app is served by Nginx and is accessible on port `8081`. It depends on the backend being up first.

---

## Troubleshooting

- **CORS Errors**: If you're getting CORS errors when trying to connect the frontend to the backend, ensure the backend API includes the necessary CORS headers.
  
  Example of CORS middleware in Go:

  ```go
  import "github.com/rs/cors"

  var c = cors.New(cors.Options{
    AllowedOrigins: []string{"http://localhost:8081"},
  })

  handler := c.Handler(router)
  ```

- **Permission Issues with Docker**: If you encounter permission issues when running Docker, ensure that your user is part of the `docker` group:

  ```bash
  sudo usermod -aG docker $USER
  ```

  Then restart your terminal or run `newgrp docker` to apply the changes.

---

## Contributing

Feel free to fork the repository and submit pull requests. Make sure to follow the project structure and add tests if necessary.

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
