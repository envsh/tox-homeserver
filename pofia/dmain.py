
import logging
import log
import time

def ismybg(): return True
from kivy.core.window import Window
if ismybg(): Window.clearcolor = (1, 1, 1, 1)

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

from kivy.app import App
#from kivy.uix.button import Button

from kivy.uix.widget import Widget
from kivy.uix.boxlayout import BoxLayout
from kivy.uix.image import Image
from kivy.graphics import Rectangle, Color

from kivy.lang import Builder
from kivy.uix.recycleview import RecycleView

from kivy.utils import get_color_from_hex, get_hex_from_color, get_random_color


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


from kivy.clock import Clock
# from kivy.uix.textinput import TextInput
import kivy
import kivy.uix
import kivy.uix.button as cbtn
import kivy.uix.label as clab
import kivy.uix.textinput as ctxip
class Button(cbtn.Button):
    def __init__(self, **kwargs):
        super(Button, self).__init__(**kwargs)
        self.font_name = appctx.font_heiti
        mybg = ismybg()
        # self.background_color = (.5, .5, .5, .5)
        # self.background_color = get_color_from_hex('#f8f8f880')
        if mybg: self.background_color = get_color_from_hex('#ffffff00')
        # log.l.debug(self.background_color)
        # self.color=(0, 0, 0, 1)
        # self.color = get_random_color()
        if mybg: self.color = get_color_from_hex('#00000088')
        # self.border = (16, 16, 16, 16)


class Label(cbtn.Label):
    def __init__(self, **kwargs):
        super(Label, self).__init__(**kwargs)
        self.font_name = appctx.font_heiti
        mybg = ismybg()
        if mybg: self.background_color = get_color_from_hex('#ffffff00')
        if mybg: self.color = get_color_from_hex('#00000088')
        self.valign = 'middle'


class TextInput(ctxip.TextInput):
    def __init__(self, **kwargs):
        super(TextInput, self).__init__(**kwargs)
        self.font_name = appctx.font_heiti
        mybg = ismybg()
        if mybg: self.background_color = get_color_from_hex('#ffffff00')
        if mybg: self.color = get_color_from_hex('#00000088')


Builder.load_string('''
<RVPeer>:
    viewclass: 'Label'
    RecycleBoxLayout:
        default_size: None, dp(56)
        default_size_hint: 1, None
        size_hint_y: None
        height: self.minimum_height
        orientation: 'vertical'
        multiselect: False
        touch_multiselect: False


<ContactView>:
    # Draw a background to indicate selection
    canvas.before:
        Color:
            rgba: (.0, 0.9, .1, .3) if self.selected else (1, 1, 1, 1)
        Rectangle:
            pos: self.pos
            size: self.size

<RVContact>:
    viewclass: 'ContactView'
    SelectableRecycleBoxLayout:
        default_size: None, dp(56)
        default_size_hint: 1, None
        size_hint_y: None
        height: self.minimum_height
        orientation: 'vertical'
        multiselect: False
        touch_multiselect: False
''')


class ContactItem:
    def __init__(self):
        self.ctid = ''
        self.text = ''
        self.stmsg = ''
        self.ctype = ''
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
        return


import os
def icopath(p):
    return os.getenv('HOME')+'/oss/src/tox-homeserver/tofia/app/src/main/res/drawable/'+p+'.png'


class RVPeer(RecycleView):
    def __init__(self, **kwargs):
        super(RVPeer, self).__init__(**kwargs)
        self.data = [{'text': self.styleText(x), 'markup': True} for x in range(3)]
        # print(self.data)

        if 'groups' in appctx.jinfo:
            for fnum in appctx.jinfo['groups'].keys():
                peers = appctx.jinfo['groups'][fnum]['members']
                for pnum in peers.keys():
                    peer = peers[pnum]
                    if 'name' in peer: self.additem(peer['name'])
                    else: self.additem('POFGP.'+peer['pubkey'][0:5])

        for x in range(5): self.additem(self.styleText(str(x)+str(x)))

    def styleText(self, txt):
        return '[color=#123456]'+str(txt)+'[/color]'

    def additem(self, txt):
        self.data.append({'text': self.styleText(txt), 'markup': True})



from kivy.properties import BooleanProperty
from kivy.uix.behaviors import FocusBehavior
from kivy.uix.behaviors.compoundselection import CompoundSelectionBehavior
from kivy.uix.recycleview.layout import LayoutSelectionBehavior
from kivy.uix.recycleview.views import RecycleDataViewBehavior
from kivy.uix.recycleboxlayout import RecycleBoxLayout
class SelectableRecycleBoxLayout(FocusBehavior, LayoutSelectionBehavior,
                                 RecycleBoxLayout):
    ''' Adds selection and focus behaviour to the view. '''
