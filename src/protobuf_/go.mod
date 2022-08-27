module protobuf_

go 1.19

require (
	github.com/golang/protobuf v1.5.0
	myprotopb v0.0.0
)

require google.golang.org/protobuf v1.28.1 // indirect

replace myprotopb => ./myprotopb
