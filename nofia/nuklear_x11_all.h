#ifndef _NUKLEAR_X11_ALL_H_
#define _NUKLEAR_X11_ALL_H_

#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>
#include <string.h>
#include <limits.h>
#include <math.h>
#include <sys/time.h>
#include <unistd.h>
#include <time.h>

#define NK_INCLUDE_FIXED_TYPES
#define NK_INCLUDE_STANDARD_IO
#define NK_INCLUDE_STANDARD_VARARGS
#define NK_INCLUDE_DEFAULT_ALLOCATOR
// #define NK_INCLUDE_FONT_BAKING
#define STB_IMAGE_IMPLEMENTATION
#define NK_XLIB_INCLUDE_STB_IMAGE
#define NK_XLIB_USE_XFT
#define NK_IMPLEMENTATION
#define NK_XLIB_IMPLEMENTATION

// 只需要包含类型信息即可
#ifdef _INCLUDED_INGO_
#include <X11/Xlib.h>
#include <X11/Xft/Xft.h>
#include <X11/Xresource.h>
#include </home/me/oss/nuklear/src/nuklear.h>

// some xlib type
struct XFont {
    int ascent;
    int descent;
    int height;
#ifdef NK_XLIB_USE_XFT
    XftFont * ft;
#else
    XFontSet set;
    XFontStruct *xfont;
#endif
    struct nk_user_font handle;
};

// typedef struct XFont XFont;
// typedef struct XSurface XSurface;

// from stb_image.h
typedef void* stbi__jpeg;
typedef unsigned char stbi_uc;

// fix unhandled relocation for stderr, part0
// put here to export via cgo
FILE *getStdout(void) { return stdout; }
FILE *getStderr(void) { return stderr; }

#else

// fix unhandled relocation for stderr, part1
// using symbol exported via cgo before
extern FILE *getStderr();
extern FILE *getStdout();
// #undef stderr
// #undef stdout
// #define stderr getStderr()
// #define stdout getStdout()

#include </home/me/oss/nuklear/nuklear.h>
// #include </home/me/oss/nuklear/example/stb_image.h>

#include </home/me/oss/nuklear/demo/x11_xft/nuklear_xlib.h>

#define INCLUDE_ALL
/*#define INCLUDE_ALL */
/*#define INCLUDE_STYLE */
/*#define INCLUDE_CALCULATOR */
/*#define INCLUDE_OVERVIEW */
/*#define INCLUDE_NODE_EDITOR */

#include </home/me/oss/nuklear/demo/style.c>

#endif


#endif
