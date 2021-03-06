build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/lukasjarosch/user-service \
		proto/auth/auth.proto
	docker build -t user-service:latest .

run:
	docker run --net="host" \
		-p 50051 \
		-e DB_HOST=localhost \
		-e DB_PASS=password \
		-e DB_USER=postgres \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		user-service


