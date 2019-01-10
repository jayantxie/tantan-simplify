#!/usr/bin/env bash

VERSION_META_FILE=${PWD}/.version.yml
VERSION=$1
BUILD_DATE=`date -u +'%Y-%m-%dT%H:%M:%SZ'`
GIT_REVISION=`git rev-parse --short HEAD`

if [ "$#" -ne 1 ]; then
    echo "Usage:"
    echo "    ./scripts/generate-version.sh v0.0.1"
    exit 1
fi

echo "version: ${VERSION}" > ${VERSION_META_FILE}
echo "build_date: ${BUILD_DATE}" >> ${VERSION_META_FILE}
echo "git_revison: ${GIT_REVISION}" >> ${VERSION_META_FILE}
