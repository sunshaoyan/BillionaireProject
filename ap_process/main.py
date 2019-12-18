
from src import log
from src import client
from h2v.wrapper import H2VWrapperThread
import time


if __name__ == "__main__":
    log.init()
    log.show_logger.info("Main Thread Started")
    config = {
        "host": "localhost:8080"
    }
    httpClient = client.HttpClient(config)

    wrapper_thread = H2VWrapperThread()
    # wrapper_thread.set_config('[{"type":"while","condition1":"left_hand_stretch","condition2":"right_hand_fold","logic":"none","action":"move_left"}]')
    wrapper_thread.start()
    while True:
        # Do somthing else
        resp = httpClient.get("config")
        if resp.connect_failed():
            log.show_logger("connection lost")
        elif resp.ok() != True:
            log.show_logger("invalid response")
        elif len(resp.data) != 0 :
            wrapper_thread.set_config(resp.data)
            log.show_logger("set data {}".format(resp.data))
        else:
            log.show_logger("keep originated config")

        time.sleep(1)




