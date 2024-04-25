**Dockerizing Your Project: A Step-by-Step Guide**

Given your project context, I'll create a Dockerfile that builds and runs your application using slim images and multi-stage builds.

### **Prerequisites**

Before we begin, make sure you have the following files in your project directory:

* `main.go`: The entry point for your Go application
* `go.mod` and `go.sum`: The Go module file and its corresponding sum file
* `public/index.html`: The HTML file for your web application

### **Dockerfile: Building the Project**

Create a new file named `Dockerfile` in the root of your project directory. This file will contain the instructions for building and running your application.

```dockerfile
# Stage 1: Build the Go Application
FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Create a Slim Image for the Web Application
FROM nginx:alpine AS web
WORKDIR /usr/share/nginx/html
COPY public/ ./

# Stage 3: Combine the Two Stages and Set the Default Command
FROM builder AS final
CMD ["main"]
```

### **Explanation**

Here's what each section of the Dockerfile does:

1. **Stage 1**: We use the `golang:alpine` image as a base for building our Go application. We set the working directory to `/app`, copy the `go.mod` and `go.sum` files, download dependencies, copy the project files, and build the application using `go build`.
2. **Stage 2**: We use the `nginx:alpine` image as a base for creating a slim image for our web application. We set the working directory to `/usr/share/nginx/html`, copy the `public/` directory, and make it the default document root.
3. **Stage 3**: We combine the two stages by using the `builder` stage as the base for the final image. We set the default command to run the `main` executable.

### **Building and Running the Docker Image**

To build the Docker image, navigate to your project directory and run the following command:
```bash
docker build -t my-app .
```

This will create a Docker image with the name `my-app`.

To run the Docker image, use the following command:
```bash
docker run -p 8080:80 my-app
```

This will start a new container from the `my-app` image and map port 8080 on your host machine to port 80 in the container. You can then access your web application by visiting `http://localhost:8080` in your browser.

That's it! You now have a Dockerized version of your project that builds and runs using slim images and multi-stage builds.
