
type XFont = pointer
type nk_context = pointer

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


proc nk_xfont_create(dpy:PXDisplay, name:cstring) : XFont {.importc.}
proc nk_xlib_init(f: XFont, dpy: PXDisplay, scrn: cint, root: TWindow, vis: PVisual, cmap: TColormap, w: int, h: int) : nk_context {.importc.}
proc nk_xlib_handle_event(dpy:PXDisplay, scrn: cint, win: TWindow, evt: PXEvent) {.importc}
proc nk_xlib_render(scrn:TWindow, clear: nk_color) {.importc.}

proc nk_input_begin(ctx:nk_context) {.importc.}
proc nk_input_end(ctx:nk_context) {.importc.}

proc nk_begin(ctx:nk_context, name:cstring, r: nk_rect, flags: cint) : cint {.importc.}
proc nk_end(ctx:nk_context) : cint {.importc.}

proc nk_layout_row_static(ctx:nk_context, height: float32, item_width: cint, cols : cint)  {.importc.}
proc nk_label(ctx:nk_context,str:cstring, alignment:cint) {.importc.}

