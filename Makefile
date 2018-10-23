clean:
	rm -rf logs

logs:
	make clean
	make dev
	mkdir logs
	cp config.json logs/
	pip3 install --user Faker
	./scripts/helper.py logs/

dev:
	go get github.com/pilu/fresh

build:
	fresh -c fresh.conf
