/*
 * Copyright (C) 2019 Intel Corporation
 * SPDX-License-Identifier: BSD-3-Clause
 */
package validation

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEnvList(t *testing.T) {
	a := assert.New(t)

	os.Setenv("GOENV_TEST1", "NOT IMPORTATN CONTENT")
	os.Setenv("GOENV_TEST2", "NOT IMPORTATN CONTENT")
	os.Setenv("GOENV_TEST3", "NOT IMPORTATN CONTENT")

	needed_env := []string{"GOENV_TEST1", "GOENV_TEST2", "GOENV_TEST3"}
	missing, err := ValidateEnvList(needed_env)

	a.Equal(nil, err)
	a.Equal(0, len(missing))

	more_than_needed_env := []string{"GOENV_TEST2", "GOENV_TEST3", "GOENV_TEST4", "GOENV_TEST5"}

	missing1, err1 := ValidateEnvList(more_than_needed_env)
	sample_err := errors.New("One or more required environment variables are missing")

	a.Equal(sample_err, err1)

	// should show [GOENV_TEST4:0 GOENV_TEST5:0]
	// run with with go test -v
	fmt.Println(missing1)
}

func TestValidateURL(t *testing.T) {
	a := assert.New(t)

	good_url_1 := "https://google.com"
	good_url_2 := "http://good.url.with.port:5566"
	good_url_3 := "https://good.url.https.with.port:5566"

	good_tests := []string{good_url_1, good_url_2, good_url_3}

	protocols := make(map[string]byte)
	protocols["http"] = 0
	protocols["https"] = 0

	for _, good_str := range good_tests {

		err := ValidateURL(good_str+"/tds/", protocols, "/tds/")
		a.Equal(nil, err)
	}

	bad_url_1 := "bad.url.without.protocol/tds/"
	bad_url_2 := "scheme://bad.url.with.wrong.protocol/tds/"
	bad_url_3 := "https://bad.url.with.path/tds/path/path/"
	bad_url_4 := "https://bad.url.with.query/tds/?query=haha"

	// err_1 := errors.New("Invalid base URL")
	err_2 := errors.New("Unsupported protocol")
	err_3 := errors.New("Invalid path in URL")
	err_4 := errors.New("Unexpected inputs")

	bad_tests := []string{bad_url_1, bad_url_2, bad_url_3, bad_url_4}
	bad_result := []error{err_2, err_2, err_3, err_4}

	for i, bad_str := range bad_tests {

		err := ValidateURL(bad_str, protocols, "/tds/")
		a.Equal(bad_result[i], err)
	}
}

func TestValidateAccount(t *testing.T) {
	// a := assert.New(t)

	// good_uname_1 := "uname.has-symbol"
	// good_uname_2 := "abcd1234"
	// good_uname_3 := "a1a2_3d4f5g6_7j8k9l1_2s3d45g6h7"

	// good_pwd_1 := "easy_guess123"
	// good_pwd_2 := "hardone_A-Za-z0-9#?!@$%^&*-"
	// good_pwd_3 := "tooshort"

	// good_uname_tests := []string{good_uname_1, good_uname_2, good_uname_3}
	// good_pwd_tests := []string{good_pwd_1, good_pwd_2, good_pwd_3}

	// for _, uname := range good_uname_tests {

	// 	for _, pwd := range good_pwd_tests {
	// 		// fmt.Println(uname + " " + pwd)
	// 		err := ValidateAccount(uname, pwd)
	// 		a.Equal(nil, err)
	// 	}
	// }

	// bad_uname_1 := "fishy_symbols \" ` '"
	// bad_uname_2 := "\" \" \" UNION SELECT * FROM"
	// bad_uname_3 := "))) OR TRUE"
	// bad_uname_4 := "12_number_start_with"

	// bad_pwd_1 := "fishy_symbols \" ` '"
	// bad_pwd_2 := "\" \" \" UNION SELECT * FROM"
	// bad_pwd_3 := "))) OR TRUE"
	// bad_pwd_4 := "tooshrt"
	// bad_pwd_5 := ""

	// bad_ret := errors.New("Invalid input for username or password")

	// bad_uname_tests := []string{bad_uname_1, bad_uname_2, bad_uname_3, bad_uname_4}
	// bad_pwd_tests := []string{bad_pwd_1, bad_pwd_2, bad_pwd_3, bad_pwd_4, bad_pwd_5}

	// for _, bad_uname := range bad_uname_tests {
	// 	fmt.Println(bad_uname + " " + good_pwd_1)
	// 	err := ValidateAccount(bad_uname, good_pwd_1)
	// 	a.Equal(bad_ret, err)
	// }
	// for _, bad_pwd := range bad_pwd_tests {
	// 	err := ValidateAccount(good_uname_1, bad_pwd)
	// 	a.Equal(bad_ret, err)
	// }
}

func TestValidateStrings(t *testing.T) {
	goodString1 := "/var/lib/nova/instances/"
	goodString2 := "abcd1234"
	goodString3 := "workload-agent"
	goodString4 := "samplestring"
	goodString5 := ""

	goodStringValueArr := []string{goodString1, goodString2, goodString3, goodString4, goodString5}
	err := ValidateStrings(goodStringValueArr)
	assert.NoError(t, err)

	badString1 := "fishy_symbols \" ` ' )))((*&^ "
	badString2 := "\" \" \" SELECT * FROM TABLE"
	badString3 := "<inputTag>"

	badStringValueArr := []string{badString1, badString2, badString3}
	err = ValidateStrings(badStringValueArr)
	assert.Error(t, err)
}

