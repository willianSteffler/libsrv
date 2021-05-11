# libsrv
implementação do server do projeto da biblioteca de sistemas distribuidos em go

## Rodar
### Linux
-  fazer download do repositorio `git clone https://github.com/willianSteffler/libsrv.git` **ou fazer o download do binario** [libsrv-linux64](https://github.com/willianSteffler/libsrv/releases/download/v0.0.4/libsrv-linux64)
-  fazer o download dos [dados](https://github.com/willianSteffler/libsrv/releases/download/v0.0.4/database.db)
 - `./bin/libsrv-linux64 ./basedados/database.db

### Outros sistemas
tem que compilar o codigo
-  fazer download do repositorio `git clone https://github.com/willianSteffler/libsrv.git`
- `go build -o libsrv`
- `./libsrv ./basedados/database.db`
