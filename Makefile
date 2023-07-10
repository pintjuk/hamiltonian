dist:
	mkdir $@

dist/routmester: dist src/*
	cd src/;go build -o ../$@

.PHONY:run
run:
	${MAKE} clean
	${MAKE} dist/routmester
	./dist/routmester

.PHONY:clean
clean:
	rm -rf dist

cover.out:
	go test -coverprofile cover.out github.com/pintjuk/routemaster/src/route
.PHONY:test
test:
	rm cover.out || true
	${MAKE} cover.out

.PHONY: cover
cover: cover.out
	go tool cover -html=cover.out -o cover.html
	open cover.html

.PHONY: setup
setup:
	go install honnef.co/go/tools/cmd/staticcheck@latest

.PHONY: check
check:

.PHONY: docker-build
docker-build:
	docker build . -t pintjuk/routemaster

.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 docker.io/pintjuk/routemaster

.PHONY: help
help:
	@printf "run\t\t\t start server\n"
	@printf "test\t\t\t run tests\n"
	@printf "cover\t\t\t opens coverage report\n"
	@printf "docker-build\t\t build docker container\n"
	@printf "docker-run\t\t run server in docker\n"