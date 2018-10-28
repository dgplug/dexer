.PHONY: clean
clean:
	rm -rf logs

.PHONY: logs
logs:
	make clean
	make dev
	mkdir logs
	cp config.json logs/
	pip3 install --user Faker
	./scripts/helper.py logs/

.PHONY: dev
dev:
	go run main.go

.PHONY: docker-build
docker-build:
	docker build -t file-indexer .

.PHONY: docker-run
docker-run: docker-build
	docker run -it -p 8000:8000 file-indexer

all:
	go install github.com/farhaanbukhsh/file-indexer
