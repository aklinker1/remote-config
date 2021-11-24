FROM node:16-alpine as frontend-builder
RUN mkdir /build
WORKDIR /build
RUN apk --no-cache add curl
RUN curl -f https://get.pnpm.io/v6.16.js | node - add --global pnpm@6
ADD package.json pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile
ADD . .
RUN pnpm vite build

FROM golang:1.17-alpine as backend-builder
RUN mkdir /build
WORKDIR /build
ADD go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=frontend-builder /build/dist ./dist
RUN go build -o bin/app main.go

FROM alpine
WORKDIR /app
COPY --from=backend-builder /build/bin/app .
CMD ./app
