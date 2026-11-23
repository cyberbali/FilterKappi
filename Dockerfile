# Stage 1: Build binary in a reproducible minimal environment
FROM golang:1.27-alpine AS builder

WORKDIR /src

# Cache dependencies
COPY go.mod ./
RUN go mod download

# Copy source tree
COPY . .

# Build hardened static binary without CGO
RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-w -s" \
    -o /bin/filterkappi ./cmd/filterkappi

# Stage 2: Minimal, non-root runtime container
FROM alpine:3.24

# Create non-root dedicated group and user (UID/GID 10001)
RUN addgroup -g 10001 -S appgroup && \
    adduser -u 10001 -S appuser -G appgroup

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /bin/filterkappi /bin/filterkappi

# Switch to non-root user
USER 10001:10001

EXPOSE 8080

ENTRYPOINT ["/bin/filterkappi"]
