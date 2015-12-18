#!/usr/bin/python2
# --coding:UTF8

import  os
import sys

makerandompath = "./makerandommarkov.py"

try:
    nstatefrom = int(sys.argv[1])
    nstateto = int(sys.argv[2])
    step = int(sys.argv[3])
    nlabels = int(sys.argv[4])
    bf = int(sys.argv[5])
    saveto = sys.argv[6]
except IndexError:
    print("Usage ./makeexperiments.y <nstatesfrom> <nstatesto> <step> <nlabels> <bf> <safeto>")
    exit(0)
os.mkdir(saveto)


def generate_state():
    for i in range(nstatefrom, nstateto, step):
        nstates = i
        filename = "{}/{}labels_{}states_{}bf.lmc".format(saveto, nlabels, nstates, bf)
        os.system("{} {} {} {} {}".format(
            makerandompath,
            nlabels,
            nstates, 
            bf, 
            filename))


generate_state()
