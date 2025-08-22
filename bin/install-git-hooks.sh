#!/bin/bash

echo " - Install pre-commit hook"

rm ./.git/hooks/pre-commit
cp ./bin/pre-commit ./.git/hooks/pre-commit
chmod +x ./.git/hooks/pre-commit

echo " - Done"

exit 0