# Daily Sleep Tracker API
Service to power a Frontend Web App to track sleep statistics: Functionality:
1. CRUD User
2. CRUD Sleep
3. User authentication
4. Password encryptions

# Built with
- Go
- Gin
- Postgress SQL
- Gorm

# Deployed with help of
- Docker 
- Railway: https://railway.app/

# Documentation
https://documenter.getpostman.com/view/22977269/2s8ZDU4P9h

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
