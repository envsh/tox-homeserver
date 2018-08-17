import QtQuick 2
import QtQuick.Controls 2
import QtQuick.Layouts 1

    RowLayout {
        width: parent.width
        Layout.fillWidth: true
        Layout.fillHeight: false

            ColumnLayout {
                VSpacerItem{}
               Button {
                                Layout.maximumWidth: 30
                            text: "ICON2"
            }
            VSpacerItem{}
    }

        ColumnLayout {
            RowLayout {
                Label {
                    text: "name" + seqno + name
            }
            HSpacerItem{}
            Label { text : "time2" }
        }

            Label { text: "lastmsg"}
            Label { text: "stmsg"}
    }

        ColumnLayout {
            VSpacerItem{}
            Button {
            text: "ICON3"
                Layout.maximumWidth: 30
                Layout.maximumHeight: 30
            }
            Label{ text:"999"}
        VSpacerItem{}
    }

    /////
    property int seqno: 0
    property string name: ""
}
