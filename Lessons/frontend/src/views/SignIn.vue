<template>
  <v-form ref="form" v-model="valid" lazy-validation>
    <v-text-field
      v-model="email"
      :rules="emailRules"
      label="Почта"
      required
    ></v-text-field>

    <v-text-field
      v-model="password"
      :rules="passwordRules"
      label="Пароль"
      required
    ></v-text-field>

    <v-btn :disabled="!valid" color="success" class="mr-4" @click="signIn">
      Войти
    </v-btn>
  </v-form>
</template>

<script>
export default {
  data: () => ({
    valid: true,
    email: '',
    emailRules: [
      v => !!v || 'E-mail обязателен',
      v => /.+@.+\..+/.test(v) || 'Введите почту',
    ],
    password: '',
    passwordRules: [v => !!v || 'Пароль обязателен'],
  }),

  methods: {
    signIn() {
      // запускаем валидацию формы
      if (!this.$refs.form.validate()) {
        return
      }

      console.log('Send sign_in request')
    },
  },
}
</script>
