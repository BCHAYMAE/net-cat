package main

import "testing"

//test the port validation function
/*func TestIsValidPort(t *testing.T) { //receives a pointer to go's testing object
									// the name must start with Test, followed by a name beginning with an uppercase letter
    //ports that should be accepted
    validPorts := []string{"8989", "80", "443", "65535"}

    //check each valid port
    for _, port := range validPorts {
        if !isValidPort(port) {
            t.Errorf("expected port %s to be valid", port) //reports formatted failure and continues
        }
    }

    //ports that should be rejected
    invalidPorts := []string{"abc", "0", "-1", "65536", "hello"}

    //check each invalid port
    for _, port := range invalidPorts {
        if isValidPort(port) {
            t.Errorf("expected port %s to be invalid", port)
        }
    }
}*/

// Table-driven tests in Go:
func TestIsValidPort(t *testing.T) {
	//define test cases
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"valid port", "8989", true},
		{"another valid port", "80", true},
		{"letters", "abc", false},
		{"zero", "0", false},
		{"negative port", "-1", false},
		{"port too high", "70000", false},
	}

	//run each test case
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//execute the function
			result := isValidPort(test.input)

			//compare actual and expected results
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// test the username validation function
func TestIsValidUsername(t *testing.T) {
	//valid usernames
	validNames := []string{"Chaymae", "user123", "Oujda"}

	//check valid usernames
	for _, name := range validNames {
		if !isValidUsername(name) {
			t.Errorf("expected username %s to be valid", name)
		}
	}

	//invalid usernames
	invalidNames := []string{"", "	", "ThisUsernameIsWayTooLong"}

	//check invalid usernames
	for _, name := range invalidNames {
		if isValidUsername(name) {
			t.Errorf("expected username %s to be invalid", name)
		}
	}
}
