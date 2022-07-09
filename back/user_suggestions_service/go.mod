module github.com/XWS-2022-Tim12/Dislinkt/back/user_suggestions_service

go 1.17

replace github.com/XWS-2022-Tim12/Dislinkt/back/common => ../common

replace github.com/XWS-2022-Tim12/Dislinkt/back/user_suggestions_service/domain => ../user_suggestions_service/domain

replace github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/user_suggestions_service => ../common/proto/user_suggestions_service

replace github.com/XWS-2022-Tim12/Dislinkt/back/user_suggestions_service/application => ../user_suggestions_service/application

replace github.com/XWS-2022-Tim12/Dislinkt/back/user_suggestions_service/infrastructure/api => ../user_suggestions_service/infrastructure/api

replace github.com/XWS-2022-Tim12/Dislinkt/back/user_suggestions_service/infrastructure/persistence => ../user_suggestions_service/infrastructure/persistence

replace github.com/XWS-2022-Tim12/Dislinkt/back/user_suggestions_service/startup/config => ../user_suggestions_service/startup/config

replace github.com/XWS-2022-Tim12/Dislinkt/back/user_suggestions_service/startup => ../user_suggestions_service/startup

require (
	github.com/neo4j/neo4j-go-driver/v4 v4.4.3
	google.golang.org/grpc v1.45.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.9.0 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
	google.golang.org/genproto v0.0.0-20220314164441-57ef72a4c106 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

require (
	github.com/XWS-2022-Tim12/Dislinkt/back/common v1.0.0
	golang.org/x/text v0.3.7 // indirect
)
