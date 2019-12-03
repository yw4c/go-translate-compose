protoc -I ./ --go_out=plugins=grpc:. ./github.com/yw4c/trans-pb/p010user/v1/user.proto
protoc -I ./ --go_out=plugins=grpc:. ./github.com/yw4c/trans-pb/p012dict/v1/dict.proto
protoc -I ./ --go_out=plugins=grpc:. ./github.com/yw4c/trans-pb/p013keep/v1/keep.proto