# Contributing

Here you'll find the docs for setting up your local environment for development.

## Required Tools

- `docker`
- `docker-compose`
- `node` v16+
- `pnpm`
- `go` v17+
- GNU `make`
- `aws` CLI v2

## First Time Setup

Please note, that local development still requires a real AWS S3 bucket.

> I'll accept a PR to use localstack or some other local S3 equivalent. Eventually I'll abstract out the storage unit and support more storage solutions

```bash
# Install front-end dependencies
pnpm install

# Setup env variables for local development
echo "AWS_REGION=<region>
AWS_ACCESS_KEY_ID=<fill-out>
AWS_SECRET_ACCESS_KEY=<fill-out>

S3_BUCKET=example-bucket-name
S3_FILE_PATH=path/to/remote-config.json

PORT=80
AUTH_TOKEN=password" > .env
```

## Running In Dev Mode

Run the application in dev mode, then open <http://localhost>!

```bash
make run
```

Dev mode will hot reload the frontend or restart the backend when relevant files are saved.

- Frontend: Vite does stateful HMR
- Backend: Super fast thanks to Go's compile time, docker multi-stage builds, and the `http` libraries insane startup time

> Saving the `main.go` file will not automatically restart, you'll need to restart `make run`
