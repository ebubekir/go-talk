'use client'
// src/app/VideoConference.tsx
import React, { useEffect, useRef, useState } from 'react'
import { useAuth } from '@/context/auth-context'
import { useSocket } from '@/context/room-socket-context'

const SIGNALING_URL = 'ws://localhost:3001'

export default function VideoConference() {
  const {
    localStream,
    toggleMic,
    toggleCam,
    leaveCall,
    camEnabled,
    micEnabled,
    participantStreams,
  } = useSocket()

  const { user } = useAuth()
  return (
    <div>
      <h2>Video Conference</h2>
      <div>
        <button onClick={toggleCam}>
          {camEnabled ? 'Turn Camera Off' : 'Turn Camera On'}
        </button>
        <button onClick={toggleMic}>{micEnabled ? 'Mute' : 'Unmute'}</button>
        <button onClick={leaveCall}>Leave</button>
      </div>
      <div>
        <h3>Participants</h3>
        <div style={{ display: 'flex', flexWrap: 'wrap' }}>
          {/* Local video */}
          <div style={{ margin: 8 }}>
            <div>{user?.email ? `Me (${user?.email})` : 'Me'}</div>
            {localStream && camEnabled ? (
              <video
                ref={(el) => {
                  if (el) el.srcObject = localStream
                }}
                autoPlay
                muted
                style={{ width: 200, height: 150, background: '#222' }}
              />
            ) : (
              <div
                style={{
                  width: 200,
                  height: 150,
                  background: '#444',
                  color: '#fff',
                  display: 'flex',
                  alignItems: 'center',
                  justifyContent: 'center',
                }}
              >
                {user?.id || 'Me'}
              </div>
            )}
          </div>
          {/* Remote videos */}
          {participantStreams?.map((p) => (
            <div key={p.id} style={{ margin: 8 }}>
              <div>{p.id}</div>
              {p.stream && p.stream.getVideoTracks().some((t) => t.enabled) ? (
                <video
                  ref={(el) => {
                    if (el) el.srcObject = p.stream
                  }}
                  autoPlay
                  style={{ width: 200, height: 150, background: '#222' }}
                />
              ) : (
                <div
                  style={{
                    width: 200,
                    height: 150,
                    background: '#444',
                    color: '#fff',
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center',
                  }}
                >
                  {p.id}
                </div>
              )}
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}
