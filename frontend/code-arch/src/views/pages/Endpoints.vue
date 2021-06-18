<template>
  <div>
    <CRow>
      <CCol>
        <CWidgetDropdown
          @click.native="$router.push('/services/endpoints/editor')"
          color="dark"
          header="Service 1"
          text="This is a sample description. Please feel free to edit it."
        >
          <template #default>
            <CDropdown color="transparent p-0" placement="bottom-end">
              <template #toggler-content>
                <CIcon name="cil-settings" />
              </template>
              <CDropdownItem>Action</CDropdownItem>
              <CDropdownItem>Another action</CDropdownItem>
              <CDropdownItem>Something else here...</CDropdownItem>
              <CDropdownItem disabled>Disabled action</CDropdownItem>
            </CDropdown>
          </template>
          <template #footer> </template>
        </CWidgetDropdown>
      </CCol>
    </CRow>
    <CRow>
      <CCol>
        <CWidgetDropdown
          @click.native="$router.push('/login')"
          color="dark"
          header="Service 1"
          text="This is a sample description. Please feel free to edit it."
        >
          <template #default>
            <CDropdown color="transparent p-0" placement="bottom-end">
              <template #toggler-content>
                <CIcon name="cil-settings" />
              </template>
              <CDropdownItem>Action</CDropdownItem>
              <CDropdownItem>Another action</CDropdownItem>
              <CDropdownItem>Something else here...</CDropdownItem>
              <CDropdownItem disabled>Disabled action</CDropdownItem>
            </CDropdown>
          </template>
          <template #footer> </template>
        </CWidgetDropdown>
      </CCol>
    </CRow>
    <CRow>
      <CCol>
        <CWidgetDropdown
          @click.native="$router.push('/login')"
          color="dark"
          header="Service 1"
          text="This is a sample description. Please feel free to edit it."
        >
          <template #default>
            <CDropdown color="transparent p-0" placement="bottom-end">
              <template #toggler-content>
                <CIcon name="cil-settings" />
              </template>
              <CDropdownItem>Action</CDropdownItem>
              <CDropdownItem>Another action</CDropdownItem>
              <CDropdownItem>Something else here...</CDropdownItem>
              <CDropdownItem disabled>Disabled action</CDropdownItem>
            </CDropdown>
          </template>
          <template #footer> </template>
        </CWidgetDropdown>
      </CCol>
    </CRow>
    <CRow>
      <CCol>
        <CWidgetDropdown
          @click.native="$router.push('/login')"
          color="dark"
          header="Service 1"
          text="This is a sample description. Please feel free to edit it."
        >
          <template #default>
            <CDropdown color="transparent p-0" placement="bottom-end">
              <template #toggler-content>
                <CIcon name="cil-settings" />
              </template>
              <CDropdownItem>Action</CDropdownItem>
              <CDropdownItem>Another action</CDropdownItem>
              <CDropdownItem>Something else here...</CDropdownItem>
              <CDropdownItem disabled>Disabled action</CDropdownItem>
            </CDropdown>
          </template>
          <template #footer> </template>
        </CWidgetDropdown>
      </CCol>
    </CRow>
    <CRow>
      <CCol>
        <CWidgetDropdown
          @click.native="$router.push('/login')"
          color="dark"
          header="Service 1"
          text="This is a sample description. Please feel free to edit it."
        >
          <template #default>
            <CDropdown color="transparent p-0" placement="bottom-end">
              <template #toggler-content>
                <CIcon name="cil-settings" />
              </template>
              <CDropdownItem>Action</CDropdownItem>
              <CDropdownItem>Another action</CDropdownItem>
              <CDropdownItem>Something else here...</CDropdownItem>
              <CDropdownItem disabled>Disabled action</CDropdownItem>
            </CDropdown>
          </template>
          <template #footer> </template>
        </CWidgetDropdown>
      </CCol>
    </CRow>
  </div>
</template>

<script>
export default {
  name: "Endpoints",
  components: {},
  methods: {
    clickMe() {
      alert("jegfj");
    },
  },
};
</script>

<style lang="scss" scoped>
.card {
  cursor: pointer;
}
</style>