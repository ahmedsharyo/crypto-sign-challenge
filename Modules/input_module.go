package Modules

//make sure the length of 'message' is 250 characters or less
func Check_input_length(input string) bool {

	return len(input) <= 250

}
