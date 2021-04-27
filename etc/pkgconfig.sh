#!/bin/bash

# These values will need changed
TF_PREFIX=/opt
TF_VERSION=2.4.1

# Create a libtensorflow.pc file suitable for consumption by cgo
install -d ${TF_PREFIX}/tensorflow-${TF_VERSION}/lib/pkgconfig
tee ${TF_PREFIX}/tensorflow-${TF_VERSION}/lib/pkgconfig/libtensorflow.pc <<EOF >/dev/null
Name: tensorflow
Description: Tensorflow C Library
Version: ${TF_VERSION}
Cflags: -I${TF_PREFIX}/tensorflow-${TF_VERSION}/include
Libs: -L${TF_PREFIX}/tensorflow-${TF_VERSION}/lib -ltensorflow.${TF_VERSION} -lstdc++
EOF

# Print out the file
cat ${TF_PREFIX}/tensorflow-${TF_VERSION}/lib/pkgconfig/libtensorflow.pc

