#!/bin/bash
# wait-for-it.sh: A script to wait for TCP services to become available

set -e

host="$1"
port="$2"
shift 2
cmd="$@"

until nc -z "$host" "$port"; do
  >&2 echo "Waiting for $host:$port to be available..."
  sleep 1
done

>&2 echo "$host:$port is available. Starting the command: $cmd"
exec $cmd