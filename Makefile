dev: 
	docker compose up server vue-ui
build-ui:
	docker compose run vue-ui npm run build
build-server:
	cd ./software/dccpi/cmd && go build ./main.go
	sudo chown root ./software/dccpi/cmd/main
	sudo chmod +s ./software/dccpi/cmd/main
bash:
	docker compose run vue-ui bash

lint-go:
	cd ./software/dccpi && golangci-lint run --fix --config .golangci.yml

compile-protobuf-go:
	protoc -I=./software/frontend/src/pb --go_out=./software/dccpi/internal/pb/build ./software/frontend/src/pb/controller.proto

test:
	cd ./software/dccpi && go test -race ./...