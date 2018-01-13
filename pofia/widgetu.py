
from kivy.uix.boxlayout import BoxLayout


class VBoxLayout(BoxLayout):
    def __init__(self, **kwargs):
        super(VBoxLayout, self).__init__(**kwargs)
        self.orientation = 'vertical'


class HBoxLayout(BoxLayout):
    def __init__(self, **kwargs):
        super(HBoxLayout, self).__init__(**kwargs)
        self.orientation = 'horizontal'

