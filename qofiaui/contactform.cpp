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
    SetQLabelElideText(ctv->uiw.label_3,stmsg,"..",true);

    connect(ctv, &ContactItem::clicked, this, &ContactForm::clicked, Qt::QueuedConnection);
    auto lo9 = uiw.verticalLayout_9;
    lo9->insertWidget(0, ctv);
}

void ContactForm::AddConferenceMessage(QString uid, QString msg) {
    AddConferenceMessage1(uid,msg);
}
void ContactForm::AddConferenceMessage1(QString uid, QString msg) {
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
}
