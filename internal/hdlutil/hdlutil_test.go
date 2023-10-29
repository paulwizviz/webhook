package hdlutil

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidatePostMethod(t *testing.T) {
	testcases := []struct {
		Method      string
		Want        int
		description string
	}{
		{
			Method:      http.MethodPost,
			Want:        http.StatusAccepted,
			description: "Method POST",
		},
		{
			Method:      http.MethodGet,
			Want:        http.StatusMethodNotAllowed,
			description: "Method Get",
		},
	}

	for i, tc := range testcases {

		var endHDL validateFunc = func(rw http.ResponseWriter, r *http.Request) {
			rw.WriteHeader(http.StatusAccepted)
		}
		wrapphdl := ValidatePostMethod(endHDL)

		req := httptest.NewRequest(tc.Method, "/", nil)
		rw := httptest.NewRecorder()
		wrapphdl.ServeHTTP(rw, req)
		resp := rw.Result()
		if tc.Want != resp.StatusCode {
			t.Fatalf("Want: %v Got: %v Description: %s", tc.Want, resp.StatusCode, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
		}
	}
}

func TestMethodPostValidate(t *testing.T) {
	testcases := []struct {
		Method      string
		Want        int
		description string
	}{
		{
			Method:      http.MethodPost,
			Want:        http.StatusAccepted,
			description: "Method POST",
		},
		{
			Method:      http.MethodGet,
			Want:        http.StatusMethodNotAllowed,
			description: "Method Get",
		},
	}

	for i, tc := range testcases {

		endHDL := func(rw http.ResponseWriter, r *http.Request) {
			rw.WriteHeader(http.StatusAccepted)
		}
		wrapphdl := MethodPostValidate(endHDL)

		req := httptest.NewRequest(tc.Method, "/", nil)
		rw := httptest.NewRecorder()
		wrapphdl(rw, req)
		resp := rw.Result()
		if tc.Want != resp.StatusCode {
			t.Fatalf("Want: %v Got: %v Description: %s", tc.Want, resp.StatusCode, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
		}
	}
}

func TestValidateGetMethod(t *testing.T) {
	testcases := []struct {
		Method      string
		Want        int
		description string
	}{
		{
			Method:      http.MethodPost,
			Want:        http.StatusMethodNotAllowed,
			description: "Method POST",
		},
		{
			Method:      http.MethodGet,
			Want:        http.StatusAccepted,
			description: "Method Get",
		},
	}

	for i, tc := range testcases {

		var endHDL validateFunc = func(rw http.ResponseWriter, r *http.Request) {
			rw.WriteHeader(http.StatusAccepted)
		}
		wrapphdl := ValidateGetMethod(endHDL)

		req := httptest.NewRequest(tc.Method, "/", nil)
		rw := httptest.NewRecorder()
		wrapphdl.ServeHTTP(rw, req)
		resp := rw.Result()
		if tc.Want != resp.StatusCode {
			t.Fatalf("Want: %v Got: %v Description: %s", tc.Want, resp.StatusCode, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
		}
	}
}

func TestMethodGetValidate(t *testing.T) {
	testcases := []struct {
		Method      string
		Want        int
		description string
	}{
		{
			Method:      http.MethodPost,
			Want:        http.StatusMethodNotAllowed,
			description: "Method POST",
		},
		{
			Method:      http.MethodGet,
			Want:        http.StatusAccepted,
			description: "Method Get",
		},
	}

	for i, tc := range testcases {

		endHDL := func(rw http.ResponseWriter, r *http.Request) {
			rw.WriteHeader(http.StatusAccepted)
		}
		wrapphdl := MethodGetValidate(endHDL)

		req := httptest.NewRequest(tc.Method, "/", nil)
		rw := httptest.NewRecorder()
		wrapphdl(rw, req)
		resp := rw.Result()
		if tc.Want != resp.StatusCode {
			t.Fatalf("Want: %v Got: %v Description: %s", tc.Want, resp.StatusCode, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
		}
	}
}
