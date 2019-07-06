#ifndef CONTACTFORM_H
#define CONTACTFORM_H

#include <QtWidgets>

#include "ui_contactform.h"

class ContactForm : public QWidget {
    Q_OBJECT;
 public:
    ContactForm(QWidget*parent = nullptr);
    Ui::ContactForm uiw;
};


#endif /* CONTACTFORM_H */
