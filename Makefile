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

test:
	go test ./...

test-detail:
	go test ./... -v

test-coverage:
	go test -cover ./...

test-cover-generate-output:
	go test -cover ./... -coverprofile=bin/testCoverage.out

test-cover-output:
	go tool cover -func=bin/testCoverage.out

test-cover-html:
	go tool cover -html=bin/testCoverage.out
