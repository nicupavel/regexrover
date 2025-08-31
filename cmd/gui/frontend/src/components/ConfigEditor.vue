<script setup>
import { ref } from "vue";
import { storeToRefs } from "pinia";
import { useSettingsStore } from "@/stores/settings";

import ConfigItem from "@/components/ConfigItem.vue";

const { settings } = storeToRefs(useSettingsStore());
const { settingsHelp } = useSettingsStore();

const selectedCrawlMode = ref(2);

const crawlModes = ref([
  { title: "Google Search", value: 1 },
  { title: "List of URLs", value: 2 },
]);

defineExpose({ settings, selectedCrawlMode });

const settingsLayout = {
  ApiKey: {
    groupUnder: 1,
    advanced: false,
  },
  Cx: {
    groupUnder: 1,
    advanced: false,
  },
  OutputDriver: {
    advanced: true,
    show: true,
  },
  KeywordsFile: {},
  CrawlIgnoreDomains: {
    advanced: true,
    show: true,
  },
  CrawlAllowedUrlsRegex: {
    advanced: true,
    show: true,
  },
  CrawlUserAgent: {
    advanced: true,
    show: true,
  },
  CrawlTag: {
    advanced: false,
    show: true,
  },
  CrawlMatchRegex: {
    advanced: false,
    show: true,
  },
  CrawlUrlsFile: {},
  CrawlCacheDir: {
    advanced: true,
    show: true,
  },
  CrawlDepth: {
    advanced: true,
    show: true,
  },
  CrawlThreads: {
    advanced: true,
    show: true,
  },
  CrawlLog: {
    advanced: false,
    show: false,
  },
  MatchOutputChunks: {
    advanced: true,
    show: true,
  },
  MaxSearchResults: {
    groupUnder: 1,
    advanced: false,
  },
  // Entries below are added by the ConfigEditor, don't show them
  urlList: {},
  keywords: {},
  crawlMode: {},
};
</script>
<template>
  <div class="surface-section">
    <ul class="list-none p-0 m-0">
      <li class="flex align-items-center py-3 px-2 border-top-1 surface-border">
        <div class="flex flex-column align-items-flexstart flex-wrap w-4">
          <div class="text-700 w-12 md:w-4 font-medium font-bold">
            Crawler Mode
          </div>
        </div>
        <div
          class="flex text-900 w-8 md:w-8 flex-column align-items-flexstart flex-wrap"
        >
          <SelectButton
            v-model="selectedCrawlMode"
            :options="crawlModes"
            optionLabel="title"
            optionValue="value"
            aria-labelledby="basic"
          />
          <template v-if="selectedCrawlMode == 1">
            <Textarea v-model="settings.keywords" rows="5" cols="30" />
            <div class="text-400 text-sm w-10 mt-2">
              Will search each keyword on google and crawl URLs from the
              results. One keyword per line.
            </div>
            <template v-for="(_, label) in settings">
              <ConfigItem
                v-if="settingsLayout[label].groupUnder == 1"
                v-model="settings[label]"
                :label="label"
                :help="settingsHelp[label]"
              />
            </template>
          </template>
          <template v-else>
            <Textarea
              v-if="selectedCrawlMode == 2"
              v-model="settings.urlList"
              rows="5"
              cols="30"
            />
            <div class="text-400 text-sm w-10 mt-2">
              Will crawl only in this list of URLs. One URL per line.
            </div>
          </template>
        </div>
      </li>
      <template v-for="(_, label) in settings">
        <ConfigItem
          v-if="settingsLayout[label].show && !settingsLayout[label].advanced"
          v-model="settings[label]"
          :label="label"
          :help="settingsHelp[label]"
        />
      </template>
      <Fieldset legend="Advanced options" toggleable :collapsed="true">
        <template v-for="(_, label) in settings">
          <ConfigItem
            v-if="settingsLayout[label].show && settingsLayout[label].advanced"
            v-model="settings[label]"
            :label="label"
            :help="settingsHelp[label]"
          />
        </template>
      </Fieldset>
    </ul>
  </div>
</template>
