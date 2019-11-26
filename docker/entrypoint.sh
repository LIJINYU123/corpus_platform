#!/bin/sh

while read line
do
  eval echo $line >> dev.env
done < $1

go run server.go dev.env

