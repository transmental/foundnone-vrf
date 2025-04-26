# Stage 1: install snarkjs (and all of its deps) into a full Node image
FROM node:23-slim AS snark
WORKDIR /snark
RUN npm install -g snarkjs

# Stage 2: build your Go listener
FROM golang:1.24-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o listener .

# Stage 3: final image, still based on node:slim so we have Node.js & npm modules
FROM snark
# install just what we need to talk HTTPS
RUN apt-get update && apt-get install -y ca-certificates --no-install-recommends \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# 1) copy your Go listener binary
COPY --from=builder /app/listener /app/listener

# tell Node where to find those global modules
ENV NODE_PATH=/usr/local/lib/node_modules

# 3) copy your circom artifacts
COPY zk ./zk

# your listener will read .env from the host
ENTRYPOINT ["/app/listener"]
