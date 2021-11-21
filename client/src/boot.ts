import { createApp as createVueApp } from 'vue'
import { Quasar, Loading, Notify, QuasarPluginOptions } from 'quasar'

// Import icon libraries
import '@quasar/extras/roboto-font-latin-ext/roboto-font-latin-ext.css'
import '@quasar/extras/material-icons/material-icons.css'

// Import Quasar css
import 'quasar/src/css/index.sass'

import Root from './layout/Main.vue'

export const createApp = async () => {
  const app = createVueApp(Root)

  app.use(Quasar, {
    plugins: {
      Loading,
      Notify,
    },
    config: {
      brand: {
        primary: '#cd5d57',
        secondary: '#6d7278',
        accent: '#5961ff',

        dark: '#1d1d1d',

        positive: '#465282',
        negative: '#C10015',
        info: '#14708d',
        warning: '#cf5815',
      },
      notify: {
        position: 'top',
      }, // default set of options for Notify Quasar plugin
      // loading: {...}, // default set of options for Loading Quasar plugin
      // loadingBar: { ... }, // settings for LoadingBar Quasar plugin
      // ..and many more (check Installation card on each Quasar component/directive/plugin)
    },
  } as Partial<QuasarPluginOptions>)

  return app
}
