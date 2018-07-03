#!/bin/bash

kubectl autoscale deployment my-nginx-deployment --min=2 --max=5 --cpu-percent=80
