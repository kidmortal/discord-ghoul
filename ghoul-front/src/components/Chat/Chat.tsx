import { Discord } from "../../models/Discord";
import { ChatHeader } from "../ChatHeader/ChatHeader";
import { ChatRow } from "../ChatRow/ChatRow";
import { ChatWriteInput } from "../ChatWriteInput/ChatWriteInput";
import styles from "./Chat.module.scss";

type ChatProps = {
  socket: WebSocket;
  selectedGuild: Discord.Guilds.Guild;
  selectedChannel: Discord.Guilds.Channel;
  channelMessages: Discord.Guilds.Message.Message[];
  connectChannel: () => void;
};

export function Chat({
  socket,
  selectedGuild,
  selectedChannel,
  channelMessages,
  connectChannel,
}: ChatProps) {
  function SendMessage(message: string) {
    socket.send(message);
  }

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <ChatHeader
          selectedChannel={selectedChannel}
          selectedGuild={selectedGuild}
          connectChannel={connectChannel}
        />
      </div>
      <div className={styles.inner}>
        {channelMessages.map((message) => (
          <ChatRow message={message} />
        ))}
      </div>
      <div className={styles.bottom}>
        <ChatWriteInput SendMessage={SendMessage} />
      </div>
    </div>
  );
}
