#!/usr/bin/python
# --coding:UTF8
import os
import logging

from common import log

from flask import request, Flask

app = Flask(__name__)


@app.route("/")
def main():
    return "something"

@app.route("/bdsolver", methods=["POST"])
def runit():
    token = request.form["token"]

    if token != "forkimandradu":
        return "Wrong token!"

    lmc = request.form["lmc"]
    
    return "bdsolver"

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
