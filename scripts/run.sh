#!/bin/zsh

cd ../backend/

rm test

go mod tidy

echo "Building the go-project"
go build -o test

echo "Running the project"
./test

echo "Deleting the binary"
rm test
