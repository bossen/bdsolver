from os import getpid
from logging.handlers import SysLogHandler
import logging


def register(appname):
    logger = logging.getLogger()
    handler =  SysLogHandler('/dev/log')
    formatter = logging.Formatter(appname + ' %(message)s', datefmt='%Y-%m-%dT%H:%M:%S')
    handler.setFormatter(formatter)
    logger.addHandler(handler)

    logging.warn("Logging started from app: {} with pid: {}".format(appname, getpid()))
