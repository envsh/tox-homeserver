package main

//  header block begin
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"
import "github.com/kitech/qt.go/qtquickwidgets"
import "github.com/kitech/qt.go/qtmock"

func init() { qtcore.KeepMe() }
func init() { qtwidgets.KeepMe() }
func init() { qtquickwidgets.KeepMe() }
func init() { qtmock.KeepMe() }

//  header block end

//  struct block begin
func NewUi_ContactItemView() *Ui_ContactItemView {
	return &Ui_ContactItemView{}
}

type Ui_ContactItemView struct {
	HorizontalLayout *qtwidgets.QHBoxLayout
	Label            *qtwidgets.QLabel
	VerticalLayout   *qtwidgets.QVBoxLayout
	Label_2          *qtwidgets.QLabel
	Label_3          *qtwidgets.QLabel
	Label_4          *qtwidgets.QLabel
	VerticalLayout_2 *qtwidgets.QVBoxLayout
	VerticalSpacer   *qtwidgets.QSpacerItem
	ToolButton       *qtwidgets.QToolButton
	VerticalSpacer_2 *qtwidgets.QSpacerItem
	ContactItemView  *qtwidgets.QWidget
	SizePolicy       *qtwidgets.QSizePolicy
	Icon             *qtgui.QIcon // 116
}

//  struct block end

//  setupUi block begin
// void setupUi(QWidget *ContactItemView)
func (this *Ui_ContactItemView) SetupUi(ContactItemView *qtwidgets.QWidget) {
	this.ContactItemView = ContactItemView
	// { // 126
	if ContactItemView.ObjectName() == "" {
		ContactItemView.SetObjectName("ContactItemView")
	}
	ContactItemView.Resize(252, 99)
	this.HorizontalLayout = qtwidgets.NewQHBoxLayout_1(this.ContactItemView) // 111
	this.HorizontalLayout.SetObjectName("HorizontalLayout")                  // 112
	this.Label = qtwidgets.NewQLabel(this.ContactItemView, 0)                // 111
	this.Label.SetObjectName("Label")                                        // 112
	this.Label.SetMaximumSize_1(30, 0)                                       // 113

	this.HorizontalLayout.Layout().AddWidget(this.Label) // 115

	this.VerticalLayout = qtwidgets.NewQVBoxLayout()            // 111
	this.VerticalLayout.SetObjectName("VerticalLayout")         // 112
	this.Label_2 = qtwidgets.NewQLabel(this.ContactItemView, 0) // 111
	this.Label_2.SetObjectName("Label_2")                       // 112
	this.SizePolicy = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Preferred, qtwidgets.QSizePolicy__Expanding, 1)
	this.SizePolicy.SetHorizontalStretch(0)                                                                    // 114
	this.SizePolicy.SetVerticalStretch(0)                                                                      // 114
	this.SizePolicy.SetHeightForWidth(this.Label_2.SizePolicy().HasHeightForWidth())                           // 114
	this.Label_2.SetSizePolicy(this.SizePolicy)                                                                // 114
	this.Label_2.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.VerticalLayout.Layout().AddWidget(this.Label_2) // 115

	this.Label_3 = qtwidgets.NewQLabel(this.ContactItemView, 0)                                                // 111
	this.Label_3.SetObjectName("Label_3")                                                                      // 112
	this.SizePolicy.SetHeightForWidth(this.Label_3.SizePolicy().HasHeightForWidth())                           // 114
	this.Label_3.SetSizePolicy(this.SizePolicy)                                                                // 114
	this.Label_3.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.VerticalLayout.Layout().AddWidget(this.Label_3) // 115

	this.Label_4 = qtwidgets.NewQLabel(this.ContactItemView, 0)                                                // 111
	this.Label_4.SetObjectName("Label_4")                                                                      // 112
	this.SizePolicy.SetHeightForWidth(this.Label_4.SizePolicy().HasHeightForWidth())                           // 114
	this.Label_4.SetSizePolicy(this.SizePolicy)                                                                // 114
	this.Label_4.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.VerticalLayout.Layout().AddWidget(this.Label_4) // 115

	this.HorizontalLayout.AddLayout(this.VerticalLayout, 0) // 115

	this.VerticalLayout_2 = qtwidgets.NewQVBoxLayout()                         // 111
	this.VerticalLayout_2.SetObjectName("VerticalLayout_2")                    // 112
	this.VerticalLayout_2.SetSizeConstraint(qtwidgets.QLayout__SetMinimumSize) // 114
	this.VerticalSpacer = qtwidgets.NewQSpacerItem(20, 40, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)

	this.VerticalLayout_2.AddItem(this.VerticalSpacer) // 115

	this.ToolButton = qtwidgets.NewQToolButton(this.ContactItemView) // 111
	this.ToolButton.SetObjectName("ToolButton")                      // 112
	this.Icon = qtgui.NewQIcon()
	this.Icon.AddFile(":/icons/online_30.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton.SetIcon(this.Icon)                                                                   // 114
	this.ToolButton.SetAutoRepeat(false)                                                                 // 114
	this.ToolButton.SetAutoExclusive(false)                                                              // 114
	this.ToolButton.SetToolButtonStyle(qtcore.Qt__ToolButtonTextUnderIcon)                               // 114
	this.ToolButton.SetAutoRaise(true)                                                                   // 114

	this.VerticalLayout_2.Layout().AddWidget(this.ToolButton) // 115

	this.VerticalSpacer_2 = qtwidgets.NewQSpacerItem(20, 40, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)

	this.VerticalLayout_2.AddItem(this.VerticalSpacer_2) // 115

	this.HorizontalLayout.AddLayout(this.VerticalLayout_2, 0) // 115

	this.RetranslateUi(ContactItemView)

	qtcore.QMetaObject_ConnectSlotsByName(ContactItemView) // 100111
	// } // setupUi // 126

}

// void retranslateUi(QWidget *ContactItemView)
//  setupUi block end

//  retranslateUi block begin
func (this *Ui_ContactItemView) RetranslateUi(ContactItemView *qtwidgets.QWidget) {
	this.ContactItemView.SetWindowTitle(qtcore.QCoreApplication_Translate("ContactItemView", "Form", "dummy123", 0))
	this.Label.SetText(qtcore.QCoreApplication_Translate("ContactItemView", "TextLabel", "dummy123", 0))
	this.Label_2.SetText(qtcore.QCoreApplication_Translate("ContactItemView", "TextLabel", "dummy123", 0))
	this.Label_3.SetText(qtcore.QCoreApplication_Translate("ContactItemView", "TextLabel", "dummy123", 0))
	this.Label_4.SetText(qtcore.QCoreApplication_Translate("ContactItemView", "TextLabel", "dummy123", 0))
	this.ToolButton.SetText(qtcore.QCoreApplication_Translate("ContactItemView", "999", "dummy123", 0))
}

//  retranslateUi block end
