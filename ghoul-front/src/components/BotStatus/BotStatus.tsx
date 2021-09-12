import { Discord } from "../../models/Discord";
import styles from "./BotStatus.module.scss";

type BotStatusProps = {
  bot: Discord.User;
};

export function BotStatus({ bot }: BotStatusProps) {
  const avatarUrl =
    bot?.id &&
    `https://cdn.discordapp.com/avatars/${bot?.id}/${bot?.avatar}.png?size=64`;

  const peepoSad =
    "https://cdn.discordapp.com/emojis/405951684339302400.png?v=1&size=64";

  function BotImage() {
    return (
      <img alt={bot?.username || "Desconectado"} src={avatarUrl || peepoSad} />
    );
  }

  function BotName() {
    return (
      <span>{`Bot name: ${bot?.username || ""}#${
        bot.discriminator || ""
      }`}</span>
    );
  }

  function BotConnection() {
    return bot.id ? (
      <span>{`Status: Connected`}</span>
    ) : (
      <span>{`Status: Disconnected`}</span>
    );
  }

  return (
    <div className={styles.container}>
      <BotImage />
      <div className={styles.info}>
        <BotName />
        <BotConnection />
      </div>
    </div>
  );
}
