{.hint[XDeclaredButNotUsed]:off.}

import macros
import unicode

macro NK_FLAGS(v:untyped) : untyped = result = quote do: 1 shl `v`

const NK_WINDOW_BORDER            = 1 shl 0
const NK_WINDOW_MOVABLE           = 1 shl 1
const NK_WINDOW_SCALABLE          = 1 shl 2
const NK_WINDOW_CLOSABLE          = 1 shl 3
const NK_WINDOW_MINIMIZABLE       = 1 shl 4
const NK_WINDOW_NO_SCROLLBAR      = 1 shl 5
const NK_WINDOW_TITLE             = 1 shl 6
const NK_WINDOW_SCROLL_AUTO_HIDE  = 1 shl 7
const NK_WINDOW_BACKGROUND        = 1 shl 8
const NK_WINDOW_SCALE_LEFT        = 1 shl 9
const NK_WINDOW_NO_INPUT          = 1 shl 10

const NK_DYNAMIC = 0
const NK_STATIC = 1

const NK_TEXT_ALIGN_LEFT        = 0x01
const NK_TEXT_ALIGN_CENTERED    = 0x02
const NK_TEXT_ALIGN_RIGHT       = 0x04
const NK_TEXT_ALIGN_TOP         = 0x08
const NK_TEXT_ALIGN_MIDDLE      = 0x10
const NK_TEXT_ALIGN_BOTTOM      = 0x20

const NK_TEXT_LEFT        = NK_TEXT_ALIGN_MIDDLE or NK_TEXT_ALIGN_LEFT
const NK_TEXT_CENTERED    = NK_TEXT_ALIGN_MIDDLE or NK_TEXT_ALIGN_CENTERED
const NK_TEXT_RIGHT       = NK_TEXT_ALIGN_MIDDLE or NK_TEXT_ALIGN_RIGHT


const NK_EDIT_DEFAULT                 = 0
const NK_EDIT_READ_ONLY               = 1 shl (0)
const NK_EDIT_AUTO_SELECT             = 1 shl (1)
const NK_EDIT_SIG_ENTER               = 1 shl (2)
const NK_EDIT_ALLOW_TAB               = 1 shl (3)
const NK_EDIT_NO_CURSOR               = 1 shl (4)
const NK_EDIT_SELECTABLE              = 1 shl (5)
const NK_EDIT_CLIPBOARD               = 1 shl (6)
const NK_EDIT_CTRL_ENTER_NEWLINE      = 1 shl (7)
const NK_EDIT_NO_HORIZONTAL_SCROLL    = 1 shl (8)
const NK_EDIT_ALWAYS_INSERT_MODE      = 1 shl (9)
const NK_EDIT_MULTILINE               = 1 shl (10)
const NK_EDIT_GOTO_END_ON_ACTIVATE    = 1 shl (11)

const NK_EDIT_SIMPLE  = NK_EDIT_ALWAYS_INSERT_MODE
const NK_EDIT_FIELD   = NK_EDIT_SIMPLE or NK_EDIT_SELECTABLE or NK_EDIT_CLIPBOARD
const NK_EDIT_BOX     = NK_EDIT_ALWAYS_INSERT_MODE or NK_EDIT_SELECTABLE or NK_EDIT_MULTILINE or NK_EDIT_ALLOW_TAB or NK_EDIT_CLIPBOARD
const NK_EDIT_EDITOR  = NK_EDIT_SELECTABLE or NK_EDIT_MULTILINE or NK_EDIT_ALLOW_TAB or NK_EDIT_CLIPBOARD

type XFont = pointer
type nk_context = pointer
type nk_window = pointer
type nk_input = pointer

type nk_color = object
    r*: uint8
    g*: uint8
    b*: uint8
    a*: uint8

type nk_rect = object
    x*: float32
    y*: float32
    w*: float32
    h*: float32

type nk_vec2 = object
    x*: float32
    y*: float32