class ContactView(RecycleDataViewBehavior, BoxLayout):
    index = None
    selected = BooleanProperty(False)
    selectable = BooleanProperty(True)

    def __init__(self):
        super(ContactView, self).__init__()
        self.orientation = 'horizontal'
        self.buildit()

    def buildit(self):
        self.name = 'PofU123'
        self.stmsg = 'Stm123'
        self.icon = 'iiicon'
        self.msgcnt = "999"
        self.ststxt = "N\nT\nU"

        # widgets
        self.nmbtn = None
        self.stmbtn = None
        self.icobtn = None
        self.mcbtn = None
        self.stsbtn = None

        self.icobtn = Label(text=self.icon, size_hint_x=None, width=30, size_hint_y=None, height=30)
        self.icobtn = Image(source=icopath('icon_avatar_40'), size_hint_x=None, width=40)
        self.add_widget(self.icobtn)
        self.nmbtn = Label(text=self.name, size_hint_x=1)
        self.stmbtn = Label(text=self.stmsg, size_hint_x=1)
        lo = BoxLayout(orientation='vertical')
        lo.add_widget(self.nmbtn)
        lo.add_widget(self.stmbtn)
        self.add_widget(lo)
        self.stsbtn = Label(text=self.ststxt, size_hint_x=None, width=30)
        self.stsbtn = Image(source=icopath('offline_30'), allow_stretch=False, size_hint_x=None, width=16)
        self.stsbtn.text = '99'  # no use
        self.add_widget(self.stsbtn)
        self.mcbtn = Label(text=str(self.msgcnt), size_hint_x=None, width=30)
        self.add_widget(self.mcbtn)
        pass

    # why call 2 time when initialized
    def refresh_view_attrs(self, rv, index, data):
        ''' Catch and handle the view changes '''
        self.index = index
        # log.l.debug(str(index)+str(data))
        self.nmbtn.text = data['text']
        self.nmbtn.markup = data['markup']
        self.stmbtn.text = data.get('stmsg')
        self.stmbtn.markup = True
        if index == 0:
            import traceback
            # traceback.print_stack()
        return super(ContactView, self).refresh_view_attrs(rv, index, data)

    def refresh_view_layout(self, rv, index, layout, viewport):
        return super(ContactView, self).refresh_view_layout(rv, index, layout, viewport)

    def on_touch_down(self, touch):
        ''' Add selection on touch down '''
        if super(ContactView, self).on_touch_down(touch):
            return True
        if self.collide_point(*touch.pos) and self.selectable and not self.selected:
            log.l.debug('selected....'+str(self.index)+str(self.selectable))
            # appctx.ctpage.data = []
            ctdata = appctx.rvct.data[self.index]
            log.l.debug(ctdata)
            if ctdata['ctid'] in appctx.ctmsgs:
                msgs = appctx.ctmsgs[ctdata['ctid']]
                appctx.ctpage.data = msgs
            else:
                appctx.ctpage.data = []

            appctx.ctform.nmbtn.text = ctdata['text']
            appctx.ctform.nmbtn.markup = True
            appctx.ctform.stmbtn.text = ctdata['stmsg']
            appctx.ctform.stmbtn.markup = True
            self.selected = True
            return True
        else:
            self.selected = False
            return False

    def apply_selection(self, rv, index, is_selected):
        ''' Respond to the selection of items in the view. '''
        self.selected = is_selected
        if is_selected:
            log.l.debug("selection changed to {0}".format(rv.data[index]))
        else:
            log.l.debug("selection removed for {0}".format(rv.data[index]))


class RVContact(RecycleView):
    def __init__(self, **kwargs):
        super(RVContact, self).__init__(**kwargs)
        self.data = [{'text': self.styleText(str(x)+str(x)), 'markup': True, 'ctid': '--', 'stmsg':'---'} for x in range(1)]
        # print(self.data)
        for fnum in appctx.jinfo['friends'].keys():
            frnd = appctx.jinfo['friends'][fnum]
            if 'name' in frnd: self.additem(frnd['name'], frnd['pubkey'], frnd.get('stmsg'))
            else: self.additem('POFU.'+frnd['pubkey'][0:5], frnd['pubkey'], frnd.get('stmsg'))
        if 'groups' in appctx.jinfo:
            for fnum in appctx.jinfo['groups'].keys():
                frnd = appctx.jinfo['groups'][fnum]
                if 'title' in frnd: self.additem(frnd['title'], frnd['groupId'], '')
                else: self.additem('POFG.'+frnd['groupId'][0:5], frnd['groupId'], '')

        for x in range(2): self.additem(self.styleText(str(x)+str(x)), '--', '')

    def styleText(self, txt):
        return '[color=#123456]'+str(txt)+'[/color]'

    def additem(self, txt, ctid, stmsg):
        if stmsg is None or stmsg == '': stmsg = '---'
        self.data.append({'text': self.styleText(txt), 'markup': True,
                          'ctid': ctid, 'stmsg': stmsg})



