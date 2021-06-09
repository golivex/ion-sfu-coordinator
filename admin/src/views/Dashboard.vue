<template>
  <v-container class="grey lighten-5">
    <!-- TOP BUTTONS ROW -->
    <v-row>
      <v-col class="text-right">
        <v-btn
          class="text-capitalize"
          :loading="isLoadTestNC"
          v-if="!isLoadTesting"
          :disabled="isLoadTestNC"
          depressed
          @click="$refs.startLoadTestDialog.open()"
        >
          start load test
        </v-btn>
        <v-btn
          class="text-capitalize"
          :loading="isLoadTestNC"
          :disabled="isLoadTestNC"
          v-else
          depressed
          color="error"
          @click="handleStopLoadTest"
        >
          stop load test
        </v-btn>
      </v-col>
      <v-col class="text-right">
        <v-btn
          class="text-capitalize"
          depressed
          :loading="isDiskSaving"
          v-if="!diskSavingStatus"
          :disabled="isDiskSaving"
          @click="$refs.startDiskDialog.open()"
        >
          start disk saving
        </v-btn>
        <v-btn
          class="text-capitalize"
          depressed
          color="error"
          v-if="diskSavingStatus"
          :loading="isDiskSaving"
          :disabled="isDiskSaving"
          @click="handleStopDisk"
        >
          stop disk saving
        </v-btn>
      </v-col>
      <v-col class="text-right">
        <v-btn
          class="text-capitalize"
          depressed
          :loading="isStreaming"
          v-if="!liveStreamingStatus"
          :disabled="isStreaming"
          @click="openSessionDialog('Stream', 'Live')"
        >
          start live stream
        </v-btn>
        <v-btn
          class="text-capitalize"
          depressed
          color="error"
          v-if="liveStreamingStatus"
          :loading="isDiskSaving"
          :disabled="isDiskSaving"
          @click="handleStopStreaming"
        >
          stop live stream
        </v-btn>
      </v-col>
      <v-col class="text-right">
        <v-btn
          class="text-capitalize"
          depressed
          :loading="isLiveRtmp"
          v-if="!liveRtmpStatus"
          :disabled="isLiveRtmp"
          @click="openSessionDialog('Rtmp','Live')"
        >
          start live rtmp
        </v-btn>
        <v-btn
          class="text-capitalize"
          depressed
          color="error"
          v-if="liveRtmpStatus"
          :loading="isLiveRtmp"
          :disabled="isLiveRtmp"
          @click="handleStopLiveRtmp"
        >
          stop live rtmp
        </v-btn>
      </v-col>
      <v-col class="text-right">
        <!-- <v-btn
          class="text-capitalize"
          depressed
          :loading="isDemoLoader"
          :disabled="isDemoLoader"
          @click="handleDemoFor"
        >
          Stop Demo
        </v-btn> -->
        <v-menu offset-y>
        <template v-slot:activator="{ on, attrs }">
         <v-btn
          color="primary"
          dark
          v-bind="attrs"
          v-on="on"
        >
          Demos
        </v-btn>
      </template>
      <v-list>
        <v-list-item
          v-for="(item, index) in items"
          :key="index"
          @click="openSessionDialog(item.title, 'Demo')"
        >
          <v-list-item-title>{{ item.title }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
      </v-col>
    </v-row>

    <!-- HOSTS SECTION -->
    <v-row>
      <v-col cols="12">
        <div class="title">Hosts</div>
      </v-col>

      <v-col v-for="(host, index) in hosts" :key="index" cols="12" sm="6">
        <v-card class="pa-2" outlined tile>
          <div class="d-flex justify-space-between">
            <div>
              <div class="subtitle">IP: {{ host.ip }}</div>
              <div class="subtitle">Port: {{ host.port }}</div>
            </div>

            <ul>
              <li>Peer: {{ host.peer }}</li>
              <li>Audio: {{ host.audio }}</li>
              <li>Video: {{ host.video }}</li>
            </ul>
          </div>

          <div class="mt-5">
            <line-chart :xaxis="host.xaxis" :series="host.series" />
          </div>

          <div v-if="host.sessions && host.sessions.length">
            <div class="caption">Session</div>
            <ul class="d-flex flex-wrap" style="gap: 2rem">
              <li v-for="(session, index) in host.sessions" :key="index">
                <div class="subtitle">Name: {{ session.Name }}</div>
                <div class="subtitle">PeerCount: {{ session.PeerCount }}</div>
                <div class="subtitle">
                  AudioTracks: {{ session.AudioTracks }}
                </div>
                <div class="subtitle">
                  VideoTracks: {{ session.VideoTracks }}
                </div>
              </li>
            </ul>
          </div>
        </v-card>
      </v-col>
    </v-row>

    <!-- ACTION HOST SECTION -->
    <v-row>
      <v-col cols="12">
        <div class="title">Action Hosts</div>
      </v-col>

      <v-col
        v-for="(actionHost, index) in actionHosts"
        :key="index"
        cols="12"
        sm="6"
      >
        <v-card class="pa-2" outlined tile>
          <div class="d-flex justify-space-between">
            <div>
              <div class="subtitle">IP: {{ actionHost.ip }}</div>
              <div class="subtitle">Port: {{ actionHost.port }}</div>
            </div>

            <div v-if="actionHost.loadStatus && actionHost.loadStatus.length">
              <div class="caption">Load Test</div>
              <ul class="d-flex flex-wrap" style="gap: 2rem">
                <li v-for="(ls, index) in actionHost.loadStatus" :key="index">
                  <div class="subtitle">Clients: {{ ls.Stats.clients }}</div>
                  <div class="subtitle">Engine: {{ ls.Stats.engine }}</div>
                  <div class="subtitle">Cpu: {{ ls.Stats.hostload.cpu }}</div>
                  <div class="subtitle">
                    Tasks: {{ ls.Stats.hostload.tasks }}
                  </div>
                </li>
              </ul>
            </div>
          </div>

          <div class="mt-5">
            <line-chart :xaxis="actionHost.xaxis" :series="actionHost.series" />
          </div>
        </v-card>
      </v-col>
    </v-row>

    <!-- DIALOGS -->
    <start-load-test-dialog
      ref="startLoadTestDialog"
      @submit="handleStartLoadTest"
    />
    
    <start-disk-save-dialog ref="startDiskDialog" @submitDisk="handleDiskSaving"/>

    <session-name-dialog ref="askSessionDialog" :sessionDialogContent="sessionDialogContent" @submitLiveStream="handleLiveStream" @submitLiveRtmp="handleLiveRtmp" @submitDemoStream="handleDemoLiveStream" @submitDemoRtmp="handleDemoLiveRtmp" />
  </v-container>
</template>

<script>
import LineChart from "@/components/LineChart.vue";
import StartLoadTestDialog from "@/components/StartLoadTestDialog.vue";
import SessionNameDialog from "@/components/AskSessionNameDialog.vue";
import StartDiskSaveDialog from "@/components/StartDiskSaveDialog.vue";
import { mapActions, mapGetters } from "vuex";

export default {
  name: "Dashboard",

  components: {
    LineChart,
    StartLoadTestDialog,
    SessionNameDialog,
    StartDiskSaveDialog
  },

  data() {
    return {
      intervalHosts: null,
      intervalLoadStats: null,
      getStatsIntervalInSeconds: 5, // should be set to 5
      fetchLoadStatsIntervalInSeconds: 10, // should be set to 10
      isLoadTestNC: false,
      isDiskSaving: false,
      isStreaming: false,
      isLiveRtmp: false,
      isDemoLoader: false,
      sessionDialogContent: {},
      items: [
        { title: 'Stream' },
        { title: 'Rtmp' }
      ],
    };
  },

  created() {
    this.intervalHosts = setInterval(
      () => this.getStatsData(),
      this.getStatsIntervalInSeconds * 1000
    );

    this.intervalLoadStats = setInterval(() => {
      if (this.isLoadTesting) this.fetchLoadStats();
    }, this.fetchLoadStatsIntervalInSeconds * 1000);
  },

  mounted() {
    this.getStatsData();
  },

  computed: {
    ...mapGetters("Stats", ["getStats"]),
    ...mapGetters("Load", ["getLoadTestingStatus", "getLoadStatus"]),
    ...mapGetters("Disk", ["getdiskSavingStatus"]),
    ...mapGetters("Stream", ["getliveStreamingStatus"]),
    ...mapGetters("Rtmp", ["getliveRtmpStatus"]),

    hosts() {
      for (const key in this.getStats.hosts) {
        if (Object.hasOwnProperty.call(this.getStats.hosts, key)) {
          let h = this.getStats.hosts[key];
          const { ip, port } = this.getStats.hosts[key];

          for (const sessionkey in this.sessions) {
            let sessions = [];
            if (Object.hasOwnProperty.call(this.sessions, sessionkey)) {
              let { Host, Port } = this.sessions[sessionkey];

              if (ip === Host && port === Port) {
                sessions.push(this.sessions[sessionkey]);
                h.sessions = sessions;
              }
            }
          }

          h.xaxis = this.getXaxis(h.Loads.length);
          h.series = this.getSeries(h.Loads);
        }
      }

      return this.getStats.hosts;
    },

    actionHosts() {
      for (const key in this.getStats.actionhosts) {
        if (Object.hasOwnProperty.call(this.getStats.actionhosts, key)) {
          let h = this.getStats.actionhosts[key];
          const { ip, port } = this.getStats.actionhosts[key];

          for (const sessionkey in this.getLoadStatus) {
            let loadStatus = [];
            if (Object.hasOwnProperty.call(this.getLoadStatus, sessionkey)) {
              let { Ip, Port } = this.getLoadStatus[sessionkey];

              if (ip === Ip && port === Port) {
                loadStatus.push(this.getLoadStatus[sessionkey]);
                h.loadStatus = loadStatus;
              }
            }
          }

          h.xaxis = this.getXaxis(h.Loads.length);
          h.series = this.getSeries(h.Loads);
        }
      }

      return this.getStats.actionhosts;
    },

    sessions() {
      return this.getStats.sessions;
    },

    isLoadTesting() {
      return this.getLoadTestingStatus;
    },

    diskSavingStatus() {
      return this.getdiskSavingStatus
    },
    liveStreamingStatus() {
      return this.getliveStreamingStatus
    },
    liveRtmpStatus() {
      return this.getliveRtmpStatus
    }
  },

  methods: {
    ...mapActions("Stats", ["fetchStats"]),
    ...mapActions("Load", ["startLoadTest", "stopLoadTest", "fetchLoadStats"]),
    ...mapActions("Disk", ["startDiskSave", "stopDiskSave"]),
    ...mapActions("Stream", ["startLiveStream", "stopLiveStream", "demoLiveStream"]),
    ...mapActions("Rtmp", ["startLiveRtmp", "stopLiveRtmp", "demoLiveRtmp"]),

    getStatsData() {
      this.fetchStats();
    },

    async handleStartLoadTest(form) {
      this.isLoadTestNC = true;
      await this.startLoadTest(form);
      await this.fetchLoadStats();
      this.isLoadTestNC = false;
    },

    async handleStopLoadTest() {
      this.isLoadTestNC = true;
      await this.stopLoadTest();
      this.isLoadTestNC = false;
    },

    getXaxis(length) {
      const xaxisArray = [];
      let date = new Date().getTime();
      let count = 0;

      while (count < length) {
        let spd = new Date(date).toLocaleString();

        spd = spd.split(",")[1].replace(/[AM/PM]/g, "");

        xaxisArray.push(spd);

        date += 1000;
        count++;
      }

      return xaxisArray;
    },

    getSeries(series) {
      return series.reduce((acc, curr) => {
        const keys = Object.keys(curr);

        if (!acc.length) {
          keys.forEach((k) => {
            acc.push({ name: k, data: [] });
          });
        }

        acc.map((a) => {
          a.data.push(curr[a.name]);
          return a;
        });

        return acc;
      }, []);
    },

    async handleDiskSaving(param) {
      this.isDiskSaving = true
      await this.startDiskSave(param)
      this.isDiskSaving = false
    },

    async handleStopDisk() {
      this.isDiskSaving = true
      await this.stopDiskSave()
      this.isDiskSaving = false
    },
    
    async handleLiveStream(param) {
      this.isStreaming = true
      await this.startLiveStream(param)
      this.isStreaming = false
    },

    async handleStopStreaming() {
      this.isStreaming = true
      await this.stopLiveStream()
      this.isStreaming = false
    },

    async handleDemoLiveStream(param) {
      this.isDemoLoader = true
      await this.demoLiveStream(param)
      this.isDemoLoader = false
    },

    async handleLiveRtmp(param) {
      this.isLiveRtmp = true
      await this.startLiveRtmp(param)
      this.isLiveRtmp = false
    },

    async handleStopLiveRtmp() {
      this.isLiveRtmp = true
      await this.stopLiveRtmp()
      this.isLiveRtmp = false
    },

    async handleDemoLiveRtmp(param) {
      this.isDemoLoader = true
      await this.demoLiveRtmp(param)
      this.isDemoLoader = false
    },

    openSessionDialog(header, type) {
      const param = {
        header: header,
        type: type
      }
      this.sessionDialogContent = param
      this.$refs.askSessionDialog.open()
    },

    handleDemoFor() {
      console.log("CLOSE DEMO")
    }
  },

  beforeDestroy() {
    clearInterval(this.intervalHosts);
    clearInterval(this.intervalLoadStats);
  },
};
</script>
