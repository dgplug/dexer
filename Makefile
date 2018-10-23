clean:
	rm -rf logs

dev:
	make clean
	mkdir logs
	go get github.com/pilu/fresh
	cp config.json logs/
	./scripts/helper.py logs/

build:
	fresh -c fresh.conf
