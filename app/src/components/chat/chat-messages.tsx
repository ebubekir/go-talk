import { ChatMessage } from '@/api/types'
import { UserAvatar } from '@/components/user-avatar'

export function ChatMessages({
  messages,
  messagesEndRef,
}: {
  messages?: ChatMessage[]
  messagesEndRef?: React.RefObject<HTMLDivElement>
}) {
  return (
    <div className="flex-1 overflow-y-auto px-6 py-4 space-y-4 bg-muted">
      {messages?.map((message, id) => (
        <div
          key={id}
          className={`flex items-end gap-2 ${message.isCurrentUser ? 'justify-end' : 'justify-start'}`}
        >
          {!message.isCurrentUser && (
            <UserAvatar name={message.user.name} email={message.user.email} />
          )}
          <div
            className={`rounded-xl px-4 py-2 max-w-xs text-sm shadow-md ${
              message.isCurrentUser
                ? 'bg-primary text-primary-foreground rounded-br-none'
                : 'bg-white text-gray-900 rounded-bl-none'
            }`}
          >
            {message.text}
          </div>
          {message.isCurrentUser && (
            <UserAvatar name={message.user.name} email={message.user.email} />
          )}
        </div>
      ))}
    </div>
  )
}
