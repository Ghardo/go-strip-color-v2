<template>
    <div :class="ccl">
        <span @click="edit(true)" v-show="!editing">
            {{v}}
        </span>
        <input v-model="v"
            v-show="editing"
            @blur="blur()"
            @keydown.enter="finish()"
            ref="input"
            type="text">
    </div>
</template>

<script>
export default {
    name: 'EditableText',
    props: ['value', 'customClass'],
    data() {
        return {
            editing: false,
            ccl: '',
            v: null
        }
    },
    mounted () {
        this.ccl = this.customClass
        this.v = this.value
    },
    methods: {
        blur () {
            this.editing = false
            this.v = this.value
            this.notify()
        },
        finish () {
            this.editing = false
            this.notify()
            this.$emit('input', this.v)
            this.$emit('change', this.v)
        },
        edit (editing) {
            this.editing = editing
            this.notify()
            setTimeout(() => {
                this.$refs.input.focus();
            }, 10);
        },
        notify() {
            this.$astor.emit('key.capture', !this.editing)
        }
    },
    watch: {
        value: function (val) {
            this.v = val
        }
    }
}
</script>

<style>
</style>