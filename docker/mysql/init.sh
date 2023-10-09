#!/bin/bash
echo $DB_USER
echo $DB_PASS
mysql -u${DB_USER} -p${DB_PASS} -e"CREATE DATABASE ${DB_NAME};"
mysql -u${DB_USER} -p${DB_PASS} -e"CREATE DATABASE ${DB_NAME}_test;"

