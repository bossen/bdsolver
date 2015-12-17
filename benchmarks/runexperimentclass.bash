#!/bin/bash

export foldertotest=$1

for file in $(find $foldertotest -type f -name "*.lmc" | sort)
do
	echo "Testing $file"
	/usr/bin/time --format="$file %U %x" ./bdsolver $file > /dev/null  2>> $foldertotest/results.text
done
