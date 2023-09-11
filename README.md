# Leveled Logger

[![Go Report Card](https://goreportcard.com/badge/github.com/BorisPlus/leveledlogger)](https://goreportcard.com/report/github.com/BorisPlus/leveledlogger)

`Leveled Logger` - журналирование с введением уровней его доскональности без рализации целевого ресурса, куда направляется лог. Данное действие находится под ответственностью конечной реализации интерфейса `io.Writer`.

## Тестирование

```bash
make test
go clean -testcache && go test -race -cover ./
ok github.com/BorisPlus/leveledlogger 0.046s coverage: 88.2% of statements
```
