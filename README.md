# nova-cli

A convenient command line tool to pipe logs to [splunknova.com](https://www.splunknova.com) and search them.

# Usage

## Sending logs
````
echo "my first log" | nova

cat /var/log/system.log | nova

tail -f /var/log/system.log | nova
````

## Searching logs
````
nova ERROR

nova ERROR -s count

nova "my_key=" -r "stats count avg(my_key)"
````

# Installation

## macOS

````
brew tap splunknova/nova-cli
brew install nova-cli
````

## Linux & Windows

You'll need to ensure `go` is installed and `GOROOT`, `GOPATH`, and `GOBIN` are set.

````
go get github.com/splunknova/nova-cli
cd $GOPATH/src/github.com/splunknova/nova-cli
go install nova.go // will install to $GOBIN
````

_Help with making Linux and Windows installations better will be appreciated!_
