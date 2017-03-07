#!/usr/bin/env bash

CODE=$(awk '/expect:/{found=0} {if(found) print} /code:/{found=1}' $1)
EXPECT=$(awk '{if(found) print} /expect:/{found=1}' $1)
RESULT=$(echo "$CODE" | ./glipso)
if [ "$EXPECT" == "$RESULT" ];  then
printf "$1 \e[92m✔\e[0m\n"
else
printf "$1 \e[31m✘\e[0m\n"
echo "expected: $EXPECT"
echo "received: $RESULT"
fi