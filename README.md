# commandwrapper

## Running multiple binaries in a scratch docker container
If you only wanted to run a single executable, your life would be simple. CMD does not require a shell, so you could simply add:

```
FROM scratch
...
...
CMD ["/file1"]
```

While running a single command is easy, the fact that you want to run two different commands in sequence means that you need something to handle that workflow. The easiest solution is a shell, because then you can simply run:

```
FROM scratch
...
...
CMD /file1;/file2
```

As we do not want to include a shell in our scratch containers we're using this commandwrapper to call other binaries:

```
/commandwrapper  \
-execute="/usr/local/bin/envscan -server-address=device-api.gridx.de:80....." \
-execute="/usr/local/bin/monitoring -server-address=device-api.gridx.de:80....."
```

Commands can also be chained more precisely using the ```stop-on-failure=true/false``` flag.

## Signal handling
The commandwrapper is forwarding any signal to the running process. in the case of ```SIGTERM``` and ```SIGINT``` the complete process chain will be terminated, no matter if ```stop-on-failure``` is set. To just terminate the currently running process and go on with the chain if ```stop-on-failure=false``` is set, use the ```SIGHUB``` signal.