#!/bin/bash

# Prompt the user for input
echo "table name in singilar: "
read  table
echo "create or alter: "
read method 
migrate create -ext sql -dir migrations/ -seq $(printf "%s_%s_table" "$method" "$table")

up_query=(
"CREATE TABLE IF NOT EXISTS ${table}s ("
"   id BIGSERIAL NOT NULL PRIMARY KEY,"    
"   ${table}_name VARCHAR(64) NOT NULL DEFAULT '',"    
"   created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),"    
"   updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),"    
"   deleted_at TIMESTAMP WITHOUT TIME ZONE"
");"
)
if [[ $method == 'alter' ]]; then
    echo "Hey"
    up_query=(
        "ALTER TABLE ${table}s ADD new_column varchar NOT NULL DEFAULT '';"
    )
fi
down_query="DROP TABLE IF EXISTS ${table}s"

if [[ $method == 'alter' ]]; then
    echo "Hey"
    down_query="ALTER TABLE ${table}s DROP COLUMN new_column;"
fi


directory=./migrations
files=($(ls -lt "$directory" | grep "^-" | head -2 | awk '{print $9}'))

for file in "${files[@]}"; do
    echo $file
    if [[ $file == *"up"* ]]; then
        for line in "${up_query[@]}"; do
            echo "$line" >> "./migrations/$file"
        done
    else
        echo "$down_query" >> "./migrations/$file"
    fi
done