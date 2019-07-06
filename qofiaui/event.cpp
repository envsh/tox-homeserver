#include "event.h"

Event Event::fromJson(QJsonDocument jdoc) {
    Event evto;
    QJsonObject jobj = jdoc.object();
    QString evtname = jobj.value("EventName").toString();
    evto.EventName = evtname;

    auto jarr = jobj.value("Args").toArray();
    auto marr = jobj.value("Margs").toArray();
    for (int i = 0; i < jarr.count(); i++) {
        evto.Args.append(jarr.at(i).toString());
    }
    for (int i = 0; i < marr.count(); i++) {
        evto.Margs.append(marr.at(i).toString());
    }

    return evto;
}