class MessageView(RecycleDataViewBehavior, BoxLayout):
    _latest_data = None
    _rv = None

    def __init__(self):
        super(MessageView, self).__init__()
        self.orientation = 'horizontal'
        self.buildit()

    def buildit(self):
        import time
        self.name = 'PofU123'
        self.time = time.asctime()
        self.icon = 'iiicon'
        self.text = 'tttext大中'
        self.icon2 = 'icon222'

        # widgets
        self.ctbtn = None
        self.nmbtn = None
        self.tmbtn = None
        self.icoPeerBtn = None
        self.icoSelfBtn = None
        self.lcbtn = None

        self.icoSelfBtn = Button(text=self.icon, size_hint_x=None, width=50, size_hint_y=None, height=50)
        self.icoSelfBtn = Image(source=icopath('icon_avatar_40'), allow_stretch=False, size_hint_x=None, width=50)
        self.add_widget(self.icoSelfBtn)
        lo1 = BoxLayout(orientation='horizontal')
        self.nmbtn = Button(text=self.name, size_hint_y=None, height=30)
        lo1.add_widget(self.nmbtn)
        self.tmbtn = Button(text=self.time, size_hint_y=None, height=30)
        lo1.add_widget(self.tmbtn)
        self.lcbtn = Label(text='000', size_hint_x=None, width=32)
        lo1.add_widget(self.lcbtn)
        lo = BoxLayout(orientation='vertical')
        lo.add_widget(lo1)
        self.ctbtn = Label(text=self.text)
        self.ctbtn.halign = 'left'
        self.ctbtn.valign = 'top'
        # self.ctbtn.text_size = (self.ctbtn.width*4, None)
        btn = self.ctbtn
        # TODO 高度还不对，文字太长的话由于折成多行然后会和上下的重叠
        btn.bind(width=lambda *x: btn.setter('text_size')(btn, (btn.width, None)),
                    texture_size=lambda *x: btn.setter('height')(btn, btn.texture_size[1]))
        log.l.debug('text size:' + str(self.ctbtn.text_size) + str(self.ctbtn.max_lines))
        lo.add_widget(self.ctbtn)
        self.add_widget(lo)
        self.icoPeerBtn = Button(text=self.icon2, size_hint_x=None, width=50, size_hint_y=None, height=50)
        self.icoPeerBtn = Image(source=icopath('icon_avatar_40'), allow_stretch=False, size_hint_x=None, width=50)
        self.add_widget(self.icoPeerBtn)
        pass

    # why call 2 time when initialized
    def refresh_view_attrs(self, rv, index, data):
        ''' Catch and handle the view changes '''
        self._rv = rv
        if self._latest_data is not None:
            self._latest_data["height"] = self.height
        self._latest_data = data

        # log.l.debug(str(index)+str(data))
        self.ctbtn.text = data['text']
        self.ctbtn.markup = True
        self.nmbtn.text = data['name']
        self.tmbtn.text = data['time']
        self.lcbtn.text = str(data['height']) if 'height' in data else '000'
        self.size_hint_y = None
        if index == 0:
            import traceback
            # traceback.print_stack()
        return super(MessageView, self).refresh_view_attrs(rv, index, data)

    def refresh_view_layout(self, rv, index, layout, viewport):
        return super(MessageView, self).refresh_view_layout(rv, index, layout, viewport)

    def on_height(self, instance, value):
        data = self._latest_data
        log.l.debug(str(data))
        log.l.debug(str(value))
        log.l.debug(str(self.height))
        if data is not None and 'height' not in data: return
        if data is not None and data["height"] != value:
            data["height"] = value
            self._rv.refresh_from_data()


Builder.load_string('''
<RVChatPage>:
    RecycleBoxLayout:
        default_size: None, dp(56)
        default_size_hint: 1, None
        size_hint_y: None
        height: self.minimum_height
        orientation: 'vertical'
        multiselect: True
        touch_multiselect: True
''')

