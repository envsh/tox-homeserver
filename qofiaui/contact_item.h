#ifndef CONTACT_ITEM_H
#define CONTACT_ITEM_H

#include <QtCore>

#include "ui_contact_item_view.h"

class ContactItem : public QWidget {
    Q_OBJECT;
 public:
    ContactItem(QWidget* parent = nullptr);

    Ui::ContactItemView uiw;
    QString uid;
    QDateTime prtime;
    bool selected;

 signals:
    void clicked(QString uid, QWidget* that);
    void reqmenu(QString uid, QWidget* that, const QPoint &pos);

 protected:
    bool event(QEvent *event);
    bool eventFilter(QObject *object, QEvent *event);

    /* void mousePressEvent(QMouseEvent *event) override; */
    /* void mouseMoveEvent(QMouseEvent *event) override; */
    /* void mouseReleaseEvent(QMouseEvent *event) override; */
    /* // void paintEvent(QPaintEvent *event) override; */
    /* void resizeEvent(QResizeEvent *event) override; */

};

#endif /* CONTACT_ITEM_H */
