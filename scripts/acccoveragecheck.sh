!#/bin/sh

set -e -u

# TODO remove run flag once acc. tests job setup is done
go test ./illumio-core -run="TestProvider" -coverprofile=cover.out

perc=`go tool cover -func=cover.out | tail -n 1 | sed -Ee 's!^[^[:digit:]]+([[:digit:]]+(\.[[:digit:]]+)?)%$!\1!'`
res=`echo "$perc >= 70.0" | bc`
test "$res" -eq 1 && exit 0
echo "Insufficient coverage: $perc" >&2
exit 1