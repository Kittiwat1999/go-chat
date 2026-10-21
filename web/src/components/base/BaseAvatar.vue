<template>
  <v-avatar class="base-avatar" :size="size" :image="avatar || undefined">
    <span v-if="!avatar">{{ initials(displayName) }}</span>
    <span
      v-if="status"
      class="avatar-status"
      :class="`avatar-status-${status}`"
      :aria-label="status"
    />
  </v-avatar>
</template>

<script setup lang="ts">
  withDefaults(defineProps<{
    displayName: string
    avatar?: string
    status?: 'online' | 'offline' | 'away'
    size?: number
  }>(), {
    avatar: '',
    size: 36,
  })

  function initials(name: string) {
    return name
      .trim()
      .split(/\s+/)
      .slice(0, 2)
      .map(part => part.charAt(0))
      .join('')
      .toUpperCase()
  }
</script>

<style scoped>
  .base-avatar { position: relative; flex: 0 0 auto; color: #fff; background: #69a8a2; font-size: 11px; font-weight: 700; }
  .avatar-status { position: absolute; right: 0; bottom: 0; width: 10px; height: 10px; border: 2px solid #17191f; border-radius: 50%; }
  .avatar-status-online { background: #59bd8d; }
  .avatar-status-away { background: #e8b75e; }
  .avatar-status-offline { background: #747b87; }
</style>
