import logging
import logging.handlers
import os

from src import utils

logger = None
show_logger = None


def init():
    global logger
    global show_logger
    log_filename = os.path.join(utils.get_logs_dir(), "log.log")
    logging.basicConfig()
    fileshandle = logging.handlers.TimedRotatingFileHandler(
        log_filename, when='D', interval=5, backupCount=10)
    fileshandle.suffix = "%Y-%m-%d.log"
    formatter = logging.Formatter('%(asctime)s - %(levelname)s - %(filename)s - %(funcName)s - %(lineno)d '
                                  '- %(message)s', datefmt='%Y-%m-%d %H:%M:%S')
    fileshandle.setFormatter(formatter)
    logger = logging.getLogger()
    show_logger = logging.getLogger("show_log")
    logger.setLevel(logging.INFO)
    fileshandle.setLevel(logging.INFO)
    logger.addHandler(fileshandle)

    log_filename = os.path.join(utils.get_logs_dir(), "show_log.log")
    fileshandle = logging.handlers.TimedRotatingFileHandler(
        log_filename, when='D', interval=5, backupCount=10)
    fileshandle.suffix = "%Y-%m-%d.log"
    fileshandle.setFormatter(formatter)
    ch = logging.StreamHandler()
    ch.setLevel(logging.DEBUG)
    ch.setFormatter(formatter)
    show_logger.setLevel(logging.INFO)
    show_logger.addHandler(ch)
    show_logger.addHandler(fileshandle)

    for hand in logger.handlers:
        if not isinstance(hand, logging.handlers.TimedRotatingFileHandler):
            logger.removeHandler(hand)


if __name__ == "__main__":
    init()
    logger.info("***************************333")
    show_logger.info("************kkkkkkk")
