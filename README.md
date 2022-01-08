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
- Check goroutines leaks.

## Run tests and linter
`make test lint`

## How to run locally:
```
docker compose build
docker compose up

curl -v -X POST "http://localhost:8080/users" -H "Content-Type: application/json" -d '{"email": "test@mail.ru", "password":"pwd"}'
```

## Kubernetes setup
```
kubectl create secret generic postgres-secret --from-literal=user=root --from-literal=pass=toor

kubectl create secret generic postgres-db-url --from-literal=db-url="postgresql://root:toor@postgres:5432/profile?sslmode=disable&application_name=profile"

kubectl create configmap backend-config-v5 --from-env-file=profile.properties

// run postgres
kubectl apply -f postgres.yaml

// prepare local image
eval $(minikube docker-env)
docker build -t profile .

// add ingress-controller (example for minikube)
minikube addons enable ingress
minikube tunnel

// run backend
kubectl apply -f service.yaml
kubectl apply -f backend.yaml
kubectl apply -f ingress.yaml

// check api
curl -v -X POST "http://127.0.0.1:80/users" -d '{"name": "test", "email": "test@mail.ru", "password": "Pass"}' -H "Content-Type: application/json"

// change minikube env back
eval $(minikube docker-env -u)
```
