package main

/*
 */
import "C"
import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"runtime"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtqml"
	"github.com/kitech/qt.go/qtrt"
	// "github.com/kitech/qt.go/qtwidgets"
	"github.com/kitech/qt.go/qtgui"
)

type DataEntryModel struct {
	*qtcore.QAbstractListModel
}

type QmlRuntime struct {
	qmlae   *qtqml.QQmlApplicationEngine
	qmle    *qtqml.QQmlEngine
	rootObj *qtcore.QObject

	objects    map[string]*qtcore.QObject // object cache, objName => obj
	zeroArgVal *qtcore.QGenericArgument
}

func NewQmlRuntimeFrom1(qmlae *qtqml.QQmlApplicationEngine) *QmlRuntime {
	this := &QmlRuntime{}
	this.qmlae = qmlae

	this.zeroArgVal = qtcore.NewQGenericArgument__()
	return this
}

func (this *QmlRuntime) RootObject() *qtcore.QObject {
	qmlae := this.qmlae
	objs := qmlae.RootObjects_1()
	objsx := qtcore.NewQObjectListxFromPointer(objs.GetCthis_())
	if objsx.Count_1() == 0 {
		// return nil
	}
	obj := objsx.At(0)
	// log.Println(objsx.Count_1(), objs.GetCthis_(), obj, obj.GetCthis())
	if obj != nil && obj.GetCthis() != nil {
		return obj
	}
	return nil
}

func (this *QmlRuntime) findObject(objName string) *qtcore.QObject {
	if this.rootObj == nil {
		this.rootObj = this.RootObject()
	}
	if objName == "" {
		return this.rootObj
	}
	if obj, ok := this.objects[objName]; ok {
		return obj
	}

	if this.rootObj == nil {
		log.Println("Maybe qml not loaded")
		return nil
	}

	// find in objects tree
	obj := this.rootObj.FindChild(objName) //("_PageMessages1")
	if obj == nil {
		log.Println("obj not found:", objName, "in", this.rootObj.ObjectName())
	} else {
		this.objects[objName] = obj
		return obj
	}

	return nil
}

// objName=="" then use rootObj
func (this *QmlRuntime) CallJSFunc(objName, funcName string, args ...interface{}) (*qtcore.QVariant, error) {
	var symobj *qtcore.QObject
	symobj = this.findObject(objName)

	return this.CallJSFunc2(symobj, funcName, args...)
}

