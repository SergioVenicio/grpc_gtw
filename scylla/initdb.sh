#!/bin/sh

echo "######### Waiting database starts... #########"
sleep 15
echo "######### Starting to execute SH script... #########"
cqlsh database -f ./scylla_scripts/scylladb.txt ;