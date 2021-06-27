<template>
  <CCard>
    <CCardHeader> <strong>Load </strong><small>Form</small> </CCardHeader>
    <CCardBody>
      <CInput
        label="Result field"
        horizontal
        v-model="load.result_field"
        placeholder="Please enter result field"
      />
      <CSelect
        label="Type"
        horizontal
        :value.sync="load.type"
        placeholder="Please select db type"
        :options="[{ value: 0, label: 'PgSQL' }]"
      />
      <div v-if="load.type === 0">
        <CSelect
          label="Connection"
          horizontal
          :value.sync="load.pg_sql_load.connection"
          placeholder="Please select connection"
          :options="['a', 'b']"
        />
        <CTextarea
          label="Query"
          horizontal
          v-model="load.pg_sql_load.query"
          placeholder="Enter PGSQL Query to load data"
        />
        <CTextarea
          label="Args"
          horizontal
          v-model="load.pg_sql_load.args"
          placeholder="Arguments line seperated"
        />
      </div>
    </CCardBody>
  </CCard>
</template>

<script>
export default {
  name: "Load",
  data() {
    return {
      load: {
        type: -1,
        result_field: "",
        pg_sql_load: {
          connection: "",
          query: "",
          args: "",
        },
        custom_errors: [],
      },
    };
  },
};
</script>

<style>
</style>