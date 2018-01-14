import logging
import log
import time

from kivy.lang import Builder
from kivy.uix.widget import Widget
from kivy.uix.boxlayout import BoxLayout
from kivy.graphics import Rectangle, Color
from kivy.uix.image import Image

from kivy.uix.recycleview import RecycleView
from kivy.uix.recycleview.views import RecycleDataViewBehavior

from kivy.clock import Clock

from appcontext import appctx, icopath

from widgetu import Button, Label, LabelSR, TextInput


class MessageItemView(RecycleDataViewBehavior, BoxLayout):
    _latest_data = None
    _rv = None

    def __init__(self):
        super(MessageItemView, self).__init__()
        self.orientation = 'horizontal'
        self.cttxusz = None
        self.pdata = None
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
        self.icoSelfBtn = Image(source=icopath('icon_avatar_40'), allow_stretch=False, size_hint_x=None, width=50, size_hint_y=None, height=50)
        lo0 = BoxLayout(orientation='vertical', size_hint_x=None, width=50)
        lo0.add_widget(self.icoSelfBtn)
        lo0.add_widget(Label(size_hint_x=None, width=1))
        self.add_widget(lo0)
        lo1 = BoxLayout(orientation='horizontal', size_hint_y=None, height=32)
        self.nmbtn = Button(text=self.name, size_hint_y=None, height=30)
        self.nmbtn = Label(text=self.name, size_hint_y=None, height=30)
        lo1.add_widget(self.nmbtn)
        self.tmbtn = Button(text=self.time, size_hint_y=None, height=30)
        self.tmbtn = Label(text=self.time, size_hint_y=None, height=30)
        lo1.add_widget(self.tmbtn)
        self.lcbtn = Label(text='000', size_hint_x=None, width=32)
        lo1.add_widget(self.lcbtn)
        lo = BoxLayout(orientation='vertical')
        lo.add_widget(lo1)
        self.ctbtn = Label(text=self.text)  # , size_hint_y=None)
        self.ctbtn.halign = 'left'
        self.ctbtn.valign = 'top'
        self.ctbtn.bind(size=self.on_set_ctbtn_size)
        self.ctbtn.bind(texture_size=self.on_content_texture_size)
        # self.ctbtn.text_size = (self.ctbtn.width*4, None)
        # TODO 高度还不对，文字太长的话由于折成多行然后会和上下的重叠
        # btn = self.ctbtn
        # btn.bind(width=lambda *x: btn.setter('text_size')(btn, (btn.width, None)),
        #            texture_size=lambda *x: btn.setter('height')(btn, btn.texture_size[1]))
        # log.l.debug('text size:' + str(self.ctbtn.text_size) + str(self.ctbtn.max_lines))
        lo.add_widget(self.ctbtn)
        lo.add_widget(Label(size_hint_y=None, height=3))
        self.add_widget(lo)
        self.icoPeerBtn = Button(text=self.icon2, size_hint_x=None, width=50, size_hint_y=None, height=50)
        self.icoPeerBtn = Image(source=icopath('icon_avatar_40'), allow_stretch=False, size_hint_x=None, width=50, size_hint_y=None, height=50)
        lo9 = BoxLayout(orientation='vertical', size_hint_x=None, width=50)
        lo9.add_widget(self.icoPeerBtn)
        lo9.add_widget(Label(size_hint_x=None, width=50))
        self.add_widget(lo9)
        pass

    # why call 2 time when initialized
    def refresh_view_attrs(self, rv, index, data):
        ''' Catch and handle the view changes '''
        self.pdata = data
        self._rv = rv
        if self._latest_data is not None:
            self._latest_data["height"] = self.height
        self._latest_data = data

        # log.l.debug(str(index)+str(data))
        self.ctbtn.markup = True
        self.ctbtn.text = data['text']
        log.l.debug(str(self.ctbtn.texture_size)+str(data))
        cttxusz = self.ctbtn.texture_size
        if 'height' in data and self.ctbtn.texture_size[1] != data['height'] + 33:
            # data['height'] = self.ctbtn.texture_size[1] + 33
            pass
        self.nmbtn.markup = True
        self.nmbtn.text = data['name']
        self.tmbtn.text = data['time']
        self.lcbtn.text = str(data['height']) if 'height' in data else '000'
        self.lcbtn.text = str(cttxusz[1]+30)
        # self.size_hint_y = None
        if index == 0:
            import traceback
            # traceback.print_stack()
        return super(MessageItemView, self).refresh_view_attrs(rv, index, data)

    def refresh_view_layout(self, rv, index, layout, viewport):
        return super(MessageItemView, self).refresh_view_layout(rv, index, layout, viewport)

    def on_height(self, instance, value):
        self.cttxusz = self.ctbtn.texture_size
        data = self._latest_data
        log.l.debug(str(value) + str(self.cttxusz) + str(self.pdata))
        log.l.debug(str(self.height) + str(self.cttxusz) + str(self.pdata))
        if data is not None and 'height' not in data: return
        # if data is not None and data["height"] != value:
        if data is not None and data["height"] != self.cttxusz[1] + 33:
            #data["height"] = self.cttxusz[1] + 33
            #self._rv.refresh_from_data()
            pass
        return

    def on_set_ctbtn_size(self, obj, sz):
        log.l.debug(str(sz) + obj.text)
        self.ctbtn.text_size = (sz[0], None)
        return

    def on_content_texture_size(self, obj, txusz):
        log.l.debug(str(txusz) + str(self.pdata))
        data = self._latest_data
        # if True: return
        self.lcbtn.text = str(txusz[1] + 33)
        if data is not None and 'height' not in data: return
        if data is not None and data["height"] != txusz[1] + 33:
            data["height"] = txusz[1] + 33
            self.pdata['height'] = txusz[1] + 33
            log.l.debug('schedule_refresh:' + self.lcbtn.text + str(self.pdata))
            Clock.schedule_once(self._rv.refresh_from_data, -1)
            # self._rv.refresh_from_data()
        return


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
        self.viewclass = 'MessageItemView'
        self.key_size = 'height'
        # relo = RecycleBoxLayout()
        # relo.orientation = 'vertical'
        # self.RecycleBoxLayout = relo
        # log.l.debug(self.RecycleBoxLayout)
        self.data = [{'text': self.styleText(x), 'markup': True,
                      'name': '---', 'time': '---',
                      'icol': '', 'icor': ''} for x in range(3)]

    def styleText(self, txt):
        return '[color=#123456]'+str(txt)+'[/color]'

    def additem(self, ename, emsg):
        import time

        linecnt = int(len(emsg)/30) + 1
        height = linecnt * 30 + 30
        # msgo['height'] = int(height)  # 'dp('+str(int(height))+')'
        log.l.debug('item height:'+str(height))

        self.data.append({'text': emsg, 'markup': True, 'height': height,
                          'name': self.styleText(ename), 'time': str(time.asctime())})
        return
