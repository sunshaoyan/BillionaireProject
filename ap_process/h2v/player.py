import math


class Skeleton:
    index_map = {
        'left_hand': 9,
        'right_hand': 10,
        'left_elbow': 7,
        'right_elbow': 8,
        'left_shoulder': 5,
        'right_shoulder': 6,
        'left_hip': 11,
        'right_hip': 12
    }
    occ_threshold = 0.3

    def __init__(self, points):
        assert len(points) >= 17, "there should be no less than 17 points to initialize the skeleton"
        self.left_hand = points[9]
        self.right_hand = points[10]
        self.left_elbow = points[7]
        self.right_elbow = points[8]
        self.left_shoulder = points[5]
        self.right_shoulder = points[6]
        self.left_hip = points[11]
        self.right_hip = points[12]
        self.points = points

    def point_less_than(self, first_point, second_point, axis, extent):
        i = Skeleton.index_map[first_point]
        j = Skeleton.index_map[second_point]
        if self.points[i]['score'] < Skeleton.occ_threshold or self.points[j]['score'] < Skeleton.occ_threshold:
            return False
        length = math.sqrt(math.pow(self.points[i]['x'] - self.points[j]['x'], 2) +
                           math.pow(self.points[i]['y'] - self.points[j]['y'], 2))
        return self.points[j][axis] - self.points[i][axis] > length * extent

    def point_nearly_equal_to(self, first_point, second_point, axis, extent):
        assert first_point in Skeleton.index_map and second_point in Skeleton.index_map
        i = Skeleton.index_map[first_point]
        j = Skeleton.index_map[second_point]
        if self.points[i]['score'] < Skeleton.occ_threshold or self.points[j]['score'] < Skeleton.occ_threshold:
            return False
        length = math.sqrt(math.pow(self.points[i]['x'] - self.points[j]['x'], 2) +
                           math.pow(self.points[i]['y'] - self.points[j]['y'], 2))
        return math.fabs(self.points[j][axis] - self.points[i][axis]) < length * extent


class PlayerStatus:
    img_width = 1920

    def __init__(self):
        self.left_hand_status = None
        self.right_hand_status = None
        self.player_in_view_status = None
        self.left_hand_previous_status = None
        self.right_hand_previous_status = None
        self.player_in_view_previous_status = None

    def update_status(self, box, points):
        half_img_width = PlayerStatus.img_width // 2
        if (box is not None) and (box[0] < half_img_width < box[2]) and \
                ((box[2] - box[0]) > PlayerStatus.img_width / 8):
            self.player_in_view_status = "in"
        else:
            self.player_in_view_status = "out"

        if points is None:
            self.left_hand_status = None
            self.right_hand_status = None
            return
        skeleton = Skeleton(points)
        if skeleton.point_nearly_equal_to("left_hand", "left_elbow", "y", 0.5) \
            and skeleton.point_nearly_equal_to("left_hand", "left_shoulder", "y", 0.5) \
            and skeleton.point_less_than("left_shoulder", "left_hand", "x", 0.7):
            self.left_hand_status = "stretch"
        elif skeleton.point_nearly_equal_to("left_hand", "left_elbow", "y", 0.5) \
            and skeleton.point_less_than("left_hand", "left_shoulder", "x", 0.1) \
            and skeleton.point_less_than("left_hand", "left_elbow", "x", 0.7):
            self.left_hand_status = "fold"
        elif skeleton.point_less_than("left_hand", "left_elbow", "y", 0.5) \
            and skeleton.point_less_than("left_hand", "left_shoulder", "y", 0.1):
            self.left_hand_status = "up"
        else:
            self.left_hand_status = None
        if skeleton.point_nearly_equal_to("right_hand", "right_elbow", "y", 0.5) \
            and skeleton.point_nearly_equal_to("right_hand", "right_shoulder", "y", 0.5) \
            and skeleton.point_less_than("right_hand", "right_shoulder", "x", 0.7):
            self.right_hand_status = "stretch"
        elif skeleton.point_nearly_equal_to("right_hand", "right_elbow", "y", 0.5) \
                and skeleton.point_less_than("right_shoulder", "right_hand", "x", 0.1) \
                and skeleton.point_less_than("right_elbow", "right_hand", "x", 0.7):
            self.right_hand_status = "fold"
        elif skeleton.point_less_than("right_hand", "right_elbow", "y", 0.5) \
                and skeleton.point_less_than("right_hand", "right_shoulder", "y", 0.1):
            self.right_hand_status = "up"
        else:
            self.right_hand_status = None

    def get_status(self):
        status = {
            "left_hand": {
                "status": self.left_hand_status,
                "is_new": self.left_hand_status != self.left_hand_previous_status
            },
            "right_hand": {
                "status": self.right_hand_status,
                "is_new": self.right_hand_status != self.right_hand_previous_status
            },
            "player_in_view": {
                "status": self.player_in_view_status,
                "is_new": self.player_in_view_status != self.player_in_view_previous_status
            }
        }
        return status

    def update_previous(self):
        self.left_hand_previous_status = self.left_hand_status
        self.right_hand_previous_status = self.right_hand_status
        self.player_in_view_previous_status = self.player_in_view_status

