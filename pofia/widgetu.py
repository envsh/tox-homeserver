
from kivy.uix.widget import Widget
from kivy.uix.label import Label
from kivy.uix.boxlayout import BoxLayout


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


class VSpacer(Label):
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


class HSpacer(Label):
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
