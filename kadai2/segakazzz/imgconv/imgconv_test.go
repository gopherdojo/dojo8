package imgconv

import (
	"testing"
)

type pattern struct {
	dirname  string
	input    string
	output   string
	expected *converter
	isError  bool
}

func TestRunConverter(t *testing.T) {
	t.Helper()
	t.Run("Success", func(t *testing.T) {
		dir, input, output := "../testdata", "jpg", "png"
		testRunConverter(t, &dir, &input, &output, false)
	})
	t.Run("Error", func(t *testing.T) {
		dir, input, output := "../testdata", "gif", "png"
		testRunConverter(t, &dir, &input, &output, true)
	})
}

func testRunConverter(t *testing.T, dir *string, input *string, output *string, isError bool) {
	t.Helper()
	if err := RunConverter(dir, input, output); isError && err == nil {
		t.Fatalf("Error is expected but not found")
	} else if !isError && err != nil {
		t.Fatalf("Error is not expected but found %s", err)
	}
}

func TestNewConverter(t *testing.T) {
	successPatterns := []pattern{
		{"testdata", "png", "jpg", &converter{dirname: "testdata", input: "png", output: "jpg"}, false},
		{"testdata", "jpg", "png", &converter{dirname: "testdata", input: "jpg", output: "png"}, false},
	}
	errorPatterns := []pattern{
		{"testdata", "jpg", "gif", nil, true},
		{"testdata", "gif", "jpg", nil, true},
		{"testdata", "jpg", "jpg", nil, true},
	}

	t.Run("successPatterns", func(t *testing.T) {
		for i, p := range successPatterns {
			testNewConverter(t, i, p)
		}
	})

	t.Run("errorPatterns", func(t *testing.T) {
		for i, p := range errorPatterns {
			testNewConverter(t, i, p)
		}
	})
}

func testNewConverter(t *testing.T, i int, p pattern) {
	t.Helper()
	actual, err := newConverter(p.dirname, p.input, p.output)
	if err != nil && p.isError == false {
		t.Fatalf("pattern[%d]: %v want: NO ERROR, actual: %v", i, p, err)
	}
	if err == nil && p.isError == true {
		t.Fatalf("pattern[%d]: %v want: ERROR, actual: NO ERROR", i, p)
	}
	if (actual == nil && p.expected != nil) || (actual != nil && p.expected == nil) {
		t.Fatalf("pattern[%d]: %v want: isNil(%t), actual: isNil(%t)", i, p, p.expected == nil, actual == nil)
	}

	if actual != nil && p.expected != nil && *actual != *p.expected {
		t.Fatalf("pattern[%d]: %v want: %v [%T], actual: %v [%T]", i, p, *p.expected, *p.expected, *actual, *actual)
	}

}

func TestGetSourceFiles(t *testing.T) {
	t.Run("Dir Exists", func(t *testing.T) {
		testGetSourceFiles(t, &converter{dirname: "../testdata/test/4", input: "jpg", output: "png"}, 4, false)
		testGetSourceFiles(t, &converter{dirname: "../testdata/test/2", input: "jpg", output: "png"}, 2, false)
	})
	t.Run("Dir Not Exists", func(t *testing.T) {
		testGetSourceFiles(t, &converter{dirname: "../notfound", input: "jpg", output: "png"}, 0, true)
	})
}

func testGetSourceFiles(t *testing.T, c *converter, expected int, isError bool) {
	t.Helper()
	files, err := c.getSourceFiles()
	if isError && err == nil {
		t.Fatalf("converter[%v] Error is expected. But not found", *c)
	}
	if !isError && err != nil {
		t.Fatalf("Error is not expected but found %s", err)
	}
	if actual := len(files); actual != expected {
		t.Fatalf("converter[%v] File count is not expected. expected: %d actual: %d", *c, expected, actual)
	}
}

func TestConvertSingle(t *testing.T) {
	t.Run("File Found", func(t *testing.T) {
		testConvertSingle(t, &converter{dirname: "../testdata", input: "jpg", output: "png"}, "001.jpg", false)
		testConvertSingle(t, &converter{dirname: "../testdata", input: "png", output: "jpg"}, "001.png", false)
	})
	t.Run("Output folder exists", func(t *testing.T) {
		testConvertSingle(t, &converter{dirname: "../testdata", input: "jpg", output: "png"}, "001.jpg", false)
		testConvertSingle(t, &converter{dirname: "../testdata/test/2", input: "png", output: "jpg"}, "001.png", false)
	})
	t.Run("File Not Found", func(t *testing.T) {
		testConvertSingle(t, &converter{dirname: "../testdata", input: "jpg", output: "png"}, "111.jpg", true)
	})

}

func testConvertSingle(t *testing.T, c *converter, fname string, isError bool) {
	t.Helper()
	err := c.convertSingle(fname)
	if err == nil && isError {
		t.Fatalf("pattern[%v] Error is expected but not found", *c)
	}
	if err != nil && !isError {
		t.Fatalf("pattern[%v] Error is not expected but found %s", *c, err)
	}
}

func TestConvertFiles(t *testing.T) {
	t.Helper()
	t.Run("Success", func(t *testing.T) {
		testConvertFiles(t, &converter{dirname: "../testdata", input: "png", output: "jpg"}, false)
	})
	// t.Run("Error", func(t *testing.T) {
	// 	testConvertFiles(t, &converter{dirname: "../testdata", input: "gif", output: "jpg"}, true)
	// })
}

func testConvertFiles(t *testing.T, c *converter, isError bool) {
	t.Helper()
	files, err := c.getSourceFiles()
	if err != nil {
		t.Fatalf("error found %s", err)
	}
	if err := c.convertFiles(files); isError && err == nil {
		t.Fatalf("pattern[%v] Error is expected but not found", c)
	} else if !isError && err != nil {
		t.Fatalf("pattern[%v] Error is not expected but found %s", c, err)
	}

}
