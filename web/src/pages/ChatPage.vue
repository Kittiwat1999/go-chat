<template>
  <main class="chat-shell">
    <ConversationSidebar
      :active-conversation-id="activeConversationId"
      :conversations="conversations"
      :current-user="currentUser"
      :users="users"
      @edit-profile="openProfileDialog"
      @select="selectConversation"
    />

    <ConversationView
      v-if="activeConversation"
      :conversation="activeConversation"
      :current-user="currentUser"
      :messages="activeMessages"
      :users="users"
      @send-message="sendMessage"
    />

    <v-dialog v-model="profileDialog" max-width="420">
      <v-card class="profile-card" rounded="xl">
        <div class="profile-icon">
          <v-icon icon="mdi-account-edit-outline" size="24" />
        </div>

        <h2>Your profile</h2>
        <p>Choose the name other people see in conversations.</p>

        <form @submit.prevent="saveDisplayName">
          <label class="profile-label" for="display-name">Display name</label>

          <v-text-field
            id="display-name"
            v-model="nameDraft"
            autocomplete="name"
            autofocus
            density="comfortable"
            hide-details
            maxlength="32"
            placeholder="e.g. Alex Morgan"
            variant="outlined"
          />

          <div class="profile-actions">
            <v-btn rounded="lg" variant="text" @click="profileDialog = false">Cancel</v-btn>

            <v-btn
              color="primary"
              :disabled="!nameDraft.trim()"
              rounded="lg"
              type="submit"
              variant="flat"
            >
              Save
            </v-btn>
          </div>
        </form>
      </v-card>
    </v-dialog>
  </main>
</template>

<script setup lang="ts">
  import type { ChatMessage, Conversation, User } from '@/types/chat'
  import { computed, ref } from 'vue'
  import ConversationSidebar from '@/components/chat/ConversationSidebar.vue'
  import ConversationView from '@/components/chat/ConversationView.vue'
  import { demoConversations, demoMessages, demoUsers } from '@/data/chat-demo'

  const users = ref<User[]>(demoUsers)
  const conversations = ref<Conversation[]>(demoConversations)
  const messages = ref<ChatMessage[]>(demoMessages)
  const currentUser = computed(() => users.value[0]!)
  const activeConversationId = ref('conv_2')
  const activeConversation = computed(() =>
    conversations.value.find(conversation => conversation._id === activeConversationId.value),
  )
  const activeMessages = computed(() =>
    messages.value.filter(message => message.conversationId === activeConversationId.value),
  )
  const profileDialog = ref(false)
  const nameDraft = ref(currentUser.value.displayName)
  let nextMessageId = 101

  function selectConversation (conversationId: string) {
    activeConversationId.value = conversationId
  }

  function openProfileDialog () {
    nameDraft.value = currentUser.value.displayName
    profileDialog.value = true
  }

  function saveDisplayName () {
    const displayName = nameDraft.value.trim()
    if (!displayName) return

    currentUser.value.displayName = displayName
    profileDialog.value = false
  }

  function sendMessage (content: string) {
    const conversation = activeConversation.value
    if (!conversation || !content.trim()) return

    const now = new Date().toISOString()
    const message: ChatMessage = {
      _id: `msg_${nextMessageId++}`,
      conversationId: conversation._id,
      senderId: currentUser.value._id,
      type: 'text',
      content: content.trim(),
      createdAt: now,
      updatedAt: now,
    }

    messages.value.push(message)
    conversation.lastMessage = {
      messageId: message._id,
      content: message.content,
      senderId: message.senderId,
      createdAt: message.createdAt,
    }
  }
</script>
