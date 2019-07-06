#ifndef CHATFORM_H
#define CHATFORM_H

#include <QtWidgets>

#include "ui_chatform.h"

class ChatForm  : public QWidget {
    Q_OBJECT;
 public:
    ChatForm(QWidget*parent = nullptr);
    Ui::ChatForm uiw;
};


#endif /* CHATFORM_H */
