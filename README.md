# web-2-biblioteca

# START DB POSTGRES IN BACKGROUND
    docker-compose up -d

# STOP DB POSTGRES
    docker-compose down

# START API
    go mod tidy   
    go run main.go 

# BUILD
    go build
