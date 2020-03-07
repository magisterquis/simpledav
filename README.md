SimpleDAV
=========
Dead simple WebDAV server.  Does nearly nothing except
serve a filesystem over WebDAV.  Configuration is all at
compile-time using `-ldflags="-X main.foo=bar"`.

Due to a bug in the underlying WebDAV library, SimpleDAV
will not work very well if not run as root.  Sorry :(

Configurable Options
--------------------
Option       | Description                                               | Default
-------------|-----------------------------------------------------------|--------
`network`    | Network passed to `net.Listen`                            | `unix`
`address`    | Address passed to `net.Listen`                            | `/tmp/.wds`
`root`       | Root of the served directory tree                         | `/`
`allowWrite` | If not-empty, allows writing to the underlying filesystem | `""`
