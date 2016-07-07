#!/bin/bash -e
FILES=`find . -name index.html`
for i in ${FILES}
do
  DIR=$(dirname $i)
  echo "Building ${DIR}..."
  pushd ${DIR}
  gopherjs build
  echo "Opening ${i}..."
  open index.html
  popd
done
