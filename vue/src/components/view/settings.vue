<template>
  <div id="settings">
      <h3>Connected Devices</h3>
      <v-data-table
      :headers="headers"
      :items="$store.state.app.connectedDevices"
      :search="search"
      item-key="id"
    />
  </div>
</template>

<script>
export default {
  name: "ViewSettings",
  data () {
    return {
      ready: false,
      search: '',
      items: [],
      headers: [
        {
          text: 'Id',
          align: 'Start',
          value: 'id'
        },
        {
          text: 'Name',
          align: 'start',
          value: 'name',
        },
        {
          text: 'Port',
          align: 'start',
          value: 'port',
        },
        {
          text: 'Led',
          align: 'center',
          value: 'numleds'
        },
        {
          text: 'Vendor',
          align: 'Center',
          value: 'vendor'
        },
        {
          text: 'Product',
          align: 'Center',
          value: 'product'
        },
        {
          text: 'Serial',
          align: 'start',
          value: 'serialnumber'
        },
      ]
    }
  },
  mounted () {
    this.$store.commit('app/setTitle', 'Settings')
  },
}
</script>

<style>

</style>