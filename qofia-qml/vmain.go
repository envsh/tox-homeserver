package main

/*
 */
import "C"
import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtqml"
	"github.com/kitech/qt.go/qtrt"
	// "github.com/kitech/qt.go/qtwidgets"
	"github.com/kitech/qt.go/qtgui"
)

func main() {
	log.Println("enter main...")
	app := qtgui.NewQGuiApplication(len(os.Args), os.Args, 0)

	qmlae := qtqml.NewQQmlApplicationEngine__()
	qtrt.Connect(qmlae, "objectCreated(QObject *, const QUrl &)", func(obj *qtcore.QObject, uo *qtcore.QUrl) {
		log.Println("hehhee", obj.GetCthis(), uo.GetCthis(), uo.ToLocalFile())
		qtrt.Connect(obj, "aaaaa(int)", func(v int) {
			log.Println("hehehhe", v)
		})
	})
	qmlae.Load_1("./qmlapp/main.qml")

	app.Exec()
}

////////
func init() {
	qtqml.RegisterModel("AListModel", NewAListModel)
}

type AListModel struct {
	mdlop *qtcore.QAbstractListModel
}

func NewAListModel(mdlop *qtcore.QAbstractListModel) qtqml.QGoListModel_ITF_RO2 {
	this := &AListModel{mdlop: mdlop}
	return this
}

func (this *AListModel) RowCount() int {
	return rand.Intn(3) + 1
}

func (this *AListModel) Data(index *qtcore.QModelIndex, role int) *qtcore.QVariant {
	log.Println("hehehhe", index.Row(), role)
	retv := qtcore.NewQVariant_15(fmt.Sprintf("retfromgo,row%d-role%d", index.Row(), role))
	return retv
}

func (this *AListModel) RoleNames() map[int]string {
	roles := map[int]string{
		qtcore.Qt__UserRole + 1: "name",
		qtcore.Qt__UserRole + 2: "hue",
		qtcore.Qt__UserRole + 3: "status",
		qtcore.Qt__UserRole + 4: "age",
	}
	return roles
}
