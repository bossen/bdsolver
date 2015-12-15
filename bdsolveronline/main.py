#!/usr/bin/python
# --coding:UTF8
import os
import logging

from common import log

from flask import request, Flask

app = Flask(__init__)


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
