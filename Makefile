BIN_FILE=whicinth-steganography-bd

hello:
	echo "Hello"

build:
	go build -o ${BIN_FILE}

run:
	./${BIN_FILE}

start:
	nohup make run > output/log.txt 2> output/error.txt &

build&run: build run
