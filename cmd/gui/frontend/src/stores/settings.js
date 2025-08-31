import { ref } from 'vue';
import { defineStore } from 'pinia';

import { GetConfig } from "../../wailsjs/go/main/App";

const config = ref(GetConfig());

export const useSettingsStore = defineStore('regexroverSettings', () => {

        const settings = ref({
            ApiKey: "",
            Cx: "",
            OutputDriver: "sqlite",
            CrawlIgnoreDomains: [
                "facebook.com",
                "twitter.com",
            ],
            CrawlAllowedUrlsRegex: "\\bhttps?://(?:www\\.)?[a-zA-Z0-9.-]+\\.ro(?:/[a-zA-Z0-9.-]+(?:/[a-zA-Z0-9.-]+)?)?b|https?://(?:www\\.)?[a-zA-Z0-9.-]+\\.ro\\b",
            CrawlUserAgent: "Mozilla/5.0 (X11; Linux i686; rv:122.0) Gecko/20100101 Firefox/122.0",
            CrawlTag: "body",
            CrawlMatchRegex: "[a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\\.[a-zA-Z0-9_-]+",
            CrawlCacheDir: "./regexrover_cache",
            CrawlDepth: 1,
            CrawlThreads: 20,
            CrawlLog: true,
            MatchOutputChunks: 50,
            MaxSearchResults: 99
        });

        // prettier-ignore
        const settingsHelp = {
            ApiKey: "API Key for the Google Custom Search v1",
            Cx: "ID of the Google Programable Engine (optional)",
            OutputDriver: "Format of the file with results (sqlite or csv)",
            CrawlIgnoreDomains: "A list separated by , of domains to ignore in crawling",
            CrawlAllowedUrlsRegex: "URLs matching this regex will be crawled, can be used to select certain tld or ignore paths with query strings",
            CrawlUserAgent: "Browser User Agent to use",
            CrawlTag: "Matching will be performed in all DOM elements with this tag. For example `body` will have the regex applied to all body content",
            CrawlMatchRegex: "The regex to match content. Text matching this regex will be saved in the output file along with the URL",
            CrawlUrlsFile: "",
            CrawlCacheDir: "Directory to store the cache for the crawler. Crawler won't visit cached pages on another run. If empty it won't keep a cache",
            CrawlDepth: "How many levels deep the crawler should go from the page URL obtained from Google or the file with links. Use 0 for infinite recursion. *Default: 1*",
            CrawlThreads: "How many threads the crawler should use. *Default: 20*",
            CrawlLog: "",
            MatchOutputChunks: "Optimize file writing and deduplicate matches. After how many matches the results are saved to disk",
            MaxSearchResults: "Maximum number of results returned per keyword. Google search has a maximum of 99 results per search",
        };
  

        async function importConfigFromEnv() {
            settings.value = await GetConfig();
        }

        //----------------------------------------------------------------------
        return {
            settings,
            settingsHelp,
            importConfigFromEnv,
        };
    },
    {
        persist: {
            storage: localStorage,
            paths: ['settings'],
        },
    },
); 