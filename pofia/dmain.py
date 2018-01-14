
import logging
import log
import time

from theme import themeh
from appcontext import appctx, ismybg
from kivy.core.window import Window
Window.clearcolor = themeh.bgcolor


from kivy.app import App
#from kivy.uix.button import Button

from kivy.clock import Clock


class ContactItem:
    def __init__(self):
        self.ctid = ''
        self.text = ''
        self.status = 0
        self.stmsg = ''
        self.ctype = ''
        self.isgroup = False
        self.mbcnt = 0  # peer/member count
        return


class MessageItem:
    def __init__(self):
        self.ctid = ''
        self.name = ''
        self.peer = ''
        self.title = ''
        self.text = ''
        self.time = ''
        self.isme = False
        return


def put_contact_message(msgo):
    if msgo['ctid'] not in appctx.ctmsgs:
        log.l.debug(msgo)
        appctx.ctmsgs[msgo['ctid']] = []

    linecnt = int(len(msgo['text'])/30) + 1
    height = linecnt * 30 + 30
    msgo['height'] = int(height)  # 'dp('+str(int(height))+')'
    log.l.debug('item height:'+str(msgo['height']))

    appctx.ctmsgs[msgo['ctid']].append(msgo)
    return


from mainform import PofiaWin


class PofiaApp(App):
    def build(self):
        import fetch
        appctx.jinfo = fetch.thc_get_base_info()
        appctx.ivtick = Clock.schedule_interval(self.becallback, 0.5)
        # return Button(text='Hello World')
        self.pw = PofiaWin()
        # pw.__ainit__()
        # pw.bind(size=self._update_rect, pos=self._update_rect)
        #with pw.canvas.before:
        #    self.rect = Rectangle(size=pw.size, pos=pw.pos)
        # return pw
        return self.pw
        # return LoginWin()
        # return LeftHead()

    def _update_rect(self, instance, value):
        #self.rect.pos = instance.pos
        #self.rect.size = instance.size
        #print(self.rect)
        pass

    def becallback(self, what):
        import fetch
        while True:
            evt = fetch.thc_poll_event()
            # log.l.debug(evt)
            if evt is None: break
            log.l.debug(evt)
            self.updateByEvent(evt)

    def updateByEvent(self, evt):
        ename = evt['name']
        margs = evt['margs']
        args = evt['args']
        if ename == 'FriendMessage':
            msgo = {'ctid': evt['margs'][1], 'name': evt['margs'][0],
                    'text': evt['args'][1], 'time': str(time.asctime())}
            put_contact_message(msgo)
            self.pw.ctform.ctpage.additem(evt['margs'][0]+'.'+evt['margs'][1][0:5],
                                          evt['args'][1])
        elif ename == 'ConferenceMessage':
            msgo = {'ctid': evt['margs'][3], 'name': evt['margs'][0]+'@'+evt['margs'][2],
                    'peer': evt['margs'][1], 'title': evt['margs'][2],
                    'text': evt['args'][3], 'time': str(time.asctime())}
            put_contact_message(msgo)
            self.pw.ctform.ctpage.additem(evt['margs'][0]+'@'+evt['margs'][2],
                                          evt['args'][3])
            pass
        pass


PofiaApp().run()

