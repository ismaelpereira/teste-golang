build: clean
		go build -v -o dist/bin/crud-people -ldflags="-s -w" main.go
		

run: build
		./dist/bin/crud-people

clean:
		rm -rf dist/