build:
	docker build . -t anime-skip.com/remote-config

prep:
	@mkdir -p dist
	@echo "<html></html>" > dist/index.html

run: prep
	docker-compose up --build -V

run-prod: prep build
	docker run \
		--env-file docker/aws.env \
		--env PORT=80 \
		--env S3_BUCKET=remote-config.anime-skip.com \
		--env S3_FILE_PATH=prod.json \
		--env AUTH_TOKEN=password \
		-p 80:80 \
		anime-skip.com/remote-config

format:
	@pnpm prettier --write .
	@go fmt anime-skip.com/remote-config
	@go fmt anime-skip.com/remote-config/src/...
