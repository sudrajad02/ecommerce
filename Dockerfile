# Base image, golang 1.20
FROM golang:1.20.5
WORKDIR /workspace
# Copy all files into the image
COPY . .
# Run go mod
RUN go mod download
# Expose ports
EXPOSE 3000
# Run Go program, just like locally
ENTRYPOINT ["go","run","main.go"]