#!/bin/bash

# setup OPA directory and check OPA executable
if [ -z "$2" ]
then
    opa_dir=/opt/fortics/fcs-worker
else
    opa_dir=$2
fi

opa_executable=$opa_dir/opa
if [ -e opa_executable ]
then
    echo "cannot locate opa executable from $opa_executable"
    exit 1
fi

# run OPA
if [ -z "$1" ]
then
    echo "running OPA on default port"
    $opa_executable run --server > /dev/null 2>&1 &
else
    echo "running OPA on port $1"
    $opa_executable --server --addr $1 > /dev/null 2>&1 &
fi

# might check OPA port with "nc" if available
# nc -z 127.0.0.1 $1
