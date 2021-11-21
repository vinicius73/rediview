<template>
  <q-layout class="bg-grey-1">
    <PageHeader />
    <q-page-container>
      <home />
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">
  import { defineComponent } from 'vue'
  import PageHeader from './Header.vue'
  import Home from '../pages/Home.vue'

  export default defineComponent({
    name: 'MainLayout',
    components: { PageHeader, Home },
    setup() {
      return {}
    },
  })
</script>
