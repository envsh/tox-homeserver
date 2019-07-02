#ifndef QOFIAUI_H
#define QOFIAUI_H

#include <QtCore>

#ifdef __cplusplus
extern "C" {
#endif // __cplusplus
    typedef struct qofiaui_context {
        void (*uion_command)(char* cmd);
    } qofiaui_context ;

    extern qofiaui_context* uictx;

    void qofiaui_main(qofiaui_context* ctx);
    void qofiaui_dmcommand(char* cmdmsg);

    // void uion_login();
#ifdef __cplusplus
};
#endif // __cplusplus

void uion_command(QString cmd);

#endif /* QOFIAUI_H */
