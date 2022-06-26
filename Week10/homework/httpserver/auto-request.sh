#!/bin/bash
for ((i=1;i<=200;i++));
do
    curl 127.0.0.1:30001/hello
    sleep 2
done
