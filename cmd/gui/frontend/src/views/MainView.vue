<script setup>
import { ref, onMounted } from "vue";
import ConfigEditor from "@/components/ConfigEditor.vue";
import {
  RunCrawler,
  StopCrawler,
  GetOutputFileName,
} from "../../wailsjs/go/main/App";

const running = ref(false);
const error = ref(null);
const status = ref({ severity: "info", text: "idle" });
const outputFileName = ref();
const runningLog = ref([]);
const runningStats = ref();
const configEditorComponent = ref(null);
const configEditorCollapsed = ref(false);
const scrollPanelEl = ref();
let statsInterval;

async function startCrawling() {
  running.value = true;
  error.value = null;
  runningLog.value = [];
  status.value = { severity: "info", text: "idle" };
  configEditorCollapsed.value = true;
  outputFileName.value = null;

  try {
    const crawlSettings = configEditorComponent.value.settings;
    const crawlMode = configEditorComponent.value.selectedCrawlMode;
    let crawlList;

    if (crawlMode == 1) {
      // Keywords search in Google
      crawlList = configEditorComponent.value.settings.keywords.split("\n");
    } else if (crawlMode == 2) {
      // List of URLs check each line to have http(s)
      crawlList = configEditorComponent.value.settings.urlList.split("\n");
      for (let i = 0; i < crawlList.length; i++) {
        if (!crawlList[i].startsWith("http")) {
          crawlList[i] = `https://${crawlList[i]}`;
        }
      }
      // Show changed values in the textarea element
      configEditorComponent.value.settings.urlList = crawlList.join("\n");
    }

    console.log(crawlSettings);
    console.log(crawlList);
    console.log(crawlMode);

    await RunCrawler(crawlSettings, crawlList, crawlMode);
    outputFileName.value = await GetOutputFileName();
  } catch (e) {
    error.value = e.toString();
    status.value = { severity: "error", text: "Error" };
  } finally {
    running.value = false;
    status.value = { severity: "info", text: "idle" };
  }
}

async function stopCrawling() {
  try {
    console.log("Stopping crawler.");
    await StopCrawler();
    running.value = false;
  } catch (e) {
    error.value = e.toString();
  }
}

onMounted(() => {
  window.runtime.EventsOn("ErrorEvent", async (e) => {
    error.value = e;
    status.value = { severity: "error", text: "Error" };
    console.error("Received error from backend", e);
  });
  window.runtime.EventsOn("StatusEvent", async (s) => {
    runningLog.value.push(s);
    scrollPanelEl.value.scrollTop = scrollPanelEl.value.scrollHeight;
    status.value = { severity: "success", text: "Running" };
  });
  window.runtime.EventsOn("StateEvent", async (s) => {
    runningStats.value = s;
  });
});
</script>
<template>
  <main>
    <Fieldset
      legend="Crawler Configuration"
      toggleable
      :collapsed="configEditorCollapsed"
    >
      <ConfigEditor ref="configEditorComponent" />
    </Fieldset>
    <br />
    <div class="flex flex-column">
      <div v-if="error" class="p-2">
        {{ error }}
      </div>
      <div class="p-2 flex flex-row gap-2">
        <div>
          <Button
            :severity="status.severity"
            outlined
            :loading="running"
            :label="` Status: ${status.text}`"
          />
        </div>
        <div>
          <Button
            v-if="!running"
            label="Start Crawling"
            @click="startCrawling()"
            icon="pi pi-search"
          />
          <Button
            v-if="running"
            label="Stop"
            @click="stopCrawling()"
            icon="pi pi-stop-circle"
          />
        </div>
        <div class="flex flex-row gap-2 ml-3" v-if="runningStats">
          <div
            class="flex flex-column justify-content-center align-items-center"
          >
            <div>Scraped</div>
            <div>{{ runningStats.Scraped }}</div>
          </div>
          <div
            class="flex flex-column justify-content-center align-items-center"
          >
            <div>Requests</div>
            <div>{{ runningStats.Requests }}</div>
          </div>
          <div
            class="flex flex-column justify-content-center align-items-center"
          >
            <div>Pending</div>
            <div>{{ runningStats.Pending }}</div>
          </div>
          <div
            class="flex flex-column justify-content-center align-items-center"
          >
            <div>Matches</div>
            <div>{{ runningStats.Matches }}</div>
          </div>
        </div>
      </div>
      <div
        class="m-2"
        v-if="outputFileName"
        style="font-size: 0.8em; color: mediumseagreen; font-weight: bold"
      >
        Output saved to: {{ outputFileName }}
      </div>
      <div class="p-2">
        <Fieldset legend="Crawl Log" :toggleable="true">
          <div
            ref="scrollPanelEl"
            style="
              width: 100%;
              height: 400px;
              overflow-y: scroll;
              font-size: 0.8em;
            "
          >
            <div v-for="line of runningLog" class="mb-1">{{ line }}</div>
          </div>
        </Fieldset>
      </div>
    </div>
  </main>
</template>
