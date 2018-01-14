import log

from kivy.lang import Builder
from kivy.uix.recycleview import RecycleView
from kivy.uix.widget import Widget
from kivy.uix.boxlayout import BoxLayout
from kivy.uix.image import Image
from kivy.graphics import Rectangle, Color

from appcontext import appctx, icopath
from theme import themeh

from widgetu import Button, Label, LabelSR, TextInput


Builder.load_string('''
#<PeerItemView>:
    # Draw a background to indicate selection
#    canvas.before:
        #Color:
        #    rgba: (.0, 0.9, .1, .3) if self.selected else (1, 1, 1, 1)
#        Rectangle:
#            pos: self.pos
#            size: self.size

<RVPeer>:
    viewclass: 'PeerItemView'
    RecycleBoxLayout:
        default_size: None, dp(26)
        default_size_hint: 1, None
        size_hint_y: None
        height: self.minimum_height
        orientation: 'vertical'
        multiselect: False
        touch_multiselect: False


# <ContactItemView>:
    # Draw a background to indicate selection
    # canvas.before:
        # Color:
            # rgba: (.0, .1, .0, .1) if self.selected else (1, 0, 1, 1)
    #    Rectangle:
    #        pos: self.pos
    #        size: self.size

<RVContact>:
    viewclass: 'ContactItemView'
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


class PeerItemView(RecycleDataViewBehavior, BoxLayout):
    index = None
    selected = BooleanProperty(False)
    selectable = BooleanProperty(True)

    def __init__(self):
        super(PeerItemView, self).__init__()
        self.orientation = 'horizontal'
        self.colorobj = Color(*themeh.bgcolor)
        self.canvas.add(self.colorobj)
        self.rectobj = Rectangle(pos=self.pos, size=self.size)
        self.canvas.add(self.rectobj)
        self.buildit()
        return

    def buildit(self):
        self.name = 'PofU123'
        self.icon = 'iiicon'

        # widgets
        self.nmbtn = None
        self.icobtn = None

        self.icobtn = Label(text=self.icon, size_hint_x=None, width=30, size_hint_y=None, height=30)
        self.icobtn = Image(source=icopath('icon_avatar_40'), size_hint_x=None, width=30)
        self.add_widget(self.icobtn)
        self.nmbtn = LabelSR(text=self.name, size_hint_x=1)
        self.add_widget(self.nmbtn)
        pass

    # why call 2 time when initialized
    def refresh_view_attrs(self, rv, index, data):
        ''' Catch and handle the view changes '''
        self.index = index
        # log.l.debug(str(index)+str(data))
        self.nmbtn.text = data['text']
        self.nmbtn.markup = data['markup']
        if index == 0:
            import traceback
            # traceback.print_stack()
        return super(PeerItemView, self).refresh_view_attrs(rv, index, data)

    def refresh_view_layout(self, rv, index, layout, viewport):
        return super(PeerItemView, self).refresh_view_layout(rv, index, layout, viewport)

    def on_touch_down(self, touch):
        ''' Add selection on touch down '''
        if super(PeerItemView, self).on_touch_down(touch):
            return True
        if self.collide_point(*touch.pos) and self.selectable and not self.selected:
            log.l.debug('selected....'+str(self.index)+str(self.selectable))
            # appctx.ctpage.data = []
            ctdata = appctx.rvct.data[self.index]
            log.l.debug(ctdata)
            if ctdata['ctid'] in appctx.ctmsgs:
                pass
            else:
                pass

            self.selected = True
            if self.colorobj in self.canvas.children:
                self.canvas.remove(self.colorobj)
            if self.rectobj in self.canvas.children:
                self.canvas.remove(self.rectobj)
            self.colorobj = Color(*themeh.selected_color)
            self.canvas.add(self.colorobj)
            self.rectobj = Rectangle(pos=self.pos, size=self.size)
            self.canvas.add(self.rectobj)
            return True
        else:
            self.selected = False
            if self.colorobj in self.canvas.children:
                self.canvas.remove(self.colorobj)
            if self.rectobj in self.canvas.children:
                self.canvas.remove(self.rectobj)
            return False

    def apply_selection(self, rv, index, is_selected):
        ''' Respond to the selection of items in the view. '''
        self.selected = is_selected
        if is_selected:
            log.l.debug("selection changed to {0}".format(rv.data[index]))
        else:
            log.l.debug("selection removed for {0}".format(rv.data[index]))


class ContactItemView(RecycleDataViewBehavior, BoxLayout):
    '''
    '''
    index = None
    selected = BooleanProperty(False)
    selectable = BooleanProperty(True)

    def __init__(self):
        super(ContactItemView, self).__init__()
        self.orientation = 'horizontal'
        self.colorobj = Color(*themeh.bgcolor)
        self.canvas.add(self.colorobj)
        self.rectobj = Rectangle(pos=self.pos, size=self.size)
        self.canvas.add(self.rectobj)
        # self.canvas.add(Color(*themeh.bgcolor))
        #with self.canvas:
        #    Color(*get_color_from_hex('#000000FF'))
        self.buildit()
        return

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
        self.nmbtn = LabelSR(text=self.name, size_hint_x=1)
        self.stmbtn = LabelSR(text=self.stmsg, size_hint_x=1)
        lo = BoxLayout(orientation='vertical')
        lo.add_widget(self.nmbtn)
        lo.add_widget(self.stmbtn)
        self.add_widget(lo)
        self.stsbtn = LabelSR(text=self.ststxt, size_hint_x=None, width=30)
        self.stsbtn = Image(source=icopath('offline_30'), allow_stretch=False, size_hint_x=None, width=12)
        self.stsbtn.text = '99'  # no use
        self.add_widget(self.stsbtn)
        self.mcbtn = Label(text=str(self.msgcnt), size_hint_x=None, width=30)
        self.add_widget(self.mcbtn)
        pass

    # why call 2 time when initialized
    def refresh_view_attrs(self, rv, index, data):
        ''' Catch and handle the view changes '''
        self.index = index
        log.l.debug(str(index)+str(data))
        self.nmbtn.text = data['text']
        self.nmbtn.markup = data['markup']
        self.stmbtn.text = data.get('stmsg')
        self.stmbtn.markup = True
        self.stsbtn.source = icopath('online_30') if data.get('status') > 0 else icopath('offline_30')
        self.icobtn.source = icopath('groupgray') if data.get('isgroup') is True else icopath('icon_avatar_40')
        if index == 0:
            import traceback
            # traceback.print_stack()
        return super(ContactItemView, self).refresh_view_attrs(rv, index, data)

    def refresh_view_layout(self, rv, index, layout, viewport):
        return super(ContactItemView, self).refresh_view_layout(rv, index, layout, viewport)

    def on_touch_down(self, touch):
        ''' Add selection on touch down '''
        if super(ContactItemView, self).on_touch_down(touch):
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
            if self.colorobj in self.canvas.children:
                self.canvas.remove(self.colorobj)
            if self.rectobj in self.canvas.children:
                self.canvas.remove(self.rectobj)
            self.colorobj = Color(*themeh.selected_color)
            self.canvas.add(self.colorobj)
            self.rectobj = Rectangle(pos=self.pos, size=self.size)
            self.canvas.add(self.rectobj)
            return True
        else:
            self.selected = False
            if self.colorobj in self.canvas.children:
                self.canvas.remove(self.colorobj)
            if self.rectobj in self.canvas.children:
                self.canvas.remove(self.rectobj)
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
        self.data = [{'text': self.styleText(str(x)+str(x)), 'markup': True, 'ctid': '--', 'stmsg':'---', 'status': 0, 'isgroup': False} for x in range(1)]
        # print(self.data)
        for fnum in appctx.jinfo['friends'].keys():
            frnd = appctx.jinfo['friends'][fnum]
            if 'name' in frnd: self.additem(frnd['name'], frnd['pubkey'], frnd.get('stmsg'), frnd.get('status'), False)
            else: self.additem('POFU.'+frnd['pubkey'][0:5], frnd['pubkey'], frnd.get('stmsg'), frnd.get('status'), False)
        if 'groups' in appctx.jinfo:
            for fnum in appctx.jinfo['groups'].keys():
                frnd = appctx.jinfo['groups'][fnum]
                if 'title' in frnd: self.additem(frnd['title'], frnd['groupId'], '', 0, True)
                else: self.additem('POFG.'+frnd['groupId'][0:5], frnd['groupId'], '', 0, True)

        for x in range(2): self.additem(self.styleText(str(x)+str(x)), '--', '', 0, False)

    def styleText(self, txt):
        return '[color=#123456]'+str(txt)+'[/color]'

    def additem(self, txt, ctid, stmsg, status, isgroup):
        if stmsg is None or stmsg == '': stmsg = '---'
        if status is None: status = 0
        if isgroup is None: isgroup = False
        self.data.append({'text': self.styleText(txt), 'markup': True,
                          'ctid': ctid, 'stmsg': stmsg, 'status': status,
                          'isgroup': isgroup})

        return
