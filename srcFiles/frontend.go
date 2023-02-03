package user_funcs

import (
	"database/sql"
	"strconv"

	"github.com/therecipe/qt/widgets"
)

var Globaldb *sql.DB

func handleAuthSection(layout *widgets.QVBoxLayout) *widgets.QGroupBox {

	AuthorSection := widgets.NewQGroupBox2("Author section", nil)
	AuthorSectionLayout := widgets.NewQVBoxLayout()
	// AuthorSectionLayout.AddStretch(2)

	AuthorSectionLabel := widgets.NewQLabel2("This is the Author section", nil, 0)
	AuthorSectionLayout.AddWidget(AuthorSectionLabel, 0, 0)

	// Create a new tab widget
	tabWidget := widgets.NewQTabWidget(nil)

	// Create a new widget for the first tab
	tab1 := widgets.NewQWidget(nil, 0)
	tab1Layout := widgets.NewQVBoxLayout()
	tab1Layout.AddStretch(2)
	tab1.SetLayout(tab1Layout)

	// Create a new scroll area
	scrollArea := widgets.NewQScrollArea(nil)

	// Create a new widget to hold the contents of the scroll area
	scrollWidget := widgets.NewQWidget(nil, 0)
	vbox := widgets.NewQVBoxLayout()
	vbox.AddStretch(4)
	scrollWidget.SetLayout(vbox)

	// Add questions UI
	questionEdit := widgets.NewQLineEdit(nil)
	questionEdit.SetPlaceholderText("Enter question")
	scrollWidget.Layout().AddWidget(questionEdit)

	// Create 4 QLineEdit for options
	option1Edit := widgets.NewQLineEdit(nil)

	option2Edit := widgets.NewQLineEdit(nil)
	option3Edit := widgets.NewQLineEdit(nil)
	option4Edit := widgets.NewQLineEdit(nil)
	correctoption := widgets.NewQLineEdit(nil)

	setPlaceholderquest := func() {
		option1Edit.SetPlaceholderText("Enter option 1")
		option2Edit.SetPlaceholderText("Enter option 2")
		option3Edit.SetPlaceholderText("Enter option 3")
		option4Edit.SetPlaceholderText("Enter option 4")
		correctoption.SetPlaceholderText("Enter correct option")
	}

	setPlaceholderquest()
	scrollWidget.Layout().AddWidget(option1Edit)
	scrollWidget.Layout().AddWidget(option2Edit)
	scrollWidget.Layout().AddWidget(option3Edit)
	scrollWidget.Layout().AddWidget(option4Edit)
	scrollWidget.Layout().AddWidget(correctoption)
	// Create a QPushButton for submission
	AddQuestionButton := widgets.NewQPushButton2("Add question", nil)
	scrollWidget.Layout().AddWidget(AddQuestionButton)

	// Connect the submit button to a function for handling the form submission
	AddQuestionButton.ConnectClicked(func(bool) {

		quest := Question{}

		// Get the text from the question and option QLineEdit

		quest.Questiontitle = questionEdit.Text()

		quest.Options = append(quest.Options, option1Edit.Text(), option2Edit.Text(), option3Edit.Text(), option4Edit.Text())
		quest.CorrectAns = correctoption.Text()

		InsertData(Globaldb, quest)
		questionEdit.SetText("")
		option1Edit.SetText("")
		option2Edit.SetText("")
		option3Edit.SetText("")
		option4Edit.SetText("")
		correctoption.SetText("")

	})

	// Set the scroll widget as the central widget of the scroll area
	scrollArea.SetWidget(scrollWidget)

	// Make the scroll area scrollable
	scrollArea.SetWidgetResizable(true)

	//````````````````````````````````````````````````````````````````/````/
	tab1Layout.AddWidget(scrollArea, 0, 0)
	tab1Layout.AddStretch(2)
	// Create a new widget for the second tab

	// Create a new widget for the second tab
	tab2 := widgets.NewQWidget(nil, 0)
	tab2Layout := widgets.NewQVBoxLayout()
	tab2.SetLayout(tab2Layout)

	// Create a new scroll area
	scrollArea2 := widgets.NewQScrollArea(nil)

	// Create a new widget to hold the contents of the scroll area
	scrollWidget2 := widgets.NewQWidget(nil, 0)
	vbox2 := widgets.NewQVBoxLayout()
	vbox2.AddStretch(4)
	scrollWidget2.SetLayout(vbox2)

	resultantrows, _ := ReadData(Globaldb)

	for _, rowresult := range resultantrows {

		questionLabel := widgets.NewQLabel2(rowresult.Questiontitle, nil, 0)

		questionUpdate := widgets.NewQLineEdit(nil)
		questionUpdate.SetText(rowresult.Questiontitle)

		option1Update := widgets.NewQLineEdit(nil)
		option1Update.SetText(rowresult.Options[0])

		option2Update := widgets.NewQLineEdit(nil)
		option2Update.SetText(rowresult.Options[1])

		option3Update := widgets.NewQLineEdit(nil)
		option3Update.SetText(rowresult.Options[2])

		option4Update := widgets.NewQLineEdit(nil)
		option4Update.SetText(rowresult.Options[3])

		Correctansupdate := widgets.NewQLineEdit(nil)
		Correctansupdate.SetText(rowresult.CorrectAns)

		DeleteQuestion := widgets.NewQPushButton2("Delete", nil)

		UpdateQuestion := widgets.NewQPushButton2("Update", nil)
		scrollWidget2.Layout().AddWidget(questionLabel)
		scrollWidget2.Layout().AddWidget(questionUpdate)
		scrollWidget2.Layout().AddWidget(option1Update)
		scrollWidget2.Layout().AddWidget(option2Update)
		scrollWidget2.Layout().AddWidget(option3Update)
		scrollWidget2.Layout().AddWidget(option4Update)
		scrollWidget2.Layout().AddWidget(Correctansupdate)
		scrollWidget2.Layout().AddWidget(UpdateQuestion)
		scrollWidget2.Layout().AddWidget(DeleteQuestion)

		DeleteQuestion.ConnectClicked(func(bool) {
			scrollWidget2.Layout().RemoveWidget(UpdateQuestion)
			scrollWidget2.Layout().RemoveWidget(questionLabel)
			scrollWidget2.Layout().RemoveWidget(questionUpdate)
			scrollWidget2.Layout().RemoveWidget(option1Update)
			scrollWidget2.Layout().RemoveWidget(option2Update)
			scrollWidget2.Layout().RemoveWidget(option3Update)
			scrollWidget2.Layout().RemoveWidget(option4Update)
			scrollWidget2.Layout().RemoveWidget(Correctansupdate)
			scrollWidget2.Layout().RemoveWidget(DeleteQuestion)
			DeleteData(Globaldb, questionUpdate.Text())
		})

		UpdateQuestion.ConnectClicked(func(bool) {

			FinalQuestion := Question{}

			// Get the text from the question and option QLineEdit

			FinalQuestion.Questiontitle = questionUpdate.Text()

			FinalQuestion.Options = append(FinalQuestion.Options, option1Update.Text(), option2Update.Text(), option3Update.Text(), option4Update.Text())
			FinalQuestion.CorrectAns = Correctansupdate.Text()

			UpdateData(Globaldb, FinalQuestion, questionLabel.Text())
			questionLabel.SetText(questionUpdate.Text())
		})

		// Set the scroll widget as the central widget of the scroll area
		scrollArea2.SetWidget(scrollWidget2)

	}

	// Make the scroll area scrollable
	scrollArea2.SetWidgetResizable(true)

	tab2Layout.AddWidget(scrollArea2, 0, 0)
	tab2Layout.AddStretch(2)
	// Create a new widget for the second tab

	// Add the tab widgets to the tab widget
	tabWidget.AddTab(tab1, "Add Questions")
	tabWidget.AddTab(tab2, "Update Questions")

	AuthorSectionLayout.AddWidget(tabWidget, 0, 0)

	AuthorSection.SetLayout(AuthorSectionLayout)
	layout.AddWidget(AuthorSection, 0, 0)

	return AuthorSection
}

