#!/usr/bin/env bash
set -e

echo "Formatting:"
go fmt ./...

echo "
Tests:"
go test ./...

echo "
Linting:"
golint ./...

echo "
Building:"
go build

echo "
Successfully built glipso."

echo "
Running Acceptance Tests:
"
FILES=acceptance/*
for f in $FILES
do

  # take action on each file. $f store current file name
  CODE=$(awk '/expect:/{found=0} {if(found) print} /code:/{found=1}' $f)
  EXPECT=$(awk '{if(found) print} /expect:/{found=1}' $f)
  RESULT=$(echo "$CODE" | ./glipso)
  if [ "$EXPECT" == "$RESULT" ];  then
    echo "$f Y"
  else
    echo "$f N"
    echo "expected: $EXPECT"
    echo "received: $RESULT"
  fi
done
