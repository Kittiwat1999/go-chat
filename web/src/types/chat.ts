export type UserStatus = 'online' | 'offline' | 'away'
export type ConversationType = 'direct' | 'group'

export interface User {
  _id: string
  username: string
  displayName: string
  avatar: string
  status: UserStatus
  createdAt: string
  updatedAt: string
}

export interface LastMessage {
  messageId: string
  content: string
  senderId: string
  createdAt: string
}

export interface Conversation {
  _id: string
  type: ConversationType
  name?: string
  participants: string[]
  admins?: string[]
  lastMessage: LastMessage
  createdAt: string
}

export interface ChatMessage {
  _id: string
  conversationId: string
  senderId: string
  type: 'text'
  content: string
  createdAt: string
  updatedAt: string
}
