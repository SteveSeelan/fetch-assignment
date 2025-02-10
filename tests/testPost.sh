#!/bin/bash

# Check if an argument is provided
if [[ $# -eq 0 ]]; then
    echo "Error: Please provide the path to the JSON file as an argument."
    exit 1
fi

# JSON file path (passed as the first argument)
JSON_FILE="$1"

# Check if the JSON file exists
if [[ ! -f "$JSON_FILE" ]]; then
    echo "Error: JSON file '$JSON_FILE' not found."
    exit 1
fi

# Send POST request with JSON data
curl -X POST "http://localhost:8080/receipts/process" \
     -H "Content-Type: application/json" \
     -d @"$JSON_FILE"

# Print new line for better output formatting
echo