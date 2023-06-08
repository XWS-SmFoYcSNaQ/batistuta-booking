module github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service

go 1.19

replace github.com/XWS-SmFoYcSNaQ/batistuta-booking/common => ../common

require (
	github.com/XWS-SmFoYcSNaQ/batistuta-booking/common v1.0.0
	github.com/google/uuid v1.3.0
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	google.golang.org/grpc v1.50.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.9.0 // indirect
	golang.org/x/net v0.1.0 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20221018160656-63c7b68cfc55 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
