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
RUN apk add --update gcc libc-dev
RUN mkdir /build
WORKDIR /build
ADD go.mod go.sum ./
RUN go mod download
COPY . .
ARG VERSION
ARG STAGE
RUN : "${VERSION:?Build argument needs to be passed and non-empty}"
RUN : "${STAGE:?Build argument needs to be passed as 'development' or 'production'}"
COPY --from=frontend-builder /build/dist ./dist
RUN go build \
    -trimpath \
    -ldflags "-X anime-skip.com/remote-config/src/backend/env.VERSION=$VERSION -X anime-skip.com/remote-config/src/backend/env.STAGE=$STAGE" \
    -o bin/app main.go

FROM alpine
WORKDIR /app
COPY --from=backend-builder /build/bin/app .
CMD ./app
