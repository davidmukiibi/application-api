package main

import (
	"testing"
	"net/url"
	"fmt"

	"github.com/davidmukiibi/controllers"
	"github.com/stretchr/testify/assert"
)

// ======================================================================
//                              Benchmarks
// ======================================================================

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, eachguy := range testCases {
			fmt.Println(eachguy.FirstName)
		}
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, name := range testCasesSlice {
			fmt.Println(string(name[:5]))
		}
	}
}

// ========================================================================
//                                 Tests
// ========================================================================

// testing healthcheck endpoint
func TestDefault(t *testing.T) {
	fmt.Println(assert.HTTPSuccess(t, controllers.Default, "GET", "localhost/", nil))
}

// testing user create endpoint
func TestCreateEndPoint(t *testing.T) {
	for _, testCase := range testCases {
		q := url.Values{}
		// q.Set("user_email", testCase.UserEmail)
		q.Add("password", testCase.Password)
		q.Add("first_name", testCase.FirstName)
		q.Add("surname", testCase.Surname)
		fmt.Println(assert.HTTPBodyContains(t, controllers.CreateEndPoint, "POST", "localhost/signup", q, "New user created successfully"))
	}
}

