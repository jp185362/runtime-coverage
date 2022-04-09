
# Dead simple makefile to track these cool commands

# Build the go binary with coverage instrumentation
bin:
	go test -coverpkg="e2e/..." -c -tags testrunmain e2e

# Run the go coverage binary and output results to coverage file
test: bin
	./e2e.test -test.coverprofile coverage.out

curl: 
	sleep 1
	curl localhost:8080/basic
	curl localhost:8080/conditional
	curl localhost:8080/conditional?key=test

# Write a report of function coverage
report: test
	go tool cover -func=coverage.out