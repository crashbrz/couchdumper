package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AllDocsResponse struct {
	Rows []struct {
		ID string `json:"id"`
	} `json:"rows"`
}

func main() {
	urlFlag := flag.String("u", "", "Couchdb URL with http(s)://")
	portFlag := flag.String("p", "", "Couchdb port")
	jqFlag := flag.Bool("j", true, "Default enabled for jq integration. Set to false to get verbose output")
	flag.Parse()

	if *urlFlag == "" || *portFlag == "" {
		log.Fatal("URL and port flags are required")
	}

	url := fmt.Sprintf("%s:%s", *urlFlag, *portFlag)

	// Make a request to the endpoint "/"
	resp, err := http.Get(url + "/")
	if err != nil {
		log.Fatal("Failed to make a request to the endpoint '/':", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read the response body:", err)
	}

	if *jqFlag == false {
		fmt.Println("Response from endpoint '/':")
	}
	fmt.Println(string(body))

	// Make a request to the endpoint "/_all_dbs"
	resp, err = http.Get(url + "/_all_dbs")
	if err != nil {
		log.Fatal("Failed to make a request to the endpoint '/_all_dbs':", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read the response body:", err)
	}

	// Parse the JSON response
	var databases []string
	err = json.Unmarshal(body, &databases)
	if err != nil {
		log.Fatal("Failed to parse the JSON response:", err)
	}

	if *jqFlag == false {
		fmt.Println("Response from endpoint '/_all_dbs':")
		fmt.Println(databases)
	}
	// Make requests to each database
	for _, db := range databases {
		resp, err = http.Get(url + "/" + db + "/_all_docs")
		if err != nil {
			log.Printf("Failed to make a request to the endpoint '/%s/_all_docs': %v\n", db, err)
			continue
		}
		defer resp.Body.Close()

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read the response body for database '%s': %v\n", db, err)
			continue
		}

		// Parse the JSON response from _all_docs
		var allDocsResponse AllDocsResponse
		err = json.Unmarshal(body, &allDocsResponse)
		if err != nil {
			log.Printf("Failed to parse the JSON response from '/%s/_all_docs': %v\n", db, err)
			continue
		}

		if *jqFlag == false {
			fmt.Printf("Response from endpoint '/%s/_all_docs':\n", db)
		}
		for _, row := range allDocsResponse.Rows {
			// Make a request for each ID
			resp, err = http.Get(url + "/" + db + "/" + row.ID)
			if err != nil {
				log.Printf("Failed to make a request to the endpoint '/%s/%s': %v\n", db, row.ID, err)
				continue
			}
			defer resp.Body.Close()

			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Failed to read the response body for document '%s' in database '%s': %v\n", row.ID, db, err)
				continue
			}

			if *jqFlag == false {
				fmt.Printf("Response from endpoint '/%s/%s':\n", db, row.ID)
			}
			fmt.Println(string(body))
		}
	}
}
