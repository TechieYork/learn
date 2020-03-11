#!/bin/bash

cd ./comments-grpc
gofra kube deployment --override; gofra kube service --override; gofra istio virtual-service --override; gofra istio destination-rule --override

cd ../goods-http
gofra kube deployment --override; gofra kube service --override; gofra istio virtual-service --override; gofra istio destination-rule --override

cd ../goods-grpc
gofra kube deployment --override; gofra kube service --override; gofra istio virtual-service --override; gofra istio destination-rule --override

cd ../stock-grpc
gofra kube deployment --override; gofra kube service --override; gofra istio virtual-service --override; gofra istio destination-rule --override

cd ../user-grpc
gofra kube deployment --override; gofra kube service --override; gofra istio virtual-service --override; gofra istio destination-rule --override
