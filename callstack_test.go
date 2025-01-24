package toolz

import (
	"fmt"
	"log/slog"
	"testing"
)

func TestCallStack(t *testing.T) {

	stack := GetCallStack(2)

	// Keep this around for those cases where the call stack function names change
	// for _, entry := range stack {
	// 	t.Logf("funcName: %s, fileName: %s, line: %d", entry.FuncName, entry.FileName, entry.Line)
	// }

	expCount := 2
	expFuncNames := []string{
		"api-service/pkg/toolz.TestCallStack",
		"testing.tRunner",
	}

	if len(stack) != expCount {
		t.Errorf("GetCallStack() returned incorrect number of entries. Expected: %d, Got: %d", expCount, len(stack))
	}

	for i, entry := range stack {
		if entry.FuncName != expFuncNames[i] {
			t.Errorf("GetCallStack() returned incorrect function name at index %d. Expected: %s, Got: %s", i, expFuncNames[i], entry.FuncName)
		}
	}

	expKey := "call_stack"
	expKind := slog.KindGroup
	slogified := stack.Slogify(expKey)

	if slogified.Key != expKey {
		t.Errorf("CallStack.Slogify() returned incorrect name. Expected: %s, Got: %s", expKey, slogified.Key)
	}

	if slogified.Value.Kind() != expKind {
		t.Errorf("CallStack.Slogify() returned incorrect kind. Expected: %s, Got: %s", expKind, slogified.Value.Kind())
	}

	grp := slogified.Value.Group()
	if len(grp) != expCount {
		t.Errorf("CallStack.Slogify() returned incorrect number of attributes. Expected: %d, Got: %d", expCount, len(grp))
	}

	for i, attr := range grp {
		expSubKey := fmt.Sprint(i)
		if attr.Key != expSubKey {
			t.Errorf("CallStack.Slogify() returned incorrect key at index %d. Expected: %s, Got: %s", i, expSubKey, attr.Key)
		}
		if attr.Value.Kind() != slog.KindGroup {
			t.Errorf("CallStack.Slogify() returned incorrect kind at index %d. Expected: %s, Got: %s", i, slog.KindGroup, attr.Value.Kind())
		}
		g := attr.Value.Group()
		if len(g) != 3 {
			t.Errorf("CallStack.Slogify() returned incorrect number of attributes in group at index %d. Expected: %d, Got: %d", i, 3, len(g))
		}
		for _, subAttr := range g {
			if subAttr.Key == "funcName" {
				if subAttr.Value.Kind() != slog.KindString {
					t.Errorf("CallStack.Slogify() returned incorrect kind for funcName at index %d. Expected: %s, Got: %s", i, slog.KindString, subAttr.Value.Kind())
				}
				if subAttr.Value.String() != expFuncNames[i] {
					t.Errorf("CallStack.Slogify() returned incorrect value for funcName at index %d. Expected: %s, Got: %s", i, expFuncNames[i], subAttr.Value.String())
				}
			}
		}
	}
}
