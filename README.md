# library_api web-2

# START DB and API IN BACKGROUND
    docker-compose up -d

# STOP DB and API
    docker-compose down

# START API
    go mod tidy  
    cd cmd/ 
    go run main.go 

# Build project
    docker build -f Dockerfile -t library_api:1.0 .
    docker run -p 8080:8080 library_api:1.0