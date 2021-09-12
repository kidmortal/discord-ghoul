import ScrollContainer from "react-indiana-drag-scroll";
import { Discord } from "../../models/Discord";
import { Member } from "../Member/Member";
import styles from "./MemberList.module.scss";

type MemberListProps = {
  members: Discord.Guilds.Member[];
};

export function MemberList({ members }: MemberListProps) {
  return (
    <div className={styles.container}>
      <div className={styles.title}>
        <h2>Membros </h2>
      </div>
      <div className={styles.guildList}>
        <ScrollContainer className={styles.scroll}>
          {members?.map((m) => (
            <Member member={m} key={m.user.id} />
          ))}
        </ScrollContainer>
      </div>
    </div>
  );
}
