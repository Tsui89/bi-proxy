package qc

import "testing"

func TestQCAuth_Authentication(t *testing.T) {
	var auth QCAuth
	auth.Signature = []byte("p6kmQ4o5sQ_oQ8SNAVOUuVvGYHomGPqZ20vsdBG5kJ4")
	auth.Payload = []byte("eyJsYW5nIjoiemgtY24iLCJhY2Nlc3NfdG9rZW4iOiJ0eWc1OVNGVjhRcGFpWjJBeWZpcVJ2VXFFV05UN2VMRSIsInVzZXJfaWQiOiJ1c3ItS1RWWXMzY2ciLCJ6b25lIjoicG9jMSIsImFjdGlvbiI6InZpZXdfYXBwIiwidGltZV9zdGFtcCI6IjIwMTgtMDEtMTJUMDg6Mzc6MzlaIiwiZXhwaXJlcyI6IjIwMTgtMDEtMTJUMTQ6Mzc6MzlaIn0")
	auth.SecretKey = []byte("cH6xY5i9YpkGNDXVQsyDDh9z9L5bZC8W7KvwTogA")
	if (!auth.Authentication()){
		t.Fatal()
	}
}
