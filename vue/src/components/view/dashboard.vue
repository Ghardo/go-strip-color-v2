<template>
  <div id="dashboard">
      <DashboardCard v-for="device in $store.state.app.connectedDevices" :device="device" :key="device.id"></DashboardCard>
  </div>
</template>

<script>

import DashboardCard from '@/components/dashboard/card.vue'
export default {
    name: "ViewDashboard",
    components: {DashboardCard},
    mounted () {
        this.$store.commit('app/setTitle', 'Dashboard')
    }
    
}
</script>

<style lang="scss" scoped>
#dashboard {
    display: flex;
    justify-content: flex-start;
    flex-flow: row wrap;

    .v-card {
        margin: 5px;
    }
}
</style>