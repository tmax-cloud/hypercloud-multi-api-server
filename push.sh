#!/bin/sh
img=tmaxcloudck/hypercloud-multi-api-server:b5.0.0.14
docker rmi $img 
docker build -t $img  . 
docker push $img 
