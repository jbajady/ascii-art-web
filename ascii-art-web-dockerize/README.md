# Go Web Server with Docker
This project demonstrates how to create a simple web server in Go, containerize it using Docker, and manage Docker objects. It adheres to best practices for Dockerfile creation and image management.

## Project Overview
- Language: Go
- Containerization: Docker
- Metadata: Applied to Docker objects
- Good Practices: Followed for Go code and Dockerfile
- Dependencies: Only standard Go packages are used

## Instructions
### Prerequisites
- Docker: Ensure Docker is installed and running on your machine. You can find installation instructions in the [Docker documentation](https://docs.docker.com/).

## Build the Docker Image
To build the Docker image for the Go web server, use the following command:

```bash
docker build -f Dockerfile -t go-web-server .
```
- -f Dockerfile: Specifies the Dockerfile to use.
- -t go-web-server: Tags the image with the name go-web-server.

## Run the Docker Container
To run the Docker container and map port 8080 on your host to port 8080 in the container, use:

```bash
docker run -p 8080:8080 --detach --name go-web-server-container go-web-server
```
- `-p 8080:8080`: Maps port 8080 on your host to port 8080 in the container.
- `--detach`: Runs the container in the background.
- `--name go-web-server-container`: Names the container go-web-server-container.

## Good Practices
- Go Code: The Go code should follow best practices, including proper error handling and efficient use of resources.
- Dockerfile: The Dockerfile should follow Dockerfile best practices, including using a minimal base image, combining commands to reduce layers, and properly setting metadata.

## Learning Outcomes
- Client Utilities: Interacting with Docker using client commands.
- Basics of Web Development: Creating a web server, handling HTTP requests, and generating HTML responses.
- Docker: Understanding Docker concepts such as images, containers, and garbage collection.
- Containerizing Applications: Building and running Docker containers.
- Image Management: Tagging and managing Docker images.

## References
- [Docker Documentation](https://docs.docker.com/)
- [Go Documentation](https://go.dev/doc/)