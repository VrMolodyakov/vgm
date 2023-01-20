proto-music-check:
	cd music && buf lint

proto-music-gen:
	cd music && buf generate

start: 
	docker-compose up --build 