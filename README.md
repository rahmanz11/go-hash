"# go-hash" 

# run
docker compose up

# postman
import the Go Hash.postman_collection.json in postman

# request url: 
http://localhost:8080/hash

# request body
{
    "data": ["hello", "mr", "smith"], 
    "algorithm": "xxhash"
}

# response body:
{
    "hash": "6523488163744647430"
}