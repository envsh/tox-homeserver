package main

//  header block begin
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"
import "github.com/kitech/qt.go/qtquickwidgets"
import "github.com/kitech/qt.go/qtmock"
import "github.com/kitech/qt.go/qtrt"

func init() { qtcore.KeepMe() }
func init() { qtgui.KeepMe() }
func init() { qtwidgets.KeepMe() }
func init() { qtquickwidgets.KeepMe() }
func init() { qtmock.KeepMe() }
func init() { qtrt.KeepMe() }

//  header block end

//  struct block begin
func NewUi_ContactItemView() *Ui_ContactItemView {
	return &Ui_ContactItemView{}
}

type Ui_ContactItemView struct {
	HorizontalLayout_3 *qtwidgets.QHBoxLayout
	VerticalLayout_3   *qtwidgets.QVBoxLayout
	VerticalSpacer_3   *qtwidgets.QSpacerItem
	ToolButton_2       *qtwidgets.QToolButton
	VerticalSpacer_4   *qtwidgets.QSpacerItem
	VerticalLayout     *qtwidgets.QVBoxLayout
	HorizontalLayout   *qtwidgets.QHBoxLayout
	Label_2            *qtwidgets.QLabel
	HorizontalSpacer   *qtwidgets.QSpacerItem
	LabelLastMsgTime   *qtwidgets.QLabel
	HorizontalLayout_2 *qtwidgets.QHBoxLayout
	Label_3            *qtwidgets.QLabel
	HorizontalSpacer_2 *qtwidgets.QSpacerItem
	Label_4            *qtwidgets.QLabel
	VerticalLayout_2   *qtwidgets.QVBoxLayout
	VerticalSpacer     *qtwidgets.QSpacerItem
	ToolButton         *qtwidgets.QToolButton
	VerticalSpacer_2   *qtwidgets.QSpacerItem
	ContactItemView    *qtwidgets.QWidget
	SizePolicy         *qtwidgets.QSizePolicy
	Icon               *qtgui.QIcon // 116
	SizePolicy1        *qtwidgets.QSizePolicy
	Font               *qtgui.QFont // 116
	SizePolicy2        *qtwidgets.QSizePolicy
	Icon1              *qtgui.QIcon // 116
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
	ContactItemView.Resize(419, 64)
	this.SizePolicy = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Preferred, qtwidgets.QSizePolicy__Fixed, 1)
	this.SizePolicy.SetHorizontalStretch(0)                                                  // 114
	this.SizePolicy.SetVerticalStretch(0)                                                    // 114
	this.SizePolicy.SetHeightForWidth(this.ContactItemView.SizePolicy().HasHeightForWidth()) // 114
	this.ContactItemView.SetSizePolicy(this.SizePolicy)                                      // 114
	this.ContactItemView.SetFocusPolicy(qtcore.Qt__ClickFocus)                               // 114
	this.ContactItemView.SetContextMenuPolicy(qtcore.Qt__DefaultContextMenu)                 // 114
	this.HorizontalLayout_3 = qtwidgets.NewQHBoxLayout_1(this.ContactItemView)               // 111
	this.HorizontalLayout_3.SetSpacing(1)                                                    // 114
	this.HorizontalLayout_3.SetObjectName("HorizontalLayout_3")                              // 112
	this.HorizontalLayout_3.SetContentsMargins(0, 1, 0, 0)                                   // 114
	this.VerticalLayout_3 = qtwidgets.NewQVBoxLayout()                                       // 111
	this.VerticalLayout_3.SetSpacing(0)                                                      // 114
	this.VerticalLayout_3.SetObjectName("VerticalLayout_3")                                  // 112
	this.VerticalLayout_3.SetContentsMargins(-1, 0, -1, 0)                                   // 114
	this.VerticalSpacer_3 = qtwidgets.NewQSpacerItem(20, 3, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)
	qtrt.ReleaseOwnerToQt(this.VerticalSpacer_3)

	this.VerticalLayout_3.AddItem(this.VerticalSpacer_3) // 115

	this.ToolButton_2 = qtwidgets.NewQToolButton(this.ContactItemView) // 111
	this.ToolButton_2.SetObjectName("ToolButton_2")                    // 112
	this.ToolButton_2.SetFocusPolicy(qtcore.Qt__NoFocus)               // 114
	this.ToolButton_2.SetContextMenuPolicy(qtcore.Qt__NoContextMenu)   // 114
	this.Icon = qtgui.NewQIcon()
	this.Icon.AddFile(":/icons/icon_avatar_40.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_2.SetIcon(this.Icon)                                                                      // 114
	this.ToolButton_2.SetIconSize(qtcore.NewQSize_1(32, 32))                                                  // 113
	this.ToolButton_2.SetAutoRaise(true)                                                                      // 114

	this.VerticalLayout_3.Layout().AddWidget(this.ToolButton_2) // 115

	this.VerticalSpacer_4 = qtwidgets.NewQSpacerItem(20, 3, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)
	qtrt.ReleaseOwnerToQt(this.VerticalSpacer_4)

	this.VerticalLayout_3.AddItem(this.VerticalSpacer_4) // 115

	this.HorizontalLayout_3.AddLayout(this.VerticalLayout_3, 0) // 115

	this.VerticalLayout = qtwidgets.NewQVBoxLayout()            // 111
	this.VerticalLayout.SetSpacing(0)                           // 114
	this.VerticalLayout.SetObjectName("VerticalLayout")         // 112
	this.HorizontalLayout = qtwidgets.NewQHBoxLayout()          // 111
	this.HorizontalLayout.SetSpacing(0)                         // 114
	this.HorizontalLayout.SetObjectName("HorizontalLayout")     // 112
	this.Label_2 = qtwidgets.NewQLabel(this.ContactItemView, 0) // 111
	this.Label_2.SetObjectName("Label_2")                       // 112
	this.SizePolicy1 = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Preferred, 1)
	this.SizePolicy1.SetHorizontalStretch(0)                                          // 114
	this.SizePolicy1.SetVerticalStretch(0)                                            // 114
	this.SizePolicy1.SetHeightForWidth(this.Label_2.SizePolicy().HasHeightForWidth()) // 114
	this.Label_2.SetSizePolicy(this.SizePolicy1)                                      // 114
	this.Font = qtgui.NewQFont()
	this.Font.SetBold(true)                                                                                    // 114
	this.Font.SetWeight(75)                                                                                    // 114
	this.Label_2.SetFont(this.Font)                                                                            // 114
	this.Label_2.SetContextMenuPolicy(qtcore.Qt__NoContextMenu)                                                // 114
	this.Label_2.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.HorizontalLayout.Layout().AddWidget(this.Label_2) // 115

	this.HorizontalSpacer = qtwidgets.NewQSpacerItem(3, 20, qtwidgets.QSizePolicy__Preferred, qtwidgets.QSizePolicy__Minimum)
	qtrt.ReleaseOwnerToQt(this.HorizontalSpacer)

	this.HorizontalLayout.AddItem(this.HorizontalSpacer) // 115

	this.LabelLastMsgTime = qtwidgets.NewQLabel(this.ContactItemView, 0) // 111
	this.LabelLastMsgTime.SetObjectName("LabelLastMsgTime")              // 112
	this.SizePolicy2 = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Preferred, qtwidgets.QSizePolicy__Preferred, 1)
	this.SizePolicy2.SetHorizontalStretch(0)                                                                       // 114
	this.SizePolicy2.SetVerticalStretch(0)                                                                         // 114
	this.SizePolicy2.SetHeightForWidth(this.LabelLastMsgTime.SizePolicy().HasHeightForWidth())                     // 114
	this.LabelLastMsgTime.SetSizePolicy(this.SizePolicy2)                                                          // 114
	this.LabelLastMsgTime.SetContextMenuPolicy(qtcore.Qt__NoContextMenu)                                           // 114
	this.LabelLastMsgTime.SetAlignment(qtcore.Qt__AlignRight | qtcore.Qt__AlignTrailing | qtcore.Qt__AlignVCenter) // 114
	this.LabelLastMsgTime.SetTextInteractionFlags(qtcore.Qt__TextBrowserInteraction)                               // 114

	this.HorizontalLayout.Layout().AddWidget(this.LabelLastMsgTime) // 115

	this.VerticalLayout.AddLayout(this.HorizontalLayout, 0) // 115

	this.HorizontalLayout_2 = qtwidgets.NewQHBoxLayout()                              // 111
	this.HorizontalLayout_2.SetSpacing(0)                                             // 114
	this.HorizontalLayout_2.SetObjectName("HorizontalLayout_2")                       // 112
	this.Label_3 = qtwidgets.NewQLabel(this.ContactItemView, 0)                       // 111
	this.Label_3.SetObjectName("Label_3")                                             // 112
	this.SizePolicy1.SetHeightForWidth(this.Label_3.SizePolicy().HasHeightForWidth()) // 114
	this.Label_3.SetSizePolicy(this.SizePolicy1)                                      // 114
	this.Label_3.SetContextMenuPolicy(qtcore.Qt__NoContextMenu)                       // 114
	this.Label_3.SetTextInteractionFlags(qtcore.Qt__TextBrowserInteraction)           // 114

	this.HorizontalLayout_2.Layout().AddWidget(this.Label_3) // 115

	this.HorizontalSpacer_2 = qtwidgets.NewQSpacerItem(3, 20, qtwidgets.QSizePolicy__Preferred, qtwidgets.QSizePolicy__Minimum)
	qtrt.ReleaseOwnerToQt(this.HorizontalSpacer_2)

	this.HorizontalLayout_2.AddItem(this.HorizontalSpacer_2) // 115

	this.VerticalLayout.AddLayout(this.HorizontalLayout_2, 0) // 115

	this.Label_4 = qtwidgets.NewQLabel(this.ContactItemView, 0)                       // 111
	this.Label_4.SetObjectName("Label_4")                                             // 112
	this.SizePolicy1.SetHeightForWidth(this.Label_4.SizePolicy().HasHeightForWidth()) // 114
	this.Label_4.SetSizePolicy(this.SizePolicy1)                                      // 114
	this.Label_4.SetContextMenuPolicy(qtcore.Qt__NoContextMenu)                       // 114
	this.Label_4.SetTextInteractionFlags(qtcore.Qt__TextBrowserInteraction)           // 114

	this.VerticalLayout.Layout().AddWidget(this.Label_4) // 115

	this.HorizontalLayout_3.AddLayout(this.VerticalLayout, 0) // 115

	this.VerticalLayout_2 = qtwidgets.NewQVBoxLayout()                         // 111
	this.VerticalLayout_2.SetSpacing(0)                                        // 114
	this.VerticalLayout_2.SetObjectName("VerticalLayout_2")                    // 112
	this.VerticalLayout_2.SetSizeConstraint(qtwidgets.QLayout__SetMinimumSize) // 114
	this.VerticalSpacer = qtwidgets.NewQSpacerItem(20, 3, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)
	qtrt.ReleaseOwnerToQt(this.VerticalSpacer)

	this.VerticalLayout_2.AddItem(this.VerticalSpacer) // 115

	this.ToolButton = qtwidgets.NewQToolButton(this.ContactItemView) // 111
	this.ToolButton.SetObjectName("ToolButton")                      // 112
	this.ToolButton.SetFocusPolicy(qtcore.Qt__NoFocus)               // 114
	this.ToolButton.SetContextMenuPolicy(qtcore.Qt__NoContextMenu)   // 114
	this.Icon1 = qtgui.NewQIcon()
	this.Icon1.AddFile(":/icons/online_30.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton.SetIcon(this.Icon1)                                                                   // 114
	this.ToolButton.SetAutoRepeat(false)                                                                  // 114
	this.ToolButton.SetAutoExclusive(false)                                                               // 114
	this.ToolButton.SetToolButtonStyle(qtcore.Qt__ToolButtonTextUnderIcon)                                // 114
	this.ToolButton.SetAutoRaise(true)                                                                    // 114

	this.VerticalLayout_2.Layout().AddWidget(this.ToolButton) // 115

	this.VerticalSpacer_2 = qtwidgets.NewQSpacerItem(20, 3, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)
	qtrt.ReleaseOwnerToQt(this.VerticalSpacer_2)

	this.VerticalLayout_2.AddItem(this.VerticalSpacer_2) // 115

	this.HorizontalLayout_3.AddLayout(this.VerticalLayout_2, 0) // 115

	this.RetranslateUi(ContactItemView)

	qtcore.QMetaObject_ConnectSlotsByName(ContactItemView) // 100111
	// } // setupUi // 126

}

