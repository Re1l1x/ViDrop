"use client";
import styles from "./Dropdown.module.css";
import { useState, useEffect } from "react";
interface Props {
    options: number[] | undefined;
}

const Dropdown = ({ options = [] }: Props) => {
    const [isOpen, setIsOpen] = useState(false);
    const [selectedOption, setSelectedOption] = useState<number | undefined>(undefined);

    useEffect(() => {
        if (options.length > 0) {
            setSelectedOption(options[options.length - 1]);
        }
    }, [options]);

    const handleOptionClick = (option: number | undefined) => {
        setSelectedOption(option);
        setIsOpen(false);
    };
    const handleDropdownClick = () => {
        setIsOpen(!isOpen);
    };

    return (
        <div className={styles.dropdown_wrapper}>
            <button onClick={handleDropdownClick} className={styles.dropdown_button}>
                {selectedOption}
            </button>
            {isOpen && (
                <div className={styles.dropdown_content}>
                    {options.map((option, index) => (
                        <div key={index} onClick={() => handleOptionClick(option)} className={styles.dropdown_item}>
                            {option}
                        </div>
                    ))}
                </div>
            )}
        </div>
    );
};
export default Dropdown;
