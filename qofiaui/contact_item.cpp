#include <QtGui>

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
        auto mevt = (QMouseEvent*)event;
        if (mevt->buttons() & Qt::LeftButton) {
            prtime = QDateTime::currentDateTime();
        }else if (mevt->buttons() & Qt::RightButton) {
            prtime = QDateTime::currentDateTime();
        }else {
            qInfo()<<"what"<<mevt->buttons();
        }
    }else if (event->type() == QEvent::MouseButtonRelease) {
        auto mevt = (QMouseEvent*)event;
        auto nowt = QDateTime::currentDateTime();
        if (mevt->button() == Qt::LeftButton) {
            if (prtime.msecsTo(nowt) < 300) {
                emit clicked(uid, this);
            }
        }else if (mevt->button() == Qt::RightButton) {
            if (prtime.msecsTo(nowt) < 300) {
                emit reqmenu(uid, this, mevt->pos());
            }
        }else{
            qInfo()<<"what"<<mevt->buttons()<<mevt->button();
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


