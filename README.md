# Receipt Processor
## Decription:
Build a webservice that processes Receipts

## Instructions
Run
```
git mod tidy
```

Then run this to start the server on port 8080
```
git mod .
```

To test the POST endpoint:
```
curl -X POST "http://localhost:8080/receipts/process" -H "Content-Type: application/json" -d "<JSON_FILE_path>"
```

To test the GET endpoint:
```
curl -X GET "http://localhost:8080/receipts/{id_returned_from_post}/points"
```
