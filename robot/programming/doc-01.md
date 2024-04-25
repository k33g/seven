# Step-by-step Guide to Create a Dockerfile for Your Project

Based on the provided file structure, you can create an efficient and reusable `Dockerfile` using markdown formatting as follows:

## Introduction

The following instructions will guide you in creating a `Dockerfile` that builds your project with specified dependencies. The given context contains a variety of files including Go source code (`main.go`), HTML, JavaScript (js folder), configuration files such as `.env`, and various assets like CSS and JS files within the 'public' directory.

## Dockerfile

Create a new file named `Dockerfile` in your project root with this content:

```markdown
# Start of the Dockerfile 

FROM golang:alpine AS builder

# Install build dependencies for Go and Node.js
RUN apk add --no-cache gcc git musl-dev
RUN curl -sL https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN echo "deb http://dl.yarnpkg.com/debian/ stable main" >> /etc/apt/sources.list.d/yarn.list
RUN curl -sL https://deb.nodesource.com/setup_14.x | bash - && \
    apt-get update && \
    apt-get install --no-install-recommends nodejs

# Set working directory and copy local code into the container image
WORKDIR /app
COPY . .

## Installing Go dependencies

RUN go mod download

## Compiling application

RUN go build -o main.out ./...

## Building front-end assets using Node.js

FROM node:alpine AS frontend
WORKDIR /app/public/js
COPY --from=builder /app/js/* ./
RUN npm install && \
    npm run build

# Installing preact and installing global packages (for CSS compilation)
FROM node:alpine AS css
WORKDIR /app/public/js
COPY --from=frontend /app/public/dist ./dist
RUN npm install -g preact-cli && \
    npx preact build dist/*.html --out ../css/

# Final image with compiled front-end assets and main Go binary
FROM alpine:latest AS final
WORKDIR /root/
COPY --from=builder /app/main.out ./
COPY --from=frontend /app/public/dist/index.html ./
COPY --from=css /appe/public/css/* .
EXPOSE 8080
CMD ["./main.out"]
```

After adding the `Dockerfile` to your project, build and run the Docker image with the following commands:

1. Build the Docker image:
   ```sh
   docker build -t my-project .
   ```
2. Run the container with port 8080 exposed and the built Go binary as entrypoint:
   ```sh
   docker run --rm -p 8080:8080 my-project
   ```

Remember to replace `my-project` with your project name, adjust file paths if necessary, and make sure all dependencies are properly installed.
