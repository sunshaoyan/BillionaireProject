from src import utils
from src import log
import os
import json


def load_config():
    conf_file = os.path.join(utils.get_conf_dir(), "config.ini")
    if os.path.exists(conf_file):
        with open(conf_file, 'r') as load_f:
            conf = json.load(load_f)
    else:
        err = "%s does not exist." % conf_file
        log.logger.error(err)
        return False, None, err

    # if not check_config(conf):
    #     err = "%s config error." % conf_file
    #     Log.logger.error(err)
    #     return False, None, err
    # flag_res, err = detect_conf_img(utils.get_conf_dir(), conf)
    # if flag_res:
    #     return True, conf, None
    # else:
    #     return False, None, err
