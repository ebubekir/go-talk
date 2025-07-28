let socket : WebSocket | null = null;


export const connectToSocket = async (roomId: string, token: string): Promise<WebSocket> => {
    if (!roomId || !token) throw new Error("roomId and token required");

    const ws = new WebSocket(`ws://localhost:8080/ws/room/${roomId}?token=${token}`);

    ws.onopen = () => {
        console.log("WebSocket connected");
        Promise.resolve(ws);
    };
    ws.onerror = (err) => {
        console.error("WebSocket error", err);
        Promise.reject(err)
    };

    ws.onclose = () => {
        console.log("WebSocket closed");
    };



    socket = ws;
    return ws;
};

export const getSocket = () => socket;
export const closeSocket = () => {
    socket?.close();
    socket = null;
};