#!/bin/bash

cqlsh
USE qh_keyspace;
DESC TABLES;
DESC tweets;
SELECT * FROM tweets;

# limpiar tabla
TRUNCATE tweets;
TRUNCATE timeline_by_user;
