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

## test 1
![TestAddBook](images/TestAddBook.png)

## test 2
![TestAddBookSQLFail](images/TestAddBookSQLFail.png)

## test 3
![TestAddBookInvalidData](images/TestAddBookInvalidData.png)

## test 4
![TestAddBookDuplicate](images/TestAddBookDuplicate.png)

## test 5
![TestUpdateBook](images/TestUpdateBook.png)

## test 6
![TestUpdateBookFail](images/TestUpdateBookFail.png)

## test 7
![TestDeleteBook](images/TestDeleteBook.png)

## test 8
![TestDeleteBookFail](images/TestDeleteBookFail.png)

## test 9
![TestDBConnection](images/TestDBConnection.png)

## test 10
![TestDBConnectionFail](images/TestDBConnectionFail.png)
