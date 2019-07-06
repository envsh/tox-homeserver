#ifndef UIUTILS_H
#define UIUTILS_H

#include <QLabel>
#include <QString>
#include <QWidget>

void SetQLabelElideText(QLabel* lab, QString txt, QString suff, bool skipTooltip);
void wdelete(QWidget* w);

#endif /* UIUTILS_H */
