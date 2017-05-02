TOKENIZER
=========

Requirements
------------
You need to have installed `docker`, `docker-compose` and `make` command.

First start
-----------
For first time start of application run `make up`. This command will:
 - Download Docker images
 - Create containers
 - Compile Go-code
 - Create Postgres database
 - Wait for DB will be ready to accept connections
 - Execute migrations
 - Apply seed data
 - Run application on 80 HTTP port
 
 
Next starts
-----------
For next time starts use `make start` command. 
This command just run exist containers.

Using
-----
1. Issue new token by user credentials:
```
curl -X POST -F "login=John" -F "password=12345" "http://localhost/token"
```

You will see issued token like a `df6c38078d191bf41a409331f4523d9e`

2. Get login by token:
```
curl -X POST -F "token=6b9b9319" "http://localhost/login" 

```
You will owner's login like a `John` and 