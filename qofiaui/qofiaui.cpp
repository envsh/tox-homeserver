
// #include "ui_add_friend.h"
// #include "ui_contact_item_view.h"
// #include "ui_create_room.h"
// #include "ui_emoji_category.h"
// #include "ui_emoji_panel.h"
// #include "ui_message_item_view.h"
// //#include "ui_scroll_widget.h"
// #include "ui_untitled.h"

#include <QtWidgets>

#include "mainwin.h"
#include "message_item.h"
#include "contact_item.h"

extern "C"
void qofiaui_main() {
    int argc = 1;
    char*argv[] = {(char*)"qofiaui",NULL};

    QApplication app(argc, argv);

    auto* mw = new MainWin();
    mw->show();

    app.exec();
}