from kivy.uix.recycleboxlayout import RecycleBoxLayout
class RVChatPage(RecycleView):
    def __init__(self, **kwargs):
        super(RVChatPage, self).__init__(**kwargs)
        self.viewclass = 'MessageView'
        self.key_size = 'height'
        # relo = RecycleBoxLayout()
        # relo.orientation = 'vertical'
        # self.RecycleBoxLayout = relo
        # log.l.debug(self.RecycleBoxLayout)
        self.data = [{'text': self.styleText(x), 'markup': True,
                      'name': '', 'time': '',
                      'icol': '', 'icor': ''} for x in range(3)]

    def styleText(self, txt):
        return '[color=#123456]'+str(txt)+'[/color]'

    def additem(self, ename, emsg):
        import time

        linecnt = int(len(emsg)/30) + 1
        height = linecnt * 30 + 30
        # msgo['height'] = int(height)  # 'dp('+str(int(height))+')'
        log.l.debug('item height:'+str(height))

        self.data.append({'text': self.styleText(emsg), 'markup': True, 'height': height,
                          'name': ename, 'time': str(time.asctime())})


class LeftHead(BoxLayout):
    def __init__(self):
        super(LeftHead, self).__init__()
        self.orientation = 'vertical'
        self.size_hint_y = None
        self.height = 70

        layout = BoxLayout(orientation='horizontal')
        btn1 = Button(text='Hello', size_hint_x=None, width=50)
        btn1 = Image(source=icopath('icon_avatar_40'), allow_stretch=False, size_hint_x=None, width=50)

        btn2 = Button(text='World')
        btn2 = Button(text=appctx.jinfo['name'])
        btn3 = Button(text='!!!!!')
        btn3 = Button(text=appctx.jinfo['stmsg'])
        layout.add_widget(btn1)
        lov = BoxLayout(orientation='vertical')
        lov.add_widget(btn2)
        lov.add_widget(btn3)
        layout.add_widget(lov)
        btn4 = Button(text='.......', size_hint_x=None, width=30)
        btn4.text = str(appctx.jinfo['connStatus'])
        btn4.text = {0: "N\nO\nN", 1: 'T', 2: "U\nD\nP"}[appctx.jinfo['connStatus']]
        layout.add_widget(btn4)

        losearch = BoxLayout(orientation='horizontal')
        searchbtn = Button(text='search text')
        searchbtn = TextInput(text='......', size_hint_y=None, height=29)
        losearch.add_widget(searchbtn)
        orderbtn = Button(text='order menu', size_hint_y=None, height=26)
        losearch.add_widget(orderbtn)
        clearbtn = Button(text='clear empty', size_hint_x=None, size_hint_y=None, width=26, height=26)
        clearbtn.text = 'X'
        losearch.add_widget(clearbtn)

        self.add_widget(layout)
        self.add_widget(losearch)

        pass


class LeftFoot(BoxLayout):
    def __init__(self):
        super(LeftFoot, self).__init__()
        self.orientation = 'horizontal'
        # self.size = (300, 100)
        self.size_hint_y = None
        self.height = 50
        loaction = self
        loaction.add_widget(Button(text='ADD'))
        loaction.add_widget(Button(text='GROP'))
        loaction.add_widget(Button(text='FILE'))
        loaction.add_widget(Button(text='PROF'))
        pass


class ChatForm(BoxLayout):
    def __init__(self):
        super(ChatForm, self).__init__()
        self.orientation = 'vertical'
        self.buildit()

    def buildit(self):
        lotop = BoxLayout(orientation='horizontal', size_hint_y=None, height=80)
        self.icobtn = Button(text='ICOO+', size_hint_x=None, width=50)
        self.icobtn = Image(source=icopath('icon_avatar_40'), allow_stretch=False, size_hint_x=None, width=50)
        lotop.add_widget(self.icobtn)
        loinfo = BoxLayout(orientation='vertical')
        self.nmbtn = Button(text='nnnnnnnnnnnnnnnn')
        loinfo.add_widget(self.nmbtn)
        lonumst = BoxLayout(orientation='horizontal')
        lonumst.add_widget(Button(text='bbbbbbbbbbb'))
        self.stmbtn = Button(text='sssssssssssssssssssssssssss')
        lonumst.add_widget(self.stmbtn)
        loinfo.add_widget(lonumst)
        lotop.add_widget(loinfo)
        lotop.add_widget(Button(text='info++', size_hint_x=None, width=50))
        lotop.add_widget(Button(text='AAA++', size_hint_x=None, width=50))
        lotop.add_widget(Button(text='VVV++', size_hint_x=None, width=50))

        locenter = BoxLayout(orientation='vertical')
        locenter.add_widget(Button(text='CCCCCC++++', size_hint_y=None, height=30))
        self.ctpage = RVChatPage()
        appctx.ctpage = self.ctpage
        locenter.add_widget(self.ctpage)

        lobtm = BoxLayout(orientation='horizontal', size_hint_y=None, height=30)
        self.msgipt = TextInput(text='input...')
        lobtm.add_widget(self.msgipt)
        self.sndmsgbtn = Button(text='send...', size_hint_x=None, width=50)
        self.sndmsgbtn.bind(on_press=self.on_sndmsgbtn_pressed)
        lobtm.add_widget(self.sndmsgbtn)

        self.add_widget(lotop)
        self.add_widget(locenter)
        self.add_widget(lobtm)

    def on_sndmsgbtn_pressed(self, obj):
        log.l.debug(obj)
        log.l.debug(self.msgipt.text)
        return


