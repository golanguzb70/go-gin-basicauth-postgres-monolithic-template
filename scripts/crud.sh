#!/bin/bash

source ./config/maker.env
echo $TEMPLATE_PATH
echo $GO_MOD_URL
sudo chmod 770 $TEMPLATE_PATH/go-gin-basicauth-postgres-monolithic-template-crud.sh
$TEMPLATE_PATH/go-gin-basicauth-postgres-monolithic-template-crud.sh
make swag_init