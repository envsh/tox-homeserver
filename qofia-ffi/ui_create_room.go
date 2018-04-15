package main

//  header block begin
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"
import "github.com/kitech/qt.go/qtquickwidgets"
import "github.com/kitech/qt.go/qtmock"

func init() { qtcore.KeepMe() }
func init() { qtgui.KeepMe() }
func init() { qtwidgets.KeepMe() }
func init() { qtquickwidgets.KeepMe() }
func init() { qtmock.KeepMe() }

//  header block end

//  struct block begin
func NewUi_Dialog() *Ui_Dialog {
	return &Ui_Dialog{}
}

type Ui_Dialog struct {
	VerticalLayout_2   *qtwidgets.QVBoxLayout
	VerticalLayout     *qtwidgets.QVBoxLayout
	Label              *qtwidgets.QLabel
	LineEdit           *qtwidgets.QLineEdit
	HorizontalLayout_2 *qtwidgets.QHBoxLayout
	Label_2            *qtwidgets.QLabel
	RadioButton        *qtwidgets.QRadioButton
	RadioButton_2      *qtwidgets.QRadioButton
	ComboBox           *qtwidgets.QComboBox
	HorizontalLayout   *qtwidgets.QHBoxLayout
	HorizontalSpacer   *qtwidgets.QSpacerItem
	ButtonBox          *qtwidgets.QDialogButtonBox
	Dialog             *qtwidgets.QDialog
}

//  struct block end

//  setupUi block begin
// void setupUi(QDialog *Dialog)
func (this *Ui_Dialog) SetupUi(Dialog *qtwidgets.QDialog) {
	this.Dialog = Dialog
	// { // 126
	if Dialog.ObjectName() == "" {
		Dialog.SetObjectName("Dialog")
	}
	Dialog.Resize(400, 190)
	this.VerticalLayout_2 = qtwidgets.NewQVBoxLayout_1(this.Dialog) // 111
	this.VerticalLayout_2.SetObjectName("VerticalLayout_2")         // 112
	this.VerticalLayout = qtwidgets.NewQVBoxLayout()                // 111
	this.VerticalLayout.SetSpacing(16)                              // 114
	this.VerticalLayout.SetObjectName("VerticalLayout")             // 112
	this.VerticalLayout.SetContentsMargins(20, -1, 20, -1)          // 114
	this.Label = qtwidgets.NewQLabel(this.Dialog, 0)                // 111
	this.Label.SetObjectName("Label")                               // 112

	this.VerticalLayout.Layout().AddWidget(this.Label) // 115

	this.LineEdit = qtwidgets.NewQLineEdit(this.Dialog) // 111
	this.LineEdit.SetObjectName("LineEdit")             // 112

	this.VerticalLayout.Layout().AddWidget(this.LineEdit) // 115

	this.HorizontalLayout_2 = qtwidgets.NewQHBoxLayout()        // 111
	this.HorizontalLayout_2.SetObjectName("HorizontalLayout_2") // 112
	this.Label_2 = qtwidgets.NewQLabel(this.Dialog, 0)          // 111
	this.Label_2.SetObjectName("Label_2")                       // 112

	this.HorizontalLayout_2.Layout().AddWidget(this.Label_2) // 115

	this.RadioButton = qtwidgets.NewQRadioButton(this.Dialog) // 111
	this.RadioButton.SetObjectName("RadioButton")             // 112
	this.RadioButton.SetChecked(true)                         // 114

	this.HorizontalLayout_2.Layout().AddWidget(this.RadioButton) // 115

	this.RadioButton_2 = qtwidgets.NewQRadioButton(this.Dialog) // 111
	this.RadioButton_2.SetObjectName("RadioButton_2")           // 112

	this.HorizontalLayout_2.Layout().AddWidget(this.RadioButton_2) // 115

	this.ComboBox = qtwidgets.NewQComboBox(this.Dialog)     // 111
	this.ComboBox.AddItem("", qtcore.NewQVariant_12("wtf")) // 115
	this.ComboBox.AddItem("", qtcore.NewQVariant_12("wtf")) // 115
	this.ComboBox.SetObjectName("ComboBox")                 // 112
	this.ComboBox.SetFrame(false)                           // 114

	this.HorizontalLayout_2.Layout().AddWidget(this.ComboBox) // 115

	this.VerticalLayout.AddLayout(this.HorizontalLayout_2, 0) // 115

	this.VerticalLayout_2.AddLayout(this.VerticalLayout, 0) // 115

	this.HorizontalLayout = qtwidgets.NewQHBoxLayout()      // 111
	this.HorizontalLayout.SetObjectName("HorizontalLayout") // 112
	this.HorizontalSpacer = qtwidgets.NewQSpacerItem(40, 20, qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Minimum)

	this.HorizontalLayout.AddItem(this.HorizontalSpacer) // 115

	this.ButtonBox = qtwidgets.NewQDialogButtonBox(this.Dialog)                                            // 111
	this.ButtonBox.SetObjectName("ButtonBox")                                                              // 112
	this.ButtonBox.SetOrientation(qtcore.Qt__Horizontal)                                                   // 114
	this.ButtonBox.SetStandardButtons(qtwidgets.QDialogButtonBox__Cancel | qtwidgets.QDialogButtonBox__Ok) // 114

	this.HorizontalLayout.Layout().AddWidget(this.ButtonBox) // 115

	this.VerticalLayout_2.AddLayout(this.HorizontalLayout, 0) // 115

	this.RetranslateUi(Dialog)
	// QObject::connect(buttonBox, SIGNAL(accepted()), Dialog, SLOT(accept())); // 126
	// QObject::connect(buttonBox, SIGNAL(rejected()), Dialog, SLOT(reject())); // 126

	qtcore.QMetaObject_ConnectSlotsByName(Dialog) // 100111
	// } // setupUi // 126

}

// void retranslateUi(QDialog *Dialog)
//  setupUi block end

//  retranslateUi block begin
func (this *Ui_Dialog) RetranslateUi(Dialog *qtwidgets.QDialog) {
	// noimpl: {
	this.Dialog.SetWindowTitle(qtcore.QCoreApplication_Translate("Dialog", "Dialog", "dummy123", 0))
	this.Label.SetText(qtcore.QCoreApplication_Translate("Dialog", "Room Name:", "dummy123", 0))
	this.Label_2.SetText(qtcore.QCoreApplication_Translate("Dialog", "Room Type:", "dummy123", 0))
	this.RadioButton.SetText(qtcore.QCoreApplication_Translate("Dialog", "&Text", "dummy123", 0))
	this.RadioButton_2.SetText(qtcore.QCoreApplication_Translate("Dialog", "&Audio", "dummy123", 0))
	this.ComboBox.SetItemText(0, qtcore.QCoreApplication_Translate("Dialog", "Text", "dummy123", 0))
	this.ComboBox.SetItemText(1, qtcore.QCoreApplication_Translate("Dialog", "Audio", "dummy123", 0))
	// noimpl:
	// noimpl: } // retranslateUi
	// noimpl:
	// noimpl: };
	// noimpl:
}

//  retranslateUi block end

//  new2 block begin
func NewUi_Dialog2() *Ui_Dialog {
	this := &Ui_Dialog{}
	w := qtwidgets.NewQDialog(nil, 0)
	this.SetupUi(w)
	return this
}

//  new2 block end

//  done block begin

func (this *Ui_Dialog) QWidget_PTR() *qtwidgets.QWidget {
	return this.Dialog.QWidget_PTR()
}

//  done block end