var ReadscrollWidget *widgets.QWidget

func handleStudSection(layout *widgets.QVBoxLayout) *widgets.QGroupBox {

	StudentSection := widgets.NewQGroupBox2("Student Section", nil)
	StudentSectionLayout := widgets.NewQVBoxLayout()

	// Create a new scroll area
	scrollArea := widgets.NewQScrollArea(nil)

	// Create a new widget to hold the contents of the scroll area
	ReadscrollWidget := widgets.NewQWidget(nil, 0)
	VerticalBox := widgets.NewQVBoxLayout()
	ReadscrollWidget.SetLayout(VerticalBox)

	StudentSectionLabel := widgets.NewQLabel2("CORRECT ANSWERS ARE CHECKED", nil, 0)

	ReadscrollWidget.Layout().AddWidget(StudentSectionLabel)

	result, _ := ReadData(Globaldb)

	for i, rowdata := range result {
		// Create the group box for the question
		groupBox := widgets.NewQGroupBox2("Qno:"+strconv.Itoa(i+1)+" "+rowdata.Questiontitle, ReadscrollWidget)

		// Create a layout for the group box
		groupBoxLayout := widgets.NewQVBoxLayout()
		groupBox.SetLayout(groupBoxLayout)

		// Create the radio buttons for the choices
		choice1 := widgets.NewQRadioButton2(rowdata.Options[0], groupBox)
		choice2 := widgets.NewQRadioButton2(rowdata.Options[1], groupBox)
		choice3 := widgets.NewQRadioButton2(rowdata.Options[2], groupBox)
		choice4 := widgets.NewQRadioButton2(rowdata.Options[3], groupBox)

		// Add the radio buttons to the group box layout
		groupBoxLayout.AddWidget(choice1, 0, 0)
		groupBoxLayout.AddWidget(choice2, 0, 0)
		groupBoxLayout.AddWidget(choice3, 0, 0)
		groupBoxLayout.AddWidget(choice4, 0, 0)

		switch rowdata.CorrectAns {
		case choice1.Text():
			choice1.SetChecked(true)
		case choice2.Text():
			choice2.SetChecked(true)
		case choice4.Text():
			choice3.SetChecked(true)
		case choice4.Text():
			choice4.SetChecked(true)
		}

		ReadscrollWidget.Layout().AddWidget(groupBox)
	}

	// Add widgets to the scroll widget

	// Set the scroll widget as the central widget of the scroll area
	scrollArea.SetWidget(ReadscrollWidget)
	// Make the scroll area scrollable
	scrollArea.SetWidgetResizable(true)

	StudentSectionLayout.AddWidget(scrollArea, 0, 0)
	StudentSection.SetLayout(StudentSectionLayout)
	layout.AddWidget(StudentSection, 0, 0)

	return StudentSection
}

