<template>
  <section class="message-list" aria-label="Messages" aria-live="polite">
    <article
      v-for="message in messages"
      :key="message._id"
      class="message-row"
      :class="{ 'message-row-own': message.senderId === currentUser._id }"
    >
      <BaseAvatar
        :display-name="sender(message.senderId)?.displayName ?? 'Unknown user'"
        :avatar="sender(message.senderId)?.avatar"
        :size="35"
      />
      <div class="message-content">
        <div class="message-heading">
          <span class="message-author">{{ sender(message.senderId)?.displayName ?? 'Unknown user' }}</span>
          <span v-if="message.senderId === currentUser._id" class="you-label">YOU</span>
          <time :datetime="message.createdAt">{{ formatTime(message.createdAt) }}</time>
        </div>
        <p>{{ message.content }}</p>
      </div>
    </article>
    <p v-if="messages.length === 0" class="empty-state">No messages yet. Start the conversation.</p>
  </section>
</template>

<script setup lang="ts">
  import BaseAvatar from '@/components/base/BaseAvatar.vue'
  import type { ChatMessage, User } from '@/types/chat'

  const props = defineProps<{
    currentUser: User
    messages: ChatMessage[]
    users: User[]
  }>()

  function sender(userId: string) {
    return props.users.find(user => user._id === userId)
  }

  function formatTime(timestamp: string) {
    return new Intl.DateTimeFormat(undefined, { hour: 'numeric', minute: '2-digit' }).format(new Date(timestamp))
  }
</script>

<style scoped>
  .message-list { display: grid; align-content: start; gap: 22px; padding: 3px 0 24px; }
  .message-row { display: flex; align-items: flex-start; gap: 12px; }
  .message-row-own { flex-direction: row-reverse; justify-content: flex-start; }
  .message-content { min-width: 0; padding-top: 1px; }
  .message-heading { display: flex; align-items: center; gap: 8px; }
  .message-row-own .message-content { text-align: right; }
  .message-row-own .message-heading { justify-content: flex-end; }
  .message-author { color: #e0e2e9; font-size: 12px; font-weight: 650; }
  .message-heading time { color: var(--muted); font-size: 10px; }
  .you-label { padding: 2px 4px; border-radius: 3px; color: var(--accent-text); background: var(--accent-soft); font-size: 8px; font-weight: 700; letter-spacing: .35px; }
  .message-content p { margin: 5px 0 0; color: var(--text-soft); font-size: 12px; line-height: 1.65; overflow-wrap: anywhere; }
  .empty-state { margin: 0; color: var(--muted); font-size: 12px; text-align: center; }
</style>
