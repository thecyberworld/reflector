# reflector

Reflector is a tool to check for reflected query parameter values.

#### Usage
Reflector is a command line tool that takes a single argument, which is the name of a file containing a list of URLs to check. The file should contain one URL per line.

```shell
reflector urls.txt
```

#### Output
Reflector will write the URLs for which query parameter values are reflected in the response body to a file named "reflected_urls.txt" in the current working directory. It will also print the URLs to the console.

#### Install

```shell
go install github.com/thecyberworld/reflector@latest
```
