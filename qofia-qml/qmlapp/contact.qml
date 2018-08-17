import QtQuick 2.1
import StratifyLabs.UI 3.0


Rectangle {
    id: root
    // color: "green"
    width: 200
    height: 300



SContainer {
  name: "Layouts";
  style: "fill";
  SPane {
    style: "block fill";
    SColumn {
      SText { style: "left text-bold"; text: "Introduction"; }
      SText { style: "left text-bold"; text: "Container"; }

      SText { style: "left text-bold"; text: "Row"; }
      SText { style: "left text-bold"; text: "Column"; }
    }

  }
}
}

