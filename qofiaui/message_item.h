#ifndef MESSAGE_ITEM_H
#define MESSAGE_ITEM_H

#include "ui_message_item_view.h"

class MessageItem : public QWidget {
    Q_OBJECT;
 public:
    MessageItem(QWidget* parent = nullptr);
    virtual ~MessageItem();
    Ui::MessageItemView uiw;

    void clear();
};

#endif /* MESSAGE_ITEM_H */
