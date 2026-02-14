package parser

import (
	"fmt"
	"testing"
)

/*
Check to make sure that the CheckVariableExistence() function reports,
correctly, whether a variable exists and has a value. The
CheckVariableExistence() function should return three values:
 1. true, false - the variable exists but has no assigned value (this would
    include placeholders for reserved variables).
 2. true, true - the variable exists and has an assigned value.
 3. false, false - the variable doesn't exist and therefore, doesn't have an
    assigned value

We check to make sure that each of those is occurring here.
*/
func TestCheckVariableExistence(t *testing.T) {
	/*
		Create some fictional variables to test against here. We have a
		variable with no value and one with a value.
	*/
	VARIABLES["novalue"] = ""
	VARIABLES["yesvalue"] = "yes"

	// Check condition 1 (ie. true, false)
	true_false_exist, true_false_value := CheckVariableExistence("novalue")
	// Check condition 2 (ie. true, true)
	true_true_exist, true_true_value := CheckVariableExistence("yesvalue")
	// Check condition 3 (ie. false, false)
	false_false_exist, false_false_value := CheckVariableExistence("value")

	// Condition 1 - throw an error if the exist is false or the value is true
	if !true_false_exist || true_false_value {
		t.Errorf(
			"[CheckVariableExistence] CheckVariableExistence returned %t, "+
				"%t, expected true, false",
			true_false_exist,
			true_false_value,
		)
	}
	// Condition 2 - throw an error if the exist is false or the value is false
	if !true_true_exist || !true_true_value {
		t.Errorf(
			"[CheckVariableExistence] CheckVariableExistence returned %t, "+
				"%t, expected true, false",
			true_true_exist,
			true_true_value,
		)
	}
	// Condition 3 - throw an error if the exist is true or the value is true
	if false_false_exist || false_false_value {
		t.Errorf(
			"[CheckVariableExistence] CheckVariableExistence returned %t, "+
				"%t, expected true, false",
			false_false_exist,
			false_false_value,
		)
	}
}

/*
Check to make sure that the removal of comments does what it should - replace
the comments with empty comment strings.
*/
func TestVariableTemplater(t *testing.T) {
	// Set up some dummy variables for the variable templater
	VARIABLES["lang"] = "TestLang"
	VARIABLES["version"] = "4"
	VARIABLES["codename"] = "CityName"

	// A simple example
	lang_and_ver := "App: #lang, Version: #version"
	// A string formatted example of the string from above
	lang_and_ver_correct := fmt.Sprintf(
		"App: %s, Version: %s",
		VARIABLES["lang"],
		VARIABLES["version"],
	)
	// A templated version
	lang_and_ver_templated := VariableTemplater(lang_and_ver)

	// A second simple example
	lang_ver_codename := "App=#lang and Version=#version " +
		"(Code Name: #codename)"
	// A string formatted example of the string from above
	lang_ver_codename_correct := fmt.Sprintf(
		"App=%s and Version=%s (Code Name: %s)",
		VARIABLES["lang"],
		VARIABLES["version"],
		VARIABLES["codename"],
	)
	// A templated version of the second example
	lang_ver_codename_templated := VariableTemplater(lang_ver_codename)

	if lang_and_ver_templated != lang_and_ver_correct {
		t.Errorf(
			"[VariableTemplater] VariableTemplater returned %s, %s"+
				", expected true, true",
			lang_and_ver_templated,
			lang_and_ver_correct,
		)
	}

	if lang_ver_codename_templated != lang_ver_codename_correct {
		t.Errorf(
			"[VariableTemplater] VariableTemplater returned %s, %s"+
				", expected true, true",
			lang_ver_codename_templated,
			lang_ver_codename_correct,
		)
	}

}
