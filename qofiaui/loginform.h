#ifndef LOGINFORM_H
#define LOGINFORM_H

#include <QtWidgets>

#include "ui_loginform.h"

class LoginForm: public QWidget
{
    Q_OBJECT;
 public:
    LoginForm(QWidget* parent = nullptr);
    Ui::LoginForm uiw;
};


#endif /* LOGINFORM_H */
