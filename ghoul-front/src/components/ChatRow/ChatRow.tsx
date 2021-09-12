import { Discord } from "../../models/Discord";
import styles from "./ChatRow.module.scss";

type ChatRowProps = {
  message: Discord.Guilds.Message.Message;
};

export function ChatRow({ message }: ChatRowProps) {
  const { author } = message;
  const { id, avatar } = author;
  const avatarUrl = `https://cdn.discordapp.com/avatars/${id}/${avatar}.png?size=32`;
  return (
    <div className={styles.container}>
      <img alt={message.author.username} src={avatarUrl} />
      <span>{`${message.author.username}: ${message.content}`}</span>
    </div>
  );
}
