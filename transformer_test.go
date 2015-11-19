package jsonresponse

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestPassthroghTransformer(t *testing.T) {
	for _, r := range []Response{
		Empty(),
		New("foo"),
		New(map[string]string{"foo": "bar"}),
	} {
		//	r := Empty()
		headers, result := PassthroughTransformer(r, http.StatusOK)
		if len(headers) != 0 {
			fmt.Println("Passthrough transformer should return no headers.")
			t.Fail()
		}
		if !reflect.DeepEqual(result, r.Data) {
			fmt.Println("Passthrough transformer did not return same data as in response. ")
			t.Fail()
		}
	}
}

func TestMessageCodeTransformer(t *testing.T) {
	status := http.StatusOK
	for dataField, codeField := range map[string]string{
		"data": "code",
	} {
		transformer := MessageCodeTransformer(dataField, codeField)

		for _, r := range []Response{
			Empty(),
			New("foo"),
			New(map[string]string{"foo": "bar"}),
		} {
			headers, res := transformer(r, status)
			result := res.(map[string]interface{})
			if len(headers) != 0 {
				fmt.Println("MessageCodeTransformer should return not headers.")
			}
			if _, ok := result[dataField]; !ok {
				fmt.Printf("Data field did not match: %s\n", dataField)
				t.Fail()
			}
			if val, ok := result[dataField]; ok {
				if !reflect.DeepEqual(val, r.Data) {
					fmt.Println("Value did not match.")
					t.Fail()
				}
			}
			if _, ok := result[codeField]; !ok {
				fmt.Printf("Code field not found in response: %s", codeField)
				t.Fail()
			}
			if val, ok := result[codeField]; ok {
				if val != status {
					fmt.Println("Code did not match!")
					t.Fail()
				}
			}
		}
	}
}


func TestMessageCodeExcuseTransformer(t *testing.T) {
	status := http.StatusOK
	for dataField, codeField := range map[string]string{
		"data": "code",
	} {
		transformer := MessageCodeExcuseTransformer(dataField, codeField)

		for _, r := range []Response{
			Empty().WithProgrammingExcuse(),
			New("foo").WithProgrammingExcuse(),
			New(map[string]string{"foo": "bar"}).WithProgrammingExcuse(),
		} {
			headers, res := transformer(r, status)
			result := res.(map[string]interface{})
			if len(headers) != 0 {
				fmt.Println("MessageCodeTransformer should return not headers.")
			}
			if _, ok := result[dataField]; !ok {
				fmt.Printf("Data field did not match: %s\n", dataField)
				t.Fail()
			}
			if val, ok := result[dataField]; ok {
				if !reflect.DeepEqual(val, r.Data) {
					fmt.Println("Value did not match.")
					t.Fail()
				}
			}

			if _, ok := result[codeField]; !ok {
				fmt.Printf("Code field not found in response: %s", codeField)
				t.Fail()
			}
			if val, ok := result[codeField]; ok {
				if val != status {
					fmt.Println("Code did not match!")
					t.Fail()
				}
			}

			if _, ok := result["programming-excuse"]; !ok {
				fmt.Println("programming-excuse not found in response.")
				t.Fail()
			}
			if val, ok := result["programming-excuse"]; ok {
				if val == "" {
					fmt.Println("Programming excuse not set.")
					t.Fail()
				}
			}
		}
	}
}
