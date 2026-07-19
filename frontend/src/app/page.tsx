"use client";

import { ChangeEvent, useState, useRef } from "react";
import styles from "./page.module.css";
import ToggleSwitch from "@/components/ToggleSwitch/ToggleSwitch";
import Dropdown from "@/components/Dropdown/Dropdown";

export default function Home() {
    const [inputUrl, setInputUrl] = useState<string>("");

    const [isUrlEntered, setIsUrlEntered] = useState<boolean>(false);
    const targetLength = 43;

    const [isVideoEnabled, setIsVideoEnabled] = useState(true);
    const [isAudioEnabled, setIsAudioEnabled] = useState(true);

    const containerRef = useRef<HTMLDivElement>(null);
    const [videoInfo, setVideoInfo] = useState<{
        title: string;
        thumbnail: string;
        resolutions: string[];
        audio_bitrates: string[];
    } | null>(null);

    const [videoStatus, setVideoStatus] = useState<{
        fileId: string;
        status: string;
        progress: number;
        error: string;
    } | null>(null);

    const [selectedResolution, setSelectedResolution] = useState<string | undefined>(undefined);
    const [selectedBitrate, setSelectedBitrate] = useState<string | undefined>(undefined);
    const [selectedExtension, setSelectedExtension] = useState<string | undefined>(undefined);

    async function downloadVideo() {
        try {
            const response = await fetch("http://localhost:8080/download", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    URL: inputUrl,
                    resolution: selectedResolution,
                    audioBitrate: selectedBitrate,
                    format: selectedExtension,
                }),
            });
            const result = await response.json();
            ProgressCheck(result.jobId);
        } catch (error) {
            const e = error as Error;
            console.error(e.message);
        }
    }

    async function ProgressCheck(jobId: string) {
        const evtSource = new EventSource(`http://localhost:8080/download/${jobId}/events`);
        evtSource.onmessage = (event) => {
            setVideoStatus(event.data);
            if (event.data.status === "done") {
                getVideo(event.data.fileId);
                evtSource.close();
            }
        };
    }

    function getVideo(fileId: string) {
        window.location.href = `http://localhost:8080${fileId}`;
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
        console.log("Getting video info for URL:", input);
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
                if (!response.ok) {
                    throw new Error(`Response status: ${response.status}`);
                }
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
                            <img src={videoInfo?.thumbnail} alt={videoInfo?.title} />
                            <div className={styles.video_name}>{videoInfo?.title}</div>
                        </div>
                        <div className={styles.control_container}>
                            <div className={styles.control_row}>
                                <div className={styles.section_name}>Video</div>
                                <ToggleSwitch checked={isVideoEnabled} onChange={() => setIsVideoEnabled(!isVideoEnabled)} />
                            </div>
                            <div className={styles.control_row}>
                                <Dropdown
                                    options={videoInfo?.resolutions}
                                    setSelectedOption={setSelectedResolution}
                                    selectedOption={selectedResolution}
                                    postfix="p"
                                />
                            </div>
                            <div className={styles.control_row}>
                                <div className={styles.section_name}>Audio</div>
                                <ToggleSwitch checked={isAudioEnabled} onChange={() => setIsAudioEnabled(!isAudioEnabled)} />
                            </div>
                            <div className={styles.control_row}>
                                <Dropdown
                                    options={videoInfo?.audio_bitrates}
                                    setSelectedOption={setSelectedBitrate}
                                    selectedOption={selectedBitrate}
                                    postfix="kbps"
                                />
                            </div>
                            <div className={styles.download_extension_button}>
                                <button className={styles.download_button} onClick={downloadVideo}>
                                    Download {videoStatus?.progress ? `(${videoStatus.progress}%)` : ""}
                                </button>
                                <Dropdown
                                    options={["mp4", "mp3", "avi"]}
                                    className={styles.extension_button}
                                    setSelectedOption={setSelectedExtension}
                                    selectedOption={selectedExtension}
                                />
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}
