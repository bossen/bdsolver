#!/usr/bin/python2
# --coding:UTF8

import  os
import subprocess
from time import time
import json

bdsolverpath = "./bdsolver -tpsolver cplex"
makerandompath = "./makerandommarkov"
saveto = "files"


def runone(nlabels, nstates, bf):
    filename = "{}/{}labels_{}states_{}bf.lmc".format(saveto, nlabels, nstates, bf)
    os.system("{} {} {} {} {}".format(makerandompath, nlabels, nstates, bf, filename))

    t = time()
    p = subprocess.Popen("{} {}".format(bdsolverpath, filename),
            shell=True, stderr=subprocess.PIPE, stdout=subprocess.PIPE)
    p.wait()
    ret = p.returncode
    t = time() -t


    result = {
        "nlabels": nlabels,
        "nstates": nstates,
        "bf": bf,
        "time": t,
        "return": ret}
    return result


def writeresult():
    global results
    with open("{}/results.json".format(saveto), "w") as f:
        f.write(json.dumps(results, indent=4))

    with open("{}/results.text".format(saveto), "w") as f:
        for result in results:
            f.write("{}\t{}\n".format(result['nstates'], result['time']))

results = []
nerrors = 0

def makeresults(gen):
    global results
    while True:
        try:
            nlabels, nstates, bf  = gen.next()
        except StopIteration:
            return results

        result = runone(nlabels, nstates, bf)

        results.append(result)
        writeresult()
        print(result)


def generator_state(nstatefrom, nstateto, nlabels, bf):
    for i in range(nstatefrom, nstateto, 50):
        yield (nlabels, i, bf)


x = generator_state(50,500, 3, 2)
makeresults(x)

msg = "No errors occured" if nerrors == 0 else "{} of errors".format(nerrors)
print("Done experiments with: " + msg)
