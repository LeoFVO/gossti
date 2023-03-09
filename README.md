# GoSSTI

GoSSTI is a SSTI scanner for web application. Developed in Go.

## Get started

### Scan a single URL

**USAGE:**  
`gossti scan --url <URL> --payload <PAYLOAD>`

**OPTIONS:**

-h, --help Print help information  
-u, --url <URL> The target IP or domain to scan  
-p, --payload <PAYLOAD> The payload to trigger potential SSTI  
-X <HTTP_METHOD> The HTTP method to use [default: GET]  
-l, --languages comma-separated languages names to tests (default "all")  
--user-agent <USER_AGENT> Custom user-agent to use [default: "gossti 1.0.0"]

### Some examples

Using GET method:

`gossti -u http://example.com/something -p "name=SSTI"`

Using POST method and only NodeJS:

`gossti -u http://example.com/something -X POST -p "name=SSTI" --languages nodeJS`

Using PUT method and custom user-agent:

`gossti -u http://example.com/something -X PUT -p "name=SSTI" --user-agent "custom-agent 1.0"`

## Easy Installation

### Binary Releases

We are now shipping binaries for each of the releases so that you don't even have to build them yourself! How wonderful is that!

If you're stupid enough to trust binaries that I've put together, you can download them from the [releases](https://github.com/LeoFVO/gossti/releases) page.

### Using `go install`

If you have a [Go](https://golang.org/) environment ready to go (at least go 1.19), it's as easy as:

```bash
go install github.com/LeoFVO/gossti@latest
```

PS: You need at least go 1.19 to compile gossti.

### Using Docker

```bash
docker pull ghcr.io/leofvo/gossti:latest
docker run gossti:latest
```

### Building From Source

Since this tool is written in [Go](https://golang.org/) you need to install the Go language/compiler/etc. Full details of installation and set up can be found [on the Go language website](https://golang.org/doc/install). Once installed you have two options. You need at least go 1.19 to compile gossti.

### Compiling

`gossti` has external dependencies, and so they need to be pulled in first:

```bash
go get && go build
```

This will create a `gossti` binary for you. If you want to install it in the `$GOPATH/bin` folder you can run:

```bash
go install
```

# License

See the [LICENSE](./LICENSE) file.

## Contributing

See the [CONTRIBUTING.md](./CONTRIBUTING.md) file

Thanks for all the contributions!
