package main

func removeOuterParentheses(S string) string {
	start := 0
	sk := make([]byte, 0)
	output := make([]byte, 0)
	// loop over input
	for i := 0; i < len(S); i++ {
		c := S[i]
		output = append(output, c)

		if len(sk) != 0 {
			if sk[len(sk)-1] == '(' && c == ')' {
				// pop stack
				sk = sk[:len(sk)-1]
			} else {
				// push stack
				sk = append(sk, c)
			}
		} else {
			// push stack, if empty
			sk = append(sk, c)
		}

		if len(sk) == 0 {
			// remove first byte (offset saved with start)
			copy(output[start:], output[start+1:]) // Shift a[i+1:] left one index.
			output[len(output)-1] = byte(0)        // Erase last element (write zero value).
			output = output[:len(output)-1]        // Truncate slice.

			// remove last byte
			output = output[:len(output)-1]
			start = len(output)
		}

	}
	return string(output)
}
