import { useState } from "react";
import { Button, Popover, PopoverHeader, PopoverBody } from "reactstrap";
import { Discord } from "../../models/Discord";
import styles from "./Member.module.scss";

type MemberProps = {
  member: Discord.Guilds.Member;
};

export function Member({ member }: MemberProps) {
  const [popoverOpen, setPopoverOpen] = useState(false);
  const { user, nick } = member;
  const { id, avatar } = user;
  return (
    <div className={styles.container}>
      {id && avatar && (
        <div>
          <img
            id={`p${id}`}
            alt={nick}
            src={`https://cdn.discordapp.com/avatars/${id}/${avatar}.png?size=64`}
            onClick={() => setPopoverOpen(!popoverOpen)}
          />
          <Popover
            placement="bottom"
            isOpen={popoverOpen}
            target={`p${id}`}
            toggle={() => setPopoverOpen(!popoverOpen)}
            className={styles.popoverBox}
          >
            <MemberPopoverInfo member={member} />
          </Popover>
        </div>
      )}
    </div>
  );
}

type MemberPopoverInfoProps = {
  member: Discord.Guilds.Member;
};

function MemberPopoverInfo({ member }: MemberPopoverInfoProps) {
  const { user } = member;
  const { username } = user;
  return (
    <>
      <div className={styles.popoverInfo}>
        <span>nome: {username}</span>
      </div>
    </>
  );
}
