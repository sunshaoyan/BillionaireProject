import hobotx2   # hobotx2.so
import bilioncar
from h2v.handler import Handler
import threading
import time

class H2VWrapperThread(threading.Thread):
    def __init__(self):
        self.handler = Handler()
        hobotx2.init_smart()
        bilioncar.init()
        threading.Thread.__init__(self)
        self.cl = threading.Lock()

    def run(self):
        while True:
            err, frame = hobotx2.read_smart_frame()
            if err is not 0:
                print("read smart frame error:", err)
            else:
                self.cl.acquire()
                self.handler.send_frame(frame)
                self.cl.release()

    def set_config(self, config_str):
        self.cl.acquire()
        self.handler.set_action_config(config_str)
        self.cl.release()
