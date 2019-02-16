# ein
a simple grpc demo

## install

```script=bash
make build  # build server
make client # build client
```

## run

* server

```script=bash
ein up # start the server
```

* client

```script=bash
$ build/ZhangYet/ein-client       
NAME:
   ein-client - A new cli application

USAGE:
   ein-client [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     getAllLastQuotes  get all quotes
     getLastQuote      get one quote
     daemon            run in daemon and waiting for update info
     help, h           Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

```