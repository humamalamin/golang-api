## About Apps

This apps is API for resolve test case fullstack from universa. I make this API use Go Language and has implemented clean architecture.

### Pre-require
- Go ^v1.17
- postgress
- Vscode

### Install Apps

```bash 
    git@gitlab.com:humamalamin/test-case-majo.git
    rename file local.env.example to .env
    config value file local.env
    make run
```

### Database

- running pgadmin or dbaver or other database management
- create database and import file dump-test_fullstack-202210140930 inside project


### How to deploy cloud
```
    cd cmd 
    go build -o <nama-file-binary>
    if you windows user:
        - env GOOS=linux GOARCH=amd64 go build -o <nama-file-binary>
    upload binary script to server cloud for execute
``` 