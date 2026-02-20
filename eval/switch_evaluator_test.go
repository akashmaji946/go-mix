/*
File    : go-mix/eval/switch_evaluator_test.go
Author  : Akash Maji
*/
package eval

import (
	"strings"
	"testing"

	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/std"
)

func TestEvalSwitchStatements(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expected     string
		expectError  bool
		errorMessage string
	}{
		{
			name:     "Basic Switch",
			input:    `var dayNum = 3; var dayName = ""; switch(dayNum) { case 1: dayName = "Mon"; break; case 3: dayName = "Wed"; break; } println(dayName);`,
			expected: "Wed\n",
		},
		{
			name:     "Switch with Default",
			input:    `var grade = "B"; var msg = ""; switch(grade) { case "A": msg="Exc"; break; case "B": msg="Good"; break; default: msg="Inv"; } println(msg);`,
			expected: "Good\n",
		},
		{
			name:     "Switch with Fallthrough",
			input:    `var month = 3; var season = ""; switch(month) { case 3: season="Spring"; case 6: season="Summer"; break; } println(season);`,
			expected: "Summer\n",
		},
		{
			name:     "Switch on Expression",
			input:    `var a = 10; var b = 5; var res = ""; switch(a-b) { case 5: res="Five"; break; } println(res);`,
			expected: "Five\n",
		},
		{
			name:     "Nested Switch",
			input:    `var cat = "fruit"; var item = "apple"; var desc = ""; switch(cat){case "fruit": desc="A fruit. "; switch(item){case "apple": desc+="An apple."; break;} break;} println(desc);`,
			expected: "A fruit. An apple.\n",
		},
		{
			name:     "Switch with No Match",
			input:    `var val = 99; var res = "init"; switch(val){case 1: res="one";} println(res);`,
			expected: "init\n",
		},
		{
			name:     "Switch with Default Only",
			input:    `var val = 1; var res = ""; switch(val){default: res="always";} println(res);`,
			expected: "always\n",
		},
		{
			name:     "Empty Switch",
			input:    `var x=10; var y="ok"; switch(x){} println(y);`,
			expected: "ok\n",
		},
		{
			name:     "Switch on Boolean Logic",
			input:    `var r=true; var e=false; var s=""; switch(r && !e){case true: s="op"; break; case false: s="fail"; break;} println(s);`,
			expected: "op\n",
		},
		{
			name:     "Switch in Function",
			input:    `func getDay(d){switch(d){case 6: return "Sat"; case 7: return "Sun";} return "Wkday";} println(getDay(6));`,
			expected: "Sat\n",
		},
		{
			name:     "Return from Switch in Function",
			input:    `func find(k){switch(k){case "A": return 100; break; case "B": return 200;} return -1;} var f = find; println(f("A"));`,
			expected: "100\n",
		},
		{
			name:     "Switch on Enum",
			input:    `enum C{R,G,B} var l=C.G; var act=""; switch(l){case C.R: act="Stop";break; case C.G: act="Go";break;} println(act);`,
			expected: "Go\n",
		},
		{
			name:     "Switch with Complex Case Expressions",
			input:    `var x=2; var y=3; var r=""; switch(x+y){case 1+1: r="2";break; case 10/2: r="5";break;} println(r);`,
			expected: "5\n",
		},
		{
			name:     "Switch with Dead Code",
			input:    `var v="A"; var r=""; switch(v){case "A": r="A"; break; r="dead";} println(r);`,
			expected: "A\n",
		},
		{
			name: "Switch inside a loop",
			input: `
				var commands = ["start", "process", "stop"];
				var log = "";
				foreach cmd in commands {
					switch (cmd) {
						case "start": log += "s"; break;
						case "process": log += "p"; break;
						case "stop": log += "t"; break;
					}
				}
				println(log);
			`,
			expected: "spt\n",
		},
		{
			name: "Switch with nested if-else",
			input: `
				var role = "admin";
				var perm = false;
				var msg = "";
				switch (role) {
					case "admin":
						if (perm) { msg = "granted"; } else { msg = "denied"; }
						break;
				}
				println(msg);
			`,
			expected: "denied\n",
		},
		{
			name: "Switch modifying outer scope variable",
			input: `
				var counter = 100;
				func process(event) {
					switch (event) {
						case "inc": counter = counter + 1; break;
					}
				}
				process("inc");
				println(counter);
			`,
			expected: "101\n",
		},
		{
			name: "Switch with case expression being a function call",
			input: `
				func getCase() { return 20; }
				var result = "";
				switch (20) {
					case 10: result = "ten"; break;
					case getCase(): result = "twenty"; break;
				}
				println(result);
			`,
			expected: "twenty\n",
		},
		{
			name: "Switch on an array element",
			input: `
				var actions = ["start", "stop"];
				var i = 1;
				var message = "";
				switch (actions[i]) {
					case "start": message = "Starting"; break;
					case "stop": message = "Stopping"; break;
				}
				println(message);
			`,
			expected: "Stopping\n",
		},
		{
			name: "Switch with fallthrough into default",
			input: `
				var val = 2;
				var result = "";
				switch (val) {
					case 1: result = "one"; break;
					case 2: result = "two";
					default: result += " and default";
				}
				println(result);
			`,
			expected: "two and default\n",
		},
		{
			name: "Switch with break inside a nested block",
			input: `
				var val = 1;
				var result = "start";
				switch (val) {
					case 1:
						result = "one";
						{ if (true) { break; } }
						result = "never assigned";
				}
				println(result);
			`,
			expected: "one\n",
		},
		{
			name: "Switch on a character literal",
			input: `
				var char_val = 'b';
				var message = "";
				switch (char_val) {
					case 'a': message = "is A"; break;
					case 'b': message = "is B"; break;
				}
				println(message);
			`,
			expected: "is B\n",
		},
		{
			name: "Switch inside an if statement",
			input: `
				var x = 20;
				var result = "";
				if (x > 10) {
					switch ("hello") {
						case "hello": result = "greeting"; break;
					}
				}
				println(result);
			`,
			expected: "greeting\n",
		},
		{
			name: "Switch with compound assignment in a case",
			input: `
				var x = 100;
				switch (1) {
					case 1: x += 50; break;
				}
				println(x);
			`,
			expected: "150\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := parser.NewParser(tt.input)
			root := p.Parse()
			if p.HasErrors() {
				t.Fatalf("parser errors: %v", p.GetErrors())
			}

			var out strings.Builder
			ev := NewEvaluator()
			ev.SetParser(p)
			ev.SetWriter(&out)

			result := ev.Eval(root)

			if tt.expectError {
				if result == nil || result.GetType() != std.ErrorType {
					t.Fatalf("expected an error, but got %T", result)
				}
				if !strings.Contains(result.ToString(), tt.errorMessage) {
					t.Errorf("expected error message to contain '%s', got '%s'", tt.errorMessage, result.ToString())
				}
			} else {
				if result != nil && result.GetType() == std.ErrorType {
					t.Fatalf("unexpected error: %s", result.ToString())
				}
				if out.String() != tt.expected {
					t.Errorf("wrong output. expected=%q, got=%q", tt.expected, out.String())
				}
			}
		})
	}
}
