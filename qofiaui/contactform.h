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
    void AddConferenceMessage(QString uid, QString msg);
    void AddConferenceMessage1(QString uid, QString msg);

 signals:
    void clicked(QString uid, QWidget* that);
};


#endif /* CONTACTFORM_H */
