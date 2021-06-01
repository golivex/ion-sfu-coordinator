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
  </v-container>
</template>

<script>
import LineChart from "@/components/LineChart.vue";
import StartLoadTestDialog from "@/components/StartLoadTestDialog.vue";
import { mapActions, mapGetters } from "vuex";

export default {
  name: "Dashboard",

  components: {
    LineChart,
    StartLoadTestDialog,
  },

  data() {
    return {
      intervalHosts: null,
      intervalLoadStats: null,
      getStatsIntervalInSeconds: 5, // should be set to 5
      fetchLoadStatsIntervalInSeconds: 10, // should be set to 10
      isLoadTestNC: false,
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
  },

  methods: {
    ...mapActions("Stats", ["fetchStats"]),
    ...mapActions("Load", ["startLoadTest", "stopLoadTest", "fetchLoadStats"]),

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
  },

  beforeDestroy() {
    clearInterval(this.intervalHosts);
    clearInterval(this.intervalLoadStats);
  },
};
</script>
