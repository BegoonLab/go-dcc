dev: 
	docker-compose up
build:
	docker-compose run vue-ui npm run build
bash:
	docker-compose run vue-ui bash