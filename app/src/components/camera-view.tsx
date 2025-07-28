'use client';

import {useEffect, useRef, useState} from "react";

export function CameraView() {
    const videoRef = useRef<HTMLVideoElement>(null);
    const [error, setError] = useState<string | null>(null);
    const [hasPermission, setHasPermission] = useState(false);

    useEffect(() => {
        let mediaStream: MediaStream;

        async function getMediaStream() {
            try {
                mediaStream = await navigator.mediaDevices.getUserMedia({
                    video: true,
                    audio: true, // Mikrofon eklendi
                });

                if (videoRef.current) {
                    videoRef.current.srcObject = mediaStream;
                }
                setHasPermission(true)
            } catch (err) {
                setError("Kamera veya mikrofona erişim sağlanamadı.");
                console.error(err);
            }
        }

        getMediaStream();

        return () => {
            mediaStream?.getTracks().forEach((track) => track.stop());
        };
    }, []);

    return hasPermission ? <video
                        ref={videoRef}
                        autoPlay
                        playsInline
                        muted
                        className="w-full h-full object-cover rounded-md shadow-md transform scale-x-[-1]"
                    />: <div className="w-full h-full flex items-center justify-center border border-dashed border-gray-400 rounded-md text-gray-500 text-center p-4">
                        You need to allow camera and microphone access to use this feature.
                    </div>

}