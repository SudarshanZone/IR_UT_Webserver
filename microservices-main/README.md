# Microservices
@ICICI Securities || IRRA Project || Application Server

### gRpc and ProtoBuffer Installation
ProtoBuffer: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

gRPC: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

Install proto from github https://github.com/protocolbuffers/protobuf/releases and set the path variable

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto

go run greeter_server/main.go

### Important Commands

Github:
1. git init
2. git remote add origin https://github.com/krishnakashyap0704/microservices.git 
3. git checkout -b "New_Branch"
4. git checkout Branch_Name (if another branch present)
5. git pull origin Branch_Name
6. git add . 
7. git commit -m "Message"
8. git push -u origin Branch_Name

Extra: 
1. git status 
2. git log
3. git remote -v
4. git remote remove origin
5. git pull --rebase origin Branch_Name

---

### Initial Project Command

1. go mod init github.com/krishnakashyap0704/microservices
2. go get google.golang.org/grpc (used to import)
3. go mod tidy (used to change indirect to direct)
4. go run to/path/main.go

### Components

Userlogin:
1. protoc --proto_path=internal --go_out=generated --go-grpc_out=generated internal\usrlogin\userlogin.proto
2. go run .\cmd\server.go
3. go run .\webserver\userlogin\client.go
   
FnoOpenPositions:
1. protoc --proto_path=internal --go_out=generated --go-grpc_out=generated internal\fnoopn\fno_pos.proto
2. go run .\cmd\server.go
3. go run .\webserver\fnoOpenPositions\client.go
   
FnoOrderDetails:
1. protoc --proto_path=internal --go_out=generated --go-grpc_out=generated internal\fnoordr\OrderDtls.proto
2. go run .\cmd\server.go
3. go run .\webserver\fnoOrderDetails\client.go
   
FnoSquareOff:
1. protoc --proto_path=internal --go_out=generated --go-grpc_out=generated internal\fnosquoff\fnoSquareOff.proto
2. go run .\cmd\server.go
3. go run .\webserver\fnoSquareOff\client.go

ComOpenPositions:
1. protoc --proto_path=internal --go_out=generated --go-grpc_out=generated internal\comopn\comOpenPositions.proto
2. go run .\cmd\server.go
3. go run .\webserver\comOpenPositions\client.go

ComOrderDetails:
1. protoc --proto_path=internal --go_out=generated --go-grpc_out=generated internal\comordr\comOrderDetails.proto
2. go run .\cmd\server.go
3. go run .\webserver\comOrderDetails\client.go

ComSquareOff:
1. protoc --proto_path=internal --go_out=generated --go-grpc_out=generated internal\comsquoff\comSquareOff.proto
2. go run .\cmd\server.go
3. go run .\webserver\comSquareOff\client.go
   
EquOpenPositions:
1. protoc --proto_path=internal --go_out=generated --go-grpc_out=generated internal\equopn\equOpenPositions.proto
2. go run .\cmd\server.go
3. go run .\webserver\equOpenPositions\client.go
   
EquOrderDetails:
1. protoc --proto_path=internal --go_out=generated --go-grpc_out=generated internal\equordr\equOrderDetails.proto
2. go run .\cmd\server.go
3. go run .\webserver\equOrderDetails\client.go

EquSquareOff:
1. protoc --proto_path=internal --go_out=generated --go-grpc_out=generated internal\equsquoff\equSquareOff.proto
2. go run .\cmd\server.go
3. go run .\webserver\equSquareOff\client.go

TradeRecord:
1. protoc --proto_path=internal --go_out=generated --go-grpc_out=generated internal\trdrec\tradeRecord.proto
2. go run .\cmd\server.go
3. go run .\webserver\tradeRecord\client.go
   
---

### Database Connection

1. go get github.com/lib/pq
2. go run to/path/db.go

---

### WebServer

1. go mod init github.com/krishnakashyap0704/microservices/webserver
2. go get google.golang.org/grpc (used to import)
3. go get github.com/gin-gonic/gin (used to import)
4. go mod tidy (used to change indirect to direct)
5. go run to/path/main.go
   
---

### Replacement Commands

1. go mod edit "replace=github.com/krishnakashyap0704/microservices/userlogin=../userlogin"
   
2. go mod edit "replace=github.com/krishnakashyap0704/microservices/fnoOrderDetails=../fnoOrderDetails"
   
3. go mod edit "replace=github.com/krishnakashyap0704/microservices/fnoOpenPositions=../fnoOpenPositions"
   
4. go mod edit "replace=github.com/krishnakashyap0704/microservices/fnoSquareOff=../fnoSquareOff"
   
5. go mod edit "replace=github.com/krishnakashyap0704/microservices/comOrderDetails=../comOrderDetails"
   
6. go mod edit "replace=github.com/krishnakashyap0704/microservices/comOpenPositions=../comOpenPositions"
   
7. go mod edit "replace=github.com/krishnakashyap0704/microservices/comSquareOff=../comSquareOff"
   
8. go mod edit "replace=github.com/krishnakashyap0704/microservices/equOrderDetails=../equOrderDetails"
   
9.  go mod edit "replace=github.com/krishnakashyap0704/microservices/equOpenPositions=../equOpenPositions"
    
10. go mod edit "replace=github.com/krishnakashyap0704/microservices/equSquareOff=../equSquareOff"
    
11. go mod edit "replace=github.com/krishnakashyap0704/microservices/tradeRecord=../tradeRecord"


### How to push code on SERVER

1. Login the ARCOS and Connect to the required Server.

2. Two ways to push the packages into the server 
    
    1. Build a packages using go build command.
   
        go build -o [PathWhereFileCreate] [PathOfServer.go]
    
        Example: go build  -o ./packages ./cmd/server.go

    2. Download Dependencies locally.
    
        go mod vendor

3. Upload the code in the specific directory.

4. Install the Packages

5. Create a Executable file in the specific folder.
   
    go build -o [NAMEOFEXECUTABLEFILE] [PathWhereServerFilePresent]

    Example:  go build -o fnoSquareOff.exe ./cmd
   
6. Run the Executable file
