
import logging
logger = logging.getLogger()
FORMAT = "[%(filename)s:%(lineno)s - %(funcName)12s() ] %(message)s"
logging.basicConfig(format=FORMAT)
logger.setLevel(logging.DEBUG)

l = logger

def debug(msg): logger.debug(msg)
def info(msg): logger.info(msg)
def error(msg): logger.error(msg)

def logln(message):
    "Automatically log the current function details."
    import inspect, logging
    # Get the previous frame in the stack, otherwise it would
    # be this function!!!
    func = inspect.currentframe().f_back.f_code
    # Dump the message + the name of this function to the log.
    logging.debug("%s: %s in %s:%i" % (
        message, 
        func.co_name, 
        func.co_filename, 
        func.co_firstlineno
    ))
