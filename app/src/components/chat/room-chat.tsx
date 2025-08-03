import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { useEffect, useRef, useState } from 'react'
import { useSocket } from '@/context/socket-context'
import { EventPayload, RoomEventName, webSocketEventHandlers } from '@/lib/room-events'
import { ChatMessages } from '@/components/chat/chat-messages'
import { Chat } from '@/api/types'
import { useRoom } from '@/context/room-context'
import { useAuth } from '@/context/auth-context'
import { api } from '@/api/api'

export function RoomChat() {
  const { chat, roomId } = useRoom()

  const [input, setInput] = useState('')
  const messagesEndRef = useRef<HTMLDivElement>(null)

  // Scroll to bottom on a new message
  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' })
  }

  // Simulate sending a message
  const handleSend = async () => {
    if (!input.trim()) return
    if (!roomId) return

    await api.chat.sendMessage(roomId, input)
    setInput('')
    setTimeout(scrollToBottom, 100)
  }

  return (
    <div className="flex flex-col h-full w-full max-w-2xl mx-auto">
      <ChatMessages messages={chat?.history} />
      <div ref={messagesEndRef} />
      <form
        className="flex items-center gap-2 px-6 py-4 border-t bg-background"
        onSubmit={(e) => {
          e.preventDefault()
          handleSend()
        }}
      >
        <Input
          className="flex-1"
          placeholder="Type your message..."
          value={input}
          onChange={(e) => setInput(e.target.value)}
          autoComplete="off"
        />
        <Button type="submit" className="shrink-0">
          Send
        </Button>
      </form>
    </div>
  )
}

export default RoomChat
