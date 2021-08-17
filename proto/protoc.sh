~/go/bin/protoc -I proto proto/auth.proto --go_out=plugins=grpc:proto
~/go/bin/protoc -I proto proto/stats.proto --go_out=plugins=grpc:proto
~/go/bin/protoc -I proto proto/push.proto --go_out=plugins=grpc:proto
