
#include "qofiaui.h"

#include "mainwin.h"
#include "contact_item.h"

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

MainWin::MainWin(QWidget* parent)
    : QMainWindow(parent) {
    uiw.setupUi(this);


    connect(uiw.pushButton_7, &QPushButton::clicked, [this](){
            QString cmd = "login/";
            cmd += this->uiw.comboBox_6->currentText();
            uion_command(cmd);
            this->uiw.stackedWidget->setCurrentIndex(UIST_CONTACTUI);
        });
}

void MainWin::AddContactItem(QString uid, QString name, QString stmsg) {
    qInfo()<<uid<<name;
    auto ctv = new ContactItem();
    ctv->uid = uid;
    ctv->uiw.label_2->setText(name);
    ctv->uiw.label_3->setText(stmsg);

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
