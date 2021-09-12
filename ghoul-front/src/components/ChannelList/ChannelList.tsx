import ScrollContainer from "react-indiana-drag-scroll";
import { Discord } from "../../models/Discord";
import { Channel } from "../Channel/Channel";
import { Guild } from "../Guild/Guild";
import styles from "./ChannelList.module.scss";

type ChannelListProps = {
  channels: Discord.Guilds.Channel[];
  selectedChannel: Discord.Guilds.Channel;
  setChannel: (channel: Discord.Guilds.Channel) => void;
};

export function ChannelList({ channels, setChannel }: ChannelListProps) {
  return (
    <div className={styles.container}>
      <div className={styles.guildList}>
        <ScrollContainer className={styles.scroll}>
          {channels?.map((c) => (
            <div key={c.id} onClick={() => setChannel(c)}>
              <Channel channel={c} />
            </div>
          ))}
        </ScrollContainer>
      </div>
    </div>
  );
}
