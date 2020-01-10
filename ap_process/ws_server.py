#!/usr/bin/python
# -*- coding: UTF-8 -*-
import hobotx2   # hobotx2.so
import base64
import time
import struct
import json
import bilioncar
from h2v.handler import Handler
import threading
import tornado.web
import tornado.websocket
import tornado.httpserver
import tornado.ioloop
from h2v.global_info import GLOBAL_VARS
from src import log
from src import client
from h2v.wrapper import H2VWrapperThread


class WSHandler(tornado.websocket.WebSocketHandler):
    def open(self):
        pass

    def check_origin(self, origin):
        return True

    def on_message(self, msg):
        GLOBAL_VARS.cl.acquire()
        result = {"points":GLOBAL_VARS.kps_to_send}
        self.write_message(json.dumps(result))
        GLOBAL_VARS.cl.release()

    def on_close(self):
        pass

class Application(tornado.web.Application):
    def __init__(self):
        handlers = [
                (r'/', WSHandler)
                ]
        settings = {}
        tornado.web.Application.__init__(self, handlers, **settings)


def serve_ws():
    ws_app = Application()
    server = tornado.httpserver.HTTPServer(ws_app)
    server.listen(8080)
    tornado.ioloop.IOLoop.instance().start()


def smart_process():
    handler = Handler()
    handler.set_action_config('[{"type":"while","condition1":"player_out","condition2":"right_hand_fold","logic":"or","action":"move_left"},{"type":"if","condition1":"left_hand_up","condition2":"right_hand_stretch","logic":"and","action":"move_right"}]')
    while True:
        err, frame = hobotx2.read_smart_frame()
        if err is not 0:
            print("read smart frame error:", err)
        else:
            handler.send_frame(frame)
    hobotx2.deinit_smart()


def http_thread():
    config = {
        "host": "192.168.1.221:8080"
    }
    httpClient = client.HttpClient(config)

    wrapper_thread = H2VWrapperThread()
    # wrapper_thread.set_config('[{"type":"while","condition1":"left_hand_stretch","condition2":"right_hand_fold","logic":"none","action":"move_left"}]')
    # wrapper_thread.set_config('[{"type":"while","condition1":"player_out","condition2":"right_hand_fold","logic":"none","action":"turn_left"},{"type":"if","condition1":"left_hand_up","condition2":"right_hand_stretch","logic":"and","action":"move_right"}]')
    wrapper_thread.start()
    while True:
        # Do somthing else
        resp = httpClient.get("config")
        if resp.connect_failed():
            log.show_logger.info("connection lost")
        elif resp.ok() != True:
            log.show_logger.info("invalid response")
        elif len(resp.data) != 0 :
            wrapper_thread.set_config(resp.data)
            log.show_logger.info("set data {}".format(resp.data))
        else:
            log.show_logger.info("keep originated config")

        time.sleep(1)


if __name__ == '__main__':
    hobotx2.init_smart()
    log.init()
    bilioncar.init()
    tr = threading.Thread(target=http_thread)
    tr.start()
    print('serving ws')
    serve_ws()
