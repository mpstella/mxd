# mxd
A deployment wrapper for gcloud

```
$> go run main.go
mxd is a gcloud wrapper that allows options to be JSON.

Usage:
  mxd [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  functions   Operations for cloud functions
  help        Help about any command
  version     Print Vlad's version number

Flags:
  -h, --help      help for mxd
  -v, --verbose   verbose output

Use "mxd [command] --help" for more information about a command.
```

## functions 

#### list
```
$> go run main.go functions list
NAME                 STATUS  TRIGGER        REGION
bob                  ACTIVE  HTTP Trigger   australia-southeast1
espv2-cf-hello-gcs   ACTIVE  Event Trigger  australia-southeast1
espv2-cf-hello-http  ACTIVE  HTTP Trigger   australia-southeast1

```

#### deploy
---
Sample config

```json
{
  "region": "australia-southeast1",
  "entry-point": "hello_http",
  "runtime": "python39",
  "opts": [
    "allow-unauthenticated",
    "trigger-http"
  ],
  "update-labels": {
    "foo": "bar",
    "baz": "bleh"
  },
  "remove-labels": ["abb", "bb"]
}
```
---

Deploying

```
$> go run main.go functions deploy
Error: required flag(s) "config", "name" not set
Usage:
  mxd functions deploy <function-name> <function-config> <source> [flags]


Flags:
  -c, --config string   path to the configuration
  -h, --help            help for deploy
  -n, --name string     name of the function
  -s, --source string   path to the source

Global Flags:
  -v, --verbose   verbose output
```
