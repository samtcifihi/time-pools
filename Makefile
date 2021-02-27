compile:
	go build -o "./bin/time-pools" ./src/main

test: compile
	./bin/time-pools

clean:
	rm ./bin/time-pools
