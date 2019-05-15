#!/bin/bash

curn=366923
while true; do
    curn=$(curl -sb -X GET "http://home.domgoodw.in/api/prime?cur=$curn")
    echo $curn
done
