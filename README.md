## Runtime Coverage Testing (POC)

This is a proof of concept for instrumenting Go services with runtime coverage testing. It works by creating a single unit test that calls `main()` and building a Go test binary to invoke it.

On exiting the `TestRunMain()` unit test, a coverage file is generated. Note that this requires a successful exit of main. If the main process exits with an error, a coverage file will not be generated

## Running the basic sample `main.go`:
```bash
# Generate the Go test binary for this service
make bin

# Run the test binary
make test

# Print a report of the output coverage file
make report
```

## Running the HTTP server example (Not yet working)
***NOTE: NOT WORKING***

#### Issues
- Service currently exits with an error & no coverage report is generated


```bash
# Replace the simple entrypoint with the HTTP server sample
cp archive/main_http.go ./main.go

# Generate the Go test binary for this service
make bin

# Run the test binary
# By default, the service stops itself after 10 seconds of execution. This can be disabled in main.go
make test

# Run a few sample curl calls to excercise the service logic
# Note that this requires the server launched in `make test` to still be running
make curl

# Print a report of the output coverage file
make report
```

