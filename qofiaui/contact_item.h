#ifndef CONTACT_ITEM_H
#define CONTACT_ITEM_H

#include "ui_contact_item_view.h"

class ContactItem : public QWidget {
 public:
    ContactItem(QWidget* parent = nullptr);

    Ui::ContactItemView uiw;
};

#endif /* CONTACT_ITEM_H */
