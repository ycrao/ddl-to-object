FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Create config directory and copy templates
RUN mkdir -p /root/.dto && \
    cp -r template/ /root/.dto/ && \
    cp config.example.json /root/.dto/config.json

# Build the web server
RUN cd web && go build -o d2o-server server.go


FROM alpine:3.23.0

WORKDIR /app

# Copy the web server binary
COPY --from=builder /app/web/d2o-server .

# Copy web files
COPY --from=builder /app/web/*.html .
COPY --from=builder /app/web/*.js .
COPY --from=builder /app/web/*.css .

# Copy config and templates
COPY --from=builder /root/.dto /root/.dto
COPY --from=builder /app/template /app/template

# Expose port
EXPOSE 8080

# Set environment variable for port
ENV PORT=8080

# Run the web server
CMD ["./d2o-server"]
