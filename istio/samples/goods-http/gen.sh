#!/bin/bash

cd ./comments
gofra kube deployment --override; gofra kube service --override; gofra istio virtual-service --override; gofra istio destination-rule --override

cd ../goods
gofra kube deployment --override; gofra kube service --override; gofra istio virtual-service --override; gofra istio destination-rule --override

cd ../stock
gofra kube deployment --override; gofra kube service --override; gofra istio virtual-service --override; gofra istio destination-rule --override

cd ../user
gofra kube deployment --override; gofra kube service --override; gofra istio virtual-service --override; gofra istio destination-rule --override
