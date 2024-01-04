### Quick start (local):

1. create `.env` file and copy the contents from `.env.example`
2. change your `.env` content by using your environment
3. start the app, run `go run cmd/api.go`
4. hit url/api/v1/health to health check
5. run this command on your terminal to seed data `go run pkg/database/gorm/seeder/run.go up`

### Quick start (docker):
    docker compose up (wait until server up)
    docker exec mezink /seeder up

you can access the app at port 3000 and db at port 3306.

### plus point this test app have:
1. delivering working RESTful API
2. clean and production code ready (docker multistage build with lean image size)
3. error handling with .log folder
4. readme quick start doc with postman doc and its success example
5. avoid over engineering (tried to implement clean architecture or at least singleton architecture)