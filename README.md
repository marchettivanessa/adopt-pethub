# adopt-pethub
Repository dedicated for future ETEC's TCC

## Configurações:
Para configuração local, é necessário criar um novo .env. Como esse arquivo não é trackeado pelo git, é necessário cria-lo manualmente, rodando:
`cp .env.example .env`

For local configuration, it's also needed to create a new .env file. The file is not tracked by git, so you have to create it manually by running:
`cp .env.example .env`

## Go version
O projeto usa a versão v1.21 de golang. Por favor, tenha certeza de que essa é a versão instalada.
We use golang v1.21, so please, make sure this is the version you download and install.

## Docker
You can run docker-compose up -d in order to start this application.

But you can also run `go run backend/main/main.go` locally and test if DB connection and http server are working as expected.

## Debugger:
VSCode launch.json config example to debug:

```
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Web App",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "main.go",
            "args": ["api"]
        }
    ]
}
```
