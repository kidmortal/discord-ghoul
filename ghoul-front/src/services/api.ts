import axios from "axios";

type apiResponse = {
  response: string;
};

export type discordBot = {
  id: string;
  email: string;
  username: string;
  avatar: string;
  locale: string;
  discriminator: string;
  token: string;
  verified: boolean;
  mfa_enabled: boolean;
  bot: boolean;
  public_flags: number;
  premium_type: number;
  system: boolean;
  flags: number;
};

async function FetchApiGet(path: string): Promise<apiResponse> {
  let response = await fetch(`${process.env.REACT_APP_API_URL}/${path}`);
  let result: apiResponse = await response.json();
  return result;
}
async function FetchApiPost<T>(path: string, body: object) {
  let response = await axios.post<T>(
    `${process.env.REACT_APP_API_URL}/${path}`,
    body
  );
  return response.data;
}

export function BotLogin(secret: string) {
  return FetchApiPost<discordBot>("bot", {
    secret,
  });
}
