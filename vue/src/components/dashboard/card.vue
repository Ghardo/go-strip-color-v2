<template>
<v-card
    max-width="400px"
    width="400px"
    :disabled="!$store.state.app.ready"
    >

    <div class="header">
        <div class="subline">
            <v-switch class="onOffSwitch" v-model="onOffSwitch" inset @change="onOffChange()"></v-switch>  
            <div class="onlineStatus online"></div>    
            <EditableText customClass="title" v-model="title"></EditableText>
        </div> 
    </div>

    <div id="colors">
      <div :style="selectedColor" class="selectedColor" @click="changeColor(0)"/>
      <div class="lastColor" :style="lastColor(index)" v-for="index in 5" :key="index" @click="changeColor(index)" />
    </div>

    <div class="colorpicker">
        <v-slider
            v-model="slider.red"
            class="align-center"
            :max="255"
            :min="0"
            hide-details
            dense
            color="#ff0000"
            thumb-color="#ff0000"
          >
            <template v-slot:append>
              <v-text-field
                v-model="slider.red"
                class="mt-0 pt-0"
                hide-details
                single-line
                type="number"
                style="width: 68px"
                dense
                solo
              ></v-text-field>
            </template>
          </v-slider>

          <v-slider
            v-model="slider.green"
            class="align-center"
            :max="255"
            :min="0"
            hide-details
            dense
            color="#00ff00"
            thumb-color="#00ff00"
          >
            <template v-slot:append>
              <v-text-field
                v-model="slider.green"
                class="mt-0 pt-0"
                hide-details
                single-line
                type="number"
                style="width: 68px"
                dense
                solo
              ></v-text-field>
            </template>
          </v-slider>

          <v-slider
            v-model="slider.blue"
            class="align-center"
            :max="255"
            :min="0"
            hide-details
            dense
            color="#0000ff"
            thumb-color="#0000ff"
          >
            <template v-slot:append>
              <v-text-field
                v-model="slider.blue"
                class="mt-0 pt-0"
                hide-details
                single-line
                type="number"
                style="width: 68px"
                dense
                solo
              ></v-text-field>
            </template>
          </v-slider>

           <v-slider
            v-model="slider.alpha"
            class="align-center"
            :max="255"
            :min="0"
            hide-details
            dense
            color="#ffffff"
            thumb-color="#ffffff"
          >
            <template v-slot:append>
              <v-text-field
                v-model="slider.alpha"
                class="mt-0 pt-0"
                hide-details
                single-line
                type="number"
                style="width: 68px"
                dense
                solo
              ></v-text-field>
            </template>
          </v-slider>
    </div>

</v-card>
</template>

