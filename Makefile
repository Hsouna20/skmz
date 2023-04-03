run:
	go run main.go
unit-tests:
	go test skmz/server
functional-tests:
	go test ./functional_tests/transformer_test.go
