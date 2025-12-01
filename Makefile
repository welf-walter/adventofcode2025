all: out/day0

clean:
	rm -rf out/

out:
	mkdir out

out/day0: out cmd/day0/day0.go
#	go run cmd/day0/day0.go   # while developing
	go run cmd/day0/day0.go > out/day0	# when done