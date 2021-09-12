import { useEffect, useState } from "react";
import { BotLoginForm } from "../components/BotLoginForm/BotLoginForm";
import { BotStatus } from "../components/BotStatus/BotStatus";
import { ChannelList } from "../components/ChannelList/ChannelList";
import { Chat } from "../components/Chat/Chat";
import { GuildList } from "../components/GuildList/GuildList";
import { Discord } from "../models/Discord";
import {
  BotLogin,
  GetBot,
  GetChannelMessages,
  GetGuilds,
} from "../services/api";
import { ConnectWS, SetupMessageSocket } from "../services/websocket";
import styles from "./Home.module.scss";

export function HomePage() {
  const [socket, setSocket] = useState({} as WebSocket);
  const [bot, setBot] = useState({} as Discord.User);
  const [guilds, setGuilds] = useState([] as Discord.Guilds.Guild[]);
  const [selectedGuild, setSelectedGuild] = useState(
    {} as Discord.Guilds.Guild
  );
  const [selectedChannel, setSelectedChannel] = useState(
    {} as Discord.Guilds.Channel
  );
  const [messages, setMessages] = useState(
    [] as Discord.Guilds.Message.Message[]
  );
  const reverseMessages = _reverseArray(messages);

  function _reverseArray(arr: any[]) {
    let ret = new Array();
    for (let i = arr.length - 1; i >= 0; i--) {
      ret.push(arr[i]);
    }
    return ret;
  }

  function handleLogin(secret: string) {
    if (secret)
      BotLogin(secret).then((bot) => {
        if (bot?.id) {
          setBot(bot);
          _GetGuilds();
        }
      });
  }

  function _GetGuilds() {
    GetGuilds().then((r) => {
      if (r?.[0]?.id) {
        setGuilds(r);
        setSelectedGuild(r[0]);
      }
    });
  }

  function channelWS(id: string) {
    ConnectWS(id).then((socket) => {
      if (socket) {
        SetupMessageSocket({ socket, messages, setMessages });
        setSocket(socket);
      }
    });
  }

  function ConnectChannel() {
    if (selectedChannel?.id) {
      channelWS(selectedChannel?.id);
    }
  }

  useEffect(() => {
    GetBot().then((bot) => {
      if (bot.id) {
        setBot(bot);
        _GetGuilds();
      }
    });
  }, []);

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <div className={styles.loginStatus}>
          <BotStatus bot={bot} />
          <BotLoginForm handleLogin={handleLogin} />
        </div>
      </div>

      <div className={styles.centerContainer}>
        <GuildList
          guilds={guilds}
          selectedGuild={selectedGuild}
          setGuild={(guild) => setSelectedGuild(guild)}
        />
        <ChannelList
          channels={selectedGuild.channels}
          selectedChannel={selectedChannel}
          setChannel={(channel) => setSelectedChannel(channel)}
        />
        <Chat
          selectedGuild={selectedGuild}
          selectedChannel={selectedChannel}
          channelMessages={reverseMessages}
          socket={socket}
          connectChannel={ConnectChannel}
        />
      </div>
    </div>
  );
}