proc nk_xfont_create(dpy:PXDisplay, name:cstring) : XFont {.importc.}
proc nk_xlib_init(f: XFont, dpy: PXDisplay, scrn: cint, root: TWindow, vis: PVisual, cmap: TColormap, w: int, h: int) : nk_context {.importc.}
proc nk_xlib_handle_event(dpy:PXDisplay, scrn: cint, win: TWindow, evt: PXEvent) {.importc}
proc nk_xlib_render(scrn:TWindow, clear: nk_color) {.importc.}

proc nk_input_begin(ctx:nk_context) {.importc.}
proc nk_input_end(ctx:nk_context) {.importc.}
proc nk_input_unicode(ctx:nk_context, r:Rune) {.importc.}
proc nk_input_is_mouse_hovering_rect(ipt:nk_input, bounds:nk_rect) : int {.importc.}
proc nk_widget_bounds(ctx:nk_context) : nk_rect {.importc.}
proc nk_get_input(ctx:nk_context) : nk_input {.importc.}
proc nk_get_window(ctx:nk_context) : nk_window {.importc.}
proc nk_curwnd_scrollto(ctx:nk_context, sbx, sby : int) {.importc.}

proc nk_begin(ctx:nk_context, name:cstring, r: nk_rect, flags: cint) : cint {.importc.}
proc nk_end(ctx:nk_context) : cint {.importc.}

proc nk_layout_row_static(ctx:nk_context, height: float32, item_width: cint, cols : cint)  {.importc.}
proc nk_layout_row_dynamic(ctx:nk_context, height: float32,  cols: cint) {.importc.}
proc nk_layout_row_begin(ctx: nk_context, fmt:cint,row_height:float32, cols:int)  {.importc.}
proc nk_layout_row_end(ctx:nk_context)  {.importc.}
proc nk_layout_row_push(ctx:nk_context, ratio_or_width:float32) {.importc.}

proc nk_selectable_label(ctx:nk_context, title:cstring, align:cint, value: ptr int) : int {.importc.}
proc nk_label(ctx:nk_context,str:cstring, alignment:cint) {.importc.}
proc nk_edit_string(ctx: nk_context, typ:cint, buffer:cstring, len:ptr int,  max:int, nk_plugin_filter:pointer) : cint {.importc.}

proc nk_button_text(ctx:nk_context, title:cstring, len : cint) : int {.importc.}
proc nk_button_label(ctx:nk_context, title: cstring) : int {.importc.}
# NK_API int nk_button_image(struct nk_context*, struct nk_image img);

proc nk_combo_begin_label(ctx:nk_context, selected: cstring, size:nk_vec2):int{.importc.}
proc nk_combo_end(ctx:nk_context) {.importc.}

proc nk_checkbox_label(ctx: nk_context, title:cstring, active:ptr int) : int {.importc.}

proc nk_popup_begin(ctx: nk_context, nk_popup_type:cint, title:cstring, nk_flags:cint, bounds: nk_rect ) : int {.importc.}

proc nk_popup_close(ctx: nk_context) {.importc.}
proc nk_popup_end(ctx: nk_context) {.importc.}

proc nk_menu_begin_label(ctx: nk_context, title:cstring, align: cint, size: nk_vec2 ): int {.importc.}
proc nk_menu_end(ctx: nk_context) {.importc.}
proc nk_menu_item_label(ctx: nk_context, title:cstring, alignment: cint) : int {.importc.}

proc nk_tooltip(ctx: nk_context, title:cstring) {.importc.}
proc nk_tooltip_begin(ctx: nk_context,  width:float32): int{.importc.}
proc nk_tooltip_end(ctx: nk_context){.importc.}

###

type
    nk_editor_state = ref object
        edbuf*: array[256,char]
        edlen*: int

proc neweditstate() : nk_editor_state =
    var e = new(nk_editor_state)
    return e

proc getstr(e: nk_editor_state) : string =
    var str = cast[string](@(e.edbuf))
    return str.substr(0, e.edlen-1)

type
    xim_state = ref object
        ximbuf*: array[64,char]
        gotlen*: int

proc newximstate() : xim_state =
    var i = new(xim_state)
    return i

proc ximstr(e:xim_state) : string =
    var str = cast[string](@(e.ximbuf))
    return str.substr(0, e.gotlen-1)

