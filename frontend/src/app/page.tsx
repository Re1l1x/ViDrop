"use client";

import { ChangeEvent, useState, useRef } from "react";
import styles from "./page.module.css";

export default function Home() {
    const [inputUrl, setInputUrl] = useState<string>("");
    const [downloadUrl, setDownloadUrl] = useState<string>("");

    const [isUrlEntered, setIsUrlEntered] = useState<boolean>(false);
    const targetLength = 43;

    const containerRef = useRef<HTMLDivElement>(null);

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
            expandContainer();
        }
    };
    function expandContainer() {
        const el = containerRef.current;
        if (!el) return;

        const start = el.scrollHeight;
        const end = window.innerHeight;

        document.body.style.overflow = "hidden";

        el.style.height = start + "px";
        // console.log("start: " + start + " end: " + end);

        // почему без этого не работает?
        requestAnimationFrame(() => {
            el.style.transition = "height 0.3s ease";
            el.style.height = end + "px";
        });

        const onEnd = (e: TransitionEvent) => {
            if (e.propertyName !== "height") return;

            document.body.style.overflow = "";
            el.removeEventListener("transitionend", onEnd);
        };

        el.addEventListener("transitionend", onEnd);
    }

    return (
        <div className={styles.layout}>
            <div ref={containerRef} className={styles.container}>
                {/* <div className={`${styles.container} ${isUrlEntered ? styles.toUp : ""}`}> */}
                <div className={styles.header_container}>
                    <div className={`${styles.title} ${isUrlEntered ? styles.urlSubmitted : ""}`}>ViDrop</div>
                    <input className={styles.input_line} onChange={onChangeUrl} type="text" placeholder="Paste Your URL..."></input>
                </div>
                <div className={styles.divider}></div>
                <div className={`${styles.info_container} ${isUrlEntered ? styles.urlSubmitted : ""}`}>
                    {/* <div className={`${styles.form_box} ${isSignIn ? "" : styles.active}`}></div> */}
                    <div className={styles.video_container}>
                        <div className={styles.section_name}>Video</div>
                        <img src={"preview final.jpg"} alt="Video Preview" />
                        <div className={styles.video_name}>Дыо против Жотары. Кабачковое противостояние</div>
                    </div>
                    <div className={styles.control_container}>
                        <div className={styles.control_row}>
                            <div className={styles.section_name}>Settings:</div>
                        </div>
                        <div className={styles.control_row}>
                            <button className={styles.control_button} onClick={Download}>
                                Video
                            </button>
                            <button className={styles.control_button}>Audio</button>
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
                            <button className={styles.download_button} onClick={getVideo}>
                                Download
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}
