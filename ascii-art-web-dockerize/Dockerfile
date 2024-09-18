# Stage 1: Build Stage
# Use the official Golang image with Alpine Linux as the base for building the application
FROM golang:alpine AS build-stage

# Set the working directory inside the container
WORKDIR /app

# Copy the local files into the container's working directory
COPY . .

# Build the Go application and output the binary named 'main'
RUN go build -o main .

# Stage 2: Final Stage
# Use the official Alpine Linux image as the base for the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the build stage to the final image
COPY --from=build-stage /app/main .

# Copy additional files and directories from the build stage to the final image
COPY --from=build-stage /app/templates/ ./templates/
COPY --from=build-stage /app/ressources/ ./ressources/
COPY --from=build-stage /app/static/ ./static/
COPY --from=build-stage /app/standard.txt ./standard.txt
COPY --from=build-stage /app/shadow.txt ./shadow.txt
COPY --from=build-stage /app/thinkertoy.txt ./thinkertoy.txt

# Install bash in the final image
RUN apk add --no-cache bash

# Expose port 8080 to allow traffic to the container
EXPOSE 8080

# Set the command to run the application when the container starts
CMD ["./main"]

# Provides metadata about the image
LABEL \
website.name="ascii-art-web-dockerize" \
    description="simple web server that serves static files, handles basic routes, and generates ASCII art based on input. The server is built using Go and provides several endpoints for viewing ASCII art and other static content." \
    authors="Fethi Abderrahmane, Bajady Jamal"