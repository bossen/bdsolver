#!/bin/bash

export foldertotest=$1
export outputfile=$2

for file in $(find $foldertotest -type f -name "*.lmc" | sort)
do
	echo "Testing $file"
	./runexperiment.bash "$file" "$outputfile"
	# /usr/bin/time --format="$file %U %x" ./bdsolver -tpsolver cplex $file > /dev/null  2>> $foldertotest/results.text
	sleep 7
done
