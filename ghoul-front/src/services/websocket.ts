import { Discord } from "../models/Discord";

type SocketMessage = {
  call: string;
  payload: [Discord.Guilds.Message.Message] | Discord.Guilds.Message.Message;
};

export async function ConnectWS(id: string) {
  let url = process.env.REACT_APP_API_URL?.replace("http", "ws").replace(
    "https",
    "ws"
  );
  let socket = await new WebSocket(`${url}/ws/${id}`);

  return socket;
}

export function ReadSocketMessage(ev: MessageEvent) {
  const { data } = ev;
  const message: SocketMessage = JSON.parse(data);
  return message;
}

type SetupMessageSocketProps = {
  socket: WebSocket;
  messages: Discord.Guilds.Message.Message[];
  setMessages: (messages: Discord.Guilds.Message.Message[]) => void;
};

export function SetupMessageSocket({
  socket,
  messages,
  setMessages,
}: SetupMessageSocketProps) {
  let socketMessages: Discord.Guilds.Message.Message[] = [];
  socket.onopen = () => {
    console.log("connected");
  };
  socket.onmessage = (ev) => {
    let message = ReadSocketMessage(ev);
    const { call, payload } = message;
    switch (call) {
      case "NEW_MESSAGES":
        Array.isArray(payload)
          ? (socketMessages = [...socketMessages, ...payload])
          : (socketMessages = [payload, ...socketMessages]);

        setMessages(socketMessages);
        break;

      default:
        break;
    }
  };
}
