#!/usr/bin/env bash

set -e

TARGET_OS=darwin
TARGET_ARCH=amd64
TARGET_NAME=tantan-simplify
REPO_PATH=tantan-simplify
BUILD_DIR=${PWD}/_build
CONFIGURATION="${PWD}/etc"
SCRIPTS="${PWD}/scripts"

VERSION_META_FILE=${PWD}/.version.yml
VERSION=`grep ^version ${VERSION_META_FILE} | head -n 1 | cut -d : -f 2 | tr -d " "`
BUILD_DATE=`grep ^build ${VERSION_META_FILE} | head -n 1 | cut -d : -f 2 | tr -d " "`
GIT_REVISION=`grep ^git ${VERSION_META_FILE} | head -n 1 | cut -d : -f 2 | tr -d " "`

print_build_info() {
    echo "Start building..."
    echo
    echo "------------------------------"
    echo "TARGET_OS: ${TARGET_OS}"
    echo "TARGET_ARCH: ${TARGET_ARCH}"
    echo "TARGET_NAME: ${TARGET_NAME}"
    echo "REPO_PATH: ${REPO_PATH}"
    echo "BUILD_DIR: ${BUILD_DIR}"
    echo "VERSION: ${VERSION}"
    echo "BUILD_DATE: ${BUILD_DATE}"
    echo "GIT_REVISION: ${GIT_REVISION}"
    echo "------------------------------"
    echo
}

build() {
    CGO_ENABLE=0 GOOS=${TARGET_OS} GOARCH=${TARGET_ARCH} go build -i -v -ldflags \
        "-X ${REPO_PATH}/version.Version=${VERSION} -X ${REPO_PATH}/version.BuildCommit=${GIT_REVISION} -X ${REPO_PATH}/version.BuildDate=${BUILD_DATE}" \
        -o ${BUILD_DIR}/${TARGET_NAME} ${REPO_PATH}/cmd/server

    if [ $? -ne 0 ]; then
        echo "Failed to build ${TARGET_NAME}"
        exit 1
    fi

    echo
    echo "Building ${TARGET_NAME} successfully."
}

configure_running_env() {
    mkdir -p ${BUILD_DIR}/bin
    mkdir -p ${BUILD_DIR}/conf
    mkdir -p ${BUILD_DIR}/log

    cp -r ${CONFIGURATION}/* ${BUILD_DIR}/conf
    cp -r ${SCRIPTS}/run-in-docker.sh ${BUILD_DIR}/
    cp -r Dockerfile ${BUILD_DIR}/
    mv ${BUILD_DIR}/${TARGET_NAME} ${BUILD_DIR}/bin
}

# main process
print_build_info
build
configure_running_env
