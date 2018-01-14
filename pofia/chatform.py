import logging
import log
import time

from kivy.uix.widget import Widget
from kivy.uix.boxlayout import BoxLayout
from kivy.graphics import Rectangle, Color
from kivy.uix.image import Image

from appcontext import appctx, icopath
from widgetu import Button, Label, LabelSR, TextInput
from chatpage import RVChatPage


class LeftHead(BoxLayout):
    def __init__(self):
        super(LeftHead, self).__init__()
        self.orientation = 'vertical'
        self.size_hint_y = None
        self.height = 90

        layout = BoxLayout(orientation='horizontal', size_hint_y=None, height=55)
        btn1 = Button(text='Hello', size_hint_x=None, width=50)
        btn1 = Image(source=icopath('icon_avatar_40'), allow_stretch=True,
                     size_hint_x=None, width=50)

        btn2 = Button(text='World')
        btn2 = Button(text=appctx.jinfo['name'], bold=True, size_hint_y=None, height=35)
        btn2 = Label(text=appctx.jinfo['name'], bold=True, size_hint_y=None, height=35)
        btn2.font_size += 5
        btn3 = Button(text='!!!!!')
        btn3 = Button(text=appctx.jinfo['stmsg'], size_hint_y=None, height=20)
        btn3 = Label(text=appctx.jinfo['stmsg'], size_hint_y=None, height=20)
        layout.add_widget(btn1)
        lov = BoxLayout(orientation='vertical')
        lov.add_widget(btn2)
        lov.add_widget(btn3)
        layout.add_widget(lov)
        btn4 = Button(text='.......', size_hint_x=None, width=30)
        btn4.text = str(appctx.jinfo['connStatus'])
        btn4.text = {0: "N\nO\nN", 1: 'T', 2: "U\nD\nP"}[appctx.jinfo['connStatus']]
        if appctx.jinfo['connStatus'] > 0:
            btn4 = Image(source=icopath('online_30'), size_hint_x=None, width=14)
        else:
            btn4 = Image(source=icopath('offline_30'), size_hint_x=None, width=14)
        layout.add_widget(btn4)

        losearch = BoxLayout(orientation='horizontal', size_hint_y=None, height=25)
        searchbtn = Button(text='search text')
        searchbtn = TextInput(text='......', size_hint_y=None, height=29, readonly=False)
        # searchbtn.background_color = (0,0,0,0)
        # searchbtn.foreground_color = (0,1,255,0.7)
        losearch.add_widget(searchbtn)
        orderbtn = Button(text='order menu')
        losearch.add_widget(orderbtn)
        clearbtn = Button(text='clear empty', size_hint_x=None, width=26)
        clearbtn.text = 'X'
        losearch.add_widget(clearbtn)

        self.add_widget(layout)
        self.add_widget(Label(size_hint_y=None, height=10))  # separate
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
        # loaction.add_widget(Button(text='ADD'))
        # loaction.add_widget(Button(text='GROP'))
        # loaction.add_widget(Button(text='FILE'))
        # loaction.add_widget(Button(text='PROF'))
        loaction.add_widget(Image(source=icopath('add-square-button_gray64')))
        loaction.add_widget(Image(source=icopath('groupgray')))
        loaction.add_widget(Image(source=icopath('transfer_gray64')))
        loaction.add_widget(Image(source=icopath('settings_gray64')))
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
        self.nmbtn = Label(text='nnnnnnnnnnnnnnnn')
        loinfo.add_widget(self.nmbtn)
        lonumst = BoxLayout(orientation='horizontal')
        self.nplbtn = Label(text='n people chat', size_hint_x=None, width=120)
        lonumst.add_widget(self.nplbtn)
        self.stmbtn = Button(text='sssssssssssssssssssssssssss')
        self.stmbtn = Label(text='sssssssssssssssssssssssssss')
        lonumst.add_widget(self.stmbtn)
        loinfo.add_widget(lonumst)
        lotop.add_widget(loinfo)
        # lotop.add_widget(Button(text='info+', size_hint_x=None, width=50))
        # lotop.add_widget(Button(text='AAA+', size_hint_x=None, width=50))
        # lotop.add_widget(Button(text='VVV+', size_hint_x=None, width=50))
        # question-mark-gray64
        lomicvol = BoxLayout(orientation='vertical', size_hint_x=None, width=30)
        lomicvol.add_widget(Image(source=icopath('phone_mic_gray64'), size_hint_x=None, width=25, size_hint_y=None, height=30))
        lomicvol.add_widget(Image(source=icopath('speaker_volume_gray64'), size_hint_x=None, width=25, size_hint_y=None, height=30))
        # lomicvol.add_widget(Label(text='dddd', size_hint_y=None, height=10))
        import widgetu
        lomicvol.add_widget(widgetu.VSpacer(size_hint_y=None, height=10))
        lotop.add_widget(lomicvol)
        # lotop.add_widget(Image(source=icopath('question-mark-gray64'), size_hint_x=None, width=50))
        lotop.add_widget(Image(source=icopath('phone_call_gray64'), size_hint_x=None, width=50))
        lotop.add_widget(Image(source=icopath('video_recorder_gray64'), size_hint_x=None, width=50))

        locenter = BoxLayout(orientation='vertical')
        locenter.add_widget(Button(text='CCCCCC++++', size_hint_y=None, height=30))
        self.ctpage = RVChatPage()
        appctx.ctpage = self.ctpage
        locenter.add_widget(self.ctpage)

        lobtm = BoxLayout(orientation='horizontal', size_hint_y=None, height=30)
        self.filebtn = Image(source=icopath('paper-clip-outline_gray64'), size_hint_x=None, width=50)
        lobtm.add_widget(self.filebtn)
        self.msgipt = TextInput(text='input...')
        lobtm.add_widget(self.msgipt)
        self.emojibtn = Image(source=icopath('smile_gray64'), size_hint_x=None, width=50)
        lobtm.add_widget(self.emojibtn)
        self.sndmsgbtn = Button(text='send...', size_hint_x=None, width=50)
        self.sndmsgbtn = Image(source=icopath('cursor_gray64'), size_hint_x=None, width=50)
        self.sndmsgbtn.bind(on_press=self.on_sndmsgbtn_pressed)
        lobtm.add_widget(self.sndmsgbtn)

        self.add_widget(lotop)
        self.add_widget(locenter)
        self.add_widget(lobtm)

    def on_sndmsgbtn_pressed(self, obj):
        log.l.debug(obj)
        log.l.debug(self.msgipt.text)
        return

