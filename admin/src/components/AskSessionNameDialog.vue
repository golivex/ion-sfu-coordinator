<template>
  <v-row justify="center">
    <v-dialog v-model="dialog" max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">Start {{sessionDialogContent.type}} {{sessionDialogContent.header}}</span>
        </v-card-title>

        <v-form ref="form">
          <v-card-text class="px-6">
            <v-text-field
              class="mb-12"
              label="Session name *"
              v-model="session"
              :rules="name"
              required
            ></v-text-field>
            <small>*indicates required field</small>
          </v-card-text>

          <v-card-actions class="px-6">
            <v-spacer></v-spacer>
            <v-btn color="blue darken-1" text @click="dialog = false">
              Close
            </v-btn>

            <v-btn color="primary" @click="handleFormSubmit"> Start </v-btn>
          </v-card-actions>
        </v-form>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
export default {
  name: "SessionNameDialog",
  props: {
      sessionDialogContent: {
          type: Object,
          default: ''
      }
  },
  data: () => ({
    dialog: false,
    session: ""
  }),

  computed: {
    name() {
      const rules = [];
      const rule = (v) =>
        (v !== "" && v.indexOf(" ") < 0 && v.length >= 4) ||
        `name should be greater than 4 charaters`;
      rules.push(rule);
      return rules;
    },
  },

  methods: {
    open() {
      this.session = '';
      this.dialog = true;
    },

    handleFormSubmit() {
      if (!this.$refs.form.validate()) return;
      if (this.sessionDialogContent.type === 'Live' && this.sessionDialogContent.header === 'Stream') {
        this.$emit("submitLiveStream", this.session)
      } else if (this.sessionDialogContent.type === 'Live' && this.sessionDialogContent.header === 'Rtmp') {
        this.$emit("submitLiveRtmp", this.session)
      } else if (this.sessionDialogContent.type === 'Demo' && this.sessionDialogContent.header === 'Stream') {
        this.$emit("submitDemoStream", this.session)
      } else if (this.sessionDialogContent.type === 'Demo' && this.sessionDialogContent.header === 'Rtmp') {
        this.$emit("submitDemoRtmp", this.session)
      }
      this.dialog = false; 
    },
  },
};
</script>