// void retranslateUi(QWidget *ContactItemView)
//  setupUi block end

//  retranslateUi block begin
func (this *Ui_ContactItemView) RetranslateUi(ContactItemView *qtwidgets.QWidget) {
	// noimpl: {
	this.ContactItemView.SetWindowTitle(qtcore.QCoreApplication_Translate("ContactItemView", "Form", "dummy123", 0))
	this.ToolButton_2.SetText(qtcore.QCoreApplication_Translate("ContactItemView", "...", "dummy123", 0))
	this.Label_2.SetText(qtcore.QCoreApplication_Translate("ContactItemView", "TextLabel", "dummy123", 0))
	this.LabelLastMsgTime.SetText(qtcore.QCoreApplication_Translate("ContactItemView", "TextLabel", "dummy123", 0))
	this.Label_3.SetText(qtcore.QCoreApplication_Translate("ContactItemView", "TextLabel", "dummy123", 0))
	this.Label_4.SetText(qtcore.QCoreApplication_Translate("ContactItemView", "TextLabel", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton.SetToolTip(qtcore.QCoreApplication_Translate("ContactItemView", "1 Online status. 2 Unread message count.", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	this.ToolButton.SetText(qtcore.QCoreApplication_Translate("ContactItemView", "999", "dummy123", 0))
	// noimpl: } // retranslateUi
	// noimpl:
	// noimpl: };
	// noimpl:
}

//  retranslateUi block end

//  new2 block begin
func NewUi_ContactItemView2() *Ui_ContactItemView {
	this := &Ui_ContactItemView{}
	w := qtwidgets.NewQWidget(nil, 0)
	this.SetupUi(w)
	return this
}

//  new2 block end

//  done block begin

func (this *Ui_ContactItemView) QWidget_PTR() *qtwidgets.QWidget {
	return this.ContactItemView.QWidget_PTR()
}

//  done block end
