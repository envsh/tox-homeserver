#ifndef MAINWIN_H
#define MAINWIN_H

#include "ui_untitled.h"


class MainWin : public QMainWindow {
    Q_OBJECT;
 public:
    class contentAreaState {
    public:
        bool isBottom;
        int curpos;
        int maxpos;
    };

    MainWin (QWidget* parent=nullptr);
    void prepui();

    Ui::MainWindow uiw;
    QStack<int> uistks;
    QString curuid;
    QVector<QWidget*> msgviews; // msg view cache
    contentAreaState ccstate;

    void SetSelfInfo(QString name, QString stmsg);
    void AddContactItem(QString uid, QString name, QString stmsg);
    void AddConferenceMessage(QString uid, QString msg);
    void AddConferenceMessage1(QString uid, QString msg);
    void AddConferenceMessage2(QString uid, QString msg);

public slots:
    void qofiaui_cmdproc(QString cmdmsg);
    void switchStackUi(int idx);
    void backStackUi();
    void ctitem_clicked(QString uid, QWidget* that);
signals:
    void cmdhandle(QString cmdmsg);
};

#endif /* MAINWIN_H */
