proto-music-check:
	cd music && buf lint

proto-music-server-gen:
	cd music && buf generate

start: 
	docker-compose up --build 

stop:
	docker-compose down

# ==============================================================================
# Make local SSL Certificate

cert:
	@echo Generating SSL certificates
	cd ./ssl && sh instructions.sh