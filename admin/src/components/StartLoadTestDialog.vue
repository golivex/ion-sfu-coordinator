<template>
  <v-row justify="center">
    <v-dialog v-model="dialog" max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">Start Load Test</span>
        </v-card-title>

        <v-form ref="form">
          <v-card-text>
            <v-container>
              <v-row>
                <v-col cols="12" sm="6">
                  <v-text-field
                    label="Session name *"
                    v-model="form.sessionName"
                    :rules="name"
                    required
                  ></v-text-field>
                </v-col>

                <v-col cols="12" sm="6">
                  <v-text-field
                    label="Number of clients"
                    type="number"
                    :rules="min"
                    v-model="form.clients"
                  ></v-text-field>
                </v-col>

                <v-col cols="12" sm="6" md="4">
                  <label for="role-radio">Role</label>
                  <v-radio-group id="role-radio" v-model="form.role">
                    <v-radio
                      v-for="n in ['pubsub', 'sub']"
                      :key="n"
                      :label="`${n}`"
                      :value="n"
                    ></v-radio>
                  </v-radio-group>
                </v-col>

                <v-col cols="12" sm="6" md="4">
                  <label for="file-radio">File</label>
                  <v-radio-group id="file-radio" v-model="form.file">
                    <v-radio
                      v-for="n in ['320p', '480p']"
                      :key="n"
                      :label="`${n}`"
                      :value="n"
                    ></v-radio>
                  </v-radio-group>
                </v-col>

                <v-col cols="12" sm="6" md="4">
                  <label for="cycle-radio">Cycle</label>
                  <v-radio-group id="cycle-radio" v-model="form.cycle">
                    <v-radio
                      v-for="n in ['0', '1000', '5000']"
                      :key="n"
                      :label="`${n}`"
                      :value="n"
                    ></v-radio>
                  </v-radio-group>
                </v-col>

                <v-col cols="12">
                  <v-text-field
                    label="Rooms"
                    type="number"
                    v-model="form.rooms"
                    :rules="minRooms"
                  ></v-text-field>
                </v-col>
              </v-row>
            </v-container>
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
  name: "StartLoadTestDialog",
  data: () => ({
    dialog: false,
    form: {
      sessionName: "test",
      clients: "1",
      role: "pubsub",
      file: "320p",
      cycle: "0",
      rooms: "-1",
    },

    allowSpaces: false,
  }),

  computed: {
    min() {
      const rules = [];
      const rule = (v) => v >= 1 || `Value should be greater than ${1}`;
      rules.push(rule);
      return rules;
    },

    minRooms() {
      const rules = [];
      const rule = (v) => v >= -1 || `Value should be greater than ${-1}`;
      rules.push(rule);
      return rules;
    },

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
      this.resetForm();
      this.dialog = true;
    },

    handleFormSubmit() {
      if (!this.$refs.form.validate()) return;
      this.$emit("submit", this.form);
      this.dialog = false;
    },

    resetForm() {
      this.form.sessionName = "test";
      this.form.clients = "1";
      this.form.role = "pubsub";
      this.form.file = "320p";
      this.form.cycle = "0";
      this.form.rooms = "-1";
    },
  },
};
</script>