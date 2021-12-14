dev: 
	docker-compose up
build-ui:
	docker-compose run vue-ui npm run build
build-server:
	cd ./dccpi && go build ./dccpi.go
	sudo chown root ./dccpi/dccpi
	sudo chmod +s ./dccpi/dccpi
bash:
	docker-compose run vue-ui bash