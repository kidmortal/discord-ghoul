import { useEffect, useState } from "react";
import { GuildList } from "../components/guildList/GuildList";
import { Discord } from "../models/Discord";
import { BotLogin, GetBot, GetGuilds } from "../services/api";
import styles from "./Home.module.scss";

export function HomePage() {
  const [secret, setSecret] = useState("");
  const [bot, setBot] = useState({} as Discord.User);
  const [guilds, setGuilds] = useState([] as Discord.Guilds.Guild[]);

  function handleLogin() {
    if (secret)
      BotLogin(secret).then((r) => {
        console.log(r);
        if (r) setBot(r);
      });
  }

  function handleGetGuilds() {
    GetGuilds().then((r) => {
      console.log(r);
      setGuilds(r);
    });
  }

  useEffect(() => {
    GetBot().then((r) => {
      if (r.id) {
        setBot(r);
      }
    });
  }, []);

  return (
    <div className={styles.container}>
      <div className={styles.status}>
        {bot?.avatar && (
          <img
            alt={bot.username}
            src={`https://cdn.discordapp.com/avatars/${bot.id}/${bot.avatar}.png?size=256`}
          />
        )}
        <p>
          Bot name: {bot?.username && `${bot.username}#${bot.discriminator}`}{" "}
        </p>
      </div>
      <div>
        <GuildList guilds={guilds} />
      </div>
      <div className={styles.inputArea}>
        <h3>Digita o secret ae</h3>
        <input
          placeholder="SECRET do BOT"
          value={secret}
          onChange={(event) => setSecret(event.target.value)}
        />

        <button onClick={() => handleLogin()}>Login :D</button>
        <button onClick={() => handleGetGuilds()}>Ver servs</button>
      </div>
    </div>
  );
}
