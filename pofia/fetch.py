from ctypes import *
from toxcore_enums_and_consts import *
from toxav import ToxAV
from libtox import LibToxCore
from tox import bin_to_string
import json

import log

libthc = LibToxCore()


def thc_get_base_info():
    length = c_int()
    binfo_ = create_string_buffer(3456)
    libthc.thc_get_base_info(binfo_)
    binfo = str(binfo_.value, 'utf-8')
    # print(length, length.value)
    # print(c_char_p(binfo_))
    # binfo = bin_to_string(binfo_, length.value)
    # binfo = str(c_char_p(binfo_).value, 'utf-8')
    log.l.info(binfo)
    log.l.info(len(binfo))
    jinfo = json.JSONDecoder().decode(binfo)
    return jinfo


def thc_poll_event():
    evt_ = create_string_buffer(3456)
    ok = libthc.thc_poll_event(evt_)
    # log.l.debug(ok)
    # log.l.debug(evt_)
    if ok == 1:
        evt = str(evt_.value, 'utf-8')
        return json.JSONDecoder().decode(evt)
    return None

