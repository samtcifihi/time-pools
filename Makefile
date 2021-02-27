compile:
	go build -o "./bin/time-pools" ./src/main

run:
	./bin/time-pools

test: compile
	./bin/time-pools

clean:
	rm ./bin/time-pools
