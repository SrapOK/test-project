# Руководство по деплою

Конфигурация приложения настраивается в config/  
Указать собственный файл с конфигурацией можно при помощи флага --config  
Значение по умолчанию для этого флага - config/dev.yaml  
В .env задать пользователя, пароль и бд.  

```
docker compose up
make run
```

Или

```
docker compose up
go run cmd/test-project/main.go --config=config/dev.yaml
```
