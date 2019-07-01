#ifndef MAINWIN_H
#define MAINWIN_H

#include "ui_untitled.h"

class MainWin : public QMainWindow {
 public:
    MainWin (QWidget* parent=nullptr);

    Ui::MainWindow uiw;
};

#endif /* MAINWIN_H */