func HandleFrontend(window *widgets.QMainWindow, db *sql.DB) {

	Globaldb = db
	// Create a vertical layout

	layout := widgets.NewQVBoxLayout()

	// Create a checkbox
	checkbox1 := widgets.NewQCheckBox2("Author", nil)
	checkbox2 := widgets.NewQCheckBox2("Student", nil)

	// Create a label
	label := widgets.NewQLabel2("", nil, 0)
	label.SetText("Check your role above!")

	checkIfBothUnchecked := func() bool {
		if !checkbox1.IsChecked() && !checkbox2.IsChecked() {
			// do something if both checkbox are unchecked
			label.SetText("Check your role above!")
			return false
		}
		return true
	}

	// Create a layout to hold the checkbox and label
	Horizontallayout := widgets.NewQHBoxLayout()
	Horizontallayout.AddWidget(checkbox1, 0, 0)
	Horizontallayout.AddWidget(checkbox2, 0, 0)

	layout.AddLayout(Horizontallayout, 0)

	// add label to Vertical layout
	layout.AddWidget(label, 0, 0)

	// Create a Authors section
	AuthorSection := handleAuthSection(layout)

	// Create a Student section
	StudentSection := handleStudSection(layout)

	// UnCheckedSection:
	UnCheckedSection := widgets.NewQGroupBox2("third Section", nil)
	UnCheckedSectionLayout := widgets.NewQVBoxLayout()
	UnCheckedSectionLabel := widgets.NewQLabel2("Make some choice", nil, 0)
	UnCheckedSectionLayout.AddWidget(UnCheckedSectionLabel, 0, 0)
	UnCheckedSection.SetLayout(UnCheckedSectionLayout)
	layout.AddWidget(UnCheckedSection, 0, 0)

	//function to maintain sections
	ManageSections := func() {
		if checkbox1.IsChecked() {
			AuthorSection.SetVisible(true)
			StudentSection.SetVisible(false)
			UnCheckedSection.SetVisible(false)

		} else if checkbox2.IsChecked() {
			AuthorSection.SetVisible(false)
			StudentSection.SetVisible(true)
			UnCheckedSection.SetVisible(false)

		} else {
			AuthorSection.SetVisible(false)
			StudentSection.SetVisible(false)
			UnCheckedSection.SetVisible(true)
		}
	}
	ManageSections()

	// Connect the checkbox to the label
	checkbox1.ConnectToggled(func(checked bool) {
		if checked {
			label.SetText("You Can edit the Questions!")
			checkbox2.SetChecked(false)
		}
		checkIfBothUnchecked()
		ManageSections()
	})
	checkbox2.ConnectToggled(func(checked bool) {
		if checked {
			label.SetText("All the Question from database!")
			checkbox1.SetChecked(false)
		}
		checkIfBothUnchecked()
		ManageSections()
	})

	// Create a central widget and set the layout
	centralWidget := widgets.NewQWidget(nil, 0)
	centralWidget.SetLayout(layout)
	window.SetCentralWidget(centralWidget)

}
