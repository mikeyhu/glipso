#!/usr/bin/env bash
set -e

echo "Formatting:"
go fmt ./...

echo "
Tests:"
go test ./...

echo "
Linting:"
golint ./... || echo "Linter not found!"

echo "
Building:"
go build

echo "
Running Acceptance Tests:
"
ERR=0
FILES=acceptance/*.glipso
for f in $FILES
do
  CODE=$(awk '/expect:/{found=0} {if(found) print} /code:/{found=1}' $f)
  EXPECT=$(awk '{if(found) print} /expect:/{found=1}' $f)
  RESULT=$(echo "$CODE" | ./glipso)
  if [ "$EXPECT" == "$RESULT" ];  then
    printf "$f \e[92m✔\e[0m\n"
  else
    printf "$f \e[31m✘\e[0m\n"
    echo "expected: $EXPECT"
    echo "received: $RESULT"
    ERR=$((ERR + 1))
  fi
done
if [ $ERR == 0 ]; then
    echo "Successful Build"
else
    echo "Failed Build"
    exit $ERR
fi
