package printUtils

func titleSeperator() string {

	var seperator string
	for i := 0; i < len(hostUser); i++ {
		seperator += "-"
	}
	return seperator
}
