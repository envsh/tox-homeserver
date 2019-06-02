import strutils
import re

# check is bot user
proc IsBotUser*(senderc : string) : bool =
    var botUsers = ["teleboto", "Orizon", "OrzGTalk", "OrzIrc2P",
                                     "xmppbot", "tg2offtopic", "tg2arch", "toxync"]
    for u in botUsers:
        if u == senderc or u == strutils.strip(senderc, leading=false, chars={'_', '^'}):
            return true
    return false

# one cycle extract
proc ExtractRealUser*(sender, message : string) : (string, string, string) =
    var new_sender = sender
    var new_message = message
    var color = ""

    var msgregs = [
        """^<FONT COLOR="([\w ]+)">\[([^\[]+)\] </FONT>""", # // teleboto
        """^\[<FONT COLOR="([\w ]+)">([^\[]+)</FONT>\]""", # // Orizon
        """^(.?[0-9]+)\[([^\[]+)\] """, # // teleboto? with 1st unprintable char
        """^()\(GTalk\) ([^\[]+):""", # // OrzGTalk
        """^()\[(.+\[m\].+)\] """, # // riot.im
        """^()\[([^\[]+)\] """, # // OrzIrc2P/tg2offtopic
    ]
    let sfxreg = "(.+)"

    for idx, reg in msgregs:
        var mats : array[5,string]
        var mok = re.match(message, re.re(reg & sfxreg), mats)
        if not mok: continue
        # new_sender = mats[0][0] // with color style
        new_sender = mats[1]
        new_message = mats[2]
        color = mats[0]
        break

    return (new_sender, new_message, color)

# multiple cycle extract
# color: 名字的颜色
# multiple depth
proc ExtractRealUserMD*(sender, message : string) : (string, string, string) =
    var (new_sender, new_message, color) = ExtractRealUser(sender, message)
    if new_sender == sender and new_message == message:
        return (new_sender, new_message, color)
    else:
        var (new_sender2, new_message2, color2) = ExtractRealUserMD(new_sender, new_message)
        return (new_sender2, new_message2, if color2 == "": color else: color2)

proc testExtractRealUser0() =
    echo IsBotUser("teleboto")

    # 带FONT的，可能已经是pidgin转换过了的。
    var msgs = [
        """<FONT COLOR="teal laet">[ngkaho1234] </FONT>後來修復了""", #// teleboto
        """[<FONT COLOR="pink knip">dant mnf</FONT>] 没有""",        #// Orizon
        """(GTalk) niconiconi: 无误""",                              #// OrzGTalk
        """[FsckGoF] 群主女装吼不吼啊？""",                             #// OrzIrc2P
        """[Lisa] \h: Lisa is here :-)""",                           #// xmppbot
        """7[Miyamizu_Mitsuha] 厉害了""",                             #// teleboto?
        """4[Abel_Abel] 这算不算父进程？""",                            #//??
        """13[Universebenzene] 缺poppler-data？""",                  #//??
        """6[KireinaHoro_] 15「Re farseerfc: wow 麗狼加油...」謝謝""",  #//??
        """[tg2offtopic@irc] [FQEgg] loop""",
        """[erhandsoME[m]@irc] 。。。""",                              #// riot.im
    ]

    for idx, m in msgs:
        var resvals = ExtractRealUser("", m)
        echo idx, " ", resvals
        if true:
            var resvals = ExtractRealUserMD("", m)
            echo idx, " ", resvals
    return

proc ExtractUrls(txt: string) : seq[string] =
    return re.findAll(txt, re(re.reURL))

if isMainModule and false:
    testExtractRealUser0()
