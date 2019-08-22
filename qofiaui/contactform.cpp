#include "uiutils.h"
#include "contact_item.h"
#include "contactform.h"

ContactForm::ContactForm(QWidget* parent)
    : QWidget(parent) {
    uiw.setupUi(this);

    prepui();
}

void ContactForm::prepui() {
    uiw.lineEdit_5->setVisible(false);
    uiw.lineEdit_6->setVisible(false);

    updateTime();
    auto tmer = new QTimer();
    connect(tmer, &QTimer::timeout, this, &ContactForm::updateTime);
    tmer->start(1000);
}

void ContactForm::updateTime() {
    auto nowt = QDateTime::currentDateTime();
    auto secstr = nowt.toString("HH:mm:ss  ");
    uiw.label->setText(secstr);
}

void ContactForm:: SetSelfInfo(QString name, QString stmsg) {
    uiw.label_2->setText(name);
    // uiw.label_3->setText(stmsg);
    SetQLabelElideText(uiw.label_3,stmsg,"",false);
}

void ContactForm::AddContactItem(QString uid, QString name, QString stmsg) {
    qInfo()<<uid<<name;
    auto ctv = new ContactItem();
    ctv->uid = uid;
    ctv->uiw.label_2->setText(name);
    // SetQLabelElideText(ctv->uiw.label_2,name,"..",true);
    // ctv->uiw.label_3->setText(stmsg);
    // SetQLabelElideText(ctv->uiw.label_3,stmsg,"..",true);
    ctv->uiw.LabelLastMsgTime->clear();
    ctv->uiw.toolButton->setText("0");

    connect(ctv, &ContactItem::clicked, this, &ContactForm::clicked, Qt::QueuedConnection);
    connect(ctv, &ContactItem::reqmenu, this, &ContactForm::showmenu);
    auto lo9 = uiw.verticalLayout_9;
    lo9->insertWidget(0, ctv);
}

void ContactForm::AddConferenceMessage(QString uid, QString msg, QString peername, QString timestr) {
    AddConferenceMessage1(uid,msg, peername, timestr);
}
void ContactForm::AddConferenceMessage1(QString uid, QString msg, QString peername, QString timestr) {
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
        // w->uiw.label_3->setText(msg); // lastmsg
        SetQLabelElideText(w->uiw.label_3,msg,"..",false);
        w->uiw.LabelLastMsgTime->setText(timestr);
        if (uid != curuid) {
            w->addUnread();
        }
        break;
    }
}

QWidget* ContactForm::findContactItem(QString uid) {
    auto lo9 = uiw.verticalLayout_9;
    int cnt = lo9->count();

    for (int i = 0; i < cnt; i++) {
        auto witem = lo9->itemAt(i);
        auto w = (ContactItem*)witem->widget();
        if (w->uid != uid) {
            continue;
        }

        return w;
    }
    return nullptr;
}

void ContactForm::showmenu(QString uid, QWidget* that, const QPoint &pos) {
    qInfo()<<uid<<that<<pos;
}

void ContactForm::setOnline(QString uid, bool on) {
    if (!uid.isEmpty()) {
        auto w = (ContactItem*)findContactItem(uid);
        w->setOnline(on);
    }else{
        if (on) {
            uiw.toolButton->setIcon(QIcon(":/icons/online_30.png"));
        }else{
            uiw.toolButton->setIcon(QIcon(":/icons/offline_2x.png"));
        }
    }
}
