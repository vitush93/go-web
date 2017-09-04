
all: 
	go build

clean:
	rm ./go-web

run: all
	./go-web