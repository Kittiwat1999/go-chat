<template>
  <form class="composer" @submit.prevent="submit">
    <label class="sr-only" for="message-input">Write a message</label>
    <textarea
      id="message-input"
      v-model="draft"
      rows="1"
      maxlength="2000"
      :placeholder="`Message ${conversationName}`"
      @keydown.enter.exact.prevent="submit"
    />
    <div class="composer-footer">
      <span class="composer-hint"><kbd>↵</kbd> to send <span>·</span> <kbd>shift + ↵</kbd> for a new line</span>
      <button class="send-button" type="submit" :disabled="!draft.trim()" aria-label="Send message">
        <v-icon icon="mdi-arrow-up" size="19" />
      </button>
    </div>
  </form>
</template>

<script setup lang="ts">
  import { ref } from 'vue'

  defineProps<{ conversationName: string }>()
  const emit = defineEmits<{ send: [content: string] }>()
  const draft = ref('')

  function submit() {
    const content = draft.value.trim()
    if (!content) return
    emit('send', content)
    draft.value = ''
  }
</script>

<style scoped>
  .composer { margin-top: auto; padding: 13px 14px 10px; border: 1px solid #343843; border-radius: 12px; background: var(--surface); box-shadow: 0 5px 20px #00000030; }
  .composer:focus-within { border-color: #7169bf; box-shadow: 0 0 0 3px #766be914; }
  .composer textarea { display: block; width: 100%; min-height: 27px; max-height: 120px; resize: vertical; border: 0; outline: 0; color: var(--ink); background: transparent; font-size: 12px; line-height: 1.6; }
  .composer textarea::placeholder { color: #818997; }
  .composer-footer { display: flex; align-items: center; justify-content: space-between; margin-top: 8px; }
  .composer-hint { color: var(--muted); font-size: 9px; }
  .composer-hint kbd { padding: 2px 4px; border: 1px solid #3b3f4a; border-radius: 4px; color: var(--text-soft); background: #292c35; font: inherit; }
  .composer-hint > span { padding: 0 4px; color: #626976; }
  .send-button { display: grid; width: 29px; height: 29px; place-items: center; border: 0; border-radius: 8px; color: #fff; background: var(--purple); transition: opacity .15s ease, transform .15s ease; }
  .send-button:hover:not(:disabled) { transform: translateY(-1px); background: #6258d6; }
  .send-button:disabled { cursor: default; opacity: .55; }
  .sr-only { position: absolute; width: 1px; height: 1px; overflow: hidden; clip: rect(0, 0, 0, 0); white-space: nowrap; clip-path: inset(50%); }
</style>
