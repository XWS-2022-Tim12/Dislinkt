module github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service

go 1.17

replace github.com/XWS-2022-Tim12/Dislinkt/back/common => ../common

replace github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/domain => ../job_suggestions_service/domain

replace github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/job_suggestions_service => ../common/proto/job_suggestions_service

replace github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/application => ../job_suggestions_service/application

replace github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/infrastructure/api => ../job_suggestions_service/infrastructure/api

replace github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/infrastructure/persistence => ../job_suggestions_service/infrastructure/persistence

replace github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/startup/config => ../job_suggestions_service/startup/config

replace github.com/XWS-2022-Tim12/Dislinkt/back/job_suggestions_service/startup => ../job_suggestions_service/startup

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/neo4j/neo4j-go-driver/v4 v4.4.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	google.golang.org/grpc v1.45.0
)

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.9.0 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.4.0 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
	google.golang.org/genproto v0.0.0-20220314164441-57ef72a4c106 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

require (
	github.com/XWS-2022-Tim12/Dislinkt/back/common v1.0.0
	golang.org/x/text v0.3.7 // indirect
)
