fmt:
	go fmt ./...
run:fmt
	go run cmd/goplantuml/main.go $(ARGS)
remote-dialer:
	go run cmd/goplantuml/main.go -show-connection-labels=true -recursive=true -output="example/remotedialer.puml" "/Users/jiahua/k8s/remotedialer"
#plantuml:remote-dialer
#	plantuml "example/remotedialer.txt"