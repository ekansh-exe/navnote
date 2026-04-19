build:
	go build -o navnote .

run:
	go run .

install:
	go install .

clean:
	rm -f navnote

.PHONY: build run install clean