# td-app
Traffic Director Demo - App

Getting started
1. Install protobuf
```brew install protobuf```
2. Install Go
```brew install go```
```go env GOPATH```
```echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.zshrc``` (only once!)
```source ~/.zshrc```
3. Install the Go protobuf plugins
```go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34```
```go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.4```
4. Make sure you are able to build the app:
```protoc --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false proto/helloworld.proto```
5. Generate go.sum
```go mod tidy```
6. Install docker
```brew install docker```
(consider Docker-Desktop if you prefer UI)
7. build the docker image (assuming docker daemon is running)
```docker build -t td-app-app -f server/Dockerfile .```
8. OPTIONAL: run the server locally
```docker run --rm -p 50051:50051 td-app-server```
9. OPTIONAL: tag and push to GCR
First Authenticate:
```gcloud auth configure-docker```
```gcloud auth login```
```gcloud config set project <project-id>```
Tag and Push:
```docker tag td-app-app gcr.io/<project-id>/td-app-app:latest```
```docker push gcr.io/<project-id>/td-app-app:latest``` 
(assuming authenticated)

