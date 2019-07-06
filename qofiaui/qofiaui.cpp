
// #include "ui_add_friend.h"
// #include "ui_contact_item_view.h"
// #include "ui_create_room.h"
// #include "ui_emoji_category.h"
// #include "ui_emoji_panel.h"
// #include "ui_message_item_view.h"
// //#include "ui_scroll_widget.h"
// #include "ui_untitled.h"

#include <QtCore>
#include <QtWidgets>
#include <QDebug>

#include "mainform.h"
#include "mainwin.h"
#include "message_item.h"
#include "contact_item.h"
#include "qofiaui.h"
#include "event.h"

// extern "C"
qofiaui_context uictxm = {0};
qofiaui_context* uictx = &uictxm;
MainWin* gmw = 0;
MainForm* gmw2 = 0;

#include <malloc.h>

void qofiaui_main(qofiaui_context* ctx) {
    int argc = 1;
    char*argv[] = {(char*)"qofiaui",NULL};
    uictx->uion_command = ctx->uion_command;
    uictx->uion_loadmsg = ctx->uion_loadmsg;

    int rv = mallopt(M_TRIM_THRESHOLD, 16000);
    rv = mallopt(M_MMAP_THRESHOLD, 16000);
    qSetMessagePattern("%{file}(%{line}): %{message}");
    QApplication app(argc, argv);

    int uiflag = 1;
    if (uiflag == 0) {
        auto btn = new QPushButton();
        btn->show();
    }else if (uiflag == 1) {
        auto mw = new MainForm();
        gmw2 = mw;
        QObject::connect(mw, &MainForm::cmdhandle, mw,
                         &MainForm::qofiaui_cmdproc, Qt::QueuedConnection);
        mw->show();
        mw->setform(UIST_LOGINUI);
    }else if (uiflag == 2){
        auto mw = new MainWin();
        gmw = mw;
        QObject::connect(mw, &MainWin::cmdhandle, mw,
                         &MainWin::qofiaui_cmdproc, Qt::QueuedConnection);
        mw->show();
    }

    app.exec();
}

void uion_command(QString cmd) {
    uictx->uion_command(cmd.toUtf8().data());
}


void MainWin::qofiaui_cmdproc(QString cmdmsg) {
    qInfo()<<cmdmsg;
    QJsonParseError perr;
    QJsonDocument jdoc = QJsonDocument::fromJson(cmdmsg.toUtf8(), &perr);
    if (perr.error != QJsonParseError::NoError) {
        qInfo()<<perr.errorString();
    }
    QJsonObject jobj = jdoc.object();
    QString evtname = jobj.value("EventName").toString();
    Event evto = Event::fromJson(jdoc);

    auto jarr = jobj.value("Args").toArray();
    auto marr = jobj.value("Margs").toArray();
    if (evtname == "SelfInfo") {
        SetSelfInfo(jarr.at(1).toString(), jarr.at(2).toString());
    }else if (evtname == "AddFriendItem") {
        AddContactItem(jarr.at(1).toString(), jarr.at(2).toString(), jarr.at(3).toString());
    }else if (evtname == "AddGroupItem") {
        AddContactItem(jarr.at(1).toString(), jarr.at(2).toString(), jarr.at(3).toString());
    }else if (evtname == "ConferenceMessage") {
        AddConferenceMessage(marr.at(3).toString(), jarr.at(3).toString());
    }else {
        qInfo()<<"todo"<<evtname;
    }
}

void qofiaui_dmcommand(char* cmdmsgc) {
    QString cmdmsg = QString::fromUtf8(cmdmsgc);
    // qInfo()<<__FILE__<<":"<<__LINE__<<": "<<cmdmsg;
    // emit gmw->cmdhandle(cmdmsg);
    emit gmw2->cmdhandle(cmdmsg);
}
