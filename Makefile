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