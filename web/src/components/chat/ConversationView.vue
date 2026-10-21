<template>
  <section class="chat-panel" aria-label="Conversation">
    <header class="topbar">
      <div class="mobile-brand">
        <span class="brand-mark"><v-icon icon="mdi-chat-processing" /></span>
        <span>loop<span class="brand-period">.</span></span>
      </div>
      <div class="conversation-heading">
        <BaseAvatar
          :display-name="conversationTitle"
          :avatar="conversation.type === 'direct' ? directUser?.avatar : ''"
          :size="34"
        />
        <div class="heading-copy">
          <span class="heading-title">{{ conversationTitle }}</span>
          <span class="heading-subtitle">
            {{ conversation.type === 'group' ? 'Group conversation' : 'Direct message' }}
          </span>
        </div>
      </div>
      <div class="topbar-actions">
        <span class="connection-status"><span /> Demo chat</span>
      </div>
    </header>

    <div class="conversation">
      <div class="conversation-intro">
        <div class="eyebrow">{{ conversation.type === 'group' ? 'GROUP CONVERSATION' : 'DIRECT MESSAGE' }}</div>
        <h1>{{ conversationTitle }}</h1>
        <p v-if="conversation.type === 'group'">{{ participantNames }}</p>
        <p v-else>Conversation between you and {{ conversationTitle }}.</p>
        <div class="conversation-meta">
          <span>
            <v-icon
              :icon="conversation.type === 'group' ? 'mdi-account-group-outline' : 'mdi-account-outline'"
              size="16"
            />
            {{ conversation.participants.length }} {{ conversation.participants.length === 1 ? 'member' : 'members' }}
          </span>
          <span class="meta-divider" />
          <span>
            <v-icon icon="mdi-message-text-outline" size="15" />
            {{ messages.length }} {{ messages.length === 1 ? 'message' : 'messages' }}
          </span>
        </div>
      </div>

      <div class="message-divider"><span>Messages</span></div>
      <ChatMessageList :current-user="currentUser" :messages="messages" :users="users" />
      <MessageComposer :conversation-name="conversationTitle" @send="emit('sendMessage', $event)" />
      <p class="demo-caption">
        <v-icon icon="mdi-information-outline" size="14" />
        Preview only — messages are not saved or shared.
      </p>
    </div>
  </section>
</template>

<script setup lang="ts">
  import { computed } from 'vue'
  import BaseAvatar from '@/components/base/BaseAvatar.vue'
  import ChatMessageList from '@/components/chat/ChatMessageList.vue'
  import MessageComposer from '@/components/chat/MessageComposer.vue'
  import type { ChatMessage, Conversation, User } from '@/types/chat'

  const props = defineProps<{
    conversation: Conversation
    currentUser: User
    messages: ChatMessage[]
    users: User[]
  }>()

  const emit = defineEmits<{ sendMessage: [content: string] }>()

  const directUser = computed(() => {
    const userId = props.conversation.participants.find(id => id !== props.currentUser._id)
    return props.users.find(user => user._id === userId)
  })

  const conversationTitle = computed(() => {
    if (props.conversation.type === 'group') return props.conversation.name ?? 'Unnamed group'
    return directUser.value?.displayName ?? 'Unknown user'
  })

  const participantNames = computed(() =>
    props.conversation.participants
      .map(id => props.users.find(user => user._id === id)?.displayName)
      .filter((name): name is string => Boolean(name))
      .join(', '),
  )
</script>

<style scoped>
  .chat-panel { display: flex; flex: 1; min-width: 0; flex-direction: column; }
  .topbar { display: flex; align-items: center; justify-content: space-between; height: 68px; padding: 0 34px; border-bottom: 1px solid var(--line); }
  .conversation-heading { display: flex; align-items: center; gap: 10px; min-width: 0; }
  .heading-copy { display: grid; gap: 3px; }
  .heading-title { color: #e0e2e9; font-size: 12px; font-weight: 650; }
  .heading-subtitle { color: var(--muted); font-size: 10px; }
  .topbar-actions { display: flex; align-items: center; gap: 13px; }
  .connection-status { display: flex; align-items: center; gap: 7px; color: var(--muted); font-size: 11px; }
  .connection-status > span { width: 6px; height: 6px; border-radius: 50%; background: #59bd8d; }
  .mobile-brand { display: none; align-items: center; gap: 9px; color: var(--ink); font-size: 19px; font-weight: 750; letter-spacing: -1.1px; }
  .brand-mark { display: grid; width: 28px; height: 28px; place-items: center; border-radius: 9px; color: #fff; background: var(--purple); font-size: 16px; }
  .brand-period { color: var(--purple); }
  .conversation { display: flex; width: min(100%, 840px); min-height: calc(100vh - 68px); flex: 1; flex-direction: column; align-self: center; padding: 49px 44px 20px; }
  .conversation-intro { padding: 0 0 25px; }
  .eyebrow { margin-bottom: 8px; color: #a1a8b2; font-size: 9px; font-weight: 700; letter-spacing: 1.2px; }
  .conversation-intro h1 { margin: 0; color: #eff0f5; font-size: 25px; font-weight: 700; letter-spacing: -.65px; }
  .conversation-intro > p { margin: 9px 0 15px; color: var(--text-soft); font-size: 13px; line-height: 1.6; }
  .conversation-meta { display: flex; align-items: center; gap: 13px; color: var(--muted); font-size: 11px; }
  .conversation-meta > span:not(.meta-divider) { display: flex; align-items: center; gap: 5px; }
  .meta-divider { width: 1px; height: 12px; background: var(--line); }
  .message-divider { position: relative; display: flex; justify-content: center; margin: 0 0 21px; border-top: 1px solid var(--line); }
  .message-divider span { position: relative; top: -9px; padding: 0 12px; color: #a1a8b2; background: #111318; font-size: 10px; }
  .demo-caption { display: flex; align-items: center; justify-content: center; gap: 5px; margin: 10px 0 0; color: #777f8d; font-size: 9px; }

  @media (max-width: 760px) {
    .topbar { padding: 0 22px; }
    .conversation { padding: 38px 26px 18px; }
  }
  @media (max-width: 580px) {
    .topbar { height: 60px; padding: 0 15px; }
    .mobile-brand { display: flex; }
    .mobile-brand .brand-mark { width: 27px; height: 27px; }
    .conversation-heading { display: none; }
    .conversation { min-height: calc(100vh - 60px); padding: 32px 18px 15px; }
    .conversation-intro h1 { font-size: 22px; }
    .conversation-intro > p { font-size: 12px; }
    .conversation-meta { gap: 9px; font-size: 10px; }
    .demo-caption { font-size: 8px; }
    .connection-status { font-size: 10px; }
  }
</style>
