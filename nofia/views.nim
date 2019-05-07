

proc nkwndhello0(nkw:PNkwindow, name : string) {.gcsafe.} =
    var cname : cstring = "hehehhehehee"
    linfo("wndname", name)
    if nk_begin(nkw.nkctx, cname, nk_rect(x:0, y:0, w:300, h:200),
                NK_WINDOW_TITLE or NK_WINDOW_BORDER or
                NK_WINDOW_MOVABLE or NK_WINDOW_SCALABLE) == ctrue:
        nk_layout_row_static(nkw.nkctx, 30, 200, 1)
        nk_label(nkw.nkctx, "hehehhehe", 1)
        discard
    discard nk_end(nkw.nkctx)
    return

proc nkwndhello1(nkw:PNkwindow, name : string) {.gcsafe.} =
    return

proc nkwndhello2(nkw:PNkwindow, name : string) {.gcsafe.} =
    return

proc nkwndrunproc(nkw:PNkwindow) {.gcsafe.} =
    for name, wp in nkw.wnds.pairs: wp(nkw, name)
    return

proc createnkwndprocs(nkw:PNkwindow) =
    nkw.wnds.add("hello0",nkwndhello0)
    nkw.wnds.add("hello1",nkwndhello1)
    nkw.wnds.add("hello2",nkwndhello1)
    nkw.wndrunner = nkwndrunproc
    return
