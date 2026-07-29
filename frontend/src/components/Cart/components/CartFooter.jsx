import {useNavigate} from "react-router";
import "./styles.css"
import {memo} from "react";

const CartFooter = ({cart}) => {
    const {user} = cart
    const navigate = useNavigate()

    return (
        <div className="cart-footer">
            <div className="cart-summary">
                <div className="summary-line">
                    <span>Количество единиц товара:</span>
                    <span>
                {user.reduce((sum, item) => sum + item.count, 0)}
              </span>
                </div>
                <div className="summary-line total">
                    <span>Примерная сумма: </span>
                    <span>от {user.reduce((sum, item) => sum + item.count * item.price, 0) + 1000}{" "} ₽
              </span>
                </div>
                <button className="checkout-btn" onClick={()=>navigate('/checkout')}>Перейти к оформлению</button>
            </div>
        </div>
    )
}

export default memo(CartFooter);