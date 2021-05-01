dev-setup: 
	@mkdir -p ./tools
	wget -O ./tools/swagger https://github.com/go-swagger/go-swagger/releases/download/v0.26.0/swagger_linux_amd64
	chmod +x ./tools/swagger
	go get -u github.com/golang/dep/cmd/dep
swagger-gen:
	@rm -Rf gen
	@mkdir -p gen
	./tools/swagger generate server -t gen -f ./swagger/daily-logs.yml ./gen -A daily-logs --exclude-main -P auth.User
clean:
	rm -Rf gen vendor
test-health:
	curl -i -H "Accept: text/plain" -H "Content-Type: text/plain" -X GET http://localhost:8080/health
run:
	go run main.go