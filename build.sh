#!/bin/sh
docker rmi hypercloud-multi-api-server:5.0 
docker build -t hypercloud-multi-api-server:5.0  . 
