all: clean  yaccdata build test

test:
	go test -v
build:
	go build
clean:
	rm -f ./yaccpragma.y.go

yaccdata:
	goyacc  -o ./yaccpragma.y.go  -p "YaccPragma"  ./yaccpragma.y
