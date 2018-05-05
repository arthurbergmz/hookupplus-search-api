protoc -I=src/ --go_out=src/ src/messaging/channels/handshake/handshake.proto
protoc -I=src/ --go_out=src/ src/messaging/channels/event/event.proto
protoc -I=src/ --go_out=src/ src/messaging/channels/user/user.proto
protoc -I=src/ --go_out=src/ src/messaging/index.proto