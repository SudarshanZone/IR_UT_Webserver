module github.com/krishnakashyap0704/microservices/webserver

go 1.22.5

toolchain go1.22.6

require (
	github.com/gin-contrib/cors v1.7.2
	github.com/gin-gonic/gin v1.10.0
	github.com/krishnakashyap0704/microservices/comOpenPositions v0.0.0-00010101000000-000000000000
	github.com/krishnakashyap0704/microservices/comOrderDetails v0.0.0-00010101000000-000000000000
	github.com/krishnakashyap0704/microservices/comSquareOff v0.0.0-00010101000000-000000000000
	github.com/krishnakashyap0704/microservices/equOpenPositions v0.0.0-00010101000000-000000000000
	github.com/krishnakashyap0704/microservices/equOrderDetails v0.0.0-00010101000000-000000000000
	github.com/krishnakashyap0704/microservices/equSquareOff v0.0.0-00010101000000-000000000000
	github.com/krishnakashyap0704/microservices/fnoOpenPositions v0.0.0-00010101000000-000000000000
	github.com/krishnakashyap0704/microservices/fnoOrderDetails v0.0.0-00010101000000-000000000000
	github.com/krishnakashyap0704/microservices/fnoSquareOff v0.0.0-00010101000000-000000000000
	github.com/krishnakashyap0704/microservices/tradeRecord v0.0.0-00010101000000-000000000000
	github.com/krishnakashyap0704/microservices/userlogin v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.66.0
)

require (
	github.com/bytedance/sonic v1.11.6 // indirect
	github.com/bytedance/sonic/loader v0.1.1 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.20.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.8.0 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.23.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240604185151-ef581f913117 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/krishnakashyap0704/microservices/userlogin => ../userlogin

replace github.com/krishnakashyap0704/microservices/fnoOrderDetails => ../fnoOrderDetails

replace github.com/krishnakashyap0704/microservices/fnoOpenPositions => ../fnoOpenPositions

replace github.com/krishnakashyap0704/microservices/fnoSquareOff => ../fnoSquareOff

replace github.com/krishnakashyap0704/microservices/comOrderDetails => ../comOrderDetails

replace github.com/krishnakashyap0704/microservices/comOpenPositions => ../comOpenPositions

replace github.com/krishnakashyap0704/microservices/comSquareOff => ../comSquareOff

replace github.com/krishnakashyap0704/microservices/equOrderDetails => ../equOrderDetails

replace github.com/krishnakashyap0704/microservices/equOpenPositions => ../equOpenPositions

replace github.com/krishnakashyap0704/microservices/equSquareOff => ../equSquareOff

replace github.com/krishnakashyap0704/microservices/tradeRecord => ../tradeRecord
