#!/bin/bash
TF_TEMP=`mktemp -t tensorflow`
SCRIPT_DIR=`dirname $0`
BASE_DIR="${SCRIPT_DIR}/.."

###############################################################################

TF_VERSION=2.4.1
TF_SOURCE="${TF_TEMP}/tensorflow-2.4.1"

# Download and extract the source code
echo "Downloading ${TF_SOURCE}"
wget --directory-prefix "${TF_TEMP}" https://github.com/tensorflow/tensorflow/archive/refs/tags/v${TF_VERSION}.tar.gz
tar -C "${TF_TEMP}" -zxf "${TF_TEMP}/v${TF_VERSION}.tar.gz" && rm "${TF_TEMP}/v${TF_VERSION}.tar.gz"

# Set toolchain
TOOLCHAIN_NAME="armv8-rpi3-linux-gnueabihf"
TOOLCHAIN_ROOT="/Volumes/xtools/armv8-rpi3-linux-gnueabihf"
TOOLCHAIN_PATCH="crosscompile.patch"
TOOLCHAIN_GCC="${TOOLCHAIN_ROOT}/bin/arm-linux-gnueabihf-gcc"
GCC_VERSION=`"${TOOLCHAIN_GCC}" --dumpversion`

echo $GCC_VERSION
exit 1

# Patch
cp "${BASE_DIR}/etc/${TOOLCHAIN_PATCH}" "${TF_SOURCE}/${TOOLCHAIN_PATCH}"
sed -i "s#%%CT_NAME%%#${TOOLCHAIN_NAME}#g" "${TF_SOURCE}/${TOOLCHAIN_PATCH}"
sed -i "s#%%CT_ROOT_DIR%%#${TOOLCHAIN_ROOT}#g" "${TF_SOURCE}/${TOOLCHAIN_PATCH}"
sed -i "s#%%CT_GCC_VERSION%%#${GCC_VERSION}#g" "${TF_SOURCE}/${TOOLCHAIN_PATCH}"


# Cross compile
echo "Cross compiling for ${TOOLCHAIN_NAME}"
cd "${TF_SOURCE}"
bazel build -c opt \
  --copt="-march=armv6" --copt="-mfpu=vfp" --copt="-funsafe-math-optimizations" --copt="-ftree-vectorize" --copt="-fomit-frame-pointer" \
  tensorflow:libtensorflow.so --cpu=armeabi --crosstool_top=//tools/arm_compiler:toolchain --verbose_failures


