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
	go get github.com/pilu/fresh

.PHONY: build
build:
	fresh -c fresh.conf

docker-build:
	docker build -t file-indexer .

docker-run: docker-build
	docker run -it -p 8000:8000 file-indexer