#include "contact_item.h"

ContactItem::ContactItem(QWidget* parent)
    : QWidget(parent) {
    uiw.setupUi(this);

    uiw.label_4->clear();

    uiw.toolButton_2->installEventFilter(this);
    uiw.label_2->installEventFilter(this);
    uiw.LabelLastMsgTime->installEventFilter(this);
    uiw.label_3->installEventFilter(this);
    uiw.label_4->installEventFilter(this);
    uiw.toolButton->installEventFilter(this);
}

bool ContactItem::event(QEvent *event) {
    return QWidget::event(event);
}

bool ContactItem::eventFilter(QObject *object, QEvent *event) {
    if (event->type() == QEvent::MouseButtonPress) {
        prtime = QDateTime::currentDateTime();
    }else if (event->type() == QEvent::MouseButtonRelease) {
        auto nowt = QDateTime::currentDateTime();
        if (prtime.msecsTo(nowt) < 300) {
            emit clicked(uid, this);
        }
    }
    return false;
}

// void ContactItem::mousePressEvent(QMouseEvent *event) {
//     prtime = QDateTime::currentDateTime();
//     QWidget::mousePressEvent(event);
// }
// void ContactItem::mouseMoveEvent(QMouseEvent *event)  {
//     QWidget::mouseMoveEvent(event);
// }
// void ContactItem::mouseReleaseEvent(QMouseEvent *event)  {
//     auto nowt = QDateTime::currentDateTime();
//     if (prtime.msecsTo(nowt) > 300) {
//         qInfo()<<"clicked"<<uid;
//     }
//     QWidget::mouseReleaseEvent(event);
// }
// // void ContactItem::paintEvent(QPaintEvent *event) override {
// // }
// void ContactItem::resizeEvent(QResizeEvent *event)  {
//     QWidget::resizeEvent(event);
// }


