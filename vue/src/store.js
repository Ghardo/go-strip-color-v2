import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
    modules: {
        app: {
            namespaced: true,
            state: {
                title: 'Dashboard',
                ready: false,
                statusText: "© 2020",
                connectedDevices: [],
                deviceNames: {},
                debug: false
            },
            mutations: {
                setWindowTitle(state, title) {
                    window.document.title = title
                },
                setTitle(state, title) {
                    state.title = title
                },
                setReady(state, ready) {
                    state.ready = ready
                },
                setStatusText(state, text) {
                    if (text === null) {
                        state.statusText = "© 2020"
                    } else {
                        state.statusText = text
                    }
                },
                setConnectedDevices(state, devices) {
                    state.connectedDevices = devices
                },
                setDeviceNames(state, names) {
                    if (names !== null) {
                        state.deviceNames = names
                    }
                },
                setDeviceName(state, name) {
                    state.deviceNames[name.id] = name.name
                },
                setDebug(state, debug) {
                    state.debug = debug
                },
                log(state, message, data = null) {
                    if (state.debug)
                        window.console.log(message, data)
                }
            }
        }
    }
    
})

export default store