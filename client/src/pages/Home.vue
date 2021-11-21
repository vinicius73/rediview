<script lang="ts">
  import { defineComponent, ref, nextTick, watch } from 'vue'
  import { useWebSocket } from '@vueuse/core'

  export default defineComponent({
    name: 'PageHome',
    setup() {
      const { status, data, send } = useWebSocket('ws://localhost:3073/ws', {
        autoReconnect: true,
      })

      const messages = ref<string[]>([])
      const newMessage = ref<string>('')

      const sendMessage = () => {
        send(newMessage.value)
        nextTick(() => (newMessage.value = ''))
      }

      watch(
        () => data.value,
        (msg) => {
          messages.value.push(msg)
        }
      )

      return {
        messages,
        newMessage,
        sendMessage,
        status,
        data,
      }
    },
  })
</script>

<template>
  <div>
    <h3>Messages</h3>
    <q-input v-model="newMessage" />
    <q-btn label="enviar" @click="sendMessage" />
    <pre>data: {{ data }}</pre>
    <pre>status: {{ status }}</pre>
    <ul>
      <li v-for="(msg, i) in messages" :key="`msg-${i}`">
        {{ msg }}
      </li>
    </ul>
  </div>
</template>
