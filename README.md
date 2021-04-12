# vaccel-go-runtime
A go package to deploy vaccelrt

Originally started as a go package to integrate vaccelrt with kata-containers

main.go test vaccel pkg


Build kata with vaccel-vsock
(dirty for now)

git clone https://github.com/nubificus/kata-containers/tree/vaccel-vsock
to kata-containers/src/runtime/vendor/github.com
git checkout vaccel-vsock

cd kata-containers/src/runtime/
make
