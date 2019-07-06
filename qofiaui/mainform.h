#ifndef MAINFORM_H
#define MAINFORM_H

#include <QtWidgets>

#include "ui_mainform.h"

class MainForm : public QWidget {
    Q_OBJECT;
 public:
    MainForm();
    Ui::MainForm uiw;

 public slots:
     QWidget* setform(int no); // return old QWidget
};

#endif /* MAINFORM_H */
