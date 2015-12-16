#!/usr/bin/python2
# --coding:UTF8
import os
import sys
import logging
import random

def _generateedgesfornode(branchingfactor):
    a = []
    for i in range(branchingfactor):
        if i+1 == branchingfactor:
            a.append(100-sum(a))
        else:
            a.append(random.randint(0, 100-sum(a)))
    return a


def generatenodes(nlabels, nstates, branchingfactor):
    output = ""
    labels = map(lambda o: "label" + str(o), range(nlabels))
    for i in range(1, nstates+1):
        output += "{} {}\n".format(i, labels[random.randint(0, nlabels-1)])
    return output

def getstatenotin(nstates, edgesto):
    while True:
        tostate = random.randint(1, nstates)
        if tostate not in edgesto:
            return tostate

def generateedges(nstates, branchingfactor):
    output = ""
    states = range(1, nstates+1)
    edges = []
    for fromstate in states:
        edgesto = []
        for edgeprob in _generateedgesfornode(branchingfactor):
            tostate = getstatenotin(nstates, edgesto)
            edgesto.append(tostate)
            output += "{} -> {} {}/100\n".format(fromstate, tostate, edgeprob)
    return output

def main(nlabels, nstates, branchingfactor): 
    output = "States\n"
    output += generatenodes(nlabels, nstates, branchingfactor)
    
    output += "Edges\n"
    output += generateedges(nstates, branchingfactor)
    return output

def printhelpmenu():
    print("usage makerandommarkov <nlabels> <nstates> <branchingfactor> <filename>")
    print("nlabels, nstates, branchingfactor has to be integers.")

if __name__ == "__main__":
    try:
        nlabels = int(sys.argv[1])
        nstates = int(sys.argv[2])
        branchingfactor = int(sys.argv[3])
        filename = sys.argv[4]
    except IndexError:
        print("Error!") 
        printhelpmenu()
        exit(1)
    except ValueError:
        print("Error!") 
        printhelpmenu()
        exit(1)


    output = main(nlabels, nstates, branchingfactor)
    with open(filename, "w") as f:
        f.write(output)
