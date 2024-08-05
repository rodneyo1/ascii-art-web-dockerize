#!/bin/bash

# Build the Docker image
docker build -t asciiartserver:latest .

# Run the Docker container in detached mode with port mapping
docker run -d -p 8080:8080 --name asciiartserver asciiartserver:latest
