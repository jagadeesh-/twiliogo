[![Build Status](https://travis-ci.org/carlosdp/twiliogo.png?branch=master)](https://travis-ci.org/carlosdp/twiliogo)
# twilio-go
Forked from github.com/carlosdp/twiliogo and added UsageRecords API.



# Installation

``` bash
go get github.com/jagadeesh-/twiliogo
```

# Documentation
[GoDoc](http://godoc.org/github.com/carlosdp/twiliogo)

# Usage






## Implemented Resources
- Usage Records

## Run Tests
Tests can be run using `go test`, as with most golang projects. This project also contains integration tests (where they can be done non-destructively using the API or the working Test Credential endpoints).

These integration tests can be run by providing the necessary environment variables to access the API, as in this Makefile:

```makefile
test:
	@export API_KEY="<API KEY>";\
	export API_TOKEN="<API TOKEN>";\
	export TEST_KEY="<TEST KEY>";\
	export TEST_TOKEN="<TEST TOKEN>";\
	export TEST_FROM_NUMBER="<DEFAULT TEST CRED NUMBER>";\
	export FROM_NUMBER="<TEST FROM NUMBER>";\
	export TO_NUMBER="<TEST TO NUMBER>";\
	go test -v
```





