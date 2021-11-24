#! /bin/bash

docker build -t registry.devous.ru/nakiner/guestcovider-frontend .
docker push registry.devous.ru/nakiner/guestcovider-frontend:latest
kubectl rollout restart deployment app-frontend -n guestcovider