> **Note**
>
> This repo is a example for this article https://arisharyanto.medium.com/best-way-to-structuring-golang-code-6e619e70ce38

## How to Run

### Run environment
```shel
$ docker-compose up -d
```

### Install GRPC
follow this http://google.github.io/proto-lens/installing-protoc.html to install protobuf in your PC\
\
and then run this in you cli
```shel
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Activate Proto Shell
open bash_profile file
```shel
$ nano ~/.bash_profile
```

add this code inside
```shel
export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
```

then save.

and then run this
```shel
$ source ~/.bash_profile
```

### Generate Your Proto file
```shel
$ protoc -I./proto --go_out=./proto --go-grpc_out=./proto ./proto/*.proto
```

### Run the code
```
$ go run cmd/*.go
```