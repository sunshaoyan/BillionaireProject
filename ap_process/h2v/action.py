import time
import threading
import json
import bilioncar


def send_command(command_name):
    print('cmd: {}'.format(command_name))
    bilioncar.lowspeed()
    if command_name == 'stop':
        bilioncar.stop()
    elif command_name == 'move_left':
        bilioncar.left()
    elif command_name == 'move_right':
        bilioncar.right()
    elif command_name == 'move_forward':
        bilioncar.forward()
    elif command_name == 'move_backward':
        bilioncar.backward()
    elif command_name == 'turn_left':
        bilioncar.anticlock()
    elif command_name == 'turn_right':
        bilioncar.clock()
    else:
        bilioncar.stop()


def send_command_and_stop(command_name, stop_time):
    send_command(command_name)
    time.sleep(stop_time)
    send_command("stop")


class Action:
    def __init__(self, action_type, condition1, condition2, logic, action):
        self.action_type = action_type
        self.condition1 = condition1
        self.condition2 = condition2
        self.logic = logic
        self.action = action
        self.while_doing = False
        self.desp = "{} - {} - {} - {} - {}".format(action_type, condition1, condition2, logic, action)

    """
    example:
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
    """
    def evoke(self, player_status):
        condition1_meet, not_new_meet1 = Action.meet_condition(self.condition1, player_status, self.action_type)
        condition2_meet, not_new_meet2 = Action.meet_condition(self.condition2, player_status, self.action_type)
        final_meet = False
        if self.logic == "and":
            final_meet = (not_new_meet1 and not_new_meet2) and (condition1_meet or condition2_meet)
        elif self.logic == "or":
            final_meet = condition1_meet or condition2_meet
        else:
            final_meet = condition1_meet
        if final_meet:
            print('{} condition meet'.format(self.desp))
            if self.action_type == 'while' and not self.while_doing:
                send_command(self.action)
                self.while_doing = True
            elif self.action_type == 'if':
                tr = threading.Thread(target=send_command_and_stop, args=(self.action, 1))
                tr.start()
        elif self.while_doing:
            print('{} condition unmeet'.format(self.desp))
            send_command("stop")
            self.while_doing = False


    @staticmethod
    def meet_condition(condition_name, status, action_type):
        can_be_not_new_status = True if action_type == "while" else False
        if condition_name.startswith("left"):
            action = condition_name[10:]
            return status["left_hand"]["status"] == action and \
                (can_be_not_new_status or status["left_hand"]["is_new"]), status["left_hand"]["status"] == action
        elif condition_name.startswith("right"):
            action = condition_name[11:]
            return status["right_hand"]["status"] == action and \
                (can_be_not_new_status or status["right_hand"]["is_new"]), status["right_hand"]["status"] == action
        elif condition_name.startswith("player"):
            action = condition_name[7:]
            return status["player_in_view"]["status"] == action and \
                (can_be_not_new_status or status["player_in_view"]["is_new"]), status["player_in_view"]["status"] == action
        else:
            return False, False


class ActionContainer:

    def __init__(self):
        self.actions = []

    """
    sample config:
    [{
        'type': 'if',  # or 'while'
        'condition1': 'left_hand_stretch',  # can be one of: '{left/right}_hand_{stretch/fold/up}' or 'player_{in/out}'
        'condition2': 'right_hand_fold',  # same as condition1, but can also be empty string ''
        'logic': 'or' ,  # can be 'and', 'or', 'none'(representing that only condition1 is in consideration)
        'action': 'move_left',  # can be one of 'move_{left/right/forward/backward}' or 'turn_{left/right}'
    }]
    """
    def set_config(self, config_str):
        self.actions = []
        configs = json.loads(config_str)
        for cfg in configs:
            action = Action(cfg['type'], cfg['condition1'], cfg['condition2'], cfg['logic'], cfg['action'])
            self.actions.append(action)

    def process(self, status):
        for action in self.actions:
            action.evoke(status)

