import type { ChatMessage, Conversation, User } from '@/types/chat'

const userCreatedAt = '2026-10-06T20:00:00Z'

export const demoUsers: User[] = [
  {
    _id: 'user_1',
    username: 'kittiwat',
    displayName: 'Kittiwat',
    avatar: '',
    status: 'online',
    createdAt: userCreatedAt,
    updatedAt: userCreatedAt,
  },
  {
    _id: 'user_2',
    username: 'maya',
    displayName: 'Maya Chen',
    avatar: '',
    status: 'online',
    createdAt: userCreatedAt,
    updatedAt: userCreatedAt,
  },
  {
    _id: 'user_3',
    username: 'jordan',
    displayName: 'Jordan Lee',
    avatar: '',
    status: 'away',
    createdAt: userCreatedAt,
    updatedAt: userCreatedAt,
  },
]

export const demoConversations: Conversation[] = [
  {
    _id: 'conv_1',
    type: 'direct',
    participants: ['user_1', 'user_2'],
    lastMessage: {
      messageId: 'msg_99',
      content: 'Hey, are we still on for the design review?',
      senderId: 'user_2',
      createdAt: '2026-10-21T09:24:00Z',
    },
    createdAt: userCreatedAt,
  },
  {
    _id: 'conv_2',
    type: 'group',
    name: 'Go Developers',
    participants: ['user_1', 'user_2', 'user_3'],
    admins: ['user_1'],
    lastMessage: {
      messageId: 'msg_100',
      content: 'Who knows Gin?',
      senderId: 'user_2',
      createdAt: '2026-10-21T10:11:00Z',
    },
    createdAt: userCreatedAt,
  },
]

export const demoMessages: ChatMessage[] = [
  {
    _id: 'msg_98',
    conversationId: 'conv_1',
    senderId: 'user_1',
    type: 'text',
    content: 'I have a few ideas for the new chat layout.',
    createdAt: '2026-10-21T09:18:00Z',
    updatedAt: '2026-10-21T09:18:00Z',
  },
  {
    _id: 'msg_99',
    conversationId: 'conv_1',
    senderId: 'user_2',
    type: 'text',
    content: 'Hey, are we still on for the design review?',
    createdAt: '2026-10-21T09:24:00Z',
    updatedAt: '2026-10-21T09:24:00Z',
  },
  {
    _id: 'msg_100',
    conversationId: 'conv_2',
    senderId: 'user_2',
    type: 'text',
    content: 'Who knows Gin?',
    createdAt: '2026-10-21T10:11:00Z',
    updatedAt: '2026-10-21T10:11:00Z',
  },
]
