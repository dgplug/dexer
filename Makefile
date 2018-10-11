clean:
	rm -rf logs

dev:
	make clean
	mkdir logs 
	cp config.json logs/
	go run main.go
