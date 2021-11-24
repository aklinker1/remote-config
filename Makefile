build:
	go build -o bin/app cmd/app/main.go

prep:
	@mkdir -p src/frontend/dist
	@echo "<html></html>" > src/frontend/dist/index.html

run: prep
	docker-compose up --build -V

run-prod: prep
	echo "TODO"
