# Nova CLI

[![Travis Status for Nova CLI](https://travis-ci.org/splunknova/nova-cli.svg?branch=master)](https://travis-ci.org/splunknova/nova-cli)

A convenient command-line tool for sending and searching logs using [splunknova.com](https://www.splunknova.com).

## Get Started

Get started by creating an account on [www.splunknova.com](https://www.splunknova.com).

## Install

You can install Nova-CLI or download the [binaries] directly.

### macOS

[Homebrew](https://brew.sh/) is a package manager for macOS that makes it easy to install Nova CLI. In a terminal window, run:

```bash
brew tap splunknova/nova-cli
brew install nova-cli
```

### Binaries

We may also have have binaries for download on the [latest release](https://github.com/splunknova/nova-cli/releases/latest).
Shout out on [slack](http://community.splunknova.com) if you need a particular binary!

### Linux & Windows

Set up your [Go] environment. Refer to the official Go documentation for more details: [https://golang.org/doc/code.html](https://golang.org/doc/code.html). Once Go is downloaded and installed, you'll need to set your `GOROOT`, `GOPATH`, and `GOBIN`.

**Linux**: By default Go is installed to directory `/usr/local/go/`, and the `GOROOT` environment variable is set to `/usr/local/go/bin`.

**Windows**: By default Go is installed in the directory `C:\Go`, the `GOROOT` environment variable is set to `C:\Go\`, and the bin directory is added to your Path (`C:\Go\bin`).

To install Nova CLI using Go, in the command-line run:

```bash
go get github.com/splunknova/nova-cli
```

Change directories into the `nova-cli` repository.

```bash
cd $GOPATH/src/github.com/splunknova/nova-cli
```

Install the `nova` binary to `$GOBIN`.

```bash
go install nova.go
```

 If it isn't in your PATH, you can run `export PATH=$PATH:$GOBIN`

## Usage

If you haven't already, [sign up or log in][nova] to obtain your Splunk Nova API credentials to get started.

### Credentials

Get started by creating an account on [splunknova.com](https://www.splunknova.com/).

API Credentials can be conveniently saved in a `~/.nova` file by running

```bash
nova login
```

You will be prompted to enter your `Client ID` and `Client Secret`:

```bash
Please enter Client ID: <Your Client ID>
Please enter Client Secret: <Your Client Secret>
```

Once your credentials are entered, you should see:

```bash
INFO[0016] Login succeeded
```

## Send logs

You can pipe logs into Splunk Nova by running:

```bash
echo "my first log" | nova
```

This sends a log string: `"my first log"` to nova. You can then search your logs from the CLI using `nova search`. For example:

```bash
nova search "my first log"
```

returns a list of `my first logs` sent to nova:

```bash
2018-1-19T19:35:01.000+00:00 my first log
2018-1-18T23:52:38.000+00:00 my first cli log
```

One example of a `cat` command for system log files would be to pipe `system.log` to Splunk Nova:

```bash
cat /var/log/system.log | nova
```

and then search:

```bash
nova search system.log
```

Returns:

```bash
2017-12-21T00:02:07.000+00:00 \tASL Module "com.apple.authd" sharing output destination "/var/log/system.log" with ASL Module "com.apple.asl".
2017-12-21T00:02:07.000+00:00 \tASL Module "com.apple.authd" sharing output destination "/var/log/system.log" with ASL Module "com.apple.asl".
2017-12-18T23:54:15.000+00:00 \tASL Module "com.apple.authd" sharing output destination "/var/log/system.log" with ASL Module "com.apple.asl".
```

You can also enter the `tail` command, followed by the file you’d like to view, which prints lines from the end of the file in reverse order:

```bash
tail -f /var/log/system.log | nova
```

Use the -f or --follow flag after tail, to show a real-time, streaming output of a changing file. It keeps watch and prints further data as it appears.

## Search logs

Search all logs containing the word "error"

```bash
nova search error
```

Returns:

```bash
count 0
```

Count only the number of lines containing the word "error"

```bash
nova search error -c
```

Returns

```bash
count 0
```

The `-s` or `stats` command calculates aggregate statistics over the
results set, such as average, count, and sum.

```bash
nova search error -s count
```

With the `stats` command you can specify a statistical function such as `count` to create a report of all errors. (How is this different than error -c?)

```bash
nova search error -r "stats count"
```

Run stats aggregations and reporting on data using Splunk Processing Language (SPL) inspired syntax. For example:

```bash
nova search "my_key=" -r "stats count avg(my_key)"
```

Returns a go routine that reports all usages of your Splunk Nova API credentials.

```bash
example here
```

Add transforming commands, a type of search command that orders the results into a data table

```bash
nova search "bytes" -t "eval mb=gb*1024" -r "stats max(mb)"
````

Returns

```bash
example here
```

## Sending Metrics

A metric is a set of measurements containing a timestamp, a metric name, a value, and a dimension. See [Overview of metrics][overview] and [Get started with metrics][getstarted].

You can post metrics to your Splunk Nova account by running

Syntax:

```bash
nova metric put <metric_name> <metric_value>
```

Example:

```bash
nova metric put cpu.usage 20
```

## Tagging with dimensions

Provide metadata about the metric. For example:

- Region: `region:us-east-1, us-west-1, us-west-2, us-central1`
- Instance Types: `t2.medium, t2.large, m3.large, n1-highcpu-2`
- Technology: `nginx, redis, tomcat`

You can think of a metric name as something that you are measuring, while dimensions are categories by which you can filter or group the results.

Example:

```bash
nova metric put cpu.usage 20 -d "region:us-east-1,role:webserver"
```

## List Metrics

List all Metrics

Example:

```bash
nova metric ls
```

## Aggregate Metrics using Statistical Functions

Simple stats

Example:

```bash
nova metric get cpu.usage -s avg,max
```

Group by dimensions

Example:

```bash
nova metric get cpu.usage -s avg -g role
```

## How to export the new releases

In order to export the new releases, you would need to:

1. create a new git tag
1. export your GITHUB_TOKEN
1. run gorealser in nova-cli folder


[getstarted]: http://docs.splunk.com/Documentation/Splunk/7.0.1/Metrics/GetStarted
[Go]: https://golang.org/dl/
[homebrew]: https://brew.sh/
[nova]: https://www.splunknova.com/
[novalogin]: https://www.splunknova.com/login
[overview]: http://docs.splunk.com/Documentation/Splunk/7.0.1/Metrics/Overview
