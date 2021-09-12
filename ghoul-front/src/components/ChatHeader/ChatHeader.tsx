import { CgArrowsExchangeAltV } from "react-icons/cg";
import { Discord } from "../../models/Discord";
import styles from "./ChatHeader.module.scss";

type ChatHeaderProps = {
  selectedGuild: Discord.Guilds.Guild;
  selectedChannel: Discord.Guilds.Channel;
  connectChannel: () => void;
};

export function ChatHeader({
  selectedGuild,
  selectedChannel,
  connectChannel,
}: ChatHeaderProps) {
  return (
    <div className={styles.container}>
      <div className={styles.title}>
        <span>
          {selectedGuild?.name} - {selectedChannel?.name}
        </span>
      </div>
      <div className={styles.connection}>
        <CgArrowsExchangeAltV onClick={connectChannel} />
      </div>
    </div>
  );
}
