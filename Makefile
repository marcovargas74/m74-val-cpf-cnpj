all: clean deps test test-web docker-build docker-deploy-up

clean:
	@go clean
	@rm -rf build

deps:
	@go get -v ./...

test:
	@go test -v ./...


# coverage:
# 	@mkdir -p build
# 	@go test -coverprofile build/cover.out ./...
# 	@go tool cover -html=build/cover.out -o build/cover.html
# 	@cd web && npm test -- --verbose --silent --coverage --watchAll=false

#docker-build:
#	@cd src/bank && go build main.go .
  

docker-deploy-up:
	@cd docker && docker-compose up -d

docker-deploy-down:
	@cd docker && docker-compose down