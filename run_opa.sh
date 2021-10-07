if [ -z "$1" ]
then
    echo "running OPA on default port"
    ./opa run --server > /dev/null 2>&1 &
else
    echo "running OPA on port $1"
    ./opa run --server --addr $1 > /dev/null 2>&1 &
fi
# might check port with "nc"
# nc -z 127.0.0.1 $1