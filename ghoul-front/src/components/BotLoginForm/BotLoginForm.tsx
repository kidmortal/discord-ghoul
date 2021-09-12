import { useEffect, useState } from "react";
import styles from "./BotLoginForm.module.scss";

type BotLoginFormProps = {
  handleLogin: Function;
};

export function BotLoginForm({ handleLogin }: BotLoginFormProps) {
  const [secret, setSecret] = useState("");
  function handleClick() {
    localStorage.setItem("@discord-ghoul/secret", secret);
    handleLogin(secret);
  }
  useEffect(() => {
    let cache = localStorage.getItem("@discord-ghoul/secret");
    if (cache) setSecret(cache);
  }, []);
  return (
    <div className={styles.container}>
      <div className={styles.inputArea}>
        <input
          placeholder="SECRET do BOT"
          value={secret}
          type="password"
          onChange={(event) => setSecret(event.target.value)}
        />
        <button onClick={handleClick}>🍰</button>
      </div>
    </div>
  );
}
