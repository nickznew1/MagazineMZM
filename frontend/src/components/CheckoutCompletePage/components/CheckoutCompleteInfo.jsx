import "../components/styles.css"
import {memo} from "react";


const CheckoutCompletedInfo = ({checkoutComplete}) =>{
    const checkout = checkoutComplete;
    return (
        <div className="checkout-complete-info-wrapper">
            <div className="checkout-complete-info">
                <span className = "header">Ваше имя: </span>
                <span>{checkout.first_name}</span>
            </div>
            <div className="checkout-complete-info">
                <span className = "header">Ваша фамилия: </span>
                <span>{checkout.second_name}</span>
            </div>
            <div className="checkout-complete-info">
                <span className = "header">Ваш город: </span>
                <span>{checkout.city}</span>
            </div>
            <div className="checkout-complete-info">
                <span className = "header">Название компании: </span>
                <span>{checkout.company}</span>
            </div>
            <div className="checkout-complete-info">
                <span className = "header">Ваш адрес: </span>
                <span>{checkout.address}</span>
            </div>
            <div className="checkout-complete-info">
                <span className = "header">Ваш email: </span>
                <span>{checkout.email}</span>
            </div>
            <div className="checkout-complete-info">
                <span className = "header">Ваш номер телефона: </span>
                <span>{checkout.phone_number}</span>
            </div>
        </div>
    )
}

export default memo(CheckoutCompletedInfo);