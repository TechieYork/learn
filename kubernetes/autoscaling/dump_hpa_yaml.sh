#!/bin/bash

kubectl get hpa.v1.autoscaling -o yaml > ./hpa.yaml
