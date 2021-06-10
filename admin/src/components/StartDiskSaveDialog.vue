<template>
  <v-row justify="center">
    <v-dialog v-model="dialog" max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">Start Disk Saving</span>
        </v-card-title>

        <v-form ref="form">
          <v-card-text>
            <v-container>
              <v-row>
                <v-col cols="12" md="6">
                  <v-text-field
                    label="Session name *"
                    v-model="session"
                    :rules="name"
                    required
                  ></v-text-field>
                </v-col>

                <v-col cols="12" md="6">
                  <v-text-field
                    label="File Name *"
                    v-model="fileName"
                    :rules="fileNameRule"
                    required
                  ></v-text-field>
                </v-col>
              </v-row>
            </v-container>
            <small>*indicates required field</small>
          </v-card-text>

          <v-card-actions>
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
  data: () => ({
    dialog: false,
    session: "",
    fileName: ""
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
    fileNameRule() {
      const rules = [];
      const rule = (v) =>
        (v !== "" && v.indexOf(" ") < 0) ||
        `should not be empty`;
      rules.push(rule);
      return rules;
    }
  },

  methods: {
    open() {
      this.session = '';
      this.fileName = ''
      this.dialog = true;
    },

    handleFormSubmit() {
      if (!this.$refs.form.validate()) return;
      const payload = {
          session: this.session,
          filename: this.fileName
      }
      this.$emit("submitDisk", payload)
      this.dialog = false;
    },
  },
};
</script>