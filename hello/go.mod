module hello

go 1.15

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.2 // indirect
	rsc.io/quote v1.5.2 // indirect
)

replace example.com/greetings => ../greetings
