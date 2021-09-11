import { Discord } from "../../models/Discord";
import styles from "./GuildList.module.scss";

type GuildListProps = {
  guilds: Discord.Guilds.Guild[];
};

export function GuildList({ guilds }: GuildListProps) {
  return (
    <div className={styles.container}>
      <h2>Servers </h2>
      <div className={styles.guildList}>
        {guilds.map((g) => (
          <img
            alt={g.name}
            src={`https://cdn.discordapp.com/icons/${g.id}/${g.icon}.webp?size=64`}
          />
        ))}
      </div>
    </div>
  );
}
