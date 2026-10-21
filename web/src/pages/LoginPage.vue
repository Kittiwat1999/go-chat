<template>
  <AuthPage @login-success="openChat" @open-chat-preview="openChatPreview" />
</template>

<script setup lang="ts">
  import { useRouter } from 'vue-router'
  import AuthPage from '@/components/auth/AuthPage.vue'

  const router = useRouter()

  function openChatPreview () {
    void router.push({ name: 'chat' })
  }

  function openChat () {
    void router.push({ name: 'chat' })
  }
</script>
