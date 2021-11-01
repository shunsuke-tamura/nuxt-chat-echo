<template>
  <div>
    <v-card class="fullheight">
      <v-card-title ref="chatBar"> RoomID on {{ roomId }} </v-card-title>
      <div class="scrollY" ref="chatBar">
        <v-card-text v-for="message in messages" v-bind:key="message.idx">
          <chat-tips
            :userId="message.userId"
            :message="message.message"
            :time="message.time"
          />
        </v-card-text>
      </div>
    </v-card>
  </div>
</template>

<script>
import ChatTips from "./ChatTips.vue";
export default {
  components: { ChatTips },
  data() {
    return {
      roomId: "",
      messages: [],
    };
  },
  computed: {
    storeRoomId() {
      return this.$store.getters.getRoomId;
    },
  },
  watch: {
    storeRoomId() {
      this.roomId = this.$store.getters.getRoomId;
      /*
      fetchChat(this.roomId).then(res => {
        console.log(res.data)
        this.messages = res.data
      })
      */
      this.ws = this.$store.getters.getWebSocket;
      this.ws.onmessage = (msg) => {
        this.messages.push(JSON.parse(msg.data));
      };
    },
  },
};
</script>

<style>
.scrollY {
  height: 90%;
  overflow-y: scroll;
  background-color: rgb(70, 161, 235);
}
</style>