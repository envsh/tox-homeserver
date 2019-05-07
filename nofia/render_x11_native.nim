##
##  include <locale.h> or the non-standard X substitutes
##  depending on the X_LOCALE compilation flag
##

type
    TRenderWindow* = ptr RenderWindow
    RenderWindow* = object
        dpy*: ptr TDisplay
        screen*: cint
        root*: TWindow
        win*: TWindow
        gc*: TGC
        gcv*: TXGCValues
        event*: TXEvent
        fontset*: TXFontSet
        im*: culong # TXIM
        ic*: culong # TXIC
        vis*: ptr TVisual
        cmap*: TColormap
        attr*: TXWindowAttributes
        swa*: TXSetWindowAttributes
        wm_delete_window*: TAtom

        im_supported_styles*: ptr TXIMStyles
        app_supported_styles*: TXIMStyle
        style*: TXIMStyle
        best_style*: TXIMStyle
        list*: TXVaNestedList
        im_event_mask*: clong
        preedit_area*: TXRectangle
        status_area*: TXRectangle
        missing_charsets*: cstringArray
        num_missing_charsets*: cint
        default_string*: cstring


proc NewRenderWindow*():  TRenderWindow
    {.importc:"NewRenderWindow", header:"render_x11_native.h".}
proc nk_x11_event_handle*(rdwin: TRenderWindow)
    {.importc:"nk_x11_event_handle", header:"render_x11_native.h".}

proc RenderWindowCmap*(rdwin:TRenderWindow): TColormap
    {.importc:"RenderWindowCmap", header:"render_x11_native.h".}

proc RenderWindowVis*(rdwin:TRenderWindow): ptr TVisual
    {.importc:"RenderWindowVis1", header:"render_x11_native.h".}