class PofiaWin(BoxLayout):
    def __init__(self):
        super(PofiaWin, self).__init__()
        # self.__init__2()
        self.orientation = 'horizontal'
        self.build2()
        pass

    def __init__2(self):
        self.bind(pos=self._callback_pos)

        #with self.canvas.before:
        #    Color(0, 0, 0, 0)

        #lo = self.abuild()
        lo = self.build2()
        lo.bind(size=self._update_rect, pos=self._update_rect)
        with lo.canvas.before:
            log.l.debug(self.size)
            log.l.debug(lo.size)
            self.rect = Rectangle(size=self.size, pos=self.pos)

        lo.size = (300, 300)
        lo.pos = (0, 0)
        self.add_widget(lo)

    def _callback_pos(self, instance, value):
        log.l.debug(instance.pos, instance.size)

    def _update_rect(self, instance, value):
        self.rect.pos = self.pos
        self.rect.size = self.size
        log.l.debug('hehehe', self.pos, self.size)

    def __ainit__(self):
        super(Widget, self).__init__()
        with self.canvas:
            self.bg_rect = Rectangle(source="cover.jpg", pos=self.pos, size=self.size)
        self.add_widget(self.abuild())
        with self.canvas:
            Color(0, 1, 0, 1)

    def abuild(self):
        layout = BoxLayout(orientation='vertical')
        btn1 = Button(text='Hello')
        btn2 = Button(text='World')
        btn3 = Button(text='!!!!!')
        layout.add_widget(btn1)
        layout.add_widget(btn2)
        layout.add_widget(btn3)
        return layout

    def build2(self):
        loleft = BoxLayout(orientation='vertical', size_hint_x=0.3)
        loleft.add_widget(LeftHead())

        locontent = BoxLayout(orientation='vertical')
        locontent.add_widget(Button(text='ccccccc', size_hint_max_y=30))
        appctx.rvct = RVContact()
        locontent.add_widget(appctx.rvct)
        loleft.add_widget(locontent)

        loaction = LeftFoot()
        loleft.add_widget(loaction)

        locenter = BoxLayout(orientation='vertical')
        locenter.add_widget(Button(text='cccchhhhha'))
        locenter = ChatForm()
        self.ctform = locenter
        appctx.ctform = locenter

        loright = BoxLayout(orientation='vertical', size_hint_x=0.23)
        tstbtn1 = Button(text='u1111+', size_hint_max_y=30)
        loright.add_widget(tstbtn1)
        tstbtn2 = Button(text='u222+', size_hint_max_y=30)
        loright.add_widget(tstbtn2)
        pvtstbtn = Button(text='u333+', size_hint_max_y=30)
        loright.add_widget(pvtstbtn)
        self.rvp = RVPeer()
        loright.add_widget(self.rvp)

        def on_pvbtn1(obj):  # test dynamic additem
            log.l.debug('bbbbbbb:'+str(obj))
            self.rvp.additem('newadded中')
        pvtstbtn.bind(on_press=on_pvbtn1)

        def on_tstbtn2(obj):  # test dynamic view data reset
            self.rvp.data = []
            pass
        tstbtn2.bind(on_press=on_tstbtn2)

        lomain = BoxLayout(orientation='horizontal')
        lomain = self
        lomain.add_widget(loleft)
        lomain.add_widget(locenter)
        lomain.add_widget(loright)
        # return lomain

    def aredraw(self, args):
        self.bg_rect.size = self.size
        self.bg_rect.pos = self.pos


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

