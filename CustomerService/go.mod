module Bitcoin/CustomerService

go 1.17

replace Bitcoin/GRPCMessage => ../GRPCMessage

require (
	Bitcoin/GRPCMessage v0.0.0-00010101000000-000000000000
	github.com/segmentio/ksuid v1.0.4
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
	google.golang.org/grpc v1.40.0
)

require (
	github.com/golang/protobuf v1.5.0 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
