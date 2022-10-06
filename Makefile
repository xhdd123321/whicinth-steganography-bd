BIN_FILE=output/whicinth-steganography-bd

hello:
	echo "Hello"

build:
	go build -o ${BIN_FILE}

run:
	./${BIN_FILE}

build&run: build run
