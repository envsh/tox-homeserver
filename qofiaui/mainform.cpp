#include "mainform.h"

#include "event.h"
#include "qofiaui.h"
#include "uiutils.h"
#include "settingform.h"
#include "loginform.h"
#include "contact_item.h"
#include "contactform.h"
#include "chatform.h"
#include "memberform.h"

MainForm::MainForm()
    : QWidget(nullptr) {
    uiw.setupUi(this);

    connect(uiw.toolButton_33, &QToolButton::clicked, this, &MainForm::onback);
}

static ContactForm* ctform = nullptr;

QWidget* MainForm::createform(int no) {
    QWidget* w = nullptr;
    SettingForm* setfrm = nullptr;
    LoginForm* loginfrm = nullptr;
    switch (no) {
    case UIST_SETTINGS:
        w = new SettingForm();
        break;
    case UIST_LOGINUI:
        loginfrm = new LoginForm();
        w = loginfrm;
        connect(loginfrm->uiw.pushButton_7, &QPushButton::clicked, this, &MainForm::onlogin);
        break;
    case UIST_CONTACTUI:
        if (ctform == nullptr) {
            ctform = new ContactForm();
            connect(ctform, &ContactForm::clicked, this, &MainForm::switchchat);
        }
        w = ctform;
        break;
    case UIST_MESSAGEUI:
        w = new ChatForm();
        break;
    default:
        assert(1==2);
        break;
    };
    return w;
}

QWidget* MainForm::setform(int no) {
    curuist = no;
    auto mlo = uiw.verticalLayout;
    int cnt = mlo->count();
    auto form = createform(no);
    QWidget* oldw = nullptr;
    if (cnt == 2) {
        auto olditem = mlo->takeAt(1);
        oldw = olditem->widget();
        mlo->removeWidget(oldw);
        oldw->setVisible(false);
        delete olditem;
        malloc_trim(0);
    }
    mlo->addWidget(form);
    form->setVisible(true);
    return oldw;
}
QWidget* MainForm::getcurform() {
    auto mlo = uiw.verticalLayout;
    int cnt = mlo->count();
    assert(cnt == 2);

    auto olditem = mlo->itemAt(1);
    auto oldw = olditem->widget();
    return oldw;
}

void MainForm::onlogin() {
    qInfo()<<"logining...";
    auto w = (LoginForm*)setform(UIST_CONTACTUI);

    QString cmd = "login/";
    cmd += w->uiw.comboBox_6->currentText();
    uion_command(cmd);

    wdelete(w);
}

void MainForm::switchchat(QString uid, QWidget* that) {
    auto cti = (ContactItem*)that;
    setform(UIST_MESSAGEUI);
    QString ctname = cti->uiw.label_2->text();
    QString ctstmsg = cti->uiw.label_3->text();

    auto msgform = (ChatForm*)getcurform();
    msgform->setandload(uid, ctname, ctstmsg);
    malloc_trim(0);
}

void MainForm::onback() {
    if (curuist == UIST_CONTACTUI) {
        return;
    }
    auto oldform = setform(UIST_CONTACTUI);
    wdelete (oldform);
}

void MainForm::qofiaui_cmdproc(QString cmdmsg) {
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
        ctform->SetSelfInfo(jarr.at(1).toString(), jarr.at(2).toString());
    }else if (evtname == "AddFriendItem") {
        ctform->AddContactItem(jarr.at(1).toString(), jarr.at(2).toString(), jarr.at(3).toString());
    }else if (evtname == "AddGroupItem") {
        ctform->AddContactItem(jarr.at(1).toString(), jarr.at(2).toString(), jarr.at(3).toString());
    }else if (evtname == "ConferenceMessage") {
        ctform->AddConferenceMessage(marr.at(3).toString(), jarr.at(3).toString());
        if (curuist == UIST_MESSAGEUI) {
            auto msgform = (ChatForm*)getcurform();
            msgform->AddConferenceMessage(marr.at(3).toString(), jarr.at(3).toString());
        }
    }else {
        qInfo()<<"todo"<<evtname;
    }
    malloc_trim(0);
}
