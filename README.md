# jbls

A simple jb license server.

## Build

```bash
# normal build
go build -ldflags "-s -w"
```

```bash
# build for mips (golang 1.10 or greater)
GOOS=linux GOARCH=mips GOMIPS=softfloat go build -ldflags "-s -w"
```

## Usage

```text
Usage of jbls:
  -host string
        Bind your ip address. (default "127.0.0.1")
  -key string
        Private key file path for the license server.
  -name string
        Give a fixed name to user. (optional)
  -port string
        Bind your port. (default "8080")
```

### Note that

The private key is pem encoded key file that you need to get it on your own :)
