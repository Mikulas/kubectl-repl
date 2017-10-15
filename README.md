Kubectl REPL
============

Wrap `kubectl` with namespace and variables.

[![asciicast](https://asciinema.org/a/142536.png)](https://asciinema.org/a/142536)


Usage
-----

`./kubectl-repl` first starts by asking you for namespace. You may enter any of the strings verbatim,
or any abbreviation that is closest. Additionally, you may use any of the variables REPL assigned ($2).

Then you are in the main REPL mode. You are presented with a prompt, into which you enter `kubectl` commands
(`kubectl -n $NS` prefix is implied).

The prompt can be exited with traditional *eof* or *sigint*, and an explicit `quit` or `exit` command.


Shell integration
-----------------

Instead of directly invoking `kubectl` with prompt as arguments, `/bin/sh -c` is used instead. This
allows for more complex usage usage as `grep` and redirects:

```console
# sentry get pods | grep app
+ kubectl -n sentry get pods | grep app
$1 	app-deployment-314667899-4r9c1       1/1       Running   0          22h
$2 	app-deployment-314667899-xr47k       1/1       Running   0          22h
# sentry get pods -o json > /tmp/pods.json
+ kubectl -n sentry get pods -o json > /tmp/pods.json
```


Variables
---------

Prompts starting with `get` return their output prefixed with `$n`. You may use those variables anywhere to
automatically substitute for the value on first column of the respective line. For example:
```console
# sentry get pods
+ kubectl -n sentry get pods
   	NAME                                 READY     STATUS    RESTARTS   AGE
$1 	app-deployment-314667899-4r9c1       1/1       Running   0          22h
$2 	app-deployment-314667899-xr47k       1/1       Running   0          22h
# sentry logs $2
+ kubectl -n sentry logs app-deployment-314667899-xr47k
```
The `$2` was substituted for `app-deployment-314667899-xr47k`.
