package validation

import "regexp"

func IsPostalCodeValid(postalCode string) bool {
	matched, _ := regexp.MatchString(`^\d{8}$`, postalCode)
	return matched
}
