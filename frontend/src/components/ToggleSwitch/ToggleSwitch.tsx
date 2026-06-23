"use client";
import styles from "./ToggleSwitch.module.css";
interface Props {
    className?: string;
    checked: boolean;
    onChange: () => void;
}
const ToggleSwitch = ({ checked, onChange }: Props) => {
    return (
        <>
            <label className={styles.toggle_switch}>
                <input type="checkbox" checked={checked} onChange={onChange}></input>
                <span className={styles.move_switch}></span>
            </label>
        </>
    );
};
export default ToggleSwitch;
