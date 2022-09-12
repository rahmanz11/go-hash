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

# supported "algorithm" value in request:
sha512
sha256
md5
xxhash
whirlpool
sha3256
sha3384
sha3512
blake2b384
blake2b512

# default "algorithm":
sha512