all: out/day0 out/day1 out/day2

clean:
	rm -rf out/

out:
	mkdir out

out/day0: out cmd/day0/day0.go
	echo "### Running Day 0 ###"
#	go run cmd/day0/day0.go   # while developing
	go run cmd/day0/day0.go > out/day0	# when done

out/day1: out cmd/day1/day1.go
	echo "### Testing Day 1 ###"
	go test adventofcode/year2025/cmd/day1
	echo "### Running Day 1 ###"
#	go run cmd/day1/day1.go   # while developing
	go run cmd/day1/day1.go > out/day1	# when done

out/day2: out cmd/day2
	echo "### Testing Day 2 ###"
	go test adventofcode/year2025/cmd/day2
	echo "### Running Day 2 ###"
#	go run cmd/day2/day2.go   # while developing
	go run cmd/day2/day2.go > out/day2	# when done
