#include "qofiaui.h"
#include "uiutils.h"
#include "message_item.h"
#include "chatform.h"

ChatForm::ChatForm(QWidget* parent)  :  QWidget(parent)
{
    uiw.setupUi(this);

    connect(uiw.toolButton_18, &QToolButton::clicked, this, &ChatForm::sendmsg);

    ccstate.isBottom = true;
    auto sa2vb = uiw.scrollArea_2->verticalScrollBar();
    connect(sa2vb, &QScrollBar::rangeChanged,
            [this,sa2vb](int min, int max) {
                int curpos = sa2vb->value();
                if (ccstate.isBottom && curpos < max) {
                    sa2vb->setValue(max);
                }
                ccstate.maxpos = max;
            });
    connect(sa2vb, &QScrollBar::valueChanged,
            [this](int value){
                ccstate.curpos = value;
                int maxval = ccstate.maxpos;
                ccstate.isBottom = value >= maxval ? true : false;
                ccstate.maxpos = value > maxval ? value : maxval;
            });
}

ChatForm::~ChatForm() { dtor(); }
void ChatForm::dtor() {
    auto vlo8 = uiw.verticalLayout_3;
    int curcnt = vlo8->count();
    for (int i = 0; i < curcnt; i++) {
        auto item = vlo8->takeAt(0);
        auto w = item->widget();
        wdelete(w);
        delete item;
    }
}

void ChatForm::sendmsg() {
    auto msg = uiw.textEdit_3->toPlainText();
    uiw.textEdit_3->clear();
    if (msg.isEmpty()) return;
    QStringList cmd = {"sendmsg", curuid};
    cmd << msg;
    uion_command(cmd.join(uicmdsep));
}

void ChatForm::scrollEnd() {
    if (!ccstate.isBottom) {return;}
    auto sb = uiw.scrollArea_2->verticalScrollBar();
    sb->setValue(ccstate.maxpos);
}

void ChatForm::AddConferenceMessage(QString uid, QString msg, QString peername, QString timestr) {
    AddConferenceMessage1(uid,msg,peername,timestr);
    if (uid == curuid) {
        AddConferenceMessage2(uid,msg,peername, timestr);
        scrollEnd();
    }
}
void ChatForm::AddConferenceMessage1(QString uid, QString msg, QString peername, QString timestr) {
    if (uid != curuid) {
        SetQLabelElideText(uiw.label_7,msg,"..",false);
    }
}
// append to message list
void ChatForm::AddConferenceMessage2(QString uid, QString msg, QString peername, QString timestr) {
    auto vlo8 = uiw.verticalLayout_3;
    int curcnt = vlo8->count();
    if (curcnt >= maxmsgcnt) {
        qWarning()<<"too many msgs"<<curcnt;
        auto item = vlo8->takeAt(0);
        wdelete(item->widget());
        delete item;
    }
    MessageItem* msgv = new MessageItem();
    vlo8->addWidget(msgv);

    msgv->uiw.label_5->setText(msg);
    msgv->uiw.LabelUserName4MessageItem->setText(peername);
    msgv->uiw.LabelMsgTime->setText(timestr);

    uiw.LabelMsgCount2->setText(QString::number(curcnt+1));
}

void ChatForm::setandload(QString uid, QString ctname, QString ctstmsg) {
    qInfo()<<uid;
    curuid = uid;

    // switched, loading new info
    uiw.label_5->setText(ctname);
    uiw.label_7->setText(ctstmsg);

    char* bcc = uictx->uion_loadmsg (uid.toUtf8().data(), maxmsgcnt);
    QString scc = QString::fromUtf8(bcc);
    free(bcc);
    // [[msg],[msg2],]
    qInfo() << scc.length() << scc.left(32);

    auto vlo8 = uiw.verticalLayout_3;
    int curcnt = vlo8->count();
    for (int i = 0; i < curcnt; i++) { }

    QJsonDocument jdoc = QJsonDocument::fromJson(scc.toUtf8());
    auto msgos = jdoc.array();
    for (int i = 0; i < msgos.count() && i < maxmsgcnt; i++) {
        auto msgo = msgos.at(i).toArray();
        auto msgv = new MessageItem();
        vlo8->addWidget(msgv);

        // msgv->uiw.label_5->setText(msgo.at(0).toString());
        msgv->uiw.label_5->setText(msgo.at(3).toString()); // html format
        msgv->uiw.LabelUserName4MessageItem->setText(msgo.at(1).toString());
        msgv->uiw.LabelMsgTime->setText(msgo.at(2).toString());
    }
    uiw.LabelMsgCount2->setText(QString::number(msgos.count()));
}

