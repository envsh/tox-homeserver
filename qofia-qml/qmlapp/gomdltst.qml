import QtQuick 2.0
import QtQuick.Window 2.0
import QtQuick.Controls 2.3
import qt.inc 1.0

Item {
    //    title: qsTr("Hello World")
    objectName: "approotwin"
    id: approotwin
    width: 340
    height: 580
    visible: true


    Component {
        id: someComponent
        GoListModel {
            id: lstmdl123
            typeName:"AListModel"
        }
    }

    function createModel(parent) {
        var newModel = someComponent.createObject(parent);
        return newModel;
    }

    Column {
        height : parent.height
        width : parent.width

        Button {text: "dtortst"
            onClicked: {
                var m = lstv123.model
                lstv123.model = createModel(null);
                m.destroy()
            }
        }

        ListView {
            height : parent.height
            width : parent.width
            id : lstv123
            model: GoListModel{
                id: lstmdl123
                typeName:"AListModel"
            }
            //             delegate: MessageItem{}
            delegate: ContactItem{seqno:index; name: model.hue+model.name+model.status}
        }
    }

}
