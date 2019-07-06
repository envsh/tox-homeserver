#ifndef QOFIAUI_H
#define QOFIAUI_H

#include <QtCore>

#ifdef __cplusplus
extern "C" {
#endif // __cplusplus
    typedef struct qofiaui_context {
        void (*uion_command)(char* cmd);
        char* (*uion_loadmsg)(char* uid, int maxcnt);
    } qofiaui_context ;

    extern qofiaui_context* uictx;

    void qofiaui_main(qofiaui_context* ctx);
    void qofiaui_dmcommand(char* cmdmsg);

    // void uion_login();
#ifdef __cplusplus
};
#endif // __cplusplus

void uion_command(QString cmd);

enum {
      UIST_QMLMCTRL = 0                   ,
      UIST_QMLORIGIN                      ,
      UIST_SETTINGS                       ,
      UIST_LOGINUI                        ,
      UIST_CONTACTUI                      ,
      UIST_MESSAGEUI                      ,
      UIST_VIDEOUI                        ,
      // UIST_PICKCALLUI // TODO video 
      UIST_ADD_GROUP                      ,
      UIST_ADD_FRIEND                     ,
      UIST_INVITE_FRIEND                  ,
      UIST_MEMBERS                        ,
      UIST_CONTACT_INFO                   ,
      UIST_TESTUI                         ,
      UIST_LOGUI                          ,
      UIST_ABOUTUI                        ,
};

#endif /* QOFIAUI_H */
