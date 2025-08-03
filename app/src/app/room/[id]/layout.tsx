'use client'

import { RoomContextProvider } from '@/context/room-context'
import { use } from 'react'
import { RoomSocketProvider } from '@/context/room-socket-context'
import { MediaProvider } from '@/context/media-context'
import { RpcProvider } from '@/context/rpc-context'

export default function RoomLayout({
  children,
  params,
}: Readonly<{
  children: React.ReactNode
  params: Promise<{ id: string }>
}>) {
  const { id } = use(params)
  return (
    <RoomContextProvider>
      <RoomSocketProvider>
        <RpcProvider>{children}</RpcProvider>
      </RoomSocketProvider>
    </RoomContextProvider>
  )
}
