<template>
  <div>
    <v-text-field
      v-model="message"
      append-icon="mdi-send"
      label="メッセージ"
      type="text"
      filled
      @click:append="send"
    />
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: ''
    }
  },
  methods: {
    send() {
      let ws = this.$store.getters.getWebSocket
      let msg = JSON.stringify({
        userId: this.$store.getters.getUserId,
        message: this.message,
        roomId: this.$store.getters.getRoomId,
      })
      ws.send(msg)
      this.message = ''
    }
  }
}
</script>