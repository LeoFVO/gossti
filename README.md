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

## Contributing

If you have any questions or suggestions, please open an issue or open a pull request.
Objectives are always open to discussion.
Objectives are prioritized and will be addressed as soon as possible.

Thanks for all the contributions!
