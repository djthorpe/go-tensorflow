# go-tensorflow

Some early efforts to get Tensorflow built. Here are the environments I am targetting:

  * Darwin 64-bit on x86, and Debian ARM 32-bit
  * Go version 1.16.X
  * Blaze 3.1.0
  * Tensorflow 2.4.1
  * Protocol Buffers 3.15.8

Compilation on a Raspberry Pi 4 (4GB) was not successful, so may require cross-compilation
on a machine with more memory. In order to compile, you have to start with the C API, then
the Go bindings.

There are a few things not working with the environment I mention above:

  * For Darwin, Tensorflow 2.4.1 is probably best installed using homebrew
    (simple as `brew install tensorflow`) which places it under `/usr/local`. However, this
    installs a dynamic library and for go code distribution, a static library would
    be a better choice so as to create a single distributable binary. Contrarily, there's a lot
    of variations on how to build Tensorflow (with GPU support, etc) so dynamic linking could
    make sense;
  * For go 1.16 with module support, you cannot use `go get` as the protocol buffers are not 
    compiled and so dependencies are not satisfied. This is documented [here](https://github.com/tensorflow/tensorflow/issues/43847) amongst other places.

Not really approached any of the other problems yet.

## Bulding the C API for Tensorflow

Before getting to the go part, you need to install the C API. Assuming you have [Bazel 3.1.0](https://docs.bazel.build/versions/3.1.0/install.html) installed (not sure if any other version will work, but latest one doesn't):

```bash
TF_TEMP=`mktemp -t go-tensorflow`
TF_VERSION=2.4.1
TF_SOURCE="${TF_TEMP}/tensorflow-2.4.1"
TF_PREFIX="/opt"
TF_DEST="${TF_PREFIX}/tensorflow-${TF_VERSION}"

# Download and extract the source code
echo "Downloading 'v${TF_VERSION}.tar.gz'"
wget --directory-prefix "${TF_TEMP}" https://github.com/tensorflow/tensorflow/archive/refs/tags/v${TF_VERSION}.tar.gz

echo "Extracting to '${TF_TEMP}'"
tar -C "${TF_TEMP}" -zxf "${TF_TEMP}/v${TF_VERSION}.tar.gz" && rm "${TF_TEMP}/v${TF_VERSION}.tar.gz"

# Build libraries
cd tensorflow-${TF_VERSION}
bazel build --config=monolithic -c opt //tensorflow/tools/lib_package:libtensorflow

# Install libraries
echo "Installing to '${TF_DEST}'"
install -d "${TF_DEST}"
tar -C "${TF_DEST}" -zxf bazel-bin/tensorflow/tools/lib_package/libtensorflow.tar.gz

# Make pkgconfig
install -d "${TF_DEST}/lib/pkgconfig"
tee "${TF_DEST}/lib/pkgconfig/tensorflow.pc" <<EOF >/dev/null
Name: tensorflow
Description: Tensorflow C Library
Version: ${TF_VERSION}
Cflags: -I${TF_DEST}/include
Libs: -L${TF_DEST}/lib -ltensorflow.${TF_VERSION} -lstdc++
EOF
```

This theoretically creates static libraries as opposed to dynamic ones, which is more
suitable for a golang environment for self-contained binaries. However, it may not work
yet, so some work is still needed.

## Go bindings

As mentioned above, the go module (under `github.com/tensorflow/tensorflow/tensorflow/go`)
do not work with modules switched on, the braindead option is simply to copy it:

```bash
TF_VENDOR="tensorflow"
TF_MODULE="github.com/tensorflow/tensorflow/tensorflow"

echo "Vendoring go bindings to '${TF_VENDOR}'"
install -d "${BASE_DIR}/${TF_VENDOR}/${TF_MODULE}"
cp -r "${TF_SOURCE}/tensorflow/go" "${BASE_DIR}/${TF_VENDOR}/${TF_MODULE}"

# Protobuf compilation
echo "Compiling protobuf to go bindings in '${TF_VENDOR}'"
go get github.com/golang/protobuf/proto
go get github.com/golang/protobuf/protoc-gen-go

for FILE in ${TF_SOURCE}/tensorflow/core/framework/*.proto \
    ${TF_SOURCE}/tensorflow/core/protobuf/*.proto \
    ${TF_SOURCE}/tensorflow/stream_executor/*.proto; do
  protoc -I "${TF_SOURCE}" --go_out="${TF_VENDOR}" "${FILE}"
done

# Adjust paths
echo "Adjusting module paths to '${THIS_MODULE}'"
TF_SUBSTITUTION="s#\(${TF_MODULE}\)#${THIS_MODULE}/${TF_VENDOR}/\1#"
find ${TF_VENDOR} -type f -name "*.go" -exec sed -i '' "${TF_SUBSTITUTION}" {} \;

echo "Running tests"
cd ${BUILD_ROOT} && go test pkg/...
```

There are two addition braindead things to do:

  * There's also a 'wrapper' generation (I think, to define the operations in go from
    the API) that needs to be done, but that part seems to have someone updating it.
    The wrapper generation is in `github.com/tensorflow/tensorflow/tensorflow/go/genops`
  * Adding a `pkgconfig` line (`#cgo pkg-config tensorflow`) in any file which has
    cgo in it. This works so far on darwin since I used the homebrew version rather than
    the version I build in `/opt/tensorflow-v2.4.1`

## Running the bindings

Here's a simple program to verify calling tensorflow library to get version, using the
forked version of tensorflow

```go
package tensorflow

import (
	tf "github.com/djthorpe/go-tensorflow/tensorflow/github.com/tensorflow/tensorflow/tensorflow/go"
)

func main() {
	fmt.Println("version=",tf.Version())
}
```

If that works, you are up and running.
