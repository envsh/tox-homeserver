package main

//  header block begin
import "qt.go/qtcore"
import "qt.go/qtgui"
import "qt.go/qtwidgets"
import "qt.go/qtquickwidgets"
import "qt.go/qtmock"

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
	if ContactItemView.ObjectName().IsEmpty() {
		ContactItemView.SetObjectName(qtcore.NewQString_5("ContactItemView"))
	}
	ContactItemView.Resize(252, 99)
	this.HorizontalLayout = qtwidgets.NewQHBoxLayout_1(qtwidgets.NewQWidgetFromPointer(this.ContactItemView.GetCthis())) // 111
	this.HorizontalLayout.SetObjectName(qtcore.NewQString_5("HorizontalLayout"))                                         // 112
	this.Label = qtwidgets.NewQLabel(this.ContactItemView, 0)                                                            // 111
	this.Label.SetObjectName(qtcore.NewQString_5("Label"))                                                               // 112
	this.Label.SetMaximumSize_1(30, 0)                                                                                   // 113

	this.HorizontalLayout.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.Label.GetCthis())) // 115

	this.VerticalLayout = qtwidgets.NewQVBoxLayout()                         // 111
	this.VerticalLayout.SetObjectName(qtcore.NewQString_5("VerticalLayout")) // 112
	this.Label_2 = qtwidgets.NewQLabel(this.ContactItemView, 0)              // 111
	this.Label_2.SetObjectName(qtcore.NewQString_5("Label_2"))               // 112
	this.SizePolicy = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Preferred, qtwidgets.QSizePolicy__Expanding, 1)
	this.SizePolicy.SetHorizontalStretch(0)                                                                    // 114
	this.SizePolicy.SetVerticalStretch(0)                                                                      // 114
	this.SizePolicy.SetHeightForWidth(this.Label_2.SizePolicy().HasHeightForWidth())                           // 114
	this.Label_2.SetSizePolicy(this.SizePolicy)                                                                // 114
	this.Label_2.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.VerticalLayout.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.Label_2.GetCthis())) // 115

	this.Label_3 = qtwidgets.NewQLabel(this.ContactItemView, 0)                                                // 111
	this.Label_3.SetObjectName(qtcore.NewQString_5("Label_3"))                                                 // 112
	this.SizePolicy.SetHeightForWidth(this.Label_3.SizePolicy().HasHeightForWidth())                           // 114
	this.Label_3.SetSizePolicy(this.SizePolicy)                                                                // 114
	this.Label_3.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.VerticalLayout.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.Label_3.GetCthis())) // 115

	this.Label_4 = qtwidgets.NewQLabel(this.ContactItemView, 0)                                                // 111
	this.Label_4.SetObjectName(qtcore.NewQString_5("Label_4"))                                                 // 112
	this.SizePolicy.SetHeightForWidth(this.Label_4.SizePolicy().HasHeightForWidth())                           // 114
	this.Label_4.SetSizePolicy(this.SizePolicy)                                                                // 114
	this.Label_4.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.VerticalLayout.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.Label_4.GetCthis())) // 115

	this.HorizontalLayout.AddLayout(qtwidgets.NewQLayoutFromPointer(this.VerticalLayout.GetCthis()), 0) // 115

	this.VerticalLayout_2 = qtwidgets.NewQVBoxLayout()                           // 111
	this.VerticalLayout_2.SetObjectName(qtcore.NewQString_5("VerticalLayout_2")) // 112
	this.VerticalLayout_2.SetSizeConstraint(qtwidgets.QLayout__SetMinimumSize)   // 114
	this.VerticalSpacer = qtwidgets.NewQSpacerItem(20, 40, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)

	this.VerticalLayout_2.AddItem(qtwidgets.NewQLayoutItemFromPointer(this.VerticalSpacer.GetCthis())) // 115

	this.ToolButton = qtwidgets.NewQToolButton(this.ContactItemView) // 111
	this.ToolButton.SetObjectName(qtcore.NewQString_5("ToolButton")) // 112
	this.Icon = qtgui.NewQIcon()
	this.Icon.AddFile(qtcore.NewQString_5(":/icons/online_30.png"), qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton.SetIcon(this.Icon)                                                                                        // 114
	this.ToolButton.SetAutoRepeat(false)                                                                                      // 114
	this.ToolButton.SetAutoExclusive(false)                                                                                   // 114
	this.ToolButton.SetToolButtonStyle(qtcore.Qt__ToolButtonTextUnderIcon)                                                    // 114
	this.ToolButton.SetAutoRaise(true)                                                                                        // 114

	this.VerticalLayout_2.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ToolButton.GetCthis())) // 115

	this.VerticalSpacer_2 = qtwidgets.NewQSpacerItem(20, 40, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)

	this.VerticalLayout_2.AddItem(qtwidgets.NewQLayoutItemFromPointer(this.VerticalSpacer_2.GetCthis())) // 115

	this.HorizontalLayout.AddLayout(qtwidgets.NewQLayoutFromPointer(this.VerticalLayout_2.GetCthis()), 0) // 115

	this.RetranslateUi(ContactItemView)

	qtcore.QMetaObject_ConnectSlotsByName(qtcore.NewQObjectFromPointer(ContactItemView.GetCthis())) // 100111
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
