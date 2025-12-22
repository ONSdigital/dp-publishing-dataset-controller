#!/bin/bash -eux

cwd=$(pwd)

pushd $cwd/dp-publishing-dataset-controller
  make lint
popd