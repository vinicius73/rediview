<script lang="ts">
  import { defineComponent } from 'vue'
  import { fasHdd } from '@quasar/extras/fontawesome-v5'

  export default defineComponent({
    name: 'PageHeader',
    setup() {
      return {
        fasHdd,
      }
    },
  })
</script>

<template>
  <q-header elevated class="text-white">
    <q-toolbar class="q-py-sm q-px-md">
      <q-btn
        round
        dense
        flat
        :ripple="false"
        :icon="fasHdd"
        color="white"
        class="q-mr-sm"
        no-caps
      />

      <q-toolbar-title>Rediview</q-toolbar-title>

      <!-- <q-space />

      <div class="q-pl-sm q-gutter-sm row items-center no-wrap">
        <q-btn dense flat no-wrap>
          <q-avatar rounded>
            <img
              src="https://avatars.dicebear.com/api/miniavs/01FN1PYD4G3V.svg"
            />
          </q-avatar>
          <q-icon name="arrow_drop_down" size="16px" />

          <q-menu auto-close>
            <q-list dense>
              <q-item>
                <q-item-section>
                  <div>Todo: <strong>Settings</strong></div>
                </q-item-section>
              </q-item>
              <q-separator />
              <q-item clickable>
                <q-item-section>Option A</q-item-section>
              </q-item>
            </q-list>
          </q-menu>
        </q-btn>
      </div> -->
    </q-toolbar>
  </q-header>
</template>
