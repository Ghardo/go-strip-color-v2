<template>
  <v-app id="inspire">
    <v-navigation-drawer
      v-model="drawer"
      mini-variant
      permanent
      clipped
      app
    >
      <v-list dense>
        <v-list-item link :to="{name: 'dashboard'}">
          <v-list-item-action>
            <v-icon>mdi-view-dashboard</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>Dashboard</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item link :to="{name: 'settings'}">
          <v-list-item-action>
            <v-icon>mdi-cog-outline</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>Settings</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar
      app
      flat
      clipped-left
    >

      <v-toolbar-title>
        <div>{{ $store.state.app.title }}</div>
        <div>
          
          <v-icon @click="dialog=true">mdi-palette</v-icon>
          <v-icon @click="reloadDevices()">mdi-refresh</v-icon>
          <v-icon @click="setPower(true)">mdi-power-on</v-icon>
          <v-icon @click="setPower(false)">mdi-power-off</v-icon>
        </div>
      </v-toolbar-title>
    </v-app-bar>

    <v-content>
      <v-container
        fluid
       >
        <ColorDialog :dialog.sync="dialog" />
        <router-view></router-view>
      </v-container>
    </v-content>

    <v-footer app>
      {{ $store.state.app.statusText }}
    </v-footer>
  </v-app>
</template>

<script>
  import ColorDialog from '@/components/dialogs/color.vue'
  export default {
    components: {ColorDialog},
    props: {
      source: String,
    },

    data: () => ({
      drawer: null,
      keyCapture: true,
      dialog: false
    }),

    created () {
      this.$vuetify.theme.dark = true
      this.$store.commit('app/setWindowTitle', 'Arduino Strip Color Changer')
      this.$astor.trigger('app-ready', {}, this.ready)
      this.$astor.listen('refresh.devices.callback', this.deviceListUpdate)
      this.$astor.listen('key.capture', this.setKeyCapture)

      this.$storage.get('db')
      .then(data => {
        console.log(data);
      })
      .catch(err => {
        console.error(err);
      });
    },
    mounted () {
      window.addEventListener("keypress", e => {
        this.onKeyPress(e)
      });
    },
    methods: {
      setKeyCapture (enabled) {
        this.keyCapture = enabled
      },
      onKeyPress(e) {
        this.$store.commit('app/log', "key press", e)

        if (!this.keyCapture) {
          this.$store.commit('app/log', "key press disabled")
          return
        }

        let deviceCount =  this.$store.state.app.connectedDevices.length
        if (e.keyCode >= 49 && e.keyCode <= 49 + deviceCount-1) {
          let devIndex = e.keyCode - 49
          let deviceId = this.$store.state.app.connectedDevices[devIndex].id
          this.$astor.emit("strip.power.toggle#" + deviceId)
          return
        }

        switch(String.fromCharCode(e.keyCode)) {
          case 'o':
            this.setPower(false)
            break
          case 'p':
            this.setPower(true)
            break
          case 'r':
            this.reloadDevices()
            break
          case 'a':
            this.dialog = true
            break;
        }
      },

      ready (payload) {
        this.$astor.debug = payload.debug
        this.$store.commit('app/setDebug', payload.debug)

        this.$astor.trigger('data.get#deviceNames', "deviceNames", this.updateDeviceNames)

        this.$astor.trigger('refresh.devices')
        this.$store.commit('app/setStatusText', 'refreshing devices...')
        
      },
      updateDeviceNames(payload) {
        this.$store.commit('app/log', "updateDeviceNames", payload)
        this.$store.commit('app/setDeviceNames', payload['value'])
      },
      reloadDevices () {
        this.$store.commit('app/setReady', false)
        this.$astor.trigger('refresh.devices')
        this.$store.commit('app/setStatusText', 'refreshing devices...')
      },
      deviceListUpdate (payload) {
        this.$store.commit('app/log', "deviceListUpdate", payload)
        let devices = []
        let self = this
        Object.entries(payload.DeviceList).forEach(function(entry) {
          if (self.$store.state.app.deviceNames[entry[1].id]) {
            entry[1].name = self.$store.state.app.deviceNames[entry[1].id]
          } else {
            entry[1].name = "Arduino Strip"
          }
          
          devices.push(entry[1])
        })

        this.$store.commit('app/setConnectedDevices', devices)
        this.$store.commit('app/setStatusText', null)
        this.$store.commit('app/setReady', true)
      },
      setPower(state) {
        this.$astor.emit('strip.power', state)
      },
      togglePower(deviceId) {
        this.$astor.emit('strip.power.toggle#' + deviceId)
      }
    },
    destroyed() {
      this.$astor.remove('app-ready', this.ready)
      this.$astor.remove('devicelist.update', this.deviceListUpdate)
    },
  }
</script>

<style lang="scss">
body::-webkit-scrollbar {
  width: 10px;
}
body::-webkit-scrollbar-button {
  width: 0px;
  height: 0px;
}
body::-webkit-scrollbar-thumb {
  background: #e1e1e1;
  border: 0px none #ffffff;
  border-radius: 10px;
}
body::-webkit-scrollbar-thumb:hover {
  background: #ffffff;
}
body::-webkit-scrollbar-track {
  background: #272727;
  border: 0px none #ffffff;
}
body::-webkit-scrollbar-track:hover {
  background: #272727;
}

.v-toolbar__title {
  div {
    text-transform: uppercase;
  }

  display: flex;
  justify-content: space-between;
  width: 100%;
}

.v-icon {
  cursor: pointer;
  padding: 4px;
}

.v-icon:hover {
  background: grey;
}

.v-footer {
    font-family: 'Courier New', Courier, monospace;
    font-size: 12px;
}



</style>