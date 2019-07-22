#ifndef CONTACTFORM_H
#define CONTACTFORM_H

#include <QtWidgets>

#include "ui_contactform.h"

class ContactForm : public QWidget {
    Q_OBJECT;
 public:
    ContactForm(QWidget*parent = nullptr);
    Ui::ContactForm uiw;
    QString curuid;

    void prepui();

    void SetSelfInfo(QString name, QString stmsg);
    void AddContactItem(QString uid, QString name, QString stmsg);
    void AddConferenceMessage(QString uid, QString msg, QString peername, QString timestr);
    void AddConferenceMessage1(QString uid, QString msg, QString peername, QString timestr);
    void setOnline(QString uid, bool on);
    QWidget* findContactItem(QString uid);

public slots:
    void showmenu(QString uid, QWidget* that, const QPoint &pos);
    void updateTime();

signals:
    void clicked(QString uid, QWidget* that);
};


#endif /* CONTACTFORM_H */
