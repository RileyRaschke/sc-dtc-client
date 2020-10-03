#!/bin/bash

host='10.2.11.233'
port='11099'
hist_port='11098'

. credentials.env

#echo $SCDTC_USERNAME
#echo $SCDTC_PASSWORD

heartbeat=""

{
#echo "{\"Type\":1,\"Username\":\"$SCDTC_USERNAME\",\"Password\":\"$SCDTC_PASSWORD\",\"Integer_1\": 2,\"HeartbeatIntervalInSeconds\": 6,\"ClientName\":\"test_shell\"}"
#echo -e \\x0\\x0\\x0\\x0
#echo -e \\0\\0\\0\\0
#printf '{\"Type\":1,\"Username\":\"$SCDTC_USERNAME\",\"Password\":\"$SCDTC_PASSWORD\",\"Integer_1\": 2,\"HeartbeatIntervalInSeconds\": 6,\"ClientName\":\"test_shell\"}%s\0'
#printf '{\"Type\":1,\"Username\":\"$SCDTC_USERNAME\",\"Password\":\"$SCDTC_PASSWORD\",\"Integer_1\": 2,\"HeartbeatIntervalInSeconds\": 6,\"ClientName\":\"test_shell\"}%b\xff'
#printf "{\"Type\":1,\"Username\":\"$SCDTC_USERNAME\",\"Password\":\"$SCDTC_PASSWORD\",\"Integer_1\": 2,\"HeartbeatIntervalInSeconds\": 6,\"ClientName\":\"test_shell\"}%b\x00"
printf "{\"Type\":1,\"Username\":\"$SCDTC_USERNAME\",\"Password\":\"$SCDTC_PASSWORD\",\"Integer_1\": true,\"HeartbeatIntervalInSeconds\": 6,\"ClientName\":\"test_shell\"}%b\x00"
#printf '{\"Type\":1,\"Username\":\"$SCDTC_USERNAME\",\"Password\":\"$SCDTC_PASSWORD\",\"Integer_1\": 2,\"HeartbeatIntervalInSeconds\": 6,\"ClientName\":\"test_shell\"}%s\0'
#echo
#printf '%s\0' * \
#  | while IFS= read -r -d '' ; do echo "$REPLY" ; done

#echo -e "{
#  \"Username\":\"$SCDTC_USERNAME\",
#  \"Password\":\"$SCDTC_PASSWORD\",
#  \"Integer_1\": 2,
#  \"HeartbeatIntervalInSeconds\": 6
#}"
#sleep 100
for x in {1..30}
do
  printf "{\"Type\":3,\"CurrentDateTime\":$(date +%s)}%b\x00"
  #echo -e "{\"Type\":3,\"CurrentDateTime\":$(date +%s)}"
  #echo -e \\x0\\x0\\x0\\x0
  #echo -e \\0\\0\\0\\0
#  printf '%s\0'
#  echo
  sleep 5
done

} | telnet ${host} ${port}
#} | xxd
