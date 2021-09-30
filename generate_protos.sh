cd protos
rm -r go
mkdir go
protoc --go_out=go *.proto
