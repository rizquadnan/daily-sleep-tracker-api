# Daily Sleep Tracker API
Service to power a Frontend Web App to track sleep statistics

# Built with
- Go
- Gin
- Postgress SQL
- Gorm

# Deployed with help of
- Docker 
- Railway: https://railway.app/

# How to run locally (for developers)
## Prerequiste
1. Install go
2. Setup go developer enviroment
3. Have a postgress running. You can have it running locally or deployed online. Make sure to have the DB url

## Steps to run app
1. Clone this repo
2. Create config.env file in the root of the project, containing this:
  - PORT=
  - DB_URL=
  - API_SECRET=
  - TOKEN_LIFE_SPAN=
2. On terminal execute: `go run cmd/main.go`
