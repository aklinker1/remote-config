build:
	docker build . -t anime-skip.com/remote-config

prep:
	@mkdir -p dist
	@echo "<html></html>" > dist/index.html

run: prep
	docker-compose up --build -V

run-prod: prep build
	docker run \
		--env-file .env \
		-p 80:80 \
		anime-skip.com/remote-config

format:
	@pnpm prettier --write .
	@go fmt anime-skip.com/remote-config
	@go fmt anime-skip.com/remote-config/src/...
