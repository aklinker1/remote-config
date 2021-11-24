# remote-config

This is a lightweight and simple remote config server with a UI for management backed by AWS S3.

The frontend is written with Vue, while the backend API written in Go serves it and the REST endpoints.

## Get Started

Make sure you have all the required tools installed

- `docker`
- `node@16`
- `pnpm`
- `go@17`
- `make`
- `aws`

Then startup the app in dev mode and visit <http://localhost>!

```bash
echo "AWS_REGION=<fill-out>
AWS_ACCESS_KEY_ID=<fill-out>
AWS_SECRET_ACCESS_KEY=<fill-out>
" > docker/aws.env

make run
```

The frontend will hot-reload when saving changes inside the `src/frontend` directory, and the backend will restart when saving any files in the `src/backend` directory!

> Saving the `main.go` file will not automatically restart, you'll need to restart `make run`

To make and run a production build of the app, run:

```bash
make build
make run-prod
```

Then visit <http://localhost:8000>!

## Configuration

You can set the `S3_BUCKET` and `S3_FILE_PATH` environment variables to specify which bucket and what file in that bucket you want to use to store the data.

In dev mode, they are hard coded according to the `docker-compose.yml` file.
