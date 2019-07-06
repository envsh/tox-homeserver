#ifndef MAINFORM_H
#define MAINFORM_H

#include <QtWidgets>

#include "ui_mainform.h"

class MainForm : public QWidget {
    Q_OBJECT;
 public:
    MainForm();
    Ui::MainForm uiw;
    int curuist = 0;

    QWidget* createform(int no);
 public slots:
    QWidget* setform(int no); // return old QWidget
    QWidget* getcurform();
    void onlogin();
    void switchchat(QString uid, QWidget* that);
    void onback();
    void qofiaui_cmdproc(QString cmdmsg);
 signals:
    void cmdhandle(QString cmdmsg);
};

#endif /* MAINFORM_H */
