package Array  

// Array is a global variable that stores a fixed-size array of strings of length 5.
var FixedArray [5]string

// GetArray returns the current state of the Array variable.
func GetArray() [5]string {
  return FixedArray
}

// UpdateArray updates the values of the Array variable with the provided strings.
// If the number of items passed in is less than 5, the remaining values in the Array will not be modified.
func UpdateArray(items ...string) [5]string {
  for index, value := range items {
    FixedArray[index] = value
  }
  return FixedArray
}