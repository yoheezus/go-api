#!/bin/bash

curn=2
while true; do
    curn=$(curl -sb -X GET "http://localhost:8080/api/v1/prime?cur=$curn")
    echo $curn
done
