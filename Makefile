build:
	@bash scripts/build.sh

create-pass: build
	@./bin/passgen generate -l=$(l)

strong-password:
	docker build -t go-passtrong ./docker
	docker run go-passtrong:latest
	@docker rmi -f go-passtrong >/dev/null 2>&1
	@docker rm $$(docker ps -a -f status=exited -q) -f >/dev/null 2>&1

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

generate-test-output:
	go test -cover ./... -coverprofile=bin/testCoverage.out

test-output:
	go tool cover -func=bin/testCoverage.out

test-html:
	go tool cover -html=bin/testCoverage.out
