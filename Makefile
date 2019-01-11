.PHONY: clean
clean:
	rm -rf logs

.PHONY: logs
logs:
	make clean
	mkdir logs
	cp config.json logs/
	pip3 install --user Faker
	./scripts/helper.py logs/

.PHONY: dev
dev:
	go run cmd/dexer/main.go

.PHONY: docker-build
docker-build:
	docker build -t dexer .

.PHONY: docker-run
docker-run: docker-build
	docker run -it -p 8000:8000 dexer

all:
	go install ./...
