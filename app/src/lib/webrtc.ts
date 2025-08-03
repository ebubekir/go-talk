// src/app/webrtc.ts
import { EventPayload } from '@/lib/room-events'

export class WebRTCManager {
  private peers: { [key: string]: RTCPeerConnection } = {}
  private localStream: MediaStream
  private signalingSend: (data: any) => void
  private onTrack: (id: string, stream: MediaStream) => void
  private onParticipantLeft: (id: string) => void

  constructor(
    localStream: MediaStream,
    signalingSend: (data: any) => void,
    onTrack: (id: string, stream: MediaStream) => void,
    onParticipantLeft: (id: string) => void,
  ) {
    this.peers = {}
    this.localStream = localStream
    this.signalingSend = signalingSend
    this.onTrack = onTrack
    this.onParticipantLeft = onParticipantLeft
  }

  async handleSignal({ type, from, payload }: EventPayload) {
    try {
      if (type === 'offer') {
        await this._createPeer(from, false)
        const pc = this.peers[from]

        // Check if we can set remote description
        if (pc.signalingState === 'stable' || pc.signalingState === 'have-local-offer') {
          await pc.setRemoteDescription(payload.data)
          const answer = await pc.createAnswer()
          await pc.setLocalDescription(answer)
          this.signalingSend({ type: 'answer', to: from, data: pc.localDescription })
        }
      } else if (type === 'answer') {
        const pc = this.peers[from]
        if (pc && pc.signalingState === 'have-local-offer') {
          await pc.setRemoteDescription(payload.data)
        }
      } else if (type === 'ice') {
        const pc = this.peers[from]
        if (pc && pc.remoteDescription) {
          try {
            await pc.addIceCandidate(payload.data)
          } catch (error) {
            console.warn('Failed to add ICE candidate:', error)
          }
        }
      } else if (type === 'leave') {
        if (this.peers[from]) {
          this.peers[from].close()
          delete this.peers[from]
          this.onParticipantLeft(from)
        }
      }
    } catch (error) {
      console.error('Error handling signal:', error, { type, from })
    }
  }

  async addParticipant(id: string) {
    // Prevent duplicate connections
    if (this.peers[id]) {
      return
    }

    try {
      await this._createPeer(id, true)
      const pc = this.peers[id]
      if (pc) {
        // @ts-ignore
        const offer = await pc.createOffer()
        // @ts-ignore
        await pc.setLocalDescription(offer)
        // @ts-ignore
        this.signalingSend({ type: 'offer', to: id, data: pc.localDescription })
      }
    } catch (error) {
      console.error('Error adding participant:', error, id)
      // Clean up on error
      this._cleanupPeer(id)
    }
  }

  private async _createPeer(id: string, _isInitiator: boolean) {
    if (this.peers[id]) {
      console.log('Peer connection already exists for:', id)
      return
    }

    const pc = new RTCPeerConnection({
      iceServers: [{ urls: 'stun:stun.l.google.com:19302' }],
      iceCandidatePoolSize: 10,
    })

    // Add local stream tracks
    this.localStream.getTracks().forEach((track) => {
      pc.addTrack(track, this.localStream)
    })

    pc.onicecandidate = (e) => {
      if (e.candidate) {
        this.signalingSend({ type: 'ice', to: id, data: e.candidate })
      }
    }

    pc.ontrack = (e) => {
      console.log('Received remote track from:', id)
      if (e.streams && e.streams[0]) {
        this.onTrack(id, e.streams[0])
      }
    }

    pc.onconnectionstatechange = () => {
      console.log(`Connection state for ${id}:`, pc.connectionState)
      if (pc.connectionState === 'disconnected' || pc.connectionState === 'failed') {
        console.log('Cleaning up failed connection for:', id)
        this._cleanupPeer(id)
      }
    }

    pc.onsignalingstatechange = () => {
      console.log(`Signaling state for ${id}:`, pc.signalingState)
    }

    this.peers[id] = pc
  }

  private _cleanupPeer(id: string) {
    if (this.peers[id]) {
      this.peers[id].close()
      delete this.peers[id]
      this.onParticipantLeft(id)
    }
  }

  removeAll() {
    Object.keys(this.peers).forEach((id) => this._cleanupPeer(id))
  }

  updateLocalStream(newStream: MediaStream) {
    this.localStream = newStream
    Object.values(this.peers).forEach((pc) => {
      pc.getSenders().forEach((sender) => {
        if (
          sender.track &&
          newStream.getTracks().some((t) => t.kind === sender.track?.kind)
        ) {
          const newTrack = newStream
            .getTracks()
            .find((t) => t.kind === sender.track?.kind)
          if (newTrack) {
            sender.replaceTrack(newTrack)
          }
        }
      })
    })
  }
}
