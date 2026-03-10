#
# Builder image using Chainguard's minimal Go environment
#
FROM cgr.dev/chainguard/go:latest AS builder

WORKDIR /app

# copy go modules manifests and download dependencies first
COPY ../go.mod .
RUN go mod download

# copy the rest of the source
COPY . .

# build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/gym ./cmd/gym

#
# Final image: also use the Chainguard base for a small, secure runtime
#
FROM cgr.dev/chainguard/go:latest

# copy the statically built binary from the builder
COPY --from=builder /out/gym /usr/local/bin/gym

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/gym"]
