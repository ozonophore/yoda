Install migration for MAC

```shell
brew install golang-migrate
```

To init migration
```shell
migrate create -ext sql -dir db/migration -seq init_schema
```