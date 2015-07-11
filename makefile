all: build-crossplatform

clean:
	@rm -rf build

build-crossplatform:
	GOOS=windows GOARCH=amd64 go build -o build/win64/serve.exe serve.go && zip -j build/win64.zip build/win64/serve.exe
	GOOS=linux GOARCH=amd64 go build -o build/linux64/serve serve.go && zip -j build/linux64.zip build/linux64/serve
	GOOS=darwin go build -o build/osx/serve serve.go && zip -j build/osx.zip build/osx/serve

build:
	go build -o build/serve serve.go
