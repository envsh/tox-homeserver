#include <QtWidgets>

#include "qofiaui.h"
#include "uiutils.h"

#include "mainwin.h"
#include "contact_item.h"
#include "message_item.h"


const int maxmsgcnt = 500; // 最多显示消息个数，每个联系人

MainWin::MainWin(QWidget* parent)
    : QMainWindow(parent) {
    uiw.setupUi(this);

    // MessageItem*
    // 3000 193M
    // 2000 151M
    // 1000 108M
    // 500 87M
    // QLabel*
    // 10000 74M
    // 3000 69M
    // 2000 66M?
    // 1000 66M
    // QWidget*
    // 10000 71M?
    // 3000 67M?
    // 2000 67M?
    // 1000 66M?
    // 0 65M
    for (int i = 0; i < maxmsgcnt; i ++) {
        // msgviews.append(new MessageItem()); // 12M
        // msgviews.append(new QLabel());
        // msgviews.append(new QWidget());
    }

    connect(uiw.pushButton_7, &QPushButton::clicked, [this](){
            QString cmd = "login/";
            cmd += this->uiw.comboBox_6->currentText();
            uion_command(cmd);
            this->switchStackUi(UIST_CONTACTUI);
        });
    connect(uiw.toolButton_33, &QToolButton::clicked, this, &MainWin::backStackUi);

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
    prepui();
}
void MainWin::prepui() {
    uiw.lineEdit_5->setVisible(false);
    uiw.lineEdit_6->setVisible(false);
}

void MainWin:: SetSelfInfo(QString name, QString stmsg) {
    uiw.label_2->setText(name);
    // uiw.label_3->setText(stmsg);
    SetQLabelElideText(uiw.label_3,stmsg,"",false);
}

void MainWin::AddContactItem(QString uid, QString name, QString stmsg) {
    qInfo()<<uid<<name;
    auto ctv = new ContactItem();
    ctv->uid = uid;
    ctv->uiw.label_2->setText(name);
    // SetQLabelElideText(ctv->uiw.label_2,name,"..",true);
    // ctv->uiw.label_3->setText(stmsg);
    SetQLabelElideText(ctv->uiw.label_3,stmsg,"..",true);

    connect(ctv, &ContactItem::clicked, this, &MainWin::ctitem_clicked, Qt::QueuedConnection);
    auto lo9 = uiw.verticalLayout_9;
    lo9->insertWidget(0, ctv);
}

void MainWin::AddConferenceMessage(QString uid, QString msg) {
    AddConferenceMessage1(uid,msg);
    AddConferenceMessage2(uid,msg);
}
void MainWin::AddConferenceMessage1(QString uid, QString msg) {
    auto lo9 = uiw.verticalLayout_9;
    int cnt = lo9->count();

    for (int i = 0; i < cnt; i++) {
        auto witem = lo9->itemAt(i);
        auto w = (ContactItem*)witem->widget();
        if (w->uid != uid) {
            continue;
        }
        if (i > 0) {
            lo9->removeWidget(w);
            lo9->insertWidget(0, w);
        }
        // w->uiw.label_4->setText(msg); // lastmsg
        SetQLabelElideText(w->uiw.label_4,msg,"..",false);
        break;
    }
    if (uid != curuid && uiw.stackedWidget->currentIndex() == UIST_MESSAGEUI) {
        SetQLabelElideText(uiw.label_7,msg,"..",false);
    }
}
// append to message list
void MainWin::AddConferenceMessage2(QString uid, QString msg) {
    if (uid != curuid) {
        return;
    }
    auto vlo8 = uiw.verticalLayout_3;
    int curcnt = vlo8->count();
    MessageItem* msgv = nullptr;
    if (curcnt >= msgviews.count()) {
        qWarning()<<"too many msgs"<<curcnt;
        // swap index 0 to newest
        msgv = (MessageItem*)msgviews.takeAt(0);
        msgv->clear();
        vlo8->removeWidget(msgv);
        msgviews.append(msgv);
    }else {
        msgv = (MessageItem*)msgviews.at(curcnt);
    }

    vlo8->addWidget(msgv);

    msgv->uiw.label_5->setText(msg);
    // msgv->uiw.LabelUserName4MessageItem->setText(msgo.at(1).toString());
    // msgv->uiw.LabelMsgTime->setText(msgo.at(2).toString());

    uiw.LabelMsgCount2->setText(QString::number(curcnt+1));
}

void MainWin:: switchStackUi(int idx) {
    auto stkw = uiw.stackedWidget;
    int curidx = stkw->currentIndex();
    if (idx == curidx) return;
    uistks.push(curidx);
    uiw.stackedWidget->setCurrentIndex(idx);
}
void MainWin:: backStackUi() {
    if (uistks.length() == 1) {
        qInfo()<<"noback"<<uistks;
        return;
    }
    auto stkw = uiw.stackedWidget;
    int curidx = stkw->currentIndex();
    int idx = uistks.pop();
    stkw->setCurrentIndex(idx);
}

extern qofiaui_context* uictx;

void MainWin::ctitem_clicked(QString uid, QWidget* that) {
    qInfo()<<uid<<that;
    auto cti = (ContactItem*)that;

    switchStackUi(UIST_MESSAGEUI);
    curuid = uid;

    // switched, loading new info
    uiw.label_5->setText(cti->uiw.label_2->text());
    uiw.label_7->setText(cti->uiw.label_3->text());

    char* bcc = uictx->uion_loadmsg (uid.toUtf8().data(), maxmsgcnt);
    QString scc = QString::fromUtf8(bcc);
    free(bcc);
    // [[msg],[msg2],]
    qInfo() << scc.length() << scc.left(32);

    auto vlo8 = uiw.verticalLayout_3;
    int curcnt = vlo8->count();
    for (int i = 0; i < curcnt; i++) {
        auto msgv = (MessageItem*)msgviews.at(i);
        msgv->clear();
        vlo8->removeWidget(msgv);
    }

    QJsonDocument jdoc = QJsonDocument::fromJson(scc.toUtf8());
    auto msgos = jdoc.array();
    for (int i = 0; i < msgos.count() && i < maxmsgcnt; i++) {
        auto msgo = msgos.at(i).toArray();
        auto msgv = (MessageItem*)msgviews.at(i);
        vlo8->addWidget(msgv);

        msgv->uiw.label_5->setText(msgo.at(0).toString());
        msgv->uiw.LabelUserName4MessageItem->setText(msgo.at(1).toString());
        msgv->uiw.LabelMsgTime->setText(msgo.at(2).toString());
    }
    uiw.LabelMsgCount2->setText(QString::number(msgos.count()));
}


