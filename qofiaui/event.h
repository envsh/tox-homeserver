#ifndef EVENT_H
#define EVENT_H

#include <QtCore>

class Event {
 public:
    QString EventName;
    QStringList Args;
    QStringList Margs;

    static Event fromJson(QJsonDocument jdoc);
};



#endif /* EVENT_H */
