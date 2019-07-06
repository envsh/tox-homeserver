#include "mainform.h"

#include "qofiaui.h"
#include "settingform.h"
#include "loginform.h"
#include "contactform.h"
#include "chatform.h"
#include "memberform.h"

MainForm::MainForm()
    : QWidget(nullptr) {
    uiw.setupUi(this);
}

static QWidget* createform(int no) {
    QWidget* w = nullptr;
    switch (no) {
    case UIST_SETTINGS:
        w = new SettingForm();
        break;
    case UIST_LOGINUI:
        w = new LoginForm();
        break;
    case UIST_CONTACTUI:
        w = new ContactForm();
        break;
    case UIST_MESSAGEUI:
        w = new ChatForm();
        break;
    default:
        assert(1==2);
    };
    return w;
}

QWidget* MainForm::setform(int no) {
    auto mlo = uiw.verticalLayout;
    int cnt = mlo->count();
    auto form = createform(no);
    QWidget* oldw = nullptr;
    if (cnt == 2) {
        auto olditem = mlo->takeAt(1);
        oldw = olditem->widget();
    }
    mlo->addWidget(form);
    return oldw;
}

