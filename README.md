## Getting Started
there are two ways to achieve this:
1. `Run locally`

Run the below command
```bash
go run main.go deck.go
```


2. `Run in a Docker environment`
### build the container using the following
```
docker build -t go_app .
```
### Running it 

```
docker run -it --rm  go_app
```

