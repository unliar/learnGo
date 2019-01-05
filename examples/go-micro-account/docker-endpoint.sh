#! /bin/sh


set -e

if [ -n $Env ]; then
    echo $Env

    echo "$Env config replacing...."
    ls -lh
    cp "./config/$Env.env.json" "./config/env.json"
else
    echo "no env"
fi

exec  ./app  -- -env $Env