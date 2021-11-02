module chat_client

go 1.15

require (
	github.com/evefox/chat/protos v0.0.1
	github.com/evefox/chat/tool v0.0.1
	github.com/golang/protobuf v1.5.2
)

replace github.com/evefox/chat/protos v0.0.1 => ../protos

replace github.com/evefox/chat/tool v0.0.1 => ../tool
