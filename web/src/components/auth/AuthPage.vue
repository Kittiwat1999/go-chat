<template>
  <main class="auth-page">
    <div class="auth-glow auth-glow-one" />
    <div class="auth-glow auth-glow-two" />

    <section :aria-labelledby="isRegister ? 'register-title' : 'login-title'" class="auth-card">
      <a aria-label="Loop home" class="auth-brand" href="#">
        <span class="brand-mark"><v-icon icon="mdi-chat-processing" /></span>
        <span>loop<span class="brand-period">.</span></span>
      </a>

      <div class="auth-heading">
        <div class="auth-eyebrow">{{ isRegister ? 'GET STARTED' : 'WELCOME BACK' }}</div>

        <h1 :id="isRegister ? 'register-title' : 'login-title'">
          {{ isRegister ? 'Create your account' : 'Sign in to Loop' }}
        </h1>

        <p>
          {{ isRegister
            ? 'Create an account to start conversations with your team.'
            : 'Enter your details to continue to your conversations.' }}
        </p>
      </div>

      <p v-if="errorMessage" class="auth-feedback auth-feedback-error" role="alert">
        {{ errorMessage }}
      </p>

      <p v-if="successMessage" class="auth-feedback auth-feedback-success" role="status">
        {{ successMessage }}
      </p>

      <form class="auth-form" @submit.prevent="submit">
        <div v-if="isRegister" class="form-field">
          <label for="display-name">Display name</label>

          <v-text-field
            id="display-name"
            v-model="displayName"
            autocomplete="name"
            density="comfortable"
            hide-details
            maxlength="32"
            placeholder="e.g. Alex Morgan"
            prepend-inner-icon="mdi-card-account-details-outline"
            required
            variant="outlined"
          />
        </div>

        <div class="form-field">
          <label for="username">Username</label>

          <v-text-field
            id="username"
            v-model="username"
            autocomplete="username"
            density="comfortable"
            hide-details
            maxlength="30"
            :minlength="isRegister ? 3 : undefined"
            placeholder="Enter your username"
            prepend-inner-icon="mdi-account-outline"
            required
            variant="outlined"
          />
        </div>

        <div class="form-field">
          <label for="password">Password</label>

          <v-text-field
            id="password"
            v-model="password"
            :append-inner-icon="showPassword ? 'mdi-eye-off-outline' : 'mdi-eye-outline'"
            :autocomplete="isRegister ? 'new-password' : 'current-password'"
            density="comfortable"
            hide-details
            :minlength="isRegister ? 6 : undefined"
            placeholder="Enter your password"
            prepend-inner-icon="mdi-lock-outline"
            required
            :type="showPassword ? 'text' : 'password'"
            variant="outlined"
            @click:append-inner="showPassword = !showPassword"
          />
        </div>

        <v-btn
          block
          class="auth-submit"
          color="primary"
          :disabled="isSubmitting"
          :loading="isSubmitting"
          rounded="lg"
          size="large"
          type="submit"
        >
          {{ isSubmitting ? (isRegister ? 'Creating account...' : 'Signing in...') : isRegister ? 'Create account' : 'Sign in' }}
          <v-icon end icon="mdi-arrow-right" />
        </v-btn>
      </form>

      <div class="auth-switch">
        <span>{{ isRegister ? 'Already have an account?' : 'New to Loop?' }}</span>

        <button :disabled="isSubmitting" type="button" @click="toggleMode">
          {{ isRegister ? 'Sign in' : 'Create an account' }}
        </button>
      </div>

      <div class="auth-demo-note">
        <v-icon icon="mdi-information-outline" size="15" />
        {{ isRegister ? 'Registration is connected.' : 'Sign in to continue to your conversations.' }}
      </div>

      <button class="preview-link" :disabled="isSubmitting" type="button" @click="$emit('open-chat-preview')">
        Open chat preview
      </button>
    </section>

    <p class="auth-footer">Simple conversations. Better together.</p>
  </main>
</template>

