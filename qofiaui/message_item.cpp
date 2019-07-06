#include <QtCore>

#include "message_item.h"

MessageItem::MessageItem(QWidget* parent)
    : QWidget(parent) {
    uiw.setupUi(this);
}

MessageItem::~MessageItem() {
    qInfo()<<this<<"dtor";
}

void MessageItem::clear() {
    QString spstr;
    uiw.toolButton_2->setText(spstr);
    uiw.LabelUserName4MessageItem->clear();
    uiw.labelSendState->clear();
    uiw.LabelMsgTime->clear();
    uiw.toolButton->setText(spstr);
    uiw.label_5->clear();
}

