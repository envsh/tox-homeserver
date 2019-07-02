
#include "qofiaui.h"

#include "mainwin.h"
#include "contact_item.h"
#include "message_item.h"

enum {
      UIST_QMLMCTRL = 0                   ,
      UIST_QMLORIGIN                      ,
      UIST_SETTINGS                       ,
      UIST_LOGINUI                        ,
      UIST_CONTACTUI                      ,
      UIST_MESSAGEUI                      ,
      UIST_VIDEOUI                        ,
      // UIST_PICKCALLUI // TODO video 
      UIST_ADD_GROUP                      ,
      UIST_ADD_FRIEND                     ,
      UIST_INVITE_FRIEND                  ,
      UIST_MEMBERS                        ,
      UIST_CONTACT_INFO                   ,
      UIST_TESTUI                         ,
      UIST_LOGUI                          ,
      UIST_ABOUTUI                        ,
};

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
    for (int i = 0; i < 500; i ++) {
        msgviews.append(new MessageItem());
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
}

void MainWin::AddContactItem(QString uid, QString name, QString stmsg) {
    qInfo()<<uid<<name;
    auto ctv = new ContactItem();
    ctv->uid = uid;
    ctv->uiw.label_2->setText(name);
    ctv->uiw.label_3->setText(stmsg);

    connect(ctv, &ContactItem::clicked, this, &MainWin::ctitem_clicked, Qt::QueuedConnection);
    auto lo9 = uiw.verticalLayout_9;
    lo9->insertWidget(0, ctv);
}

void MainWin::AddConferenceMessage(QString uid, QString msg) {
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
        w->uiw.label_4->setText(msg); // lastmsg
        break;
    }
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

    // switched, loading new info
    uiw.label_5->setText(cti->uiw.label_2->text());
    uiw.label_7->setText(cti->uiw.label_3->text());

    char* bcc = uictx->uion_loadmsg (uid.toUtf8().data(), maxmsgcnt);
    QString scc = QString::fromUtf8(bcc);
    free(bcc);
    // [[msg],[msg2],]
    qInfo() << scc.length() << scc.left(32);

    auto vlo8 = uiw.verticalLayout_3;
    for (int i = 0; i < vlo8->count(); i++) {
        vlo8->removeWidget(msgviews.at(i));
    }

    QJsonDocument jdoc = QJsonDocument::fromJson(scc.toUtf8());
    auto msgos = jdoc.array();
    for (int i = 0; i < msgos.count() && i < maxmsgcnt; i++) {
        auto msgo = msgos.at(i).toArray();
        auto msgv = (MessageItem*)msgviews.at(i);
        msgv->uiw.label_5->setText(msgo.at(0).toString());
        msgv->uiw.LabelUserName4MessageItem->setText(msgo.at(1).toString());
        msgv->uiw.LabelMsgTime->setText(msgo.at(2).toString());
        vlo8->addWidget(msgv);
    }
}


