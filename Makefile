clean:
	rm -rf logs

dev:
	make clean
	mkdir logs 
	cp config.json logs/
	fresh -c fresh.conf
