# Vuln-go


Intentionally vulnerable go (golang) application to test coverage of SAST tools.

All vulnerabilities are marked with `// vulnerability` in code.

Go web frameworks and libraries has been intentionally skipped. Custom helper functions are created using [http](https://pkg.go.dev/net/http) standard library since some SAST tools might not support a web framework like [gin](https://github.com/gin-gonic/gin).

## Vulnerabilities

* SQL Injection (SQLi)
* Command Injection (RCE)
* LFI
* Hardcoded secret

## Run
> Ensure docker compose is installed.

Run the application with `docker-compose up`
### Vulnerability Testing

[Thunder Client](https://marketplace.visualstudio.com/items?itemName=rangav.vscode-thunder-client) is used to document HTTP requests for test cases as well as vulnerabilities. Folder [thunder-tests](https://github.com/vuln-go/blob/master/thunder-tests) in the repo contains these test cases. This makes it convenient to test various vulnerabilities.

## Development

In development mode [Gow](https://github.com/mitranim/gow) is used to watch for file changes and rebuild the app.

To run in dev mode run:
```bash
docker-compose -f docker-compose-dev.yml up --build
```

Stop and delete volume for DB to recreate DB:
```bash
docker-compose down --remove-orphans --volumes --rmi local
```