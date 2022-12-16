# Dependency Injection in Golang with Wire

Please refer to [main.go](./cmd/user/main.go) file for explanation of the usage.

## Requirements

- Go v1.19
- Wire v5.0.0
  
## How To Run

Install dependency
  
```
make setup
```

Build the application
  
```
make build
```

Run the application
  
```
./bin/user
```

## Generating Dependency Injection Code

Install dependency
  
```
make setup
```

Generate dependency injection code
  
```
make wire
```