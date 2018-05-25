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
func NewUi_AddFriendDialog() *Ui_AddFriendDialog {
	return &Ui_AddFriendDialog{}
}

type Ui_AddFriendDialog struct {
	VerticalLayout_2 *qtwidgets.QVBoxLayout
	VerticalLayout   *qtwidgets.QVBoxLayout
	Label            *qtwidgets.QLabel
	LineEdit         *qtwidgets.QLineEdit
	Label_2          *qtwidgets.QLabel
	TextEdit         *qtwidgets.QTextEdit
	HorizontalLayout *qtwidgets.QHBoxLayout
	HorizontalSpacer *qtwidgets.QSpacerItem
	ButtonBox        *qtwidgets.QDialogButtonBox
	AddFriendDialog  *qtwidgets.QDialog
}

//  struct block end

//  setupUi block begin
// void setupUi(QDialog *AddFriendDialog)
func (this *Ui_AddFriendDialog) SetupUi(AddFriendDialog *qtwidgets.QDialog) {
	this.AddFriendDialog = AddFriendDialog
	// { // 126
	if AddFriendDialog.ObjectName() == "" {
		AddFriendDialog.SetObjectName("AddFriendDialog")
	}
	AddFriendDialog.Resize(606, 246)
	this.VerticalLayout_2 = qtwidgets.NewQVBoxLayout_1(this.AddFriendDialog)                                 // 111
	this.VerticalLayout_2.SetObjectName("VerticalLayout_2")                                                  // 112
	this.VerticalLayout = qtwidgets.NewQVBoxLayout()                                                         // 111
	this.VerticalLayout.SetObjectName("VerticalLayout")                                                      // 112
	this.Label = qtwidgets.NewQLabel(this.AddFriendDialog, 0)                                                // 111
	this.Label.SetObjectName("Label")                                                                        // 112
	this.Label.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.VerticalLayout.Layout().AddWidget(this.Label) // 115

	this.LineEdit = qtwidgets.NewQLineEdit(this.AddFriendDialog) // 111
	this.LineEdit.SetObjectName("LineEdit")                      // 112

	this.VerticalLayout.Layout().AddWidget(this.LineEdit) // 115

	this.Label_2 = qtwidgets.NewQLabel(this.AddFriendDialog, 0) // 111
	this.Label_2.SetObjectName("Label_2")                       // 112

	this.VerticalLayout.Layout().AddWidget(this.Label_2) // 115

	this.TextEdit = qtwidgets.NewQTextEdit(this.AddFriendDialog) // 111
	this.TextEdit.SetObjectName("TextEdit")                      // 112
	this.TextEdit.SetAcceptRichText(false)                       // 114

	this.VerticalLayout.Layout().AddWidget(this.TextEdit) // 115

	this.VerticalLayout_2.AddLayout(this.VerticalLayout, 0) // 115

	this.HorizontalLayout = qtwidgets.NewQHBoxLayout()      // 111
	this.HorizontalLayout.SetObjectName("HorizontalLayout") // 112
	this.HorizontalSpacer = qtwidgets.NewQSpacerItem(40, 20, qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Minimum)

	this.HorizontalLayout.AddItem(this.HorizontalSpacer) // 115

	this.ButtonBox = qtwidgets.NewQDialogButtonBox(this.AddFriendDialog)                                   // 111
	this.ButtonBox.SetObjectName("ButtonBox")                                                              // 112
	this.ButtonBox.SetOrientation(qtcore.Qt__Horizontal)                                                   // 114
	this.ButtonBox.SetStandardButtons(qtwidgets.QDialogButtonBox__Cancel | qtwidgets.QDialogButtonBox__Ok) // 114

	this.HorizontalLayout.Layout().AddWidget(this.ButtonBox) // 115

	this.VerticalLayout_2.AddLayout(this.HorizontalLayout, 0) // 115

	this.RetranslateUi(AddFriendDialog)
	// QObject::connect(buttonBox, SIGNAL(accepted()), AddFriendDialog, SLOT(accept())); // 126
	// QObject::connect(buttonBox, SIGNAL(rejected()), AddFriendDialog, SLOT(reject())); // 126

	qtcore.QMetaObject_ConnectSlotsByName(AddFriendDialog) // 100111
	// } // setupUi // 126

}

// void retranslateUi(QDialog *AddFriendDialog)
//  setupUi block end

//  retranslateUi block begin
func (this *Ui_AddFriendDialog) RetranslateUi(AddFriendDialog *qtwidgets.QDialog) {
	// noimpl: {
	this.AddFriendDialog.SetWindowTitle(qtcore.QCoreApplication_Translate("AddFriendDialog", "Dialog", "dummy123", 0))
	this.Label.SetText(qtcore.QCoreApplication_Translate("AddFriendDialog", "Tox ID (either 76 hexadecimal characters or name@example.com )", "dummy123", 0))
	this.Label_2.SetText(qtcore.QCoreApplication_Translate("AddFriendDialog", "Message", "dummy123", 0))
	this.TextEdit.SetPlaceholderText(qtcore.QCoreApplication_Translate("AddFriendDialog", "I am %1 ! Let's chat with Tox?", "dummy123", 0))
	// noimpl: } // retranslateUi
	// noimpl:
	// noimpl: };
	// noimpl:
}

//  retranslateUi block end

//  new2 block begin
func NewUi_AddFriendDialog2() *Ui_AddFriendDialog {
	this := &Ui_AddFriendDialog{}
	w := qtwidgets.NewQDialog(nil, 0)
	this.SetupUi(w)
	return this
}

//  new2 block end

//  done block begin

func (this *Ui_AddFriendDialog) QWidget_PTR() *qtwidgets.QWidget {
	return this.AddFriendDialog.QWidget_PTR()
}

//  done block end
