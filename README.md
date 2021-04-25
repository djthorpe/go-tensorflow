# go-tensorflow

Some early efforts to get Tensorflow built. Mostly targetting Darwin on x64 and ARM32.

## Bulding the C API for Tensorflow

Assuming you have [Bazel 3.1.0](https://docs.bazel.build/versions/3.1.0/install.html) installed (not sure if any other version will work, but latest one doesn't):

```
TF_VERSION=2.4.1
TF_PREFIX=/opt
wget https://github.com/tensorflow/tensorflow/archive/refs/tags/v${TF_VERSION}.tar.gz
tar -zxvf v${TF_VERSION}.tar.gz && rm v${TF_VERSION}.tar.gz
cd tensorflow-${TF_VERSION}
bazel build --config=monolithic -c opt //tensorflow/tools/lib_package:libtensorflow
install -d ${TF_PREFIX}
tar -C ${TF_PREFIX} -zf bazel-bin/tensorflow/tools/lib_package/libtensorflow.tar.gz
```

This theoretically creates static libraries as opposed to dynamic ones, which is more
suitable for a golang environment for self-contained binaries.

After this, create a `pkgconfig` file:

```
install -d ${TF_PREFIX}/tensorflow-${TF_VERSION}/lib/pkgconfig
tee ${TF_PREFIX}/tensorflow-${TF_VERSION}/lib/pkgconfig/libtensorflow.pc <<EOF >/dev/null
Name: tensorflow
Description: Tensorflow C Library
Version: ${TF_VERSION}
Cflags: -I${TF_PREFIX}/tensorflow-${TF_VERSION}/include
Libs: -L${TF_PREFIX}/tensorflow-${TF_VERSION}/lib -ltensorflow -lstdc++
EOF
```

## Go bindings

You should then be able to run the following command, which prints out version number:

```
PKG_CONFIG_PATH=${TF_PREFIX}/tensorflow-${TF_VERSION}/lib/pkgconfig go run ./cmd/version
```

You will need to do some additional work to get the go bindings up and running.

