clean:
	rm -rf logs

dev:
	make clean
	mkdir logs 
	cp config.json logs/
	./scripts/helper.py logs/

build:
	fresh -c fresh.conf
