<template>
  <v-form ref="form" v-model="valid" lazy-validation>
    <v-text-field v-model="email" :rules="emailRules" label="Почта" required></v-text-field>

    <v-text-field
      v-model="password"
      :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
      :rules="passwordRules"
      :type="showPassword ? 'text' : 'password'"
      name="input-10-1"
      label="Пароль"
      counter
      @click:append="showPassword = !showPassword"
      required
    ></v-text-field>

    <v-btn :disabled="!valid" color="success" class="mr-4" @click="signIn">
      Войти
    </v-btn>
    <v-snackbar v-model="snackbar" color="red" :timeout="5 * 1000">
      {{ text }}
      <template v-slot:action="{ attrs }">
        <v-btn color="black" text v-bind="attrs" @click="snackbar = false">
          Закрыть
        </v-btn>
      </template>
    </v-snackbar>
  </v-form>
</template>

<script>
export default {
  data: () => ({
    valid: true,
    email: 'new_email@gmail.com',
    emailRules: [v => !!v || 'E-mail обязателен', v => /.+@.+\..+/.test(v) || 'Введите почту'],
    password: '987654',
    passwordRules: [v => !!v || 'Пароль обязателен'],
    showPassword: false,
    // snackbar
    snackbar: false,
    text: 'Неправильный логин или пароль',
  }),

  methods: {
    async signIn() {
      // запускаем валидацию формы
      if (!this.$refs.form.validate()) {
        return
      }

      try {
        await this.$store.dispatch('auth/signIn', {
          email: this.email,
          password: this.password,
        })
        // редирект на страницу меню
        this.$router.push({ name: 'UserOrder' })
      } catch (err) {
        console.error(err)
        this.snackbar = true
      }
    },
  },
}
</script>