func TestValidatePemEncodedKey(t *testing.T) {

	goodString1 := "-----BEGIN CERTIFICATE----MIIEoDCCA4igAwIBAgIIHZR9rDPTS9IwDQYJKoZIhvcNAQELBQAwGzEZMBcGA1UEAxMQbXR3aWxzb24tcGNhLWFpazAeFw0xOTAzMDgwOTQ0MDJaFw0yOTAzMDUwOTQ0MDJaMCUxIzAhBgNVBAMMGkNOPUJpbmRpbmdfS2V5X0NlcnRpZmljYXRlMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvarJulWowk4Mr0CwJ8SMrGNvbHKryZmTgLFftnj7bwd7VUx0zACjwv0jObFAonocBMEO39hxhHeQWuyUB1OEuvt3Kzg40EG11PvsHttWM2rgYk88z+E7vsQrdHx08FLeN4T+SK9ML+uKDPFuQWLrzZ2irQxGznokn2aAz+zl8vIDkVjmVDw4I9D6/6hMAEaE4WGiznoPS1sEe6vyqPhOL9dQOrFzW5qC6JKNwPtuYIWg5VrhrxguTHe5vK7TpZxgbxjWbeSdiDAnNU7UdZM5QboJ8Ao7Oz54OtIqpUuIEXSzIaJ/Dbo0yXrYCJzcSDTpSDimitqfTiot9p6+pQSM8wIDAQABo4IB3DCCAdgwDgYDVR0PAQH/BAQDAgUgMIGdBgdVBIEFAwIpBIGR/1RDR4AXACIAC0y1Y5XeuPW1DaT1Y11paGPfMQ3JGXM61ail1pp716pUAAQA/1WqAAAAAAgwHmMAAAAGAAAAAQEABwA+AAw2AAAiAAvKQw0fWfgw6nQE4MTsGY7tTwepogu1YED+3ZSs+rkBbgAiAAvZHMHbKCiUnL7jxkTMqhbMNtoyKqi7eg5Tm7eQFQM3+jCCARQGCFUEgQUDAikBBIIBBgAUAAsBADiGza+SuXvKAkfBWbYLXXmneXfAUaUtynArnhey7icOvBNqNVVROXx9iaH5wxgmzshJRDXIkJhYpZ8xfNkFqfPkAtHSw3IugrTe4bokzVwVMe+a0c+OT9ptDEkb7TiUKNc0hX2O0jc4dPSzYBdAn/3HSnV1f5DQBKT/QTk/Io+7F0Vu2phz96Ax7d73mjufehE920hq77v/mVXNOtmyQ6Q9OXQCgDwgfAQwBc6iuBLsXPAK1GUN2D/8N29+CQQfa0/KzHrHc/Ioh+PTOgg7TaYmJhB+NasdzV7t3YlAa0x6ZPPZwagpSnbrGpnIw2H0Wvy/YebfvGskymvvsiSwUGQwDgYIVQSBBQMCKQIEAgAAMA0GCSqGSIb3DQEBCwUAA4IBAQBh6LiD+zec5Tp0qZnpNniw519/JHOVN3HDcF1mv/BKSEeu7BmWp62c3Agf6+8anWsrpTg936DAgCKUgjJN2+m5fdJhNOUs662/PlLE17FjeAYMIxfNlHlP5nNRq5F6T8I1BDaCsT4dgmQlqUkCrHRaPB9nqajPOjYihYpSxNnUT/b4NFnK8T7gEvrfGM1EF6m1dPx4IQSxznBD4XwjBU/KjjVEuPFks/bPwf/sz5t3itnpJTFazrnBp4Wr2kzxKLoPDKmhtgStE0mArpLWnEbgMXGmkm5mgzHCFs5phEghhkAHfyT/lLjMKzzkN7BmjL6qM9TIA67g8LIMuI/wyKXe-----END CERTIFICATE-----"
	goodString2 := "tNPEL9e3O97nQ2nTlzwyPxxNLAjwwSvTnGcVeOmIfbboqBZkixDx3TR6/n13s7eury6LVqOjEOsItExy/BvXE+g6WTP61dEnenM1fNlmNY/O6bNqEwgxVEfamdNBN9FYuB2inzwueEiB5J2WtkCB+0c2pr86oMl8eYlUve1r0+VRbOhQVhMKskVLKZonBxvLrup1geqIDD8DuK/EVt8aCYfjY5n0B1NtfSwQpNzWjmTc0g9H0aH/tYerfB6Q+KqVo2be1wWQscNBfVvVGjwB3brT6k2ODUSCDlGJV5IhVNq+ooM5s+o21c8+fX1bXiFOdA7T/TnuCKJyEMxrAvzypQ=="
	
	err := ValidatePemEncodedKey(goodString1)
	assert.NoError(t, err)

	err = ValidatePemEncodedKey(goodString2)
	assert.NoError(t, err)

	badString1 := "fishy_symbols \" ` ' )))((*&^ "
	badString2 := "\" \" \" SELECT * FROM TABLE"
	badString3 := "<inputTag>"

	err = ValidatePemEncodedKey(badString1)
	assert.Error(t, err)

	err = ValidatePemEncodedKey(badString2)
	assert.Error(t, err)
	
	err = ValidatePemEncodedKey(badString3)
	assert.Error(t, err)
}