func (this *QmlRuntime) CallJSFunc2(obj *qtcore.QObject, funcName string, args ...interface{}) (*qtcore.QVariant, error) {
	var symobj *qtcore.QObject
	symobj = obj

	qargc := len(args)
	qargas := this.convArgs(args...)
	for len(qargas) < 10 {
		qargas = append(qargas, this.zeroArgVal)
	}
	_ = qargc

	arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9 :=
		qargas[0], qargas[1], qargas[2], qargas[3], qargas[4],
		qargas[5], qargas[6], qargas[7], qargas[8], qargas[9]

	qretv := qtcore.NewQVariant()
	qreta := qtcore.NewQGenericReturnArgument_fix("QVariant", qretv.GetCthis())
	bok := qtcore.QMetaObject_InvokeMethod(symobj, funcName, qtcore.Qt__AutoConnection, qreta,
		arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	// log.Println(qreta.Name(), qreta.Data())
	// log.Println(qretv.Type(), qretv.TypeName(), qretv.ToInt__())

	// no return mode
	// bok := qtcore.QMetaObject_InvokeMethod_3(symobj, funcName,
	//	arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	// log.Println(bok, obj.ObjectName(), funcName, qargc)

	var err error
	if !bok {
		err = errors.New("meta call failed")
	}
	return qretv, err
}

// go variable to QVariant, then put in QGenericArgument
func (this *QmlRuntime) convArgs(args ...interface{}) (qargas []*qtcore.QGenericArgument) {
	for idx, argx := range args {
		qargv := this.convArg2QVariant(argx, idx)
		if qargv == nil {
			qargv = qtcore.NewQVariant()
		}
		qarga := qtcore.NewQGenericArgument_fix("QVariant", qargv.GetCthis())
		qargas = append(qargas, qarga)
	}
	return
}

// go variable to QVariant
func (this *QmlRuntime) convArg2QVariant(arg interface{}, idx int) *qtcore.QVariant {
	argv := reflect.ValueOf(arg)
	argty := argv.Type()

	var qargv *qtcore.QVariant
	switch argty.Kind() {
	case reflect.Int:
		qargv = qtcore.NewQVariant_5(arg.(int))
	case reflect.String:
		qargv = qtcore.NewQVariant_12(arg.(string))
	default:
		log.Println("Unknown type:", idx, argty.String(), arg)
	}
	return qargv
}

// depth, calc as js file:line info.
func (this *QmlRuntime) EvalJS(jscode string, depth ...int) *qtqml.QJSValue {
	_, file, line, _ := runtime.Caller(1)
	if len(depth) > 0 {
		_, file, line, _ = runtime.Caller(depth[0])
	}
	jsval := this.qmlae.Evaluate(jscode, file, line)
	// log.Println(jsval, jsval.IsNull())
	// log.Println(jsval.ToString(), file, line)
	return jsval
}

func main() {
	log.Println("enter main...")
	qtrt.SetDebugDynSlot(false)
	app := qtgui.NewQGuiApplication(len(os.Args), os.Args, 0)

	lstmdl := qtcore.NewQAbstractListModel__()
	log.Println(lstmdl)

	// qmle := qtqml.NewQQmlEngine__()
	qmlae := qtqml.NewQQmlApplicationEngine__()
	qmlrt := NewQmlRuntimeFrom1(qmlae)
	qmlrt.RootObject()
	qtrt.Connect(qmlae, "objectCreated(QObject *, const QUrl &)", func(obj *qtcore.QObject, uo *qtcore.QUrl) {
		log.Println("hehhee", obj.GetCthis(), uo.GetCthis(), uo.ToLocalFile())
		qtrt.Connect(obj, "testsig123(int)", func(v int) {
			log.Println("hehehhe", v)
		})
		qmlrt.CallJSFunc("", "testb123", "12345s")
		qmlrt.RootObject()
		log.Println(qmlrt.EvalJS("1+2").ToString())

		/*
			obj.DumpObjectInfo()
			obj.DumpObjectTree()

			subobjs := qmlae.RootObjects_1()
			subobjsx := qtcore.NewQObjectListxFromPointer(subobjs.GetCthis_())
			log.Println(subobjsx.Count_1(), subobjs.GetCthis_(), obj.ObjectName(), obj.MetaObject().ClassName())
			winobj := qtquick.NewQQuickWindowFromPointer(obj.GetCthis())
			subitms := qtcore.NewQObjectListxFromPointer(winobj.ContentItem().ChildItems().GetCthis_())
			log.Println(winobj, subitms.Count_1())
			for i := 0; i < subitms.Count_1(); i++ {
				subitm := subitms.At(i)
				log.Println(i, subitm, subitm.GetCthis())
				log.Println(i, subitm, subitm.Property("aaaaaccc"))
				log.Println(i, subitm, subitm.ObjectName())
				log.Println(i, subitm, subitm.MetaObject().ClassName())
			}
			mwobj := obj.FindChild("_PageMessages1")
			log.Println(mwobj, mwobj.GetCthis(), mwobj.ObjectName())
			qtcore.QMetaObject_InvokeMethod_3(mwobj, "msgmdl_add", argx, argx, argx, argx, argx, argx, argx, argx, argx, argx)
		*/
	})
	qmlae.Load_1("./qmlapp/main.qml")
	// qv := qtquick.NewQQuickView__()
	// qv.SetSource(qtcore.NewQUrl_1_("./qmlapp/main.qml"))
	// qv.Show()

	qmle := qmlae.QQmlEngine_PTR()
	log.Println(qmle)
	log.Println(qmle.ObjectName(), qmle.RootContext().ObjectName(), qmlae.ObjectName())
	// qmle.DumpObjectInfo()

	objlst := qmlae.Children()
	log.Println(objlst.Count_1())
	// qmlae.DumpObjectTree()
	ctxobj := qmlae.RootContext().ContextObject()
	log.Println(ctxobj)
	// ctxobj.DumpObjectTree()

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
