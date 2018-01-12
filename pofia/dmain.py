
import logging
import log


from kivy.core.window import Window
Window.clearcolor = (1, 1, 1, 1)

from kivy.app import App
#from kivy.uix.button import Button

from kivy.uix.widget import Widget
from kivy.uix.boxlayout import BoxLayout
from kivy.graphics import Rectangle, Color

from kivy.lang import Builder
from kivy.uix.recycleview import RecycleView

from kivy.utils import get_color_from_hex, get_hex_from_color, get_random_color

from kivy.clock import Clock
from kivy.uix.textinput import TextInput
import kivy
import kivy.uix
import kivy.uix.button as cbtn
class Button(cbtn.Button):
    def __init__(self, **kwargs):
        super(Button, self).__init__(**kwargs)
        self.background_color = (.5, .5, .5, .5)
        self.background_color = get_color_from_hex('#f8f8f880')
        self.background_color = get_color_from_hex('#ffffff00')
        log.l.debug(self.background_color)
        self.color=(0, 0, 0, 1)
        self.color = get_random_color()
        self.color = get_color_from_hex('#00000088')
        self.border = (16, 16, 16, 16)


from libthc import LibThc
thc = LibThc()


class AppContext():
    def __init__(self):
        self.jinfo = None  # json object
        self.ivtick = None  # Clock.Event

appctx = AppContext()


Builder.load_string('''
<RVPeer>:
    viewclass: 'Label'
    RecycleBoxLayout:
        default_size: None, dp(56)
        default_size_hint: 1, None
        size_hint_y: None
        height: self.minimum_height
        orientation: 'vertical'

<RVContact>:
    viewclass: 'Label'
    RecycleBoxLayout:
        default_size: None, dp(56)
        default_size_hint: 1, None
        size_hint_y: None
        height: self.minimum_height
        orientation: 'vertical'
''')


class RVPeer(RecycleView):
    def __init__(self, **kwargs):
        super(RVPeer, self).__init__(**kwargs)
        self.data = [{'text': self.styleText(x), 'markup': True} for x in range(3)]
        # print(self.data)

    def styleText(self, txt):
        return '[color=#123456]'+str(txt)+str(txt)+'[/color]'

    def additem(self, txt):
        self.data.append({'text': self.styleText(txt), 'markup': True})


class RVContact(RecycleView):
    def __init__(self, **kwargs):
        super(RVContact, self).__init__(**kwargs)
        self.data = [{'text': '[color=#123456]'+str(x)+str(x)+'[/color]', 'markup': True} for x in range(3)]
        # print(self.data)
        for fnum in appctx.jinfo['friends'].keys():
            frnd = appctx.jinfo['friends'][fnum]
            if 'name' in frnd: self.additem(frnd['name'])
            else: self.additem('POFU.'+frnd['pubkey'][0:5])

    def styleText(self, txt):
        return '[color=#123456]'+str(txt)+'[/color]'

    def additem(self, txt):
        self.data.append({'text': self.styleText(txt), 'markup': True})


class LeftHead(BoxLayout):
    def __init__(self):
        super(LeftHead, self).__init__()
        self.orientation = 'vertical'
        self.size_hint_y = None
        self.height = 100

        layout = BoxLayout(orientation='horizontal')
        btn1 = Button(text='Hello', size_hint_x=None, width=50)
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
        layout.add_widget(btn4)

        losearch = BoxLayout(orientation='horizontal')
        searchbtn = Button(text='search text')
        searchbtn = TextInput(text='......')
        losearch.add_widget(searchbtn)
        orderbtn = Button(text='order menu')
        losearch.add_widget(orderbtn)
        clearbtn = Button(text='clear empty', size_hint_x=None, width=30)
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
        loaction.add_widget(Button(text='ADD+'))
        loaction.add_widget(Button(text='GROUP+'))
        loaction.add_widget(Button(text='FILE+'))
        loaction.add_widget(Button(text='PROFILE+'))
        pass


class ChatForm(BoxLayout):
    def __init__(self):
        super(ChatForm, self).__init__()
        self.orientation = 'vertical'
        self.buildit()

    def buildit(self):
        lotop = BoxLayout(orientation='horizontal', size_hint_y=None, height=80)
        lotop.add_widget(Button(text='ICOO+', size_hint_x=None, width=50))
        loinfo = BoxLayout(orientation='vertical')
        loinfo.add_widget(Button(text='nnnnnnnnnnnnnnnn'))
        lonumst = BoxLayout(orientation='horizontal')
        lonumst.add_widget(Button(text='bbbbbbbbbbb'))
        lonumst.add_widget(Button(text='sssssssssssssssssssssssssss'))
        loinfo.add_widget(lonumst)
        lotop.add_widget(loinfo)
        lotop.add_widget(Button(text='info++', size_hint_x=None, width=50))
        lotop.add_widget(Button(text='AAA++', size_hint_x=None, width=50))
        lotop.add_widget(Button(text='VVV++', size_hint_x=None, width=50))

        locenter = BoxLayout(orientation='horizontal')
        locenter.add_widget(Button(text='CCCCCC++++'))

        lobtm = BoxLayout(orientation='horizontal', size_hint_y=None, height=80)
        lobtm.add_widget(Button(text='input...'))
        lobtm.add_widget(Button(text='send...'))

        self.add_widget(lotop)
        self.add_widget(locenter)
        self.add_widget(lobtm)


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
        loleft = BoxLayout(orientation='vertical', size_hint_x=0.23)
        loleft.add_widget(LeftHead())

        locontent = BoxLayout(orientation='vertical')
        locontent.add_widget(Button(text='ccccccc', size_hint_max_y=30))
        locontent.add_widget(RVContact())
        loleft.add_widget(locontent)

        loaction = LeftFoot()
        loleft.add_widget(loaction)

        locenter = BoxLayout(orientation='vertical')
        locenter.add_widget(Button(text='cccchhhhha'))
        locenter = ChatForm()

        loright = BoxLayout(orientation='vertical', size_hint_x=0.23)
        loright.add_widget(Button(text='u1111+', size_hint_max_y=30))
        loright.add_widget(Button(text='u222+', size_hint_max_y=30))
        pvtstbtn = Button(text='u333+', size_hint_max_y=30)
        loright.add_widget(pvtstbtn)
        self.rvp = RVPeer()
        loright.add_widget(self.rvp)

        def on_pvbtn0():
            log.l.debug('aaaaaaaa:')
        def on_pvbtn1(obj):
            log.l.debug('bbbbbbb:', obj)
            self.rvp.additem('newadded中')

        pvtstbtn.on_press = on_pvbtn0
        pvtstbtn.bind(on_press=on_pvbtn1)

        lomain = BoxLayout(orientation='horizontal')
        lomain = self
        lomain.add_widget(loleft)
        lomain.add_widget(locenter)
        lomain.add_widget(loright)
        # return lomain

    def aredraw(self, args):
        self.bg_rect.size = self.size
        self.bg_rect.pos = self.pos



class PofiaApp(App):
    def build(self):
        import fetch
        appctx.jinfo = fetch.thc_get_base_info()
        appctx.ivtick = Clock.schedule_interval(self.becallback, 0.5)
        # return Button(text='Hello World')
        pw = PofiaWin()
        # pw.__ainit__()
        # pw.bind(size=self._update_rect, pos=self._update_rect)
        #with pw.canvas.before:
        #    self.rect = Rectangle(size=pw.size, pos=pw.pos)
        # return pw
        return pw
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
        if ename == 'FriendMessage':
            pass
        pass


PofiaApp().run()

