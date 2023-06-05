#!/bin/bash

# Prompt the user for input
echo "table name: "
read  table
echo "alter or create: "
read method 
migrate create -ext sql -dir migrations/ -seq $(printf "%s_%s_table" "$method" "$table")