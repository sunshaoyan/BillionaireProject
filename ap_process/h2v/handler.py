from h2v.player import PlayerStatus
from h2v.pack_pb2 import MessagePack
from h2v.x2_pb2 import FrameMessage
from h2v.action import ActionContainer
from h2v.global_info import GLOBAL_VARS
import base64
import math
import json
import threading

def send_kps(points):
    GLOBAL_VARS.cl.acquire()
    GLOBAL_VARS.kps_to_send = points
    GLOBAL_VARS.cl.release()

class Handler:
    MIN_KPS_AREA = 70 * 100
    MIN_BOX_AREA = 70 * 200
    MIN_TIME_INTERVAL = 1000
    TIME_PER_FRAME = 40

    def __init__(self):
        self.player_status = PlayerStatus()
        self.last_frame_id = 0
        self.action_container = ActionContainer()

    def set_action_config(self, config_str):
        self.action_container.set_config(config_str)

    def send_frame(self, frame_data):
        # raw_frame = base64.b64decode(frame_data)
        mp = MessagePack()
        mp.ParseFromString(frame_data)
        fm = FrameMessage()
        fm.ParseFromString(mp.content_)

        current_frame_id = fm.smart_msg_.timestamp_
        if (current_frame_id - self.last_frame_id) * Handler.TIME_PER_FRAME > Handler.MIN_TIME_INTERVAL:
            self.last_frame_id = current_frame_id

            box, points = Handler._get_largest(fm.smart_msg_.targets_)
            tr = threading.Thread(target=send_kps, args=(points,))
            tr.start()
            if box is None or points is None:
                print('nobody')
            self.player_status.update_status(box, points)
            status = self.player_status.get_status()
            print('status: {}'.format(status))
            # if (status["left_hand"]["status"] != None or status["right_hand"]["status"] != None) \
            #        and (status["player_in_view"]["status"] == "in"):
            #    print('status: {}'.format(status))
            self.action_container.process(status)
            self.player_status.update_previous()

    @staticmethod
    def _get_largest(targets):
        largest_area = 0
        points = None
        for target in targets:
            if target.type_ == "kps":
                kps_pts = target.points_[0].points_
                length = math.fabs(kps_pts[5].x_ - kps_pts[6].x_)
                height = math.fabs(kps_pts[11].y_ - kps_pts[5].y_)
                area = length * height
                if area > largest_area:
                    largest_area = area
                    points = []
                    for pt in target.points_[0].points_:
                        points.append({"x": pt.x_, "y": pt.y_, "score": pt.score_})
                    assert len(points) == 17
        if largest_area < Handler.MIN_KPS_AREA:
            points = None
        largest_area = 0
        box = None
        for target in targets:
            if target.type_ == "body":
                x1 = target.boxes_[0].top_left_.x_
                y1 = target.boxes_[0].top_left_.y_
                x2 = target.boxes_[0].bottom_right_.x_
                y2 = target.boxes_[0].bottom_right_.y_
                area = (y2 - y1) * (x2 - x1)
                if area > largest_area:
                    largest_area = area
                    box = [x1, y1, x2, y2]
        if largest_area < Handler.MIN_BOX_AREA:
            box = None
        return box, points

