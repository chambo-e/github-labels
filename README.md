# github-labels

Small utility to manage github labels

## Installation

```bash
$ go get github.com/chambo-e/github-labels
```

## Usage

```bash
$ github-labels --help
NAME:
   github-labels - Manage github labels easily

USAGE:
   github-labels [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     list       Print labels from a repository
     set        Set repository labels from a file
     import     Import labels from a repository to another

GLOBAL OPTIONS:
   --token value, -t value      Github Token (Mandatory)
   --help, -h                   show help (default: false)
   --version, -v                print the version (default: false)
```

#### List labels

```bash
$ github-labels --token XXXXXXXXX list google/grpc
label: bug              color: #fc2929
label: duplicate        color: #cccccc
label: enhancement      color: #84b6eb
...
```

#### Import labels

```bash
$ github-labels --token XXXXXXXXX list google/grpc chambo-e/github-labels
1/30 (bug): Ok
2/30 (duplicate): Ok
3/30 (enhancement): Ok
...
```

#### Set labels

```json
// labels.json
[
    {"name": "label1", "color": "234567"}
    {"name": "label2", "color": "345678"},
    {"name": "label3", "color": "456789"},
    ...
]
```

```bash
$ github-labels --token XXXXXXXXX set --labels labels.json chambo-e/github-labels
1/30 (label1): Ok
2/30 (label2): Ok
3/30 (label3): Ok
...
```
