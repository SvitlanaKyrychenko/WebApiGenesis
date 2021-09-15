module Bitcoin/Api

go 1.15

require (
	Bitcoin/GRPCMessage v0.0.0-00010101000000-000000000000
	github.com/segmentio/ksuid v1.0.4
	google.golang.org/grpc v1.40.0
)

replace Bitcoin/GRPCMessage => ../GRPCMessage
