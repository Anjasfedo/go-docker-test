# Go Web Server with Docker 🚀

This repository contains a simple Go web server that serves different responses based on different routes. Additionally, a Dockerfile is included to containerize the Go application.

## Overview 📝

### `main.go`

-   Entry point of the program.
-   Defines a basic web server using the `net/http` package.
-   Two routes: `/` and `/hello`.
-   `/` responds with a formatted string containing the escaped URL path.
-   `/hello` responds with a simple "Hello" message.

### `Dockerfile`

-   Uses the specified Go version as the base image.
-   Sets the working directory inside the container.
-   Copies the source code, builds the Go application, and defines the command to run the executable.

## Building and Running 🛠️

To build the Docker image and run the container:

-   This command builds the Docker image for the Go web server based on the provided Dockerfile:

```bash
docker build -t go-docker-test .
```

-   Use this command to start the Docker container:

```bash
docker run -p 8000:8080 go-docker-test
```

It maps port 8000 on your local machine to port 8080 on the container, allowing you to access the web server at `http://localhost:8000/`.

## **Note for Windows**

If you are using Git Bash on Windows, you might encounter issues with the terminal. In such cases, use the following command with winpty:

```bash
winpty docker run -p 8000:8080 go-docker-test
```

## Closing Notes📝

If you find any issues or have suggestions for improvement, please feel free to open an issue or submit a pull request.

Happy coding!🚀👨‍💻
