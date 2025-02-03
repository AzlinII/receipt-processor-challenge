# Setting up the project

- Install golang
- Run the command `go mod download` to download the necessary packages specified 
  in `go.mod`. 

# Running the project locally

Run the command `go run cmd/api/main.go` or if VSCode is used, it can be run 
in debug mode with the included `.vscode/launch.json`

Sample requests
POST request
```bash 
curl -H 'Content-Type: application/json' \
      --data @examples/readme-sample1.json \                                     
      -X POST \
      localhost:8080/api/1.0.0/receipts/process
```

GET request
```bash 
curl -H 'Content-Type: application/json' \
      localhost:8080/api/1.0.0/receipts/{id}/points
```

