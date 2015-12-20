#!/bin/bash


/usr/bin/time --format="$1 %U %x" ./bdsolver -tpsolver cplex $1 > /dev/null  2>> $2

