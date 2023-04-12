proto-music-check:
	cd music && buf lint

proto-music-server-gen:
	cd music && buf generate

start: 
	docker-compose up --build 

stop:
	docker-compose down

cert:
	@echo Generating SSL certificates
	cd ./cert && sh instructions.sh

.PHONY: cert 
