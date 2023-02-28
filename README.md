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
-X <HTTP_METHOD> The http method to use [default: GET]  
--user-agent <USER_AGENT> Custom user-agent to use [default: "gossti 1.0.0"]  
-q, --quiet Less output per occurrence  
-v, --verbose More output per occurrence  
-V, --version Print version information

### Some examples

Using GET method:

`gossti -u http://example.com/something -p "name=SSTI"`

Using POST method:

`gossti -u http://example.com/something -X POST -p "name=SSTI"`

Using PUT method:

`gossti -u http://example.com/something -X PUT -p "name=SSTI"`

Using PATCH method:

`gossti -u http://example.com/something -X PATCH -p "name=SSTI"`

Using DELETE method:

`gossti -u http://example.com/something -X DELETE -p "name=SSTI"`

## Objectives

- [x] Make a CLI using clap.rs
- [x] Verbose output
- [x] Add .gitignore
- [x] Add ci/cd
- [x] allow custom user-agent
- [ ] Make decision tree algorithm
- [ ] Parallel scanning
- [ ] allow custom headers
- [ ] allow custom cookies
- [ ] command executor or os-shell like

## Contributing

If you have any questions or suggestions, please open an issue or open a pull request.
Objectives are always open to discussion.
Objectives are prioritized and will be addressed as soon as possible.

Thanks for all the contributions!
