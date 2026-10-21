<template>
  <aside class="sidebar">
    <a aria-label="Loop home" class="brand" href="#">
      <span class="brand-mark"><v-icon icon="mdi-chat-processing" /></span>
      <span>loop<span class="brand-period">.</span></span>
    </a>

    <div class="section-heading">
      <span>CONVERSATIONS</span>
      <span class="conversation-count">{{ conversations.length }}</span>
    </div>

    <nav aria-label="Conversations" class="conversation-list">
      <button
        v-for="conversation in conversations"
        :key="conversation._id"
        :aria-current="conversation._id === activeConversationId ? 'page' : undefined"
        class="conversation-button"
        :class="{ 'conversation-button-active': conversation._id === activeConversationId }"
        type="button"
        @click="$emit('select', conversation._id)"
      >
        <BaseAvatar
          :avatar="conversationAvatar(conversation)"
          :display-name="conversationTitle(conversation)"
          :size="36"
          :status="conversation.type === 'direct' ? directUser(conversation)?.status : undefined"
        />

        <span class="conversation-copy">
          <span class="conversation-heading">
            <span class="conversation-title">{{ conversationTitle(conversation) }}</span>
            <time>{{ formatTime(conversation.lastMessage.createdAt) }}</time>
          </span>

          <span class="conversation-preview">{{ conversation.lastMessage.content }}</span>

          <span class="conversation-type">
            <v-icon
              :icon="conversation.type === 'group' ? 'mdi-account-group-outline' : 'mdi-account-outline'"
              size="13"
            />
            {{ conversation.type === 'group' ? `${conversation.participants.length} members` : 'Direct message' }}
          </span>
        </span>
      </button>
    </nav>

    <div class="sidebar-spacer" />

    <button class="profile-button" type="button" @click="$emit('edit-profile')">
      <BaseAvatar
        :avatar="currentUser.avatar"
        :display-name="currentUser.displayName"
        :size="34"
        :status="currentUser.status"
      />

      <span class="profile-copy">
        <span class="profile-name">{{ currentUser.displayName }}</span>
        <span class="profile-username">@{{ currentUser.username }}</span>
      </span>

      <v-icon icon="mdi-cog-outline" size="19" />
    </button>
  </aside>
</template>

<script setup lang="ts">
  import type { Conversation, User } from '@/types/chat'
  import BaseAvatar from '@/components/base/BaseAvatar.vue'

  const props = defineProps<{
    activeConversationId: string
    conversations: Conversation[]
    currentUser: User
    users: User[]
  }>()

  defineEmits<{
    'select': [conversationId: string]
    'edit-profile': []
  }>()

  function directUser (conversation: Conversation) {
    const otherUserId = conversation.participants.find(userId => userId !== props.currentUser._id)
    return props.users.find(user => user._id === otherUserId)
  }

  function conversationTitle (conversation: Conversation) {
    if (conversation.type === 'group') return conversation.name ?? 'Unnamed group'
    return directUser(conversation)?.displayName ?? 'Unknown user'
  }

  function conversationAvatar (conversation: Conversation) {
    return conversation.type === 'direct' ? directUser(conversation)?.avatar ?? '' : ''
  }

  function formatTime (timestamp: string) {
    return new Intl.DateTimeFormat(undefined, { hour: 'numeric', minute: '2-digit' }).format(new Date(timestamp))
  }
</script>

<style scoped>
  .sidebar { display: flex; flex: 0 0 290px; flex-direction: column; min-height: 100vh; padding: 27px 15px 15px; border-right: 1px solid var(--line); background: var(--sidebar); }
  .brand { display: flex; align-items: center; gap: 9px; margin: 0 0 34px 5px; color: var(--ink); font-size: 22px; font-weight: 750; letter-spacing: -1.1px; text-decoration: none; }
  .brand-mark { display: grid; width: 31px; height: 31px; place-items: center; border-radius: 10px; color: #fff; background: var(--purple); font-size: 18px; }
  .brand-period { color: var(--purple); }
  .section-heading { display: flex; align-items: center; justify-content: space-between; margin: 0 7px 10px; color: #a1a8b2; font-size: 10px; font-weight: 700; letter-spacing: 1.05px; }
  .conversation-count { color: var(--muted); font-size: 10px; letter-spacing: 0; }
  .conversation-list { display: grid; gap: 5px; }
  .conversation-button { display: flex; align-items: flex-start; gap: 10px; width: 100%; padding: 10px 8px; border: 0; border-radius: 9px; color: var(--text-soft); background: transparent; text-align: left; transition: background .15s ease, color .15s ease; }
  .conversation-button:hover { color: var(--ink); background: var(--surface-hover); }
  .conversation-button-active { color: var(--ink); background: var(--accent-soft); }
  .conversation-copy { display: grid; flex: 1; min-width: 0; gap: 5px; padding-top: 1px; }
  .conversation-heading { display: flex; align-items: center; justify-content: space-between; gap: 5px; }
  .conversation-title { overflow: hidden; color: #e0e2e9; font-size: 12px; font-weight: 650; text-overflow: ellipsis; white-space: nowrap; }
  .conversation-heading time { flex: 0 0 auto; color: var(--muted); font-size: 9px; }
  .conversation-preview { overflow: hidden; color: var(--muted); font-size: 10px; text-overflow: ellipsis; white-space: nowrap; }
  .conversation-type { display: flex; align-items: center; gap: 4px; color: #858d9b; font-size: 9px; }
  .sidebar-spacer { flex: 1; min-height: 35px; }
  .profile-button { display: flex; align-items: center; gap: 10px; width: 100%; padding: 10px 6px 3px; border: 0; border-top: 1px solid var(--line); color: var(--ink); background: transparent; text-align: left; }
  .profile-copy { display: grid; flex: 1; min-width: 0; gap: 4px; }
  .profile-name { overflow: hidden; font-size: 11px; font-weight: 650; text-overflow: ellipsis; white-space: nowrap; }
  .profile-username { color: var(--muted); font-size: 10px; }
  .profile-button > .v-icon { color: var(--muted); }

  @media (max-width: 760px) {
    .sidebar { flex-basis: 230px; padding-right: 11px; padding-left: 11px; }
  }
  @media (max-width: 580px) {
    .sidebar { display: none; }
  }
</style>
