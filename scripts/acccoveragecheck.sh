#!/bin/bash

set -e -u

go mod vendor

# TODO remove run flag once acc. tests job setup is done
go test ./illumio-core -run="TestProvider" -coverprofile=cover.out

perc=`go tool cover -func=cover.out | tail -n 1 | sed -Ee 's!^[^[:digit:]]+([[:digit:]]+(\.[[:digit:]]+)?)%$!\1!'`
res=${perc%.*}
if [[ "$res" -ge "70" ]];
  then
    echo "Coverage: $perc PASS"
  else
    echo "Coverage: $perc FAIL" >&2
    exit 1
fi