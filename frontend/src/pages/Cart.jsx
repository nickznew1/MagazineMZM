
import { memo } from "react";
import {UserCart} from "../components/Cart/hooks/userCart";
import CartHeader from "../components/Cart/components/CartHeader";
import CartItems from "../components/Cart/components/CartItems";
import PageNotFound from "../features/404NotFound/PageNotFound";
import NotAuthorized from "../features/NotAuthorized/NotAuthorized";
import LoadingScreen from "../features/LoadingScreen/LoadingScreen";


const Cart = () => {

    const cart = UserCart()
    if (cart.loading){
        return <div>loading...</div>

    }
    if (cart.error){
        return(
            <NotAuthorized/>
        )
    }

    return (

        <div className="user-cart-wrapper">
            <div className="user-cart">
            <CartHeader/>
                <CartItems
                    cart = {cart}
                    setCart = {cart.setCart}
                />
       </div>
        </div>
    )

}

export default memo(Cart);