import {useNavigate} from "react-router";
import "./styles.css"
import {memo} from "react";
import NotAuthorized from "../../../features/NotAuthorized/NotAuthorized";


const CartEmpty = ({cart}) => {
    const navigate = useNavigate()



    return (
        <>
        {cart ? (
        <div className="cart-empty">
            <h1>Ваша корзина пуста. Перейдите в каталог!</h1>
            <button onClick={() => navigate("/")}>В каталог</button>
        </div>
            ) : (
                <NotAuthorized />
            )}
            </>
    )
}

export default memo(CartEmpty)