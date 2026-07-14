"use client";
import styles from "./Dropdown.module.css";
import { useState, useEffect } from "react";
interface Props {
    className?: string;
    options: string[] | undefined;
}

const Dropdown = ({ options = [], className }: Props) => {
    const [isOpen, setIsOpen] = useState(false);
    const [selectedOption, setSelectedOption] = useState<string | undefined>(undefined);

    useEffect(() => {
        if (options.length > 0) {
            setSelectedOption(options[options.length - 1]);
        }
    }, [options]);

    const handleOptionClick = (option: string | undefined) => {
        setSelectedOption(option);
        setIsOpen(false);
    };
    const handleDropdownClick = () => {
        setIsOpen(!isOpen);
    };

    return (
        <div className={`${styles.dropdown_wrapper} ${isOpen ? styles.menu_opened : ""}`}>
            <button onClick={handleDropdownClick} className={className || styles.dropdown_button}>
                {selectedOption}
            </button>
            <div className={`${styles.dropdown_mask} ${isOpen ? styles.menu_opened : ""}`}>
                <div className={styles.dropdown_content}>
                    {options.map((option, index) => (
                        <div key={index} onClick={() => handleOptionClick(option)} className={styles.dropdown_item}>
                            {option}
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
};
export default Dropdown;
