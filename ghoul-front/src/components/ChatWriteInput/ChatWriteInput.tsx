import { useState } from "react";
import { AiOutlineSend } from "react-icons/ai";
import styles from "./ChatWriteInput.module.scss";

type ChatWriteInputProps = {
  SendMessage: (message: string) => void;
};

export function ChatWriteInput({ SendMessage }: ChatWriteInputProps) {
  const [message, setMessage] = useState("");
  return (
    <div className={styles.container}>
      <div className={styles.inputArea}>
        <input
          placeholder={"Message"}
          value={message}
          onChange={(e) => setMessage(e.target.value)}
        />
      </div>
      <div className={styles.buttonArea}>
        <AiOutlineSend
          className={styles.sendButton}
          onClick={() => {
            SendMessage(message);
            setMessage("");
          }}
        />
      </div>
    </div>
  );
}
