run:
	go run ./main.go

test:
	go test

benchmark:
	go test -bench=. -benchmem