
<template>
  <div>
    {{ getCurrentComponent }}
    <div v-if="getCurrentComponent == -1">
      <load></load>
    </div>
  </div>
</template>

<script>
import Load from "./Load.vue";
export default {
  name: "EditorComponent",
  components: {
    Load,
  },
  computed: {
    getCurrentComponent() {
      return this.$store.state.editor.currentComponentId;
    },
  },
};
</script>

<style>
</style>