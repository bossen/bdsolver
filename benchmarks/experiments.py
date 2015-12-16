#!/usr/bin/python2
# --coding:UTF8

import  os
import subprocess
from time import time
import json

bdsolverpath = "./bdsolver -tpsolver cplex"
makerandompath = "./makerandommarkov.py"
saveto = "files"


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


def generator_state(nstatefrom, nstateto, nlabels, bf):
    for i in range(nstatefrom, nstateto, 50):
        args =  {
            "nlabels":  nlabels,
            "nstates": i,
            "bf": bf,
            "filename": "{}/{}labels_{}states_{}bf.lmc".format(saveto, nlabels, i, bf)
        }
        generatemarkov(args)
        yield args


x = generator_state(50,300, 3, 2)
rungenerator(x)

msg = "No errors occured" if nerrors == 0 else "{} of errors".format(nerrors)
print("Done experiments with: " + msg)
