package iteration

func repeat(repeatValue string, numberOfRepeat int) string {
	var repeated string
	for i := 0; i < numberOfRepeat; i++ {
		repeated += repeatValue
	}
	return repeated
}
