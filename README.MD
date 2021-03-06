# QUASAR FIRE OPERATION API

To run the project locally (needs docker installed):
1. you should add the next environment variables into your preference ide:
```
APP_DB_USERNAME=postgres
APP_DB_PASSWORD=secret
APP_DB_NAME=quasar_fire_operation
APP_DB_HOST=localhost
APP_DB_PORT=5432
```
2. then you must run `docker-compose -f docker-compose-local.yml up -d` in order to get the database up and running.
3. launch the main.go in your IDE or using the `go run main.go` on the project root folder

To run the project without the IDE or executing any additional command or configuration:
- just run `docker-compose up -d` on the project root folder

In order to run the tests in any environment you should run `docker-compose -f docker-compose.yml -f docker-compose.test.yml up`