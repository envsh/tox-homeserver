import logging
import log
import time

from theme import themeh

import kivy
import kivy.resources
kivy.resources.resource_add_path("/usr/share/fonts/truetype/")
kivy.resources.resource_add_path("/usr/share/fonts/truetype/xmind")
kivy.resources.resource_add_path("/usr/share/fonts/")
kivy.resources.resource_add_path("/usr/share/fonts/TTF")
kivy.resources.resource_add_path("/usr/share/fonts/wenquanyi/wqy-microhei")
kivy.resources.resource_add_path("/usr/share/fonts/adobe-source-code-pro")
font_heiti = kivy.resources.resource_find('SourceCodePro-Regular.otf')  # 无中文
font_heiti = kivy.resources.resource_find('simsun.ttc')  # 可中文
font_heiti = kivy.resources.resource_find('wqy-microhei.ttc')  # 可中文
if font_heiti is None:
    log.l.error('font not found')
    exit()


class AppContext():
    def __init__(self):
        self.jinfo = None  # json object
        self.ivtick = None  # Clock.Event
        self.ctmsgs = {}  # contact id => MessageItem{}
        self.curmsgs = []
        self.grpmems = {}  # contact id => MemberItem{}/ContactItem{}
        self.ctpage = None
        self.ctform = None
        self.rvct = None
        global font_heiti
        self.font_heiti = font_heiti


appctx = AppContext()


def ismybg(): return False

import os
def icopath(p):
    return os.getenv('HOME')+'/oss/src/tox-homeserver/tofia/app/src/main/res/drawable/'+p+'.png'


