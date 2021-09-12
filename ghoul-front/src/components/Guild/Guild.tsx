import { Discord } from "../../models/Discord";
import styles from "./Guild.module.scss";

type GuildProps = {
  guild: Discord.Guilds.Guild;
  select: Function;
};

export function Guild({ guild, select }: GuildProps) {
  return (
    <div
      className={styles.container}
      onClick={() => {
        select();
      }}
    >
      <img
        alt={guild.name}
        src={`https://cdn.discordapp.com/icons/${guild.id}/${guild.icon}.webp?size=64`}
      />
    </div>
  );
}
