<template>
  <div>
    <div v-if="this.userId != this.$store.getters.getUserId">
      <div>
        {{userId}}
      </div>
      <v-card style="display: inline-block;" max-width="50%">
        <v-card-text>
          {{message}}
        </v-card-text>
        <v-card-subtitle class="align-right">
          {{time}}
        </v-card-subtitle>
      </v-card>
    </div>
    <v-row class="justify-end" v-else>
      <v-card style="display: inline-block;" max-width="50%">
        <v-card-text>
          {{message}}
        </v-card-text>
        <v-card-subtitle class="align-right">{{time}}</v-card-subtitle>
      </v-card>
    </v-row>
  </div>
</template>

<script>
export default {
  props: ["userId", "message", "time"]
}
</script>