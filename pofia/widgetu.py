
from kivy.uix.widget import Widget
import kivy.uix.label as kvlabel
from kivy.uix.boxlayout import BoxLayout
from kivy.utils import get_color_from_hex, get_hex_from_color, get_random_color


class VBoxLayout(BoxLayout):
    def __init__(self, **kwargs):
        super(VBoxLayout, self).__init__(**kwargs)
        self.orientation = 'vertical'
        return


class HBoxLayout(BoxLayout):
    def __init__(self, **kwargs):
        super(HBoxLayout, self).__init__(**kwargs)
        self.orientation = 'horizontal'
        return


class VSpacer(kvlabel.Label):
    def __init__(self, **kwargs):
        print(kwargs)
        super(VSpacer, self).__init__(**kwargs)
        self.text = ''
        if 'size_hint_y' in kwargs:
            self.size_hint_y = kwargs['size_hint_y']
        else:
            self.size_hint_y = 1

        self.size_hint_x = None
        self.width = 1
        return


class HSpacer(kvlabel.Label):
    def __init__(self, **kwargs):
        super(VSpacer, self).__init__(**kwargs)
        self.text = ''
        if 'size_hint_x' in kwargs:
            self.size_hint_x = kwargs['size_hint_x']
        else:
            self.size_hint_x = 1

        self.size_hint_y = None
        self.height = 1
        return


from kivy.graphics.svg import Svg


class SvgImage(Widget):
    def __init__(self, filename):
        super(SvgImage, self).__init__()
        with self.canvas:
            svg = Svg(filename)
        self.size = svg.width, svg.height
        return

from theme import themeh
from appcontext import appctx, ismybg

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
        # if mybg: self.background_color = get_color_from_hex('#ffffff00')
        # self.background_color = themeh.bgcolor
        # log.l.debug(self.background_color)
        # self.color=(0, 0, 0, 1)
        # self.color = get_random_color()
        # if mybg: self.color = get_color_from_hex('#00000088')
        self.color = themeh.fgcolor
        self.backgroup_normal = ''
        self.backgroup_down = ''
        self.background_color = themeh.bgcolor
        # self.border = (16, 16, 16, 16)
        return

    def on_press(self):
        self.background_color = get_color_from_hex('#2ac3cecc')
        #self.background_color = themeh.bgcolor
        return

    def on_release(self):
        self.background_color = themeh.bgcolor
        return

    def on_touch_downx(self, touch):
        super(Button, self).on_touch_down(touch)
        self.background_color = get_color_from_hex('#f1f1f1')
        return True

    def on_touch_upx(self, touch):
        super(Button, self).on_touch_up(touch)
        self.background_color = themeh.bgcolor
        return True


class Label(cbtn.Label):
    def __init__(self, **kwargs):
        super(Label, self).__init__(**kwargs)
        self.font_name = appctx.font_heiti
        # always set color
        self.color = get_color_from_hex('#000000FF')
        self.color = themeh.fgcolor
        self.valign = 'middle'
        self.halign = 'left'
        # self.size_hint = (1, 1)
        self.bind(size=self.setter('text_size'))
        return


class LabelSR(Label):
    def __init__(self, **kwargs):
        super(LabelSR, self).__init__(**kwargs)
        self.shorten = True
        self.shorten_from = 'right'
        return


class TextInput(ctxip.TextInput):
    def __init__(self, **kwargs):
        super(TextInput, self).__init__(**kwargs)
        self.font_name = appctx.font_heiti
        mybg = ismybg()
        if mybg: self.background_color = get_color_from_hex('#ffffff00')
        if mybg: self.color = get_color_from_hex('#00000088')
        return

