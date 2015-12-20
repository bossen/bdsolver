#!/bin/bash


/usr/bin/time --format="$1 %U %x" ./bdsolver $1 > /dev/null  2>> $2

