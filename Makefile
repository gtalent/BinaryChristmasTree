main:
	go build -o main
clean:
	go clean
run: main
	./main
