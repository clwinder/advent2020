package main

import (
	"reflect"
	"testing"
)

const testInput = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

var testPassports = []passport{
	passport{
		fields: map[string]string{
			"ecl": "gry",
			"pid": "860033327",
			"eyr": "2020",
			"hcl": "#fffffd",
			"byr": "1937",
			"iyr": "2017",
			"cid": "147",
			"hgt": "183cm",
		},
	},
	passport{
		fields: map[string]string{
			"iyr": "2013",
			"ecl": "amb",
			"cid": "350",
			"eyr": "2023",
			"pid": "028048884",
			"hcl": "#cfa07d",
			"byr": "1929",
		},
	},
	passport{
		fields: map[string]string{
			"hcl": "#ae17e1",
			"iyr": "2013",
			"eyr": "2024",
			"ecl": "brn",
			"pid": "760753108",
			"byr": "1931",
			"hgt": "179cm",
		},
	},
	passport{
		fields: map[string]string{
			"hcl": "#cfa07d",
			"eyr": "2025",
			"pid": "166559648",
			"iyr": "2011",
			"ecl": "brn",
			"hgt": "59in",
		},
	},
}

func Test_getPassportsFromString(t *testing.T) {
	actual := getPassportsFromString(testInput)
	if !reflect.DeepEqual(actual, testPassports) {
		t.Errorf("Expected %v, got %v instead", testPassports, actual)
	}
}

func Test_day04(t *testing.T) {
	expected := 2
	actual := day04(testPassports)
	if actual != expected {
		t.Errorf("Expected %d, got %d instead", expected, actual)
	}
}
