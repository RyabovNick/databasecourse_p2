<template>
  <v-data-table :headers="headers" :items="order" :items-per-page="5" class="elevation-1">
    <template v-slot:[`item.created_at`]="{ item }">
      {{ item.created_at | formatDate }}
    </template>
  </v-data-table>
</template>

<script>
export default {
  data() {
    return {
      order: [],
      headers: [
        { text: 'ID', value: 'id' },
        { text: 'ID Клиента', value: 'client_id' },
        { text: 'Дата заказа', value: 'created_at' },
      ],
    }
  },

  async created() {
    this.init()
  },

  methods: {
    async init() {
      const res = await this.$axios.get('/user_order')
      this.order = res.data
    },
  },
}
</script>
