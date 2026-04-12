import styles from "./page.module.css";

export default function Home() {
    return (
        <div className={styles.layout}>
            <div className={styles.container}>
                <div className={styles.url_container}>
                    <input type="text" placeholder="Paste Your URL..."></input>
                </div>

                <div className={styles.info_container}>
                    <div className={styles.video_container}>
                        <img src={"favicon.ico"} alt="Video Preview" />
                        <h1>Дыо против Жотары. Кабачковое противостояние</h1>
                    </div>
                    <div className={styles.control_container}>
                        <button className={styles.button_control}>Video</button>
                        <button className={styles.button_control}>Audio</button>
                        <button className={styles.button_download}>Download</button>
                    </div>
                </div>
            </div>
        </div>
    );
}
