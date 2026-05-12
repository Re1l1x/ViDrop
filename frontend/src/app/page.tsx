"use client";

import { ChangeEvent, useState, useRef } from "react";
import styles from "./page.module.css";

export default function Home() {
    const [inputUrl, setInputUrl] = useState<string>("");
    const [downloadUrl, setDownloadUrl] = useState<string>("");

    const [isUrlEntered, setIsUrlEntered] = useState<boolean>(false);
    const targetLength = 43;

    const containerRef = useRef<HTMLDivElement>(null);
    const [videoInfo, setVideoInfo] = useState<{
        title: string;
        thumbnail_url: string;
        resolutions: number[];
        bitrates: number[];
    } | null>(null);

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
            getVideoInfo(input);
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
            el.style.height = "100vh";
            el.style.transition = "";
            el.removeEventListener("transitionend", onEnd);
        };

        el.addEventListener("transitionend", onEnd);
    }

    async function getVideoInfo(input: string) {
        fetch("http://localhost:8080/info", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                url: input,
            }),
        })
            .then((response) => {
                console.log(response);
                return response.json();
            })
            .then((data) => {
                console.log(data);
                setVideoInfo(data);
            })
            .catch((error) => {
                const e = error as Error;
                console.error(e.message);
            });
    }
    // try {
    //     const response = await fetch("http://localhost:8080/info", {
    //         method: "POST",
    //         headers: {
    //             "Content-Type": "application/json",
    //         },
    //         body: JSON.stringify({
    //             url: inputUrl,
    //         }),
    //     });

    //     const result = await response.json();

    // } catch (error) {
    //     const e = error as Error;
    //     console.error(e.message);
    // }

    return (
        <div className={styles.layout}>
            <div ref={containerRef} className={styles.container}>
                <div className={styles.header_container}>
                    <div className={`${styles.title} ${isUrlEntered ? styles.urlSubmitted : ""}`}>ViDrop</div>
                    <input className={styles.input_line} onChange={onChangeUrl} type="text" placeholder="Paste Your URL..."></input>
                </div>
                <div className={`${styles.main_page} ${isUrlEntered ? styles.urlSubmitted : ""}`}>
                    <div className={styles.info_container}>
                        <div className={styles.video_container}>
                            <img src={videoInfo?.thumbnail_url} alt={videoInfo?.title} />
                            <div className={styles.video_name}>{videoInfo?.title}</div>
                        </div>
                        <div className={styles.control_container}>
                            <div className={styles.control_row}>
                                <div className={styles.section_name}>Video</div>
                                <label className={styles.toggle_switch}>
                                    <input type="checkbox"></input>
                                    <span className={styles.move_switch}></span>
                                </label>
                            </div>
                            <div className={styles.control_row}>
                                <select className={styles.quality_selector} defaultValue="">
                                    {/* <option value="" disabled>
                                        Select Video Quality
                                    </option> */}

                                    {videoInfo?.bitrates.map((res) => (
                                        <option key={res} value={res}>
                                            {res}p
                                        </option>
                                    ))}
                                </select>
                                {/* <button className={styles.control_button} onClick={Download}>
                                    Video
                                </button>
                                <button className={styles.control_button}>Audio</button> */}
                            </div>
                            <div className={styles.control_row}>
                                <div className={styles.section_name}>Audio</div>
                                <label className={styles.toggle_switch}>
                                    <input type="checkbox"></input>
                                    <span className={styles.move_switch}></span>
                                </label>
                            </div>
                            <div className={styles.control_row}>
                                <select className={styles.quality_selector} defaultValue="">
                                    {/* <option value="" disabled>
                                        Select Audio Quality
                                    </option> */}

                                    {videoInfo?.resolutions.map((bps) => (
                                        <option key={bps} value={bps}>
                                            {bps}p
                                        </option>
                                    ))}
                                </select>
                            </div>
                            <div className={styles.download_extension_button}>
                                <button className={styles.download_button} onClick={getVideo}>
                                    Download
                                </button>
                                <select className={styles.extension_button} defaultValue="">
                                    <option> Prepfej </option>
                                </select>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}
