module github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service

go 1.17

replace github.com/XWS-2022-Tim12/Dislinkt/back/common => ../common

replace github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/domain => ../authentification_service/domain

replace github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/authentification_service => ../common/proto/authentification_service

replace github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/application => ../authentification_service/application

replace github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/infrastructure/api => ../authentification_service/infrastructure/api

replace github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/infrastructure/persistence => ../authentification_service/infrastructure/persistence

replace github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/startup/config => ../authentification_service/startup/config

replace github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/startup => ../authentification_service/startup

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	go.mongodb.org/mongo-driver v1.8.4
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.9.0 // indirect
	github.com/nats-io/nats.go v1.13.1-0.20220308171302-2f2f6968e98d // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.4.0 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
	google.golang.org/genproto v0.0.0-20220314164441-57ef72a4c106 // indirect
)

require (
	github.com/XWS-2022-Tim12/Dislinkt/back/common v1.0.0
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce // indirect
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9 // indirect
	golang.org/x/text v0.3.7 // indirect
)
