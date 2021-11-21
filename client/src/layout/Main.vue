<template>
  <q-layout class="bg-grey-1">
    <PageHeader />
    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">
  import { defineComponent } from 'vue'
  import PageHeader from './Header.vue'

  export default defineComponent({
    name: 'MainLayout',
    components: { PageHeader },
    setup() {
      return {}
    },
  })
</script>
