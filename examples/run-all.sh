#!/bin/bash -e
# Start Python server in the background...
ps ux | grep SimpleHTTPServer | grep python || python -m SimpleHTTPServer &
# Open up each example in the browser...
FILES=`find . -name index.html`
for i in ${FILES}
do
  DIR=$(dirname $i)
  echo "Building ${DIR}..."
  pushd ${DIR}
  gopherjs build
  echo "Opening ${i}..."
  open http://localhost:8000/${i}
  popd
done
