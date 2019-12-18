#!/usr/bin/python
# -*- coding: UTF-8 -*-
from h2v.wrapper import H2VWrapperThread
import time

if __name__ == '__main__':
    wrapper_thread = H2VWrapperThread()
    wrapper_thread.set_config('[{"type":"while","condition1":"left_hand_stretch","condition2":"right_hand_fold","logic":"none","action":"move_left"}]')
    wrapper_thread.start()
    while True:
        # Do somthing else
        time.sleep(1)
