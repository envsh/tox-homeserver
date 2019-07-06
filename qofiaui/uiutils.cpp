#include <QtWidgets>

#include "uiutils.h"

void SetQLabelElideText(QLabel* lab, QString txt, QString suff, bool skipTooltip) {
    auto font = lab->font();
    auto rect = lab->rect();
    auto sz1 = lab->size();
    auto sz2 = lab->sizeHint();
    auto gm = lab->geometry();

    // qInfo()<<rect<<sz1<<sz2;
    int elwidth = qMax(rect.width(), sz2.width());
    elwidth = elwidth > 500 ? rect.width() : elwidth;
    elwidth = rect.width();
    // qInfo()<<elwidth;
    elwidth = elwidth < 150 ? qMax(sz2.width(), elwidth) : elwidth;
    // qInfo()<<elwidth;
    elwidth = elwidth > 150 ? (elwidth-10) : elwidth;
    // qInfo()<<elwidth;
    if (false) qInfo()<<"";
    if (elwidth > 1000) qWarning()<<"something wrong?"<<elwidth;

    auto fm = QFontMetrics(font);
    // elwidth -= elwidth < 150 ? 0 : fm.width(suff);
    auto etxt = fm.elidedText(txt, Qt::ElideRight, elwidth);
    if (false) qInfo()<<txt.length()<<etxt.length()<<elwidth<<fm.width(suff)<<suff.length()<<suff;

    if (etxt.length() < txt.length()) {
        // etxt += suff;
    }
    lab->setText(etxt);
    if (!skipTooltip) {
        lab->setToolTip(txt);
    }
}

#include <malloc.h>
void wdelete(QWidget* w) {
    w->setParent(nullptr); // important for free memory!!!
    delete w;
    int rv = malloc_trim(0);
    // malloc_stats();
}
