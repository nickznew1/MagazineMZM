

import { memo } from "react";
import HeaderCart from "../components/Header/components/HeaderCart";
import HeaderLinks from "../components/Header/components/HeaderLinks";
import HeaderProfile from "../components/Header/components/HeaderProfile";
import HeaderMain from "../components/Header/components/HeaderMain";
import "../components/Header/components/styles.css"



const Header = () => {
    return (
        <header>
            <div className = "header__controls">
                <HeaderMain/>
                <HeaderLinks />
                <HeaderCart />
               <HeaderProfile/>
            </div>
        </header>
    )
}

export default memo(Header)