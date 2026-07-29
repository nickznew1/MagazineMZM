import "./styles.css"
import {memo} from "react";

const CartHeader = () => {


    return (
       <div className="cart-header">
            <h1>Корзина</h1>
        </div>
)
}

export default memo(CartHeader);