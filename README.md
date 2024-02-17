# RegexRover

Crawl URLs, find matches with regex and save results

## Features

- Can start URL crawling from Google Custom Search results using a file with keywords or by using a file with a list of URLs
- Can define regex to skip URL, Paths or query
- Can define regex to match content
- Can define DOM elements on which the regex to be applied
- Outputs to CSV file with caching and deduplication of matches
- Can use cache to prevent revisiting URLs that had already been crawled in a different run

## Using

`git clone https://github.com/nicupavel/regexrover && cd regexrover`

`cp .env.default .env`

`go run .`

## Running modes

1. **Google Custom Search** - will search for keywords from a file and crawl the URLs from the search results
2. **List of URLs** - will crawl a list of URLs from a file

### 1. Google Custom Search

- Create a [Programable Search Engine](https://programmablesearchengine.google.com/u/1/controlpanel/create)
- Copy the Search Engine Code / ID to `GOOGLE_SEARCH_ID` in `.env` file
- Get an API key from [Google Search JSON API](https://developers.google.com/custom-search/v1/introduction)
- Put this key in `GOOGLE_SEARCH_API_KEY` in `.env` file
- Create a file with keywords (can be multiple words per line) and put the file name in `KEYWORDS_FILE` in `.env` file
- go run

Note: All search results from google will be saved to a file named with `<keywords>_search_links.txt`. This file can be used
later in the mode below.

### 2. List of URLs (default mode)
- Create a file with a list of URLs (1 URL per line) and put the file name in `CRAWL_URLS_FILE` in `.env` file
- go run

Note: Both modes will output a CSV file named `found_matches_<run_date_time>.csv`

## .env config

1. (optional) `GOOGLE_SEARCH_ID` ID of the Google Programable Engine 
2. (optional) `GOOGLE_SEARCH_API_KEY` API Key for the Google Custom Search v1 
3. (optional) `KEYWORDS_FILE` the file that has your keywords to search on each line will do a Google search. Can be multiple words separated by space per line
4. `MATCH_OUTPUT_CHUNKS` Optimize file writing and deduplicate matches. After how many matches the results are saved to CVS file. *Default: 5*
5. `CRAWL_CACHE_DIR` Directory to store the cache for the crawler. Crawler won't visit cached pages on another run. If empty it won't keep a cache
6. `CRAWL_DEPTH` How many levels deep the crawler should go from the page URL obtained from Google or the file with links. *Default: 1*
7. `CRAWL_THREADS` How many threads the crawler should use. *Default: 20*
8. `CRAWL_IGNORE_DOMAINS` A list separated by , of domains to ignore in crawling
9. `CRAWL_ALLOWED_URLS_REGEX` URLs matching this regex will be crawled, can be used to select certain tld or ignore paths with query strings
10. `CRAWL_USER_AGENT` Browser User Agent to use
11. `CRAWL_TAG` Regex matching will be performed in all DOM elements with this tag. For example `body` will have the regex applied to all body content.
12. `CRAWL_MATCH_REGEX` The regex to match content. Text matching this regex will be saved in the output file along with the URL
13. `CRAWL_URLS_FILE` The file with the list of URLs to start crawling. If this is defined Google Search mode will be **ignored**



