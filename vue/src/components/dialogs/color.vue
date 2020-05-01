<template>
  <v-row justify="center">
    <v-dialog v-model="dialog" persistent max-width="400">
      <v-card>
        <v-card-title class="headline">Select a color</v-card-title>
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
        
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="red darken-1" text @click="close(false)">Cancel</v-btn>
          <v-btn color="green darken-1" text @click="close(true)">Set to all</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
  export default {
    name: 'ColorDialog',
    props: {
        dialog: {
            default: false
      }
    },
     data() {
        return {
            slider: {
                red: 0,
                green: 0,
                blue: 0,
                alpha: 255
            },
            lastColors: [],
        }
    },
    methods: {
        rgba (r,g,b,a) {
            return `rgba(${r}, ${g}, ${b}, ${a/255})`
        },
        lastColor (index) {
          if(typeof this.lastColors[index] === 'undefined') {
            return { backgroundColor: this.rgba(0, 0, 0, 0) }
          } else {
            let lastColor = this.lastColors[index]
            return { backgroundColor: this.rgba(lastColor[0], lastColor[1], lastColor[2], lastColor[3]) }
          }
        },
        close(state) {
            if (state) {
                this.$astor.emit('strip.color', [this.slider.red, this.slider.green, this.slider.blue, this.slider.alpha])
                this.lastColors = [[this.slider.red, this.slider.green, this.slider.blue, this.slider.alpha]].concat(this.lastColors)
                let data = {key: "globalColor", value: this.lastColors}
                this.$astor.trigger("data.set", data)

            }
            this.$emit('update:dialog', false)
        },
        loadColor(payload) {
            this.$store.commit('app/log', 'loadColor', payload)
            if (payload.value) {
               this.lastColors = payload.value
               this.updateSlider()
            }
        },
        changeColor (key) {
          if (key > 0) {
            this.slider = {red: this.lastColors[key][0], green: this.lastColors[key][1], blue: this.lastColors[key][2], alpha: this.lastColors[key][3] }
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
    },
    created () {
        this.$astor.trigger("data.get#globalColor", "globalColor", this.loadColor)
    },
    computed : {
        selectedColor () {
            let style = { backgroundColor: this.rgba(this.slider.red, this.slider.green, this.slider.blue, this.slider.alpha ) }
            this.$store.commit('app/log', 'selectedColor', style)
            return style
        },
  },
  }
</script>

<style lang="scss">
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