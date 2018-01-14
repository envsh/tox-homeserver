import logging
import log
import time

from kivy.uix.widget import Widget
from kivy.uix.boxlayout import BoxLayout
from kivy.graphics import Rectangle, Color

from appcontext import appctx
from widgetu import Button, Label, LabelSR, TextInput
from peerform import RVPeer, RVContact
from chatform import LeftHead, LeftFoot, ChatForm


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
        locontent.add_widget(Button(text='n contacts', size_hint_max_y=30))
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
        self.nplbtn = Button(text='n people chat', size_hint_max_y=30)
        loright.add_widget(self.nplbtn)
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

