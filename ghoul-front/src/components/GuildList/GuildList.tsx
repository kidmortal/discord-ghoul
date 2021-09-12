import ScrollContainer from "react-indiana-drag-scroll";
import { Discord } from "../../models/Discord";
import { Guild } from "../Guild/Guild";
import styles from "./GuildList.module.scss";

type GuildListProps = {
  guilds: Discord.Guilds.Guild[];
  selectedGuild: Discord.Guilds.Guild;
  setGuild: (guild: Discord.Guilds.Guild) => void;
};

export function GuildList({ guilds, setGuild, selectedGuild }: GuildListProps) {
  return (
    <div className={styles.container}>
      <div className={styles.guildList}>
        <ScrollContainer className={styles.scroll}>
          {guilds?.map((g) => (
            <div
              key={g.id}
              className={g.id === selectedGuild.id ? styles.selected : ""}
            >
              <Guild
                guild={g}
                select={() => {
                  setGuild(g);
                }}
              />
            </div>
          ))}
        </ScrollContainer>
      </div>
    </div>
  );
}
