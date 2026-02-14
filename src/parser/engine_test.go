/*
Tests for the engine, largely for the Tokenise() function for each statement.
*/
package parser

import (
	"reflect"
	"testing"
)

func TestValidAskCall(t *testing.T) {
	results := Tokenise("ask \"Greeting: \" to greeting", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_ASK)

	if !tokenisation_equal {
		t.Errorf(
			"[ask stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_ASK,
		)
	}
}

func TestValidCopyDirCall(t *testing.T) {
	results := Tokenise("copydirectory \"/home/user/test\" to \"/home/user/test2\"", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_COPYDIR)

	if !tokenisation_equal {
		t.Errorf(
			"[copydirectory stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_COPYDIR,
		)
	}
}

func TestValidCopyFileCall(t *testing.T) {
	results := Tokenise("copyfile \"/home/user/test.txt\" to \"/home/user/test2.txt\"", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_COPYFILE)

	if !tokenisation_equal {
		t.Errorf(
			"[copyfile stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_COPYFILE,
		)
	}
}

func TestValidDeleteDirCall(t *testing.T) {
	results := Tokenise("deletedirectory \"/home/user/test/\"", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_DELETEDIR)

	if !tokenisation_equal {
		t.Errorf(
			"[deletedirectory stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_DELETEDIR,
		)
	}
}

func TestValidDeleteFileCall(t *testing.T) {
	results := Tokenise("deletefile \"/home/user/test.txt\"", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_DELETEFILE)

	if !tokenisation_equal {
		t.Errorf(
			"[deletefile stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_DELETEFILE,
		)
	}
}

func TestValidDownloadCall(t *testing.T) {
	results := Tokenise(
		"download \"http://upload.wikimedia.org/wikipedia/commons/0/02/"+
			"La_Libert%C3%A9_guidant_le_peuple_-_Eug%C3%A8ne_Delacroix_-_Mus%C3"+
			"%A9e_du_Louvre_Peintures_RF_129_-_apr%C3%A8s_restauration_2024."+
			"jpg\" to \"#b_home/Desktop/del.jpg\"", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_DOWNLOADFILE)

	if !tokenisation_equal {
		t.Errorf(
			"[download stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_DOWNLOADFILE,
		)
	}
}

func TestValidExecuteCall(t *testing.T) {
	results := Tokenise("execute \"ls -l\"", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_EXECUTE)

	if !tokenisation_equal {
		t.Errorf(
			"[execute stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_DOWNLOADFILE,
		)
	}
}

func TestExit(t *testing.T) {
	results := Tokenise("exit", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_EXIT)

	if !tokenisation_equal {
		t.Errorf(
			"[exit stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_EXIT,
		)
	}
}

func TestValidMakeDirCall(t *testing.T) {
	results := Tokenise("makedirectory \"#b_home/Downloads/testdir2\"", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_MAKEDIR)

	if !tokenisation_equal {
		t.Errorf(
			"[set stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_MAKEDIR,
		)
	}
}

func TestValidMakeFileCall(t *testing.T) {
	results := Tokenise("makefile \"#b_home/Downloads/testdir2.txt\"", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_MAKEFILE)

	if !tokenisation_equal {
		t.Errorf(
			"[set stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_MAKEFILE,
		)
	}
}

func TestValidMinVerCall(t *testing.T) {
	results := Tokenise("minver 1", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_MINVER)

	if !tokenisation_equal {
		t.Errorf(
			"[minver stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_SET,
		)
	}
}

func TestValidMinverCalls(t *testing.T) {
	simple_line_sets := []string{
		"minver 1",
		"writeln \"Hello World!\"",
	}

	duplicate_minvers := []string{
		"minver 1",
		"minver 3",
		"write \"Hello World!\"",
	}

	incorrect_order_minvers := []string{
		"write \"Hello World!\"",
		"pause 3",
		"minver 1",
	}

	// Check a basic minver check
	simple_minver_check, _ := CheckValidMinverLocationAndCount(
		simple_line_sets,
	)
	// Check that it calls duplicates of minver calls
	duplicate_minver_check, _ := CheckValidMinverLocationAndCount(
		duplicate_minvers,
	)
	/* Check for incorrectly ordered statement (ie. wrongly placed minver
	calls)
	*/
	incorrect_order_minver_check, _ := CheckValidMinverLocationAndCount(
		incorrect_order_minvers,
	)

	if !simple_minver_check {
		t.Errorf(
			"MinverCheck failed (simple), got false, " +
				"expected true",
		)
	}

	if duplicate_minver_check {
		t.Errorf(
			"MinverCheck failed (duplicate), got false, " +
				"expected true",
		)
	}

	if incorrect_order_minver_check {
		t.Errorf(
			"MinverCheck failed (not first line), got false, " +
				"expected true",
		)
	}
}

func TestValidMoveDirCall(t *testing.T) {
	results := Tokenise("movedirectory \"/home/user/test\" to \"/home/user/test2\"", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_MOVEDIR)

	if !tokenisation_equal {
		t.Errorf(
			"[movedirectory stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_MOVEDIR,
		)
	}
}

func TestValidMoveFileCall(t *testing.T) {
	results := Tokenise("movefile \"/home/user/test.txt\" to \"/home/user/test\"", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_MOVEFILE)

	if !tokenisation_equal {
		t.Errorf(
			"[movefile stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_MOVEFILE,
		)
	}
}

func TestValidPauseCall(t *testing.T) {
	results := Tokenise("pause 3", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_PAUSE)

	if !tokenisation_equal {
		t.Errorf(
			"[pause stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_PAUSE,
		)
	}
}

func TestValidRunCall(t *testing.T) {
	results := Tokenise("run \"../samples/write.apt\"", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_RUN)

	if !tokenisation_equal {
		t.Errorf(
			"[run stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_RUN,
		)
	}
}

func TestValidSetCall(t *testing.T) {
	results := Tokenise("set name = \"Hello World!\"", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_SET)

	if !tokenisation_equal {
		t.Errorf(
			"[set stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_SET,
		)
	}
}

func TestValidWriteCall(t *testing.T) {
	results := Tokenise("write \"Hello World!\"", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_WRITE)

	if !tokenisation_equal {
		t.Errorf(
			"[write stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_WRITE,
		)
	}
}

func TestValidWriteLnCall(t *testing.T) {
	results := Tokenise("writeln \"Hello World!\"", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_WRITELN)

	if !tokenisation_equal {
		t.Errorf(
			"[writeln stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_WRITELN,
		)
	}
}

func TestValidZipDirCall(t *testing.T) {
	results := Tokenise("zipdirectory \"/home/user/test\" to \"/home/user/test2.zip\"", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_ZIPDIR)

	if !tokenisation_equal {
		t.Errorf(
			"[zipdirectory stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_ZIPDIR,
		)
	}
}

func TestValidZipFileCall(t *testing.T) {
	results := Tokenise("zipfile \"/home/user/test.txt\" to \"/home/user/test2.zip\"", 1, 1)
	tokenisation_equal := reflect.DeepEqual(results, TEST_ZIPFILE)

	if !tokenisation_equal {
		t.Errorf(
			"[zipfile stmt] Tokenisation failed, got %v, expected %v",
			results,
			TEST_ZIPFILE,
		)
	}
}
