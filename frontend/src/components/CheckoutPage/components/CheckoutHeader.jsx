import "./styles.css"
import {memo} from "react";


const CheckoutHeader = () => {
    return(
        <div className = "checkout-items-header">
            <h1 id ="header">Ваши товары</h1>
        </div>
    )
}

export default memo(CheckoutHeader);