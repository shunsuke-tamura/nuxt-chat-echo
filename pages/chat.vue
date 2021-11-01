<template>
    <div>
        <v-container fluid>
            <v-row>
                <v-col cols=2>
                    <chat-config-form />
                </v-col>
                <v-col cols=10 class="fullheight">
                    <chat-view />
                    <chat-form />
                </v-col>
            </v-row>
        </v-container>
    </div>
</template>

<script>
import ChatView from '../components/ChatView.vue'
import ChatConfigForm from '../components/ChatConfigForm.vue'
import ChatForm from '../components/ChatForm.vue'

export default {
    components: {
        ChatView,
        ChatConfigForm,
        ChatForm
    },
}
</script>

<style>
.fullheight { 
    height: 80vh;
}
</style>
