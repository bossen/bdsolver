#!/usr/bin/python2
# --coding:UTF8
import os
import logging
import tempfile
import subprocess

from common import log

from flask import request, Flask, render_template

app = Flask(__name__)

def runbd(lmc, l, v, tpsolver):
    tmp = tempfile.NamedTemporaryFile(delete=False)
    tmp.write(lmc)
    tmp.close()

    try:
        return subprocess.check_output(["./bdsolver", "-l", str(l), "-v" if v else "", "-tpsolver", tpsolver, tmp.name])
    except:
        return "An error has occurred!"

@app.route("/")
def main():
    return render_template('index.html')

@app.route("/bdsolver", methods=["POST"])
def runit():
    token = request.form["token"]

    if token != "forkimandradu":
        return "Wrong token!"

    lmc = request.form["lmc"]

    tpsolver = "default"
    if request.form["tpsolver"] != "default":
        tpsolver = "cplex"      #tmp secure

    ilambda = request.form["lambda"]
    try:
        l = float(ilambda)
    except:
        l = 1.0

    selected = request.form.getlist('verbose') 
    verbose = bool(selected)

    return runbd(lmc, l, verbose, tpsolver).replace('\n', '<br />')

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
