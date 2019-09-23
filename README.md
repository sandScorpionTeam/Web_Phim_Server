# filmserver


## Project structure
  /api : Define API logic   
  /dao : Define MongoDB data access   
  /models : Define database models    
  /config : Define config reader    
  main.go : Init API route    
  
## Install package in Go
  go get -v {package_name}    
  Ex    
  ```
  go get -v "github.com/gorilla/mux"   
  go get -v "gopkg.in/mgo.v2/bson"   
  ```
## Run server
  
  ```
  go build main.go 
  -> run main.exe file   
  ```
  or
  ```
  go run main.go  
  ```
