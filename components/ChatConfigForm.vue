<template>
    <div>
        <form>
            <v-text-field v-model="roomId" />
            <v-text-field v-model="userId" />
            <v-btn
                class="mr4"
                @click="set"
            >
                設定
            </v-btn>
        </form>
    </div>
</template>
<script>
export default {
    data() {
        return {
            roomId: '',
            userId: ''
        }
    },
    methods: {
        set() {
            this.$store.dispatch('setConfig', {
                userId: this.userId,
                roomId: this.roomId
            })
        }
    }
}
</script>
