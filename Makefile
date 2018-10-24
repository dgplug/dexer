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