<script>
import EditableText from '@/components/EditableText'
export default {
    name: "DashboardCard",
    props: ['device'],
    components: {EditableText},
    data() {
        return {
            onOffSwitch: false,
            slider: {
                red: 0,
                green: 0,
                blue: 0,
                alpha: 255
            },
            styleObject: {
                backgroundColor: 'rgb(255,0,0)',
             },
             lastColors: [],
             title: ''
        }
    },
    mounted () {
      this.title = this.device.name
      this.$astor.trigger("data.get#"+ this.device.id, "lastColors_" + this.device.id, this.loadLastColors)
      this.$astor.listen("strip.power", this.setPower)
      this.$astor.listen("strip.power.toggle#" + this.device.id, this.togglePower)
      this.$astor.listen('strip.color', this.setColor)
    },
    methods : {
        rgba (r,g,b,a) {
            return `rgba(${r}, ${g}, ${b}, ${a/255})`
        },
        hex (r, g, b) {
          let hr = Number(r).toString(16)
          let hg = Number(g).toString(16)
          let hb = Number(b).toString(16)

          hr = "000000".substr(0, 2 - hr.length) + hr;
          hg = "000000".substr(0, 2 - hg.length) + hg;
          hb = "000000".substr(0, 2 - hb.length) + hb;

          return '#' + hr + hg +hb
        },
        lastColor (index) {
          if(typeof this.lastColors[index] === 'undefined') {
            return { backgroundColor: this.rgba(0, 0, 0, 0) }
          } else {
            let lastColor = this.lastColors[index]
            return { backgroundColor: this.rgba(lastColor[0], lastColor[1], lastColor[2], lastColor[3]) }
          }
        },
        changeColor (key) {
          if (key === 0) {
            this.lastColors = [[this.slider.red, this.slider.green, this.slider.blue, this.slider.alpha]].concat(this.lastColors)
          } 

          if (this.lastColors.length > 6) {
            this.lastColors.pop()
          }
          
          let data = {key: "lastColors_" + this.device.id, value: this.lastColors}
          this.$astor.trigger("data.set#"  + this.device.id, data)

          if ( this.lastColors[key][3] > 0) {
            this.onOffSwitch = true
          }
    
          this.triggerChangeColor(key)

          if (key > 0) {
            this.lastColors = [[this.lastColors[key][0], this.lastColors[key][1], this.lastColors[key][2], this.lastColors[key][3]]].concat(this.lastColors)
            this.updateSlider()
          }
        },
        loadLastColors (payload) {
            if (payload.value) {
              this.lastColors = payload.value
              this.updateSlider()
            }
        },
        updateSlider() {
          this.slider = {
            red: this.lastColors[0][0],
            green: this.lastColors[0][1],
            blue: this.lastColors[0][2],
            alpha: this.lastColors[0][3]
          }
        },
        onOffChange() {
          this.triggerChangeColor(0)
        },
        triggerChangeColor (key) {
          let color = this.hex(this.lastColors[key][0], this.lastColors[key][1], this.lastColors[key][2])

          var colorModel = {color: color, brightness:  this.lastColors[key][3] , power: this.onOffSwitch, start: 0, end: this.device.numleds};
          let d = this.device
          delete d.name

          this.$astor.trigger('change.color', {device: d, color: colorModel});
        },
        setPower(state) {
          this.onOffSwitch = state
          this.triggerChangeColor(0)
        },
        togglePower() {
          this.onOffSwitch = !this.onOffSwitch
          this.triggerChangeColor(0)
        },
        setColor(color) {
          this.onOffSwitch = true
          this.lastColors = [[color[0], color[1], color[2], color[3]]].concat(this.lastColors)
          if (this.lastColors.length > 6) {
            this.lastColors.pop()
          }
          let data = {key: "lastColors_" + this.device.id, value: this.lastColors}
          this.$astor.trigger("data.set#"  + this.device.id, data)
          this.updateSlider()
          this.triggerChangeColor(0)
        }
    },
    computed : {
        selectedColor () {
            let style = { backgroundColor: this.rgba(this.slider.red, this.slider.green, this.slider.blue, this.slider.alpha ) }
            this.$store.commit('app/log', 'selectedColor', style)
            return style
        },
  },
  watch  : {
    title: function () {
      this.$store.commit('app/setDeviceName', {id: this.device.id, name: this.title})
      this.$astor.trigger('data.set', {key: 'deviceNames', value: this.$store.state.app.deviceNames})
    }
  }
}
</script>
b
<style lang="scss">
    .header {
        display: flex;
        justify-content: space-between;
        padding-left: 10px;
        padding-right: 10px;

        .subline {
            display: flex;
            justify-content: start;

            .title {
                line-height: 4rem;

                input {
                  background-color:  #333333;
                }
            }

            .onlineStatus {
              background-color: #333333;
              border-radius: 50%;
              height: 8px;
              width: 8px;
              margin-top: 30px;
              margin-right: 10px;
            }

            .online {
              background-color: green;
            }

        } 
    }

    .colorpicker{
        margin: 5px 10px;
    }

    #colors {
      margin-left: 20px;
      margin-right: 20px;
      display: flex;
      justify-content: space-between;

      .selectedColor, .lastColor {
        width: 30px;
        height: 30px;
        cursor: pointer;
        border: 1px solid black;
        box-shadow: 1px 1px 1px 0px rgba(0,0,0,0.75);
      }

       .selectedColor:hover, .lastColor:hover {
         border: 1px solid white;
       }

      .lastColor {
        border-radius: 5px;
      }

      .selectedColor {
        border-radius: 50%;
      }
      
    }

    

</style>