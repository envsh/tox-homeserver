#ifndef MEMBERFORM_H
#define MEMBERFORM_H

#include <QtWidgets>

#include "ui_memberform.h"

class MemberForm : public QWidget
{
    Q_OBJECT;
 public:
    MemberForm(QWidget*parent=nullptr);
    Ui::MemberForm uiw;
};

#endif /* MEMBERFORM_H */
