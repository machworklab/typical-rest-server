

# typical-rest-server

[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)


Example of typical and scalable RESTful API Server for Go

### Usage

| Usage | Description |
|---|---|
|`typical-rest-server`|Run the application|
|`typical-rest-server route`|Print available API Routes|


### Configuration

| Name | Type | Default | Required |
|---|---|---|:---:|
|APP_ADDRESS|string|:8089|Yes|
|PG_DBNAME|string||Yes|
|PG_HOST|string|localhost||
|PG_PASSWORD|string|pgpass|Yes|
|PG_PORT|int|5432||
|PG_USER|string|postgres|Yes|
|REDIS_DB|int|0||
|REDIS_DIAL_TIMEOUT|Duration|5s|Yes|
|REDIS_HOST|string|localhost|Yes|
|REDIS_IDLE_CHECK_FREQUENCY|Duration|1m|Yes|
|REDIS_IDLE_TIMEOUT|Duration|5m|Yes|
|REDIS_MAX_CONN_AGE|Duration|30m|Yes|
|REDIS_PASSWORD|string|redispass||
|REDIS_POOL_SIZE|int|20|Yes|
|REDIS_PORT|string|6379|Yes|
|REDIS_READ_WRITE_TIMEOUT|Duration|3s|Yes|
|SERVER_DEBUG|bool|false||

----

## Development Guide

### Prerequisite

Install [Go](https://golang.org/doc/install) (It is recommend to install via [Homebrew](https://brew.sh/) `brew install go`)

### Quick Start

```bash
# equivalent with `docker-compose up -d` (if infrastructure not up)
./typicalw docker up 

# drop, create and migrate postgres database (if database not ready)
./typicalw pg reset 

# generate readme (if there is readme update)
./typicalw readme 

# generate mock (if require mock)
./typicalw mock 

# run test 
./typicalw test

# run the application
./typicalw run 

# release the distribution
./typicalw release 
```

### Annotation

TypicalGo support java-like annotation for code generation before build
- [x] `[constructor]` to add constructor in service-locator 
- [x] `[mock]` is target to be mock for `./typicalw mock`
- [ ] `[api(method=<METHOD>, path=<PATH>)]` to generate route function (TODO)
- [ ] `[cacheable]` to generate cached function (TODO)


### Commands
| Command | Description |
|---|---|
|`./typicalw build`|Build the binary|
|`./typicalw clean`|Clean the project from generated file during build time|
|`./typicalw run`|Run the binary|
|`./typicalw test`|Run the testing|
|`./typicalw mock`|Generate mock class|
|`./typicalw release`|Release the distribution|
|`./typicalw docker`|Docker utility|
|`./typicalw docker compose`|Generate docker-compose.yaml|
|`./typicalw docker up`|Spin up docker containers according docker-compose|
|`./typicalw docker down`|Take down all docker containers according docker-compose|
|`./typicalw docker wipe`|Kill all running docker container|
|`./typicalw readme`|Generate README Documentation|
|`./typicalw postgres`|Postgres Database Tool|
|`./typicalw postgres create`|Create New Database|
|`./typicalw postgres drop`|Drop Database|
|`./typicalw postgres migrate`|Migrate Database|
|`./typicalw postgres rollback`|Rollback Database|
|`./typicalw postgres seed`|Data seeding|
|`./typicalw postgres reset`|Reset Database|
|`./typicalw postgres console`|PostgreSQL Interactive|
|`./typicalw redis`|Redis Tool|
|`./typicalw redis console`|Redis Interactive|
|`./typicalw rails`|Rails-like generation|
|`./typicalw rails scaffold`|Generate CRUD API|
|`./typicalw rails repository`|Generate Repository from tablename|
