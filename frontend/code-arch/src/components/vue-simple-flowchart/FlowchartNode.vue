<template>
  <div
    class="flowchart-node"
    :style="nodeStyle"
    @mousedown="handleMousedown"
    @mouseover="handleMouseOver"
    @mouseleave="handleMouseLeave"
    :class="{ selected: options.selected === id }"
  >
    <div
      v-if="showTopPort"
      class="node-port node-input"
      @mousedown="inputMouseDown"
      @mouseup="inputMouseUp"
    ></div>
    <div class="node-main">
      <div v-text="type" class="node-type"></div>
      <div v-text="label" class="node-label"></div>
    </div>
    <div
      v-if="showBottomPort"
      class="node-port node-output"
      @mousedown="outputMouseDown"
    ></div>
    <div v-show="show.delete && type != 'internal'" class="node-delete">X</div>
  </div>
</template>

<script>
export default {
  name: "FlowchartNode",
  props: {
    id: {
      type: Number,
      default: 1000,
      validator(val) {
        return typeof val === "number";
      },
    },
    x: {
      type: Number,
      default: 0,
      validator(val) {
        return typeof val === "number";
      },
    },
    y: {
      type: Number,
      default: 0,
      validator(val) {
        return typeof val === "number";
      },
    },
    type: {
      type: String,
      default: "Default",
    },
    label: {
      type: String,
      default: "input name",
    },
    options: {
      type: Object,
      default() {
        return {
          centerX: 1024,
          scale: 1,
          centerY: 140,
        };
      },
    },
  },
  data() {
    return {
      show: {
        delete: false,
      },
    };
  },
  mounted() {},
  computed: {
    showTopPort() {
      if (this.label === "request") {
        return false;
      }
      return true;
    },
    showBottomPort() {
      if (this.label === "response") {
        return false;
      }
      return true;
    },
    nodeStyle() {
      return {
        top: this.options.centerY + this.y * this.options.scale + "px", // remove: this.options.offsetTop +
        left: this.options.centerX + this.x * this.options.scale + "px", // remove: this.options.offsetLeft +
        transform: `scale(${this.options.scale})`,
      };
    },
  },
  methods: {
    handleMousedown(e) {
      const target = e.target || e.srcElement;
      if (
        target.className.indexOf("node-input") < 0 &&
        target.className.indexOf("node-output") < 0
      ) {
        this.$emit("nodeSelected", e);
      }
      e.preventDefault();
    },
    handleMouseOver() {
      this.show.delete = true;
    },
    handleMouseLeave() {
      this.show.delete = false;
    },
    outputMouseDown(e) {
      this.$emit("linkingStart");
      e.preventDefault();
    },
    inputMouseDown(e) {
      e.preventDefault();
    },
    inputMouseUp(e) {
      this.$emit("linkingStop");
      e.preventDefault();
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
$themeColor: rgb(99, 111, 131);
$portSize: 12;
$linkColor: rgb(46, 184, 92);
$dangerColor: rgb(229, 83, 83);
$warningColor: rgb(249, 177, 21);

.flowchart-node {
  margin: 0;
  width: 80px;
  height: 50px;
  position: absolute;
  box-sizing: border-box;
  border: none;
  background: white;
  z-index: 1;
  opacity: 0.9;
  cursor: move;
  transform-origin: top left;
  .node-main {
    text-align: center;
    .node-type {
      background: $themeColor;
      color: white;
      font-size: 13px;
      padding: 3px;
    }
    .node-label {
      font-size: 13px;
    }
  }
  .node-port {
    position: absolute;
    width: #{$portSize}px;
    height: #{$portSize}px;
    left: 50%;
    transform: translate(-50%);
    border: 1px solid #ccc;
    border-radius: 100px;
    background: white;
    &:hover {
      background: $linkColor;
      border: 1px solid $linkColor;
    }
  }
  .node-input {
    top: #{-2 + $portSize/-2}px;
  }
  .node-output {
    bottom: #{-2 + $portSize/-2}px;
  }
  .node-delete {
    position: absolute;
    right: -6px;
    top: -6px;
    font-size: 6px;
    width: 12px;
    height: 12px;
    color: $dangerColor;
    cursor: pointer;
    background: white;
    border: 1px solid $dangerColor;
    border-radius: 100px;
    text-align: center;
    &:hover {
      background: $dangerColor;
      color: white;
    }
  }
}
.selected {
  box-shadow: 0 0 0 2px $linkColor;
}
</style>