<script setup lang="ts">
  import { ref } from 'vue'

  const emit = defineEmits<{
    'login-success': []
    'open-chat-preview': []
  }>()

  const isRegister = ref(false)
  const showPassword = ref(false)
  const displayName = ref('')
  const username = ref('')
  const password = ref('')
  const isSubmitting = ref(false)
  const errorMessage = ref('')
  const successMessage = ref('')

  function toggleMode () {
    if (isSubmitting.value) return

    errorMessage.value = ''
    if (!isRegister.value) {
      successMessage.value = ''
    }
    isRegister.value = !isRegister.value
    showPassword.value = false
  }

  async function submitRegister () {
    if (!isRegister.value || isSubmitting.value) return

    errorMessage.value = ''
    successMessage.value = ''
    isSubmitting.value = true

    try {
      let response: Response
      try {
        response = await fetch('/api/users/register', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            username: username.value.trim(),
            displayName: displayName.value.trim(),
            password: password.value,
          }),
        })
      } catch (error) {
        console.error('Registration request failed:', error)
        errorMessage.value = 'Could not connect to the server. Please try again.'
        return
      }

      if (!response.ok) {
        if (response.status === 409) {
          errorMessage.value = 'That username is already taken. Please choose another.'
        } else if (response.status === 400) {
          errorMessage.value = 'Check your username, display name, and password, then try again.'
        } else {
          errorMessage.value = `Registration failed (HTTP ${response.status}). Please try again.`
        }
        return
      }

      password.value = ''
      successMessage.value = 'Account created. Sign in with your new username and password.'
      isRegister.value = false
    } finally {
      isSubmitting.value = false
    }
  }

  async function submitLogin () {
    if (isRegister.value || isSubmitting.value) return

    errorMessage.value = ''
    successMessage.value = ''
    isSubmitting.value = true

    try {
      let response: Response
      try {
        response = await fetch('/api/users/login', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            username: username.value.trim(),
            password: password.value,
          }),
        })
      } catch (error) {
        console.error('Login request failed:', error)
        errorMessage.value = 'Could not connect to the server. Please try again.'
        return
      }

      if (!response.ok) {
        if (response.status === 401) {
          errorMessage.value = 'Invalid username or password.'
        } else if (response.status === 400) {
          errorMessage.value = 'Enter your username and password, then try again.'
        } else {
          errorMessage.value = `Login failed (HTTP ${response.status}). Please try again.`
        }
        return
      }

      let payload: unknown
      try {
        payload = await response.json()
      } catch (error) {
        console.error('Could not read login response:', error)
        errorMessage.value = 'The server returned an invalid response. Please try again.'
        return
      }

      if (
        typeof payload !== 'object'
        || payload === null
        || !('token' in payload)
        || typeof payload.token !== 'string'
        || payload.token.length === 0
      ) {
        errorMessage.value = 'The server returned an invalid response. Please try again.'
        return
      }

      try {
        sessionStorage.setItem('authToken', payload.token)
      } catch (error) {
        console.error('Could not store login token:', error)
        errorMessage.value = 'Could not save your session in this browser.'
        return
      }

      password.value = ''
      emit('login-success')
    } finally {
      isSubmitting.value = false
    }
  }

  function submit () {
    if (isRegister.value) {
      void submitRegister()
      return
    }

    void submitLogin()
  }
</script>

<style scoped>
  .auth-page {
    position: relative;
    display: grid;
    min-height: 100vh;
    min-height: 100dvh;
    overflow: hidden;
    place-items: center;
    padding: 48px 20px 62px;
    background: #111318;
  }

  .auth-glow {
    position: absolute;
    width: 340px;
    height: 340px;
    border-radius: 50%;
    background: #746bda;
    filter: blur(150px);
    opacity: .13;
    pointer-events: none;
  }

  .auth-glow-one { top: -190px; left: calc(50% - 370px); }
  .auth-glow-two { right: calc(50% - 390px); bottom: -220px; background: #4a9e9c; }

  .auth-card {
    position: relative;
    width: min(100%, 420px);
    padding: 34px;
    border: 1px solid #2a2d37;
    border-radius: 18px;
    background: #17191f;
    box-shadow: 0 22px 80px #00000035;
  }

  .auth-brand {
    display: inline-flex;
    align-items: center;
    gap: 9px;
    margin-bottom: 35px;
    color: #e6e8ef;
    font-size: 22px;
    font-weight: 750;
    letter-spacing: -1.1px;
    text-decoration: none;
  }

  .brand-mark {
    display: grid;
    width: 31px;
    height: 31px;
    place-items: center;
    border-radius: 10px;
    color: #fff;
    background: #958bff;
    font-size: 18px;
  }

  .brand-period { color: #958bff; }
  .auth-heading { margin-bottom: 27px; }
  .auth-eyebrow { margin-bottom: 8px; color: #b2aaff; font-size: 9px; font-weight: 700; letter-spacing: 1.2px; }
  .auth-heading h1 { margin: 0; color: #eff0f5; font-size: 25px; font-weight: 700; letter-spacing: -.6px; }
  .auth-heading p { margin: 9px 0 0; color: #aeb4c0; font-size: 12px; line-height: 1.6; }
  .auth-form { display: grid; gap: 19px; }
  .form-field label { display: block; margin-bottom: 7px; color: #c2c6d0; font-size: 11px; font-weight: 600; }
  .auth-feedback { margin: 0 0 18px; font-size: 11px; line-height: 1.5; }
  .auth-feedback-error { color: #ff9c9c; }
  .auth-feedback-success { color: #91d8b1; }

  .auth-submit {
    height: 46px;
    margin-top: 4px;
    color: #fff;
    font-size: 12px;
    font-weight: 650;
    letter-spacing: 0;
    text-transform: none;
  }

  .auth-switch { display: flex; justify-content: center; gap: 5px; margin-top: 24px; color: #969dac; font-size: 11px; }
  .auth-switch button { padding: 0; border: 0; color: #b2aaff; background: none; font: inherit; font-weight: 650; }
  .auth-switch button:hover { color: #d0caff; }
  .auth-switch button:disabled, .preview-link:disabled { cursor: not-allowed; opacity: .6; }
  .auth-demo-note { display: flex; align-items: center; justify-content: center; gap: 5px; margin-top: 26px; padding-top: 17px; border-top: 1px solid #2a2d37; color: #858d9b; font-size: 9px; text-align: center; }
  .preview-link { display: block; margin: 14px auto 0; padding: 4px; border: 0; color: #969dac; background: none; font: inherit; font-size: 10px; }
  .preview-link:hover { color: #b2aaff; }
  .auth-footer { position: absolute; bottom: 22px; margin: 0; color: #777f8d; font-size: 10px; letter-spacing: .2px; }

  @media (max-width: 480px) {
    .auth-page { padding: 24px 16px 56px; }
    .auth-card { padding: 27px 22px; border-radius: 15px; }
    .auth-brand { margin-bottom: 29px; }
    .auth-heading h1 { font-size: 23px; }
    .auth-footer { bottom: 16px; }
  }
</style>
