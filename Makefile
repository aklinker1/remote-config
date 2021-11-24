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
		--env PORT=8000 \
		--env S3_BUCKET=remote-config.anime-skip.com \
		--env S3_FILE_PATH=prod.json \
		-p 80:8000 \
		anime-skip.com/remote-config
