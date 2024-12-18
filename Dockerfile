# Use an official Go image
FROM golang:1.23

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0

# Set the working directory inside the container
WORKDIR /app

# Install `bee` globally
RUN go install github.com/beego/bee/v2@latest

# Copy the entire project into the container
COPY . .

# Expose the application port
EXPOSE 8080

# Default command to run
CMD ["bee", "run"]
