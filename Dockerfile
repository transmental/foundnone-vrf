# Stage 1: install snarkjs into a full Node image
FROM node:23-slim AS snark
WORKDIR /snark
RUN npm install -g snarkjs express

# Stage 2: build your Go fulfiller
FROM golang:1.24-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY ./fulfiller/go.mod ./fulfiller/go.sum ./
RUN go mod download
COPY . ./
RUN cd fulfiller && go build -o fulfiller .

# Stage 3: final image (both Go binary + Node server)
FROM snark

# install just what we need for HTTPS
RUN apt-get update && apt-get install -y ca-certificates --no-install-recommends \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# copy Go fulfiller binary
COPY --from=builder /app/fulfiller/fulfiller /app/fulfiller

RUN chmod +x /app/fulfiller

# copy circom zk artifacts
COPY ./fulfiller/zk ./zk

# copy Node prover server
COPY ./fulfiller/prover.js ./prover.js

# make sure Node can see global modules
ENV NODE_PATH=/usr/local/lib/node_modules

# start both processes: Node and Go fulfiller
CMD node /app/prover.js & /app/fulfiller
