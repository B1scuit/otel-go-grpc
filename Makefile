GO = CGO_ENABLED=0 GOARCH=amd64 go build -a -gcflags='all=-N -l' -ldflags="-w -s" -tags timetzdata
GOOGLEAPIS=protovendor/

SERVICE_ONE=service_one
SERVICE_TWO=service_two
SERVICE_THREE=service_three

all: proto build_all_docker deploy

# Checking if the applications we need are installed
check: GO-Exists PROTOC-Exists DOCKER-Exists MINIKUBE-Exists
GO-Exists: ; @which go > /dev/null
MINIKUBE-Exists: ; @which minikube > /dev/null
DOCKER-Exists: ; @which docker > /dev/null
# If you need to install protoc:
# https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers
PROTOC-Exists: ; @which protoc > /dev/null 

# Building of the services, these are mainly going to be used by the dockerfile
service_one:
	$(GO) cmd/$(SERVICE_ONE)/main.go
service_two:
	$(GO) cmd/$(SERVICE_TWO)/main.go
service_three:
	$(GO) cmd/$(SERVICE_THREE)/main.go

# Some quick handling around building the dockerfiles
build_all_docker: service_one_docker service_two_docker service_three_docker
service_one_docker: check mini-eval
	docker build -t $(SERVICE_ONE) --build-arg SERVICE=$(SERVICE_ONE) -f deploy/Dockerfile .
service_two_docker: check mini-eval
	docker build -t $(SERVICE_TWO) --build-arg SERVICE=$(SERVICE_TWO) -f deploy/Dockerfile .
service_three_docker: check mini-eval
	docker build -t $(SERVICE_THREE) --build-arg SERVICE=$(SERVICE_THREE) -f deploy/Dockerfile .


# Some other random stuff
deploy_otel:
	kubectl apply -f deploy/otel/base.yaml

deploy: deploy_otel


# A small helper function to compile out the protobuf definitions
# You'll need 
.PHONY: proto
proto: check
	[ -d "protovendor" ] || git clone https://github.com/googleapis/googleapis $(GOOGLEAPIS)
	protoc -I. \
	--include_imports \
    --include_source_info \
	--go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	--proto_path=$(GOOGLEAPIS) \
	--proto_path=. \
	proto/*.proto

mini-eval:
	eval $$(minikube -p minikube docker-env)