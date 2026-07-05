#!/bin/bash

if [ "linux" != "linux" ]
then
    echo "true"
else
    echo "false"
fi
fname=/lib/systemd/system/cron.service
if [ -f $fname ]
then   
    head -5 $fname
else   
    echo "cron server is not installed"
fi

exit 0