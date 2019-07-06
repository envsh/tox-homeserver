#ifndef SETTINGFORM_H
#define SETTINGFORM_H

#include <QtWidgets>
#include "ui_settingform.h"

class SettingForm : public QWidget {
    Q_OBJECT;
 public:
    SettingForm(QWidget* parent = nullptr);
    Ui::SettingForm uiw;
};

#endif /* SETTINGFORM_H */
