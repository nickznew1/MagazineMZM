import {useNavigate} from "react-router";
import {memo, useState} from "react";
import {AnimatePresence, motion} from "framer-motion";
import CartFooter from "./CartFooter";
import CartEmpty from "./CartEmpty";
import "./styles.css"
import {config} from "../../../config";

const CartItems = ({cart}) => {
    const {user, setCart} = cart
    const [isDeleting, setIsDeleting] = useState(false)
    const [activeSpec, setActiveSpec] = useState(false)

    const handleSpec =(id) =>{
        setActiveSpec(activeSpec === id ? null : id)
    }

    const handleSubmitIncrementCalc = async (item) =>{
        try {
            const info ={
                item_spec_id:item.item_spec_id,
                item_id: item.item_id,
                count: item.count,
                id:item.id
            }
            const response = await fetch (`${config.apiUrl}/cart/calc/`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(info),
            })

            if (response.ok){
                const updatedUser = user.map(i =>
                    i.item_spec_id === item.item_spec_id ? { ...i, count: i.count + 1 } : i
                );
                setCart(updatedUser)
            }
        } catch (err) {
            throw new Error(err)
        }
    }

    const handleSubmitDecrementCalc = async (item) =>{
        if (item.count <= 1){
            return
        }
        try {
            const info ={
                item_spec_id:item.item_spec_id,
                item_id: item.item_id,
                count: item.count,
                id:item.id
            }
            const response = await fetch (`${config.apiUrl}/cart/calc/`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(info),
            })

            if (response.ok){
                const updatedUser = user.map(i =>
                    i.item_spec_id === item.item_spec_id ? { ...i,count: i.count - 1 } : i
                );
                setCart(updatedUser)
            }
        } catch (err) {
           throw new Error(err)
        }
    }

    const handleSubmitDelete = async (item) =>{
        setIsDeleting(true)
        try {
            const info = {
                item_spec_id:item.item_spec_id,
                item_id: item.item_id,
                id: item.id,
            }
            const response = await fetch(`${config.apiUrl}/cart/delete/`, {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(info),
            })

            const result = await response.json();

            if (response.ok) {
                setCart(result);

            }

        } catch (error) {
            throw new Error(error)
        } finally{
            setIsDeleting(false)
        }
    }

    return (
        <>
        {user && user.length > 0 ? (
            <>
                <div className="cart-items-list">
                    {user.map((item) => (
                        <div className="cart-item" key={item.item_spec_id}>
                            <button
                                className="cart-item__remove"
                                onClick={() => !isDeleting && handleSubmitDelete(item)}
                                aria-label="Удалить товар"
                            >
                                ✕
                            </button>
                            <div className="cart-item__image">
                                <img
                                    src={`${config.imgServerUrl}${item.item_picture}`}
                                    alt={item.name}
                                />
                            </div>
                            <div className="cart-item__info">
                                <h2>{item.name}</h2>
                                <span className="cart-item__type">{item.item_type}</span>
                                <p className = "cart_item__secondary-type">Тип: {item.item_secondary_type}</p>
                                <button className = {activeSpec === item.item_spec_id ? "active" : ""} onClick = {()=>handleSpec(item.item_spec_id)}>Показать выбранную спецификацию
                                    <svg className = "filter-dropdown-icon" viewBox="0 0 24 24" width="16" height="16" stroke="currentColor"
                                         strokeWidth="2" fill="none" strokeLinecap="round" strokeLinejoin="round">
                                        <polyline points="6 9 12 15 18 9"></polyline>
                                    </svg>
                                </button>
                                <AnimatePresence>
                                    {activeSpec === item.item_spec_id &&
                                        <motion.div
                                            initial = {{ height:0, opacity:0}}
                                            animate ={{height:"auto", opacity:1}}
                                            exit={{height:0, opacity:0}}
                                            transition={{duration: 0.3}}
                                            className = "cart_item__specs-wrapper">
                                            <div className = "cart_item__specs">
                                                {Object.entries(item.props).map(([key,value]) => (
                                                    <>
                                                        <p key ={key} id ="key">{key}:</p>
                                                        <p key ={value} id ="value">{value}</p>
                                                    </>
                                                ))}
                                            </div>
                                        </motion.div>
                                    }
                                </AnimatePresence>
                            </div>
                            <div className="cart-item__controls">
                                <div className="quantity-selector">
                                    <button
                                        className="qty-btn qty-minus"
                                        onClick={() => handleSubmitDecrementCalc(item)}
                                    >
                                        -
                                    </button>
                                    <span className="qty-value">{item.count}</span>
                                    <button
                                        className="qty-btn qty-plus"
                                        onClick={() => handleSubmitIncrementCalc(item)}
                                    >
                                        +
                                    </button>
                                </div>
                            </div>
                        </div>
                    ))}
                </div>

                <CartFooter
                    cart = {cart}
                />
            </>
        ) : (
            <CartEmpty
            cart = {cart}/>
        )}
        </>
    )
}

export default memo(CartItems);