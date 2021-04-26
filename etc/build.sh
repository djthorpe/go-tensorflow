#!/bin/bash
TF_TEMP=`mktemp -t go-tensorflow`
SCRIPT_DIR=`dirname $0`
BASE_DIR="${SCRIPT_DIR}/.."
THIS_MODULE=`head -1 ${BASE_DIR}/go.mod | cut -d ' ' -f 2`

###############################################################################

TF_VERSION=2.4.1
TF_SOURCE="${TF_TEMP}/tensorflow-2.4.1"
TF_PREFIX="/opt"
TF_DEST="${TF_PREFIX}/tensorflow-${TF_VERSION}"
TF_VENDOR="tensorflow"
TF_MODULE="github.com/tensorflow/tensorflow/tensorflow"

# Download and extract the source code
echo "Downloading 'v${TF_VERSION}.tar.gz'"
wget --directory-prefix "${TF_TEMP}" https://github.com/tensorflow/tensorflow/archive/refs/tags/v${TF_VERSION}.tar.gz

echo "Extracting to '${TF_TEMP}'"
tar -C "${TF_TEMP}" -zxf "${TF_TEMP}/v${TF_VERSION}.tar.gz" && rm "${TF_TEMP}/v${TF_VERSION}.tar.gz"

# Build libraries
#cd tensorflow-${TF_VERSION}
#bazel build --config=monolithic -c opt //tensorflow/tools/lib_package:libtensorflow

# Install libraries
#echo "Installing to '${TF_DEST}'"
#install -d "${TF_DEST}"
#tar -C "${TF_DEST}" -zxf bazel-bin/tensorflow/tools/lib_package/libtensorflow.tar.gz
#TODO: make tensorflow.pc

# Make vendor folder
echo "Vendoring go bindings to '${TF_VENDOR}'"
install -d "${BASE_DIR}/${TF_VENDOR}/${TF_MODULE}"
cp -r "${TF_SOURCE}/tensorflow/go" "${BASE_DIR}/${TF_VENDOR}/${TF_MODULE}"

# Protobuffers
echo "Compiling protobuf to go bindings in '${TF_VENDOR}'"
go get github.com/golang/protobuf/proto
go get github.com/golang/protobuf/protoc-gen-go

for FILE in ${TF_SOURCE}/tensorflow/core/framework/*.proto \
    ${TF_SOURCE}/tensorflow/core/protobuf/*.proto \
    ${TF_SOURCE}/tensorflow/stream_executor/*.proto; do
  protoc -I "${TF_SOURCE}" --go_out="${TF_VENDOR}" "${FILE}"
done

# Remove the temporary folder
echo "Cleaning up '${TF_TEMP}'"
rm -fr "${TF_TEMP}"

# Adjust paths
echo "Adjusting module paths to '${THIS_MODULE}'"
TF_SUBSTITUTION="s#\(${TF_MODULE}\)#${THIS_MODULE}/${TF_VENDOR}/\1#"
find ${TF_VENDOR} -type f -name "*.go" -exec sed -i '' "${TF_SUBSTITUTION}" {} \;

echo "Running tests"
cd ${BUILD_ROOT} && go test pkg/...
