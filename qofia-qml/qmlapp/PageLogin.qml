import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.0

ColumnLayout {
    Layout.fillWidth: true

    VSpacerItem{}

    RowLayout{

        Label{
            text: "Home server URL:"
        }
        HSpacerItem{}

        Label{
            text: "..."
        }
    }

    ComboBox{
        Layout.fillWidth: true
        model: ["toxhs.fixlan.tk:2080"]
        editable: true

        contentItem: TextField {
            text: parent.editText
            selectByMouse: true
        }
    }

    RowLayout {
        HSpacerItem{}
        RadioButton{
            text: "Remote Server"
            checked: true
        }
        HSpacerItem{}
        RadioButton{
            text: "Embed Server"
            enabled: false
        }
        HSpacerItem{}
    }

    Button{
        Layout.fillWidth: true
        text: "Sign in"
    }

    Label{
        Layout.fillWidth: true
        text: "..."
    }

    VSpacerItem{}
}
