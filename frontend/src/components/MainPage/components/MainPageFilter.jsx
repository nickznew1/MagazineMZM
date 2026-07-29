import {memo, useState} from "react";
import {motion, AnimatePresence} from "framer-motion";
import "./style.css"


const MainPageFilter = ({ changeActive, activeCategory, changeSecondary, activeSecondary, secondaryCategory, types, user }) => {


    const [isActive, setIsActive] = useState(false);

    const handleFilter = (category) => {
        if (category && !isActive) {
            changeActive(category);
            changeSecondary(null);
            setIsActive(true);
        }
        else {
            changeActive("all");
            setIsActive(false);
        }
    };

    return (
        <div className="filter">
            <div className="filter__controls">
                {Object.entries(types).map(([key, value]) => (
                    <div className="filter-item" key={key}>
                        <button
                            onClick={() => handleFilter(key)}
                            className={activeCategory === key ? "active" : ""}
                        >
                            {value.title}
                            <svg className = "filter-dropdown-icon"
                                 viewBox="0 0 24 24"
                                 width="24" height="24"
                                 stroke="currentColor"
                                 strokeWidth="2" fill="none"
                                 strokeLinecap="round"
                                 strokeLinejoin="round">
                                <polyline points="6 9 12 15 18 9"></polyline>
                            </svg>
                        </button>

                        <AnimatePresence>
                            {activeCategory === key && secondaryCategory.length > 0 && (
                                <motion.div
                                    initial={{ height: 0, opacity: 0 }}
                                    animate={{ height: "auto", opacity: 1 }}
                                    exit={{ height: 0, opacity: 0 }}
                                    transition={{ duration: 0.3 }}
                                    className="secondary-filter"
                                >
                                    {secondaryCategory.map((secondary) => (
                                        <button
                                            key={secondary}
                                            onClick={() => changeSecondary(secondary)}
                                            className={
                                                activeSecondary === secondary
                                                    ? "active"
                                                    : ""
                                            }
                                        >
                                            {secondary}
                                        </button>
                                    ))}
                                </motion.div>
                            )}
                        </AnimatePresence>
                    </div>
                ))}
            </div>
        </div>
    );

};
export default memo(MainPageFilter);