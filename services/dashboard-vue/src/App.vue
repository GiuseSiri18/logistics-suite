<template>
  <v-app>
    <v-app-bar app color="indigo" dark>
      <v-toolbar-title>Logistics Enterprise Suite</v-toolbar-title>
    </v-app-bar>

    <v-main>
      <v-container>
        <v-row mt-5>
          <v-col cols="12">
            <h1 class="display-1">Microservices Status</h1>
          </v-col>
          
          <v-col cols="12" md="6">
            <v-card color="#FF2D20" dark>
              <v-card-title>Laravel Core API</v-card-title>
              <v-card-text class="white--text">
                Status: {{ core.status || 'Connecting...' }}<br>
                Service: {{ core.service }}
              </v-card-text>
            </v-card>
          </v-col>

          <v-col cols="12" md="6">
            <v-card color="#00ADD8" dark>
              <v-card-title>Go Analytics Engine</v-card-title>
              <v-card-text class="white--text">
                Status: {{ stats.status || 'Connecting...' }}<br>
                Process Time: {{ stats.processing_time }}
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import axios from 'axios';

export default {
  data: () => ({
    core: {},
    stats: {}
  }),
  async mounted() {
    try {
      const coreRes = await axios.get('/api/core/status');
      this.core = coreRes.data;
      const statsRes = await axios.get('/api/stats');
      this.stats = statsRes.data;
    } catch (e) {
      console.error("API Error", e);
    }
  }
}
</script>