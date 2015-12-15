#!/usr/bin/python
# --coding:UTF8
import os
import logging

debug = os.environ.get("debug", "0") == "1"

log = logging.getLogger()
logging.basicConfig()
if debug:
    log.setLevel(logging.DEBUG)
else:
    log.setLevel(logging.INFO)
