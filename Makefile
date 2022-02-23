build: clean
		go build -o /crud-people 
		

run: build
		./dist/bin/crud-people

clean:
		rm -rf dist/