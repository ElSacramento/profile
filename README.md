# Profile
Simple CRUD service for users with the ability to notify other subscribed services.

## Assumptions:
- User email and password are required fields.
- Only `up` migrations are supported.

## ToDos:
- Proper password storing: hash function.
- Prometheus metrics.
- TraceID/RequestID in logs and other logging improvements, that can be useful. 
- Postgres tests.
- REST API generated specification. 
- Write country service or find the better one.
- GithubActions to run linter, tests, build a docker image.
- Limit param for /users.

## Run tests and linter
`make test lint`

## How to run locally:
```
docker compose build
docker compose up

curl -v -X POST "http://localhost:8080/users" -H "Content-Type: application/json" -d '{"email": "test@mail.ru", "password":"pwd"}'
```
