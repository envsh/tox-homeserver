#ifndef CONTACT_ITEM_H
#define CONTACT_ITEM_H

#include "ui_contact_item_view.h"

class ContactItem : public QWidget {
    Q_OBJECT;
 public:
    ContactItem(QWidget* parent = nullptr);

    Ui::ContactItemView uiw;
    QString uid;
};

#endif /* CONTACT_ITEM_H */
