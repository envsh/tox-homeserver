from kivy.utils import get_color_from_hex, get_hex_from_color, get_random_color
from kivy.graphics import Rectangle, Color


class Theme:
    def __init__(self, thname):
        self.thname = thname
        self.bgcolor = None
        self.fgcolor = None
        self.hint_color = None  # for status, hint like text
        self.selected_color = None

        if thname == 'dark': self.dark()
        else: self.light()
        return

    def dark(self):
        self.bgcolor = (0, 0, 1, 1)
        self.bgcolor = get_color_from_hex('#000000FF')
        self.fgcolor = get_color_from_hex('#FFFFFF88')
        self.hint_color = get_color_from_hex('#acacac88')
        self.selected_color = get_color_from_hex('#00450888')
        return

    def light(self):
        self.bgcolor = (1, 1, 1, 1)
        self.bgcolor = get_color_from_hex('#FFFFFF00')
        self.fgcolor = get_color_from_hex('#000000FF')
        self.hint_color = get_color_from_hex('#acacac88')
        self.selected_color = get_color_from_hex('#00450888')
        return


themeh = Theme('light')
#themeh = Theme('dark')


def themeo():
    global themeh
    return themeh


if __name__ == '__main__':
    print(themeh.bgcolor)
    print(tuple(themeh.bgcolor))
    print(Color(0, 0, 0, 1))
    print(Color(*themeh.bgcolor))
