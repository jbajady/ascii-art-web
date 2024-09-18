#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status

# Stop all running containers
if [ "$(docker ps -q)" ]; then
    echo "Stopping all running containers..."
    docker stop $(docker ps -q)
else
    echo "No running containers to stop."
fi

# Remove all containers
if [ "$(docker ps -a -q)" ]; then
    echo "Removing all containers..."
    docker rm $(docker ps -a -q)
else
    echo "No containers to remove."
fi

# Build the Docker image
if [ -f Dockerfile ]; then
    echo "Building Docker image..."
    docker image build -f Dockerfile -t ascii-art-web .
else
    echo "Dockerfile not found. Exiting."
    exit 1
fi

# Run the new container
echo "Running the new container..."
docker container run -p 8080:8080 --detach --name ascii-art-web-container ascii-art-web

echo "Script completed successfully."
