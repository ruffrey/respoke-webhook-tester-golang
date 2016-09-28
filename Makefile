clean:
	rm -rf build/
build-all: build-osx build-linux build-win
	chmod +x build/*

build-osx:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o build/osx/amd64/rhook
	chmod +x build/osx/amd64/rhook
	zip -j build/osx-64-rhook.zip build/osx/amd64/rhook
	rm -rf build/osx

build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o build/linux/amd64/rhook
	GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o build/linux/386/rhook
	chmod +x build/linux/amd64/rhook
	chmod +x build/linux/386/rhook
	zip -j build/linux-64-rhook.zip build/linux/amd64/rhook
	zip -j build/linux-86-rhook.zip build/linux/386/rhook
	rm -rf build/linux

build-win:
	GOOS=windows GOARCH=amd64 go build -o build/windows/amd64/rhook.exe
	zip -j build/windows-64-rhook.zip build/windows/amd64/rhook.exe
	rm -rf build/windows
