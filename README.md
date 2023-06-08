# GoSSTI

GoSSTI is a SSTI scanner for web application. Developed in Go.

## Get started

### First usage and/or update

````bash
gossti plugins update

### Basic Usage

**USAGE:**
`gossti detect -u <URL>`

**OPTIONS:**

-h, --help help for detect command
-u, --url string The target IP or domain to scan
-C, --cookies strings Cookies to use (e.g. -C 'cookie1=value1; cookie2=value2')
-X, --method string The HTTP method to use (default "GET")
--user-agent string Custom user-agent to use (default "gossti 1.0.0")
--timeout duration Timeout for HTTP requests (e.g. 10s)

### Advanced Usage with forms

**USAGE:**
`gossti detect -u <URL> -X POST --form 'field1=value1,field2=value2'`
**OPTIONS:**
--form strings Form fields to use (e.g. --form 'field1=value1,field2=value2')
--form-item stringToString Form field to use (e.g. --form 'field1=value1' --form 'field2=value2') (default [])
--form-type string Form type to use (e.g. urlencoded, multipart) (default "urlencoded")

### Some examples

Using GET method:

`gossti -u http://example.com/something?name=SSTI`

Using POST method and only NodeJS:

`gossti -u http://example.com/something?name=SSTI -X POST --form 'field1=value1,field2=value2' --form-type multipart`

Using PUT method and custom user-agent:

`gossti -u http://example.com/something?name=SSTI -X PUT --user-agent "custom-agent 1.0"`

## Using custom payloads

### Document format

Each plugins concern a single language, all plugins can be found in the plugins folder named like this: `<language>.yml`.

The document format is the following:

```yaml
version: 1.0.0
name: Python
engines:
  - name: Mako
    payloads:
      - id: python_mako_ecA9Ba9885
        payload: ${21389+219839}
        response:
          expected: 241228
          invalidate: ${241228}
          error: false
````

- `version`: The version of the plugin
- `name`: The name of the language
- `engines`: The template engines of the language
  - `name`: The name of the template engine
  - `payloads`: The payloads to use
    - `id`: The ID of the payload
    - `payload`: The payload to use
    - `response`: The response to expect
      - `expected`: The expected response
      - `invalidate`: The payload to use to invalidate the cache
      - `error`: If the response is an error

To add a new payload, you can add a new payload in the `payloads` section of the plugin.

Define the payload ID like this: `<language>_<template_engine>_<random_string>`

You can generate a random string with the following command:

```bash
python3 -c "import random; import string; print(''.join(random.choice(string.hexdigits) for i in range(10)))"
```

## Installation

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

#### Prerequisites

Since this tool is written in [Go](https://golang.org/) you need to install the Go language/compiler/etc. Full details of installation and set up can be found [on the Go language website](https://golang.org/doc/install). Once installed you have two options. You need at least go 1.19 to compile gossti.

#### Clone the repository

```bash
git clone git@github.com:LeoFVO/gossti.git
```

#### Compiling

`gossti` has external dependencies, and so they need to be pulled in first:

```bash
go get && go build
```

This will create a `gossti` binary for you. If you want to install it in the `$GOPATH/bin` folder you can run:

```bash
go install
```

## Setup

### Documentation

The documentation is available at [https://leofvo.github.io/gossti/](https://leofvo.github.io/gossti/).

In order to deploy documentation for your project, you need to allow github actions to deploy github pages. To do so, go to your repository settings > Pages, and in the `Build and deployment` section, select `Github Actions` as the source.

# License

See the LICENSE file.
