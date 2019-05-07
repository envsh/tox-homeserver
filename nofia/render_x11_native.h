#ifndef _RENDER_X11_NATIVE_H_
#define _RENDER_X11_NATIVE_H_

#include <stdio.h>
#include <stdlib.h>

#include <malloc.h>

#include <X11/Xlib.h>

#include <X11/keysym.h>

/*
 * include <locale.h> or the non-standard X substitutes
 * depending on the X_LOCALE compilation flag
 */
#include <X11/Xlocale.h>

typedef struct {
    Display *dpy;
    int screen;
    Window root;
    Window win;
    GC gc;
    XGCValues gcv;
    XEvent event;
    XFontSet fontset;
    XIM im;
    XIC ic;
    Visual *vis;
    Colormap cmap;
    XWindowAttributes attr;
    XSetWindowAttributes swa;
    Atom wm_delete_window;

    XIMStyles *im_supported_styles;
    XIMStyle app_supported_styles;
    XIMStyle style;
    XIMStyle best_style;
    XVaNestedList list;
    long im_event_mask;
    XRectangle preedit_area;
    XRectangle status_area;
    char **missing_charsets;
    int num_missing_charsets;
    char *default_string;

} RenderWindow;


RenderWindow* NewRenderWindow();
void nk_x11_event_handle(RenderWindow* rdwin);
Colormap RenderWindowCmap(RenderWindow* rdwin);
Visual* RenderWindowVis1(RenderWindow* rdwin);

#endif
