<template>
  <div>
    <CRow>
      <CCol lg="6"> </CCol>
      <CCol lg="6" class="d-none d-md-block">
        <CButton class="ml-1 float-right" size="lg" color="success"
          >Save</CButton
        >
        <CButton
          @click="showAddComponentModal"
          class="float-right"
          size="lg"
          color="info"
          >Add</CButton
        >
      </CCol>
    </CRow>
    <CRow class="mt-1">
      <CCol lg="6">
        <simple-flowchart
          :scene="scene"
          @nodeClick="nodeClick"
          @nodeDelete="nodeDelete"
          @linkBreak="linkBreak"
          @linkAdded="linkAdded"
          @canvasClick="canvasClick"
          :height="400"
        />
      </CCol>
      <CCol lg="6" class="editor-form">
        <EditorComponent></EditorComponent>
        <!-- <CCard>
          <CCardHeader>
            <strong>Company </strong><small>Form</small>
          </CCardHeader>
          <CCardBody>
            <CInput label="Company" placeholder="Enter your company name" />
            <CInput label="Company" placeholder="Enter your company name" />
            <CInput label="Company" placeholder="Enter your company name" />
            <CInput label="Company" placeholder="Enter your company name" />
            <CInput label="Company" placeholder="Enter your company name" />
            <CInput label="VAT" placeholder="PL1234567890" />
            <CInput label="Street" placeholder="Enter street name" />
            <CRow>
              <CCol sm="8">
                <CInput label="City" placeholder="Enter your city" />
              </CCol>
              <CCol sm="4">
                <CInput label="Postal code" placeholder="Postal code" />
              </CCol>
            </CRow>
            <CInput label="Country" placeholder="Country name" />
          </CCardBody>
        </CCard> -->
      </CCol>
    </CRow>
    <CModal
      :show.sync="darkModal"
      :no-close-on-backdrop="true"
      :centered="true"
      title="Modal title 2"
      size="lg"
      color="dark"
    >
      <CCard>
        <CCardBody>
          <CForm>
            <CInput
              label="Name"
              size="lg"
              v-model="newComponentName"
              placeholder="Name your component"
            />
            <CSelect
              :value.sync="selectedComponent"
              size="lg"
              label="Component"
              placeholder="Please select component"
              :options="['load', 'store', 'code']"
            />
          </CForm>
        </CCardBody>
      </CCard>
      <template #header>
        <h6 class="modal-title">Choose Component</h6>
        <CButtonClose @click="darkModal = false" class="text-white" />
      </template>
      <template #footer>
        <CButton @click="addNode" color="success">Add</CButton>
      </template>
    </CModal>
  </div>
</template>

<script>
import SimpleFlowchart from "@/components/vue-simple-flowchart";
import EditorComponent from "@/components/editor/EditorComponent";
export default {
  name: "Editor",
  components: {
    SimpleFlowchart,
    EditorComponent,
  },
  mounted() {},
  data() {
    return {
      newComponentName: "",
      selectedComponent: "",
      darkModal: false,
      scene: {
        centerX: 1024,
        centerY: 140,
        scale: 1,
        nodes: [
          {
            id: -1,
            x: -812,
            y: -138,
            type: "internal",
            label: "request",
          },
          {
            id: -2,
            x: -812,
            y: 207,
            type: "internal",
            label: "response",
          },
        ],
        links: [],
      },
    };
  },
  methods: {
    showAddComponentModal() {
      this.newComponentName = "";
      this.selectedComponent = "";
      this.darkModal = true;
    },
    canvasClick(e) {
      // console.log("canvas Click, event:", e);
    },
    addNode() {
      let maxID = Math.max(
        0,
        ...this.scene.nodes.map((link) => {
          return link.id;
        })
      );
      this.scene.nodes.push({
        id: maxID + 1,
        x: -665,
        y: -138,
        type: this.selectedComponent,
        label: this.newComponentName,
      });
      this.selectedComponent = "";
      this.newComponentName = "";
      this.darkModal = false;
    },
    nodeClick(id) {
      console.log(id);
      this.$store.commit("setCurrentComponent", id);
      // console.log(JSON.stringify(this.scene));
      // console.log("node click", id);
    },
    nodeDelete(id) {
      // console.log(JSON.stringify(this.scene));
      // console.log("node delete", id);
    },
    linkBreak(id) {
      // console.log("link break", id);
    },
    linkAdded(link) {
      // console.log("new link added:", link);
    },
  },
};
</script>

<style lang="scss">
.editor-form {
  overflow-y: scroll;
  overflow-x: hidden;
  height: 420px;
}
</style>
