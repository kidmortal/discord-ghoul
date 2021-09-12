import axios from "axios";
import { Discord } from "../models/Discord";

async function FetchApiGet<T>(path: string) {
  let response = await axios.get<T>(`${process.env.REACT_APP_API_URL}/${path}`);
  return response.data;
}
async function FetchApiPost<T>(path: string, body: object) {
  let response = await axios.post<T>(
    `${process.env.REACT_APP_API_URL}/${path}`,
    body
  );
  return response.data;
}

export function BotLogin(secret: string) {
  return FetchApiPost<Discord.User>("bot", {
    secret,
  });
}

export function GetGuilds() {
  return FetchApiGet<Discord.Guilds.Guild[]>("guilds");
}

export function GetBot() {
  return FetchApiGet<Discord.User>("bot");
}
<<<<<<< HEAD

export function GetChannelMessages(channelId: string) {
  return FetchApiGet<Discord.Guilds.Message.Message[]>(`messages/${channelId}`);
}
=======
>>>>>>> 1ab8ec363a16fd463e56112578e07570abbdf02b
