#!/bin/bash

xorm reverse postgres "host=127.0.0.1 port=5432 user=postgres password=lzf123 dbname=test sslmode=disable" templates/goxorm sqlstruct