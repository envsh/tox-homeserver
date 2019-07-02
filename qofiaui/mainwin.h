#ifndef MAINWIN_H
#define MAINWIN_H

#include "ui_untitled.h"

class MainWin : public QMainWindow {
    Q_OBJECT;
 public:
    MainWin (QWidget* parent=nullptr);

    Ui::MainWindow uiw;
    void AddContactItem(QString uid, QString name, QString stmsg);
    void AddConferenceMessage(QString uid, QString msg);

public slots:
    void qofiaui_cmdproc(QString cmdmsg);
signals:
    void cmdhandle(QString cmdmsg);
};

#endif /* MAINWIN_H */
