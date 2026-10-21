users 
{
  "_id": "user_1",

  "username": "kittiwat",

  "displayName": "Kittiwat",

  "avatar": "avatar.png",

  "status": "online",

  "createdAt": "2026-10-06T20:00:00Z",

  "updatedAt": "2026-10-06T20:00:00Z"
}

_id
username
displayName
avatar
status
createdAt
updatedAt


conversations
direct
{
  "_id": "conv_1",

  "type": "direct",

  "participants": [
    "user_1",
    "user_2"
  ],

  "lastMessage": {
    "messageId": "msg_99",
    "content": "hello",
    "senderId": "user_1",
    "createdAt": "2026-10-06T20:10:00Z"
  },

  "createdAt": "2026-10-06T20:00:00Z"
}

group 
{
  "_id": "conv_2",

  "type": "group",

  "name": "Go Developers",

  "participants": [
    "user_1",
    "user_2",
    "user_3"
  ],

  "admins": [
    "user_1"
  ],

  "lastMessage": {
    "messageId": "msg_100",
    "content": "Who knows Gin?",
    "senderId": "user_2",
    "createdAt": "2026-10-06T20:11:00Z"
  },

  "createdAt": "2026-10-06T20:00:00Z"
}

_id
type
name
participants
admins
lastMessage
createdAt


messages
{
  "_id": "msg_100",

  "conversationId": "conv_2",

  "senderId": "user_2",

  "type": "text",

  "content": "Who knows Gin?",

  "createdAt": "2026-10-06T20:11:00Z",

  "updatedAt": "2026-10-06T20:11:00Z"
}

_id
conversationId
senderId
type
content
createdAt
updatedAt


**relational
users
│
├── user_1
├── user_2
└── user_3

         │

         ▼

conversations

conv_1
(DM)

participants
├── user_1
└── user_2

conv_2
(Group)

participants
├── user_1
├── user_2
└── user_3

         │

         ▼

messages

msg_1
msg_2
msg_3
msg_4
...


Query ที่ใช้บ่อย
โหลดห้องทั้งหมดของ User
db.conversations.find({
participants: "user_1"
})

โหลดข้อความในห้อง
db.messages.find({
conversationId: "conv_2"
})

ส่งข้อความใหม่
Plain Text
insert message
↓
update conversation.lastMessage
↓
broadcast websocket

Index ที่ควรสร้างตั้งแต่วันแรก
conversations

{
participants: 1
}
messages

{
conversationId: 1,
createdAt: -1
}