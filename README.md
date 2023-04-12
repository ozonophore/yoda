# yoda

[![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov]

## Nginx configuration

| **Parameter** | **Value**                      |
|---------------|--------------------------------|
| BACKEND_HOST  | Host webapp(example: frontend) |
| BACKEND_PORT  | Port of webapp (example: 8080) |

## Webapp configuration

| **Parameter**    | **Value**                                                                             |
|------------------|---------------------------------------------------------------------------------------|
| YODA_SERVER_PORT | Port of webapp (example: 8080)                                                        |
| YODA_DB_DSN      | DSN of database (example: postgres://user:password@localhost:5432/db?sslmode=disable) |
| YODA_MQ_URL      | URL of message queue (example: amqp://guest:guest@localhost:5672/)                    |

## App configuration

| **Parameter**     | **Value**                                                                             |
|-------------------|---------------------------------------------------------------------------------------|
| YODA_DB_DSN       | DSN of database (example: postgres://user:password@localhost:5432/db?sslmode=disable) |
| YODA_MQ_CONSUMER  | Queue name for consumer (example: yoda-consumer)                                      |
| YODA_MQ_PUBLISHER | Queue name for publisher (example: yoda-publisher)                                    |

[ci-img]: https://github.com/ozonophore/yoda/actions/workflows/go.yml/badge.svg

[ci]: https://github.com/ozonophore/yoda/actions/workflows/go.yml

[cov-img]: https://codecov.io/gh/ozonophore/yoda/branch/master/graph/badge.svg

[cov]: https://codecov.io/gh/ozonophore/yoda
