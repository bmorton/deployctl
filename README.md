# deployctl

`deployctl` is the command line client for Deployster.

```ShellSession
$ deployctl help
NAME:
   deployctl - A simple command line client for deployster.

USAGE:
   deployctl [global options] command [command options] [arguments...]

VERSION:
   v0.1.0

AUTHOR:
  Brian Morton - <brian@mmm.hm>

COMMANDS:
   deploy, d   deploy a given service and version to the cluster
   run, r   run a task for a given service and version
   destroy, x  destroy all instances of a given service and version running on the cluster
   list, l  list all instances of a given service running on the cluster
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --url "http://localhost:3000" a URL to the deployster instance [$DEPLOYSTER_URL]
   --username "deployster"    username for authenticating to deployster [$DEPLOYSTER_USERNAME]
   --password           password for authenticating to deployster [$DEPLOYSTER_PASSWORD]
   --help, -h           show help
   --version, -v        print the version
```
