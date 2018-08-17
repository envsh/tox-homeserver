import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.0

ColumnLayout {
    Layout.fillWidth: true

    RowLayout{
        Button{
            text:"ICON"
            Layout.maximumWidth: 40
        }
        ColumnLayout{
            RowLayout{
                Label{
                    text: "ctname"
                        Layout.fillWidth: true
                }
                Label{
                    text: "num0"
                }
            }
            Label{
                text: "ctstmsg"
                Layout.fillWidth: true
            }
        }
        ColumnLayout{
            Button{
                text: "mic"
                Layout.maximumWidth: 20
                Layout.maximumHeight: 20
            }
            Button{
                text: "mute"
                Layout.maximumWidth: 20
                Layout.maximumHeight: 20
            }
        }
        Button{
            text: "ACall"
            Layout.maximumWidth: 40
        }
        Button{
            text:"VCall"
            Layout.maximumWidth: 40
        }
        Button{
            text: "Opti"
            Layout.maximumWidth: 40

            Menu {
                id: optimenu
                y: parent.height

                MenuItem{
                    text: "111"
                }
                MenuItem {
                    text: "222"
                }
                MenuItem{
                    text: "333"
                }
            }
        }
    }

    // content
    RowLayout{
        HSpacerItem{}
        Label{
            text: "num1"
        }
        Button{
            text: "load more"
            Layout.maximumWidth: 40
            Layout.maximumHeight: 20
        }
        HSpacerItem{}
    }

    RowLayout{
        Layout.fillHeight: true
        Rectangle{
            Layout.fillHeight: true
            Layout.fillWidth: true
            color: "blue"
        }
    }

    // footer
    RowLayout{
        Button{
            text:"Attac"
            Layout.maximumWidth: 40
        }
        Button{
            text: "Snap"
            Layout.maximumWidth: 40
        }
        TextEdit{
            text: "aaa"
            selectByMouse: true
            Layout.fillWidth: true
        }
        Button{
            text:"emoji"
            Layout.maximumWidth: 40
        }
        Button{
            text: "send"
            Layout.maximumWidth: 40
        }
    }

}
