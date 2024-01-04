### Quick start (local):

1. create `.env` file and copy the contents from `.env.example`
2. change your `.env` content by using your environment
3. start the app, run `go run cmd/api.go`
4. hit url/api/v1/health to health check
5. run this command on your terminal to seed data `go run pkg/database/gorm/seeder/run.go up`

### Quick start (docker):
    docker compose up (wait until server up)
    docker exec mezink /seeder up