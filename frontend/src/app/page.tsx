"use client";

import { ChangeEvent, useState } from "react";
import styles from "./page.module.css";

export default function Home() {
    const [inputUrl, setInputUrl] = useState<string>("");
    const [downloadUrl, setDownloadUrl] = useState<string>("");

    async function Download() {
        try {
            const response = await fetch("http://localhost:8080/download", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    url: inputUrl,
                }),
            });

            const result = await response.json();
            setDownloadUrl(result.download_url);
        } catch (error) {
            const e = error as Error;
            console.error(e.message);
        }
    }

    function getVideo() {
        window.location.href = `http://localhost:8080${downloadUrl}`;
    }

    const onChangeUrl = (e: ChangeEvent<HTMLInputElement>) => {
        const input = e.target.value;

        setInputUrl(input);
    };

    return (
        <div className={styles.layout}>
            <div className={styles.container}>
                <div className={styles.url_container}>
                    <input onChange={onChangeUrl} type="text" placeholder="Paste Your URL..."></input>
                </div>

                <div className={styles.info_container}>
                    <div className={styles.video_container}>
                        <img src={"favicon.ico"} alt="Video Preview" />
                        <h1>Дыо против Жотары. Кабачковое противостояние</h1>
                    </div>
                    <div className={styles.control_container}>
                        <button className={styles.button_control} onClick={Download}>
                            Video
                        </button>
                        <button className={styles.button_control}>Audio</button>
                        <button className={styles.button_download} onClick={getVideo}>
                            Download
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
}
