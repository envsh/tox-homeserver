#ifndef CHATFORM_H
#define CHATFORM_H

#include <QtWidgets>

#include "ui_chatform.h"

class ChatForm  : public QWidget {
    Q_OBJECT;
 public:
    ChatForm(QWidget*parent = nullptr);
    Ui::ChatForm uiw;
    QString curuid;

    virtual ~ChatForm();
    void dtor();

    void AddConferenceMessage(QString uid, QString msg, QString peername, QString timestr);
    void AddConferenceMessage1(QString uid, QString msg, QString peername, QString timestr);
    void AddConferenceMessage2(QString uid, QString msg, QString peername, QString timestr);

    void setandload(QString uid, QString ctname, QString ctstmsg);
};


#endif /* CHATFORM_H */
