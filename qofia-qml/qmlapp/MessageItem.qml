import QtQuick 2
import QtQuick.Controls 2
import QtQuick.Layouts 1

    RowLayout {
        width: parent.width
        Layout.fillWidth: true
        Layout.fillHeight: false

        ColumnLayout {
                            Button {
                                Layout.maximumWidth: 30
                            text: "ICON1"
            }
            VSpacerItem{}
    }

        ColumnLayout {
            RowLayout {
                Label {
                    text: "name"
            }
            HSpacerItem{}
            Label { text : "time" }
                Button {
                text: "stindi"
                Layout.maximumWidth: 30
                }
        }

            Label { text: "contt"}
    }

    ColumnLayout {
            Button {
            text: "ICON"
            Layout.maximumWidth: 30
            }
        VSpacerItem{}
    }
}
