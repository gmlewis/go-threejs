#!/bin/bash -e
DIRS=`find . -name index.html | xargs dirname`
for i in ${DIRS}
do
  echo "Building ${i}..."
  pushd $i
  gopherjs build
  echo "Opening ${i}..."
  google-chrome index.html
  popd
done
