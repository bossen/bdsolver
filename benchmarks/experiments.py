#!/usr/bin/python2
# --coding:UTF8

import  os
import sys
import subprocess
from time import time
import json

bdsolverpath = "./bdsolver"
makerandompath = "./makerandommarkov.py"

try:
    nstatefrom = int(sys.argv[1])
    nstateto = int(sys.argv[2])
    step = int(sys.argv[3])
    nlabels = int(sys.argv[4])
    bf = int(sys.argv[5])
    saveto = sys.argv[6]
except IndexError:
    print("Usage ./experiments <nstatesfrom> <nstatesto> <step> <nlabels> <bf> <safeto>")
    exit(0)
os.mkdir(saveto)


"""
    nstatefrom = 500
    nstateto = 700
    step = 100
    nlabels = 3
    bf = 2
"""

def generatemarkov(args):
    os.system("{} {} {} {} {}".format(
        makerandompath,
        args['nlabels'],
        args['nstates'], 
        args['bf'], 
        args['filename']))


def runone(args):
    t = time()
    p = subprocess.Popen("{} {}".format(bdsolverpath, args['filename']),
            shell=True, stderr=subprocess.PIPE, stdout=subprocess.PIPE)
    p.communicate()
    p.wait()
    ret = p.returncode
    t = time() -t


    args['time'] = t
    args['return'] =  ret


def writeresult():
    global results
    with open("{}/results.json".format(saveto), "w") as f:
        f.write(json.dumps(results, indent=4))

    with open("{}/results.text".format(saveto), "w") as f:
        for result in results:
            f.write("{}\t{}\n".format(result['nstates'], result['time']))

results = []
nerrors = 0

def rungenerator(gen):
    global results
    while True:
        try:
            args  = gen.next()
        except StopIteration:
            return results

        runone(args)

        results.append(args)
        writeresult()
        print(args)


def generator_state():
    for i in range(nstatefrom, nstateto, step):
        args =  {
            "nlabels":  nlabels,
            "nstates": i,
            "bf": bf,
            "filename": "{}/{}labels_{}states_{}bf.lmc".format(saveto, nlabels, i, bf)
        }
        generatemarkov(args)
        yield args


rungenerator(generator_state())

msg = "No errors occured" if nerrors == 0 else "{} of errors".format(nerrors)
print("Done experiments with: " + msg)
