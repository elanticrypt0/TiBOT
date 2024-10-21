#/bin/bash

# generate build for linux
go build -ldflags "-s -w" -o tibot.bin
chmod +x tibot.bin
mv ./tibot.bin ./build

# copy config into build directory
cp -R ./.env ./build
# copy config into build directory
cp -R ./config ./build
# copy scripts into build directory
cp -R ./scripts ./build
