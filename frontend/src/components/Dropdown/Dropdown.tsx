"use client";
import styles from "./Dropdown.module.css";
import { useState, useEffect, useRef } from "react";

function useOutsideClick(ref: React.RefObject<HTMLDivElement | null>, onOutsideClick: () => void) {
    useEffect(() => {
        function handleClickOutside(event: MouseEvent) {
            if (ref.current && !ref.current.contains(event.target as Node)) {
                onOutsideClick();
            }
        }
        document.addEventListener("mousedown", handleClickOutside);
        return () => {
            document.removeEventListener("mousedown", handleClickOutside);
        };
    }, [ref, onOutsideClick]);
}

interface Props {
    className?: string;
    options: string[] | undefined;
    setSelectedOption: (value: string | undefined) => void;
    selectedOption?: string | undefined;
    postfix?: string | undefined;
}

const Dropdown = ({ options = [], className, setSelectedOption, selectedOption, postfix }: Props) => {
    const [isOpen, setIsOpen] = useState(false);
    const wrapperRef = useRef<HTMLDivElement | null>(null);

    useOutsideClick(wrapperRef, () => {
        if (isOpen) {
            setIsOpen(false);
        }
    });

    useEffect(() => {
        if (!selectedOption && options.length > 0) {
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
        <div ref={wrapperRef} className={`${styles.dropdown_wrapper} ${isOpen ? styles.menu_opened : ""}`}>
            <button onClick={handleDropdownClick} className={className || styles.dropdown_button}>
                {selectedOption}
                {postfix}
            </button>
            <div className={`${styles.dropdown_mask} ${isOpen ? styles.menu_opened : ""}`}>
                <div className={styles.dropdown_content}>
                    {options.map((option, index) => (
                        <div key={index} onClick={() => handleOptionClick(option)} className={styles.dropdown_item}>
                            {option}
                            {postfix}
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
};
export default Dropdown;
