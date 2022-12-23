> **Note**
> This repo is a example for this article [https://arisharyanto.medium.com/best-way-to-structuring-golang-code-6e619e70ce38](https://arisharyanto.medium.com/best-way-to-structuring-golang-code-6e619e70ce38){:target="_blank"

## How to Run

### Run environment
```shel
$ docker-compose up -d
```

### Install GRPC
```shel
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Activate Proto Shell
```shel
#open file
nano ~/.bash_profile

#write this code
export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin

#save
```

### Type this
```shel
source ~/.bash_profile
```

### Generate Your Proto file
```shel
protoc -I./proto --go_out=./proto --go-grpc_out=./proto ./proto/*.proto
```

### Run the code
```
go run cmd/*.go
```