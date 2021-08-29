create-pass:
	build \\
	@./bin/passgen create_password

build:
	@bash scripts/build.sh

build-docker:
	docker-compose up --build

run-docker:
	docker-compose up

up:
	docker-compose up -d

down:
	docker-compose down

clean-builds:
	@rm -f bin/*
