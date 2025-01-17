// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START eventarc_audit_storage_handler]

// Sample audit_storage is a Cloud Run service which handles Cloud Audit Log events with Cloud Storage data.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// HelloEventsStorage receives and processes a Cloud Audit Log event with Cloud Storage data.
func HelloEventsStorage(w http.ResponseWriter, r *http.Request) {
	s := fmt.Sprintf("Detected change in Cloud Storage bucket: subject : %s, source : %s, type :%s, ", string(r.Header.Get("Ce-Subject")), string(r.Header.Get("Ce-Source")), string(r.Header.Get("Ce-Type")))
	log.Printf(s)
	fmt.Fprintln(w, s)
}

// [END eventarc_audit_storage_handler]
// [START eventarc_audit_storage_server]

func main() {
	http.HandleFunc("/", HelloEventsStorage)
	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Start HTTP server.
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// [END eventarc_audit_storage_server]
