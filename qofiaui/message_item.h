#ifndef MESSAGE_ITEM_H
#define MESSAGE_ITEM_H

#include "ui_message_item_view.h"

class MessageItem : public QWidget {
 public:
    MessageItem(QWidget* parent = nullptr);

    Ui::MessageItemView uiw;
};

#endif /* MESSAGE_ITEM_H */
