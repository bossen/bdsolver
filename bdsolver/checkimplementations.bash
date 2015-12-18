#!/bin/bash

export foldertotest=$1
export outputfolder=$2

mkdir $outputfolder
for file in $(find $foldertotest -type f -name "*.lmc") 
do
	echo "Checking $file"
	filename=$(basename "$file")
	./bdsolver1 $file > $outputfolder/"$filename"bdsolver1
	./bdsolver2 $file > $outputfolder/"$filename"bdsolver2
	diff -q $outputfolder/"$filename"bdsolver1 $outputfolder/"$filename"bdsolver2
done
