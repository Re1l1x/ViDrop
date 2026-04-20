"use client";

import { ChangeEvent, useState } from "react";
import styles from "./page.module.css";

export default function Home() {
    const [inputUrl, setInputUrl] = useState<string>("");
    const [downloadUrl, setDownloadUrl] = useState<string>("");

    const [isUrlEntered, setIsUrlEntered] = useState<boolean>(false);
    const targetLength = 43;

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
        if (input.length >= targetLength) {
            setIsUrlEntered(true);
        }
    };

    return (
        <div className={styles.layout}>
            <div className={styles.container}>
                <div className={styles.input_line_container}>
                    <div className={styles.title}>ViDrop</div>
                    <input onChange={onChangeUrl} type="text" placeholder="Paste Your URL..."></input>
                </div>

                <div className={`${styles.info_container} ${isUrlEntered ? styles.expanded : styles.collapsed}`}>
                    {/* <div className={`${styles.form_box} ${isSignIn ? "" : styles.active}`}></div> */}
                    <div className={styles.video_container}>
                        {/* <div className={styles.section_name}>Video</div> */}
                        <img src={"favicon.ico"} alt="Video Preview" />
                        <div className={styles.video_name}>Дыо против Жотары. Кабачковое противостояние</div>
                    </div>
                    <div className={styles.control_container}>
                        <div className={styles.control_row}>
                            <div className={styles.section_name}>Settings:</div>
                        </div>
                        <div className={styles.control_row}>
                            <button className={styles.button_control} onClick={Download}>
                                Video
                            </button>
                            <button className={styles.button_control}>Audio</button>
                        </div>
                        <div className={styles.control_row}>
                            <div className={styles.section_name}>Quality:</div>
                            <select className={styles.quality_selector} name="Quality">
                                <option value="" disabled selected>
                                    Select Quality
                                </option>
                                <option value="">144p</option>
                                <option value="">240p</option>
                                <option value="">360p</option>
                                <option value="">480p</option>
                                <option value="">720p</option>
                                <option value="">1080p</option>
                            </select>
                        </div>
                        <div className={styles.control_row}>
                            <button className={styles.button_download} onClick={getVideo}>
                                Download
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}
