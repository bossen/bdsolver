#!/usr/bin/python
# --coding:UTF8
import os
import logging
import tempfile
from subprocess import call

from common import log

from flask import request, Flask

app = Flask(__init__)

def runbd(lmc, l, v, tpsolver):
    tmp = tempfile.NamedTemporaryFile(delete=False)
    tmp.write(lmc)
    tmp.close()

    call(["./bdsolver", "-l", str(l), "-v" if v else "", "-tpsolver", tpsolver, tmp.name])

@app.route("/")
def main():
    return "something"

@app.route("/bdsolver", methods=["POST"])
def runit():
    token = request.form["token"]

    if token != "forkimandradu":
        return "Wrong token!"

    lmc = request.form["lmc"]


if __name__ == "__main__":
    main()
