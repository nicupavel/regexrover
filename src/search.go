// Copyright 2024 Nicu Pavel <npavel@linuxconsulting.ro>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package regexrover

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const googleAPIURL = "https://www.googleapis.com/customsearch/v1"

type SearchResults struct {
	nextCount int
	Links     []string
}

// GoogleSearch returns a SearchResults struct with Links as a list of links found
func GoogleSearch(query string, apiKey string, cx string, startIndex int) (SearchResults, error) {
	// Prepare the request URL with query parameters
	requestURL := fmt.Sprintf("%s?q=%s&key=%s&cx=%s&start=%d", googleAPIURL, url.QueryEscape(query), apiKey, cx, startIndex)
	// Make the HTTP request
	response, err := http.Get(requestURL)
	if err != nil {
		return SearchResults{0, nil}, err
	}
	defer response.Body.Close()

	// Decode the JSON response
	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return SearchResults{0, nil}, err
	}

	// Extract next count of results
	var nextCount int = 0

	if queries, ok := result["queries"].(map[string]interface{}); ok {
		if nextPage, ok := queries["nextPage"].([]interface{}); ok {
			if len(nextPage) > 0 {
				if countObj, ok := nextPage[0].(map[string]interface{}); ok {
					if count, ok := countObj["count"].(int); ok {
						nextCount = count
						//infoLog("Next Count: %d\n", nextCount)
					} else {
						nextCount = 0
					}
				}
			}
		}
	}

	// Extract search results
	var links []string
	items, ok := result["items"].([]interface{})
	if !ok {
		return SearchResults{0, nil}, fmt.Errorf("unable to extract search results")
	}

	for _, item := range items {
		itemMap, ok := item.(map[string]interface{})
		if ok {
			link, linkOk := itemMap["link"].(string)
			if linkOk {
				links = append(links, link)
			}
		}
	}

	return SearchResults{nextCount, links}, nil
}

// SearchAllKeywords take each keyword (can be multi word) and perform GoogleSearch retrieving
// all results that Google Custom Search returns
func SearchAllKeywords(keywords []string, config Config) []string {
	keywordIndex := 0
	startIndex := 0
	var links []string

	for {
		query := url.QueryEscape(keywords[keywordIndex])
		debugLog("Looking for %s query results index %d\n", query, startIndex)
		results, err := GoogleSearch(query, config.ApiKey, config.Cx, startIndex)
		if err != nil {
			errorLog("Error/End of results:", err)
			startIndex = config.MaxSearchResults
		}

		links = append(links, results.Links...)

		startIndex += 10 + 1
		if startIndex >= config.MaxSearchResults {
			keywordIndex++
			startIndex = 0
			if keywordIndex >= len(keywords) {
				break
			}
		}
	}

	return links
}
