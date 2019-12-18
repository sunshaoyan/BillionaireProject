# -*- coding: UTF-8 -*-

import os


def get_current_dir():
    current_dir = os.path.dirname(os.path.realpath(__file__))
    return current_dir


def get_project_dir():
    project_dir = os.path.abspath(os.path.dirname(os.path.dirname(__file__)))
    return project_dir


def get_conf_dir():
    project_dir = get_project_dir()
    conf_dir = os.path.join(project_dir, "conf")
    return conf_dir


def get_logs_dir():
    project_dir = get_project_dir()
    logs_dir = os.path.join(project_dir, "logs")
    return logs_dir



if __name__ == "__main__":
    print(get_project_dir())
    print(get_conf_dir())
    print(get_logs_dir())
