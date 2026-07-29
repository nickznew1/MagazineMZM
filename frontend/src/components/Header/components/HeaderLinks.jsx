import {memo} from "react";
import "./styles.css"


const HeaderLinks = () => {
    return (
        <div className = "header__links">
            <div className ="links__email">
                <span className ="header-links-email">nickznew1@gmail.com</span>
                <span className ="header-links-email-misc">Для заявок и обратной связи</span>
            </div>
            <div className ="links__phone">
                <span className ="header-links-phone">7(999)999-99-99</span>
                <span className ="header-links-phone-misc">Звонок по России бесплатный</span>
            </div>
        </div>
    )
}

export default memo(HeaderLinks);