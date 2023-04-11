---
title: "Introduction"
description: "Docs intro"
---

**Welcome to ${import.meta.env.PUBLIC_REPOSITORY}!**

This is the `docs` starter template. It contains all of the features that you need to build a Markdown-powered documentation site, including:

- ✅ **Full Markdown support**
- ✅ **Responsive mobile-friendly design**
- ✅ **Sidebar navigation**
- ✅ **Search (powered by Algolia)**
- ✅ **Multi-language i18n**
- ✅ **Automatic table of contents**
- ✅ **Automatic list of contributors**
- ✅ (and, best of all) **dark mode**

## Getting Started

To get started with this theme, check out the `README.md` in your new project directory. It provides documentation on how to use and customize this template for your own project. Keep the README around so that you can always refer back to it as you build.

Found a missing feature that you can't live without? Please suggest it [on our Discord](https://astro.build/chat) and even consider adding it yourself on GitHub! Astro is an open source project and contributions from developers like you are how we grow!

Good luck out there, Astronaut. 🧑‍🚀
# Template for Go tools working with cli

## Usage

## Easy Installation

### Binary Releases

We are now shipping binaries for each of the releases so that you don't even have to build them yourself! How wonderful is that!

If you're stupid enough to trust binaries that I've put together, you can download them from the [releases](https://github.com/LeoFVO/<tool_name>/releases) page.

### Using `go install`

If you have a [Go](https://golang.org/) environment ready to go (at least go 1.19), it's as easy as:

```bash
go install github.com/LeoFVO/<tool_name>@latest
```

PS: You need at least go 1.19 to compile <tool_name>.

### Using Docker

```bash
docker pull ghcr.io/leofvo/<tool_name>:latest
docker run <tool_name>:latest
```

### Building From Source

Since this tool is written in [Go](https://golang.org/) you need to install the Go language/compiler/etc. Full details of installation and set up can be found [on the Go language website](https://golang.org/doc/install). Once installed you have two options. You need at least go 1.19 to compile <tool_name>.

### Compiling

`<tool_name>` has external dependencies, and so they need to be pulled in first:

```bash
go get && go build
```

This will create a `<tool_name>` binary for you. If you want to install it in the `$GOPATH/bin` folder you can run:

```bash
go install
```

# License

See the LICENSE file.
