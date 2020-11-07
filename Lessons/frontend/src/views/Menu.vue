<template>
  <v-data-table
    :headers="headers"
    :items="menu"
    :items-per-page="5"
    class="elevation-1"
  >
    <template v-slot:[`item.created_at`]="{ item }">
        {{ item.created_at | formatDate }}
    </template>
    <template v-slot:[`item.updated_at`]="{ item }">
        {{ item.updated_at | formatDate }}
    </template>
  </v-data-table>
</template>

<script>
import axios from 'axios'

export default {
  data () {
    return {
      menu: [],
      headers: [
        { text: 'ID', value: 'id' },
        { text: 'Название', value: 'name' },
        { text: 'Стоимость', value: 'price' },
        { text: 'Описание', value: 'description' },
        { text: 'Вес (г.)', value: 'weight' },
        { text: 'Создано', value: 'created_at' },
        { text: 'Изменено', value: 'updated_at' },
      ]
    }
  },

  async created () {
    this.init()
  },

  methods: {
    async init() {
      const res = await axios.get('http://localhost:80/menu')
      this.menu = res.data
      console.log('we are in init function!!!')
    }
  }
}
</script>