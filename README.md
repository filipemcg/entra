## Entra

Simple utility to get access tokens from Azure EntraId.  

### How to use

The program assumes that you have a directory `~/.entra` with a file `credentials` 
Credentials should have the following structure:  
```
[myapp]
tenantId=<azuread-tenant>
clientId=<app-registration-clientid>
clientSecret=<app-registration-secretkey>
scope=<app-registration-scope>
```

It's recommended to build the program and add it to your user `bin` directory (e.g.: `~/.local/bin/`).  

Build with `go build -o ~/.local/bin/entra ./src`

You can then get the access token using the command:  
```sh
entra --app myapp
